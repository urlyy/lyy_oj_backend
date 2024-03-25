package route

import (
	"backend/model"
	"backend/util"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func getDomainsByUserID(c *gin.Context) {
	var domains []model.Domain
	userID, _ := c.Get("userID")
	util.GetDB().Select(&domains, `
	SELECT id,name from domain
	WHERE id IN (
		SELECT domain_id FROM domain_user
		WHERE user_id=$1
	)`, userID)
	ret_domains := make([]map[string]interface{}, len(domains))
	for idx, d := range domains {
		ret_domains[idx] = map[string]interface{}{
			"id":   d.ID,
			"name": d.Name,
		}
	}
	NewResult(c).Success(
		"",
		map[string]interface{}{
			"domains": ret_domains,
		},
	)
}

func getDomainByID(c *gin.Context) {
	domainID, err := getPathInt(c, "id")
	if err != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	var domain model.Domain
	err = util.GetDB().Get(&domain, "SELECT * FROM domain WHERE id=$1", domainID)
	if err != nil {
		NewResult(c).Fail("无权限获取该域信息或域不存在")
		return
	}
	var permission int
	userID, _ := c.Get("userID")
	err = util.GetDB().Get(&permission,
		`SELECT permission FROM role WHERE id = (
					SELECT role_id FROM domain_user WHERE user_id=$1 AND domain_id=$2
				)`,
		userID, domainID,
	)
	if err != nil {
		NewResult(c).Fail("无权限获取该域信息或域不存在")
	} else {
		NewResult(c).Success(
			"",
			map[string]interface{}{
				"domain": map[string]interface{}{
					"id":         domain.ID,
					"name":       domain.Name,
					"announce":   domain.Announce,
					"permission": permission,
				},
			},
		)
	}
}

// 只有域的拥有者一人可以管理
func changeDomainProfile(c *gin.Context) {
	domainID, err := getPathInt(c, "id")
	if err != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	type ReqData struct {
		Announce string `json:"announce"  binding:"required"`
		Name     string `json:"name"  binding:"required"`
	}
	reqData := ReqData{}
	if err := c.ShouldBindJSON(&reqData); err != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	userID, _ := c.Get("userID")
	fmt.Println(reqData.Announce)
	util.GetDB().MustExec("UPDATE domain SET announce=$1,name=$2 WHERE id=$3 AND owner_id=$4", reqData.Announce, reqData.Name, domainID, userID)
	if err != nil {
		NewResult(c).Fail("数据库错误")
		return
	}
	NewResult(c).Success("", nil)
}

func getDomainUsers(c *gin.Context) {
	domainID, err := getPathInt(c, "id")
	if err != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	type UserRole struct {
		UserID   int    `db:"user_id" json:"userID"`
		RoleID   int    `db:"role_id" json:"roleID"`
		Username string `db:"username" json:"username"`
	}
	var users []UserRole
	err = util.GetDB().Select(&users, `
		SELECT "user".id AS user_id,username,role_id FROM "user"
		INNER JOIN (
			SELECT user_id,role_id
			FROM domain_user
			WHERE domain_id=$1 AND is_deleted=false
		)t
		ON "user".id = t.user_id`,
		domainID,
	)
	if err != nil {
		NewResult(c).Fail("数据库错误")
	}
	NewResult(c).Success("", map[string]interface{}{
		"users": users,
	})
}

func getDomainRoles(c *gin.Context) {
	domainID, err := getPathInt(c, "id")
	if err != nil {
		fmt.Println(err)
		NewResult(c).Fail("参数错误")
		return
	}

	var roles []model.Role
	err = util.GetDB().Select(&roles, `
		SELECT * FROM role
		WHERE (domain_id=0 OR domain_id=$1) AND is_deleted=false
		ORDER BY id
		`,
		domainID,
	)
	if err != nil {
		NewResult(c).Fail("数据库错误")
		return
	}
	ret_roles := make([]map[string]interface{}, len(roles))
	for i, role := range roles {
		ret_roles[i] = map[string]interface{}{
			"id":         role.ID,
			"name":       role.Name,
			"desc":       role.Desc,
			"domainID":   role.DomainID,
			"permission": role.Permission,
		}
	}
	NewResult(c).Success("", map[string]interface{}{
		"roles": ret_roles,
	})
}

func removeDomainUsers(c *gin.Context) {
	domainID, err := getPathInt(c, "id")
	if err != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	type ReqData struct {
		UserIDs []int `json:"userIDs"  binding:"required"`
	}
	reqData := ReqData{}
	if err := c.ShouldBindJSON(&reqData); err != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	sql, args, _ := sqlx.In(`UPDATE domain_user SET is_deleted=true WHERE domain_id=? AND user_id IN (?)`, domainID, reqData.UserIDs)
	util.GetDB().MustExec(util.GetDB().Rebind(sql), args...)
	NewResult(c).Success("", nil)
}

// func removeDomainRoles(c *gin.Context) {
// 	domainID, err := getPathInt(c, "domainID")
// 	if err != nil {
// 		NewResult(c).Fail("参数错误")
// 		return
// 	}
// 	type ReqData struct {
// 		RoleIDs []int `json:"roleIDs"  binding:"required"`
// 	}
// 	reqData := ReqData{}
// 	if err := c.ShouldBindJSON(&reqData); err != nil {
// 		NewResult(c).Fail("参数错误")
// 		return
// 	}
// 	sql, args, _ := sqlx.In(`UPDATE role SET is_deleted=true WHERE domain_id=? AND id IN (?);`, domainID, reqData.RoleIDs)
// 	util.GetDB().MustExec(util.GetDB().Rebind(sql), args...)
// 	NewResult(c).Success("", nil)
// }

func changeDomainUsersRole(c *gin.Context) {
	domainID, err1 := getPathInt(c, "id")
	type ReqData struct {
		UserIDs []int `json:"userIDs"  binding:"required"`
		RoleID  int   `json:"roleID" binding:"required"`
	}
	reqData := ReqData{}
	err2 := c.ShouldBindJSON(&reqData)
	if err1 != nil || err2 != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	fmt.Println(reqData.RoleID, domainID, reqData.UserIDs)
	sql, args, _ := sqlx.In(`UPDATE domain_user SET role_id=? WHERE domain_id=? AND user_id IN (?)`, reqData.RoleID, domainID, reqData.UserIDs)
	util.GetDB().MustExec(util.GetDB().Rebind(sql), args...)
	NewResult(c).Success("", nil)
}

func changeDomainRole(c *gin.Context) {
	domainID, err1 := getPathInt(c, "id")
	roleID, err2 := getPathInt(c, "rid")
	if err1 != nil || err2 != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	roleDesc := c.Query("desc")
	roleName := c.Query("name")
	util.GetDB().MustExec("UPDATE role SET description=$1 AND name=$2 WHERE id=$3 AND domain_id=$4", roleDesc, roleName, roleID, domainID)
	NewResult(c).Success("", nil)
}

func upsertDomainRole(c *gin.Context) {
	domainID, err1 := getPathInt(c, "id")
	roleName := c.Query("name")
	if err1 != nil || roleName == "" {
		NewResult(c).Fail("参数错误")
		return
	}
	roleDesc := c.DefaultQuery("desc", "角色描述")
	roleIDStr := c.DefaultQuery("id", "")
	var roleID int
	if roleIDStr == "" {
		util.GetDB().QueryRow(`
			INSERT INTO role(name,description,domain_id,permission,is_deleted,create_time,update_time)
			VALUES($1,$2,$3,$4,$5,$6,$6)
			RETURNING id`,
			roleName, roleDesc, domainID, 0, false, time.Now(),
		).Scan(&roleID)
	} else {
		roleID, err := strconv.Atoi(roleIDStr)
		if err != nil {
			NewResult(c).Fail("参数错误")
			return
		}
		util.GetDB().MustExec(`
		UPDATE role SET name=$1,description=$2,update_time=$3
		WHERE id=$4
	`, roleName, roleDesc, time.Now(), roleID)
	}
	NewResult(c).Success("", map[string]interface{}{
		"role": map[string]interface{}{
			"id":   roleID,
			"name": roleName,
			"desc": roleDesc,
		},
	})
}

func getPermissions(c *gin.Context) {
	var permissions []model.Permission
	util.GetDB().Select(&permissions, `
		SELECT * FROM permission 
		ORDER BY bit
	`)
	ret_permissions := make(map[string]interface{}, len(permissions))
	for _, permission := range permissions {
		ret_permissions[permission.Name] = permission.Bit
	}
	NewResult(c).Success("", map[string]interface{}{
		"permissions": ret_permissions,
	})
}

func changeDomainRolePermission(c *gin.Context) {
	domainID, err1 := getPathInt(c, "id")
	roleID, err2 := getPathInt(c, "rid")
	bit, err3 := getPathInt(c, "bit")
	have, err3 := getPathInt(c, "have")
	if err1 != nil || err2 != nil || err3 != nil || bit < 0 || (have != 0 && have != 1) {
		NewResult(c).Fail("参数错误")
		return
	}
	var role model.Role
	err := util.GetDB().Get(&role, "SELECT * FROM role WHERE id=$1", roleID)
	if err != nil || (domainID != role.DomainID && role.DomainID != 0) {
		fmt.Println(err)
		NewResult(c).Fail("参数错误")
		return
	}
	newPermission := role.Permission & (^(1 << bit))
	if have == 1 {
		newPermission += 1 << bit
	}
	util.GetDB().MustExec(`
		UPDATE role SET permission=$1 WHERE id=$2
	`, newPermission, roleID)
	role.Permission = newPermission
	NewResult(c).Success("", map[string]interface{}{
		"permission": newPermission,
	})
}

func addDomainRoute(r *gin.Engine) {
	r.GET("/permission", getPermissions)
	api := r.Group("/domain")
	api.GET("/:id", getDomainByID)
	api.GET("/list", getDomainsByUserID)
	api.POST("/:id/profile", changeDomainProfile)
	api.GET("/:id/user", getDomainUsers)
	api.GET("/:id/role", getDomainRoles)
	api.POST("/:id/user/delete", removeDomainUsers)

	api.POST("/:id/user/role", changeDomainUsersRole)
	api.POST("/:id/role", upsertDomainRole)
	api.POST("/:id/role/:rid/permission/:bit/:have", changeDomainRolePermission)

}
