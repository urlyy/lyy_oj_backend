package route

import (
	"backend/model"
	"backend/util"
	"fmt"

	"github.com/gin-gonic/gin"
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
			"id":       role.ID,
			"name":     role.Name,
			"desc":     role.Desc,
			"domainID": role.DomainID,
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
	util.GetDB().MustExec(`UPDATE domain_user SET is_deleted=true WHERE domain_id=$1 AND user_id IN ($2)`, domainID, reqData.UserIDs)
	NewResult(c).Success("", nil)
}

func removeDomainRoles(c *gin.Context) {
	domainID, err := getPathInt(c, "domainID")
	if err != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	type ReqData struct {
		RoleIDs []int `json:"roleIDs"  binding:"required"`
	}
	reqData := ReqData{}
	if err := c.ShouldBindJSON(&reqData); err != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	util.GetDB().MustExec(`UPDATE role SET is_deleted=true WHERE domain_id=$1 AND id IN ($2)`, domainID, reqData.RoleIDs)
	NewResult(c).Success("", nil)
}

func updateDomainUsersRole(c *gin.Context) {
	domainID, err := getPathInt(c, "id")
	if err != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	type ReqData struct {
		UserIDs []int `json:"userIDs"  binding:"required"`
		RoleID  int   `json:"roleID" binding:"required"`
	}
	reqData := ReqData{}
	if err := c.ShouldBindJSON(&reqData); err != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	util.GetDB().MustExec(`UPDATE domain_user SET role_id=$1 WHERE domain_id=$2 AND user_id IN ($3)`, reqData.RoleID, domainID, reqData.UserIDs)
	NewResult(c).Success("", nil)
}

func addDomainRoute(r *gin.Engine) {
	api := r.Group("/domain")
	api.GET("/:id", getDomainByID)
	api.GET("/list", getDomainsByUserID)
	api.POST("/:id/profile", changeDomainProfile)
	api.GET("/:id/user", getDomainUsers)
	api.GET("/:id/role", getDomainRoles)
	api.POST("/:id/user/delete", removeDomainUsers)
	api.POST("/:id/role/delete", removeDomainRoles)
	api.POST("/:id/user/role", updateDomainUsersRole)
}
