package route

import (
	"backend/model"
	"backend/util"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

const (
	ROLE_OWNER_ID   = 1
	ROLE_DEFAULT_ID = 2
)

func getDomainsByUserID(c *gin.Context) {
	userID, _ := c.Get("userID")
	domains := make([]model.Domain, 0)
	util.GetDB().Select(&domains, `
	SELECT id,name from domain
	WHERE id IN (
		SELECT domain_id FROM domain_user
		WHERE user_id=$1 AND is_deleted=false
	) ORDER BY id DESC`, userID)
	retDomains := make([]map[string]interface{}, len(domains))
	for idx, d := range domains {
		retDomains[idx] = map[string]interface{}{
			"id":   d.ID,
			"name": d.Name,
		}
	}
	NewResult(c).Success(
		"",
		map[string]interface{}{
			"domains": retDomains,
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
		fmt.Println(err)
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
		fmt.Println(err)
		NewResult(c).Fail("无权限获取该域信息或域不存在")
	} else {
		NewResult(c).Success(
			"",
			map[string]interface{}{
				"domain": map[string]interface{}{
					"id":         domain.ID,
					"name":       domain.Name,
					"announce":   domain.Announce,
					"recommend":  domain.Recommend,
					"permission": permission,
					"ownerID":    domain.OwnerID,
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
		Announce  string `json:"announce"`
		Recommend string `json:"recommend"`
		Name      string `json:"name"  binding:"required"`
	}
	reqData := ReqData{}
	if err := c.ShouldBindJSON(&reqData); err != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	userID, _ := c.Get("userID")
	util.GetDB().MustExec("UPDATE domain SET announce=$1,name=$2,recommend=$3 WHERE id=$4 AND owner_id=$5", reqData.Announce, reqData.Name, reqData.Recommend, domainID, userID)
	NewResult(c).Success("", nil)
}

const (
	DOMAIN_USER_PAGE_SIZE = 25
)

func getDomainUsers(c *gin.Context) {
	domainID, err1 := getPathInt(c, "id")
	curPageStr := c.DefaultQuery("page", "1")
	trueID := c.DefaultQuery("trueID", "")
	username := c.DefaultQuery("username", "")
	curPage, err2 := strconv.Atoi(curPageStr)
	if err1 != nil || err2 != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	if curPage < 1 {
		curPage = 1
	}
	type UserRole struct {
		UserID   int    `db:"user_id" json:"userID"`
		TrueID   int    `db:"true_id" json:"trueID"`
		RoleID   int    `db:"role_id" json:"roleID"`
		Username string `db:"username" json:"username"`
	}
	params := []interface{}{domainID, (curPage - 1) * DOMAIN_USER_PAGE_SIZE, DOMAIN_USER_PAGE_SIZE}
	where := ""
	countParams := []interface{}{domainID}
	countWhere := ""
	if trueID != "" {
		params = append(params, "%"+trueID+"%")
		where += fmt.Sprintf(" AND true_id LIKE $%d", len(params))
		countParams = append(countParams, "%"+trueID+"%")
		countWhere += fmt.Sprintf(" AND true_id LIKE $%d", len(countParams))
	}
	if username != "" {
		params = append(params, "%"+username+"%")
		where += fmt.Sprintf(" AND username LIKE $%d", len(params))
		countParams = append(countParams, "%"+username+"%")
		countWhere += fmt.Sprintf(" AND username LIKE $%d", len(countParams))
	}
	users := make([]UserRole, 0)
	sql := fmt.Sprintf(`
		SELECT s.user_id,s1.true_id,s1.username,s.role_id FROM (
			SELECT user_id,role_id,id
			FROM domain_user
			WHERE domain_id=$1 AND is_deleted=false
			ORDER BY id DESC
		)s INNER JOIN (
			SELECT true_id,id,username
			FROM "user"
			WHERE 1=1 %s
		)s1 ON s1.id = s.user_id
		OFFSET $2 LIMIT $3`, where)
	err := util.GetDB().Select(&users, sql, params...)
	if err != nil {
		NewResult(c).Fail("数据库错误")
	}
	var count int
	countSql := fmt.Sprintf(`
		SELECT COUNT(*) FROM (
			SELECT user_id,role_id,id
			FROM domain_user
			WHERE domain_id=$1 AND is_deleted=false
			ORDER BY id DESC
		)s INNER JOIN (
			SELECT true_id,id,username
			FROM "user"
			WHERE 1=1 %s
		)s1 ON s1.id = s.user_id`, where)
	util.GetDB().Get(&count, countSql, countParams...)
	pageNum := math.Ceil(float64(count) / float64(DOMAIN_USER_PAGE_SIZE))
	NewResult(c).Success("", map[string]interface{}{
		"users":   users,
		"pageNum": pageNum,
	})
}

func getDomainRoles(c *gin.Context) {
	domainID, err := getPathInt(c, "id")
	if err != nil {
		fmt.Println(err)
		NewResult(c).Fail("参数错误")
		return
	}

	roles := make([]model.Role, 0)
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

func addDomainUser(c *gin.Context) {
	domainID, err1 := getPathInt(c, "id")
	userID, err2 := getPathInt(c, "uid")
	if err1 != nil || err2 != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	var count int
	util.GetDB().Get(&count, `
	SELECT COUNT(*)
	FROM domain_user WHERE domain_id=$1 AND user_id=$2 `, domainID, userID)
	if count == 0 {
		util.GetDB().MustExec(`INSERT INTO domain_user(user_id,domain_id,role_id,is_deleted) VALUES($1,$2,$3,$4)`, userID, domainID, ROLE_DEFAULT_ID, false)
	} else {
		util.GetDB().MustExec(`UPDATE domain_user SET is_deleted=false WHERE domain_id=$1 AND user_id=$2 AND role_id=$3`, domainID, userID, ROLE_DEFAULT_ID)
	}
	NewResult(c).Success("", map[string]interface{}{
		"roleID": ROLE_DEFAULT_ID,
	})
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
	permissions := make([]model.Permission, 0)
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
	have, err4 := getPathInt(c, "have")
	if err1 != nil || err2 != nil || err3 != nil || err4 != nil || bit < 0 || (have != 0 && have != 1) {
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

func removeDomain(c *gin.Context) {
	domainID, err := getPathInt(c, "id")
	if err != nil {
		NewResult(c).Fail(err.Error())
		return
	}
	util.GetDB().MustExec(`
		UPDATE domain SET is_deleted=true WHERE id=$1
	`, domainID)
	NewResult(c).Success("", nil)
}

func removeDomainRole(c *gin.Context) {
	domainID, err1 := getPathInt(c, "id")
	roleID, err2 := getPathInt(c, "rid")
	if err1 != nil || err2 != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	util.GetDB().MustExec(`
		UPDATE role SET is_deleted=true WHERE domain_id=$1 AND id=$2
	`, domainID, roleID)
	util.GetDB().MustExec(`
		UPDATE domain_user 
		SET role_id=$1 WHERE domain_id=$2 AND role_id=$3
	`, ROLE_DEFAULT_ID, domainID, roleID)
	NewResult(c).Success("", nil)
}

func addDomain(c *gin.Context) {
	type ReqData struct {
		Name    string `json:"name"`
		OwnerID int    `json:"ownerID"`
	}
	reqData := ReqData{}
	if err := c.ShouldBindJSON(&reqData); err != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	var config model.Config
	util.GetDB().Get(&config, `SELECT * FROM config`)
	var newDomainID int
	err := util.GetDB().QueryRow(`
		INSERT INTO domain(name,owner_id,announce,recommend,create_time,update_time,is_deleted)
		VALUES($1,$2,$3,$4,$5,$5,false)
		RETURNING id
	`, reqData.Name, reqData.OwnerID, config.Announce, config.Recommend, time.Now(),
	).Scan(&newDomainID)
	if err != nil {
		fmt.Println(err)
		NewResult(c).Fail("服务端异常")
		return
	}
	util.GetDB().MustExec(`
		INSERT INTO domain_user(user_id,domain_id,role_id,is_deleted)
		VALUES($1,$2,$3,false)
	`, reqData.OwnerID, newDomainID, ROLE_OWNER_ID)
	NewResult(c).Success("", nil)
}

func addDomainRoute(r *gin.Engine) {
	r.GET("/permission", getPermissions)
	api := r.Group("/domain")
	api.GET("/:id", getDomainByID)
	api.GET("/list", getDomainsByUserID)
	api.POST("/:id/profile", changeDomainProfile)
	api.DELETE("/:id", removeDomain)
	api.POST("", addDomain)
	api.GET("/:id/users", getDomainUsers)
	api.GET("/:id/role", getDomainRoles)
	api.POST("/:id/user/delete", removeDomainUsers)
	api.POST("/:id/user/:uid", addDomainUser)
	api.POST("/:id/user/role", changeDomainUsersRole)
	api.POST("/:id/role", upsertDomainRole)
	api.DELETE("/:id/role/:rid", removeDomainRole)
	api.POST("/:id/role/:rid/permission/:bit/:have", changeDomainRolePermission)

}
