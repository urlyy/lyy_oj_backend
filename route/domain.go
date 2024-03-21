package route

import (
	"backend/model"
	"backend/util"

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
	type Params struct {
		DomainID string `uri:"id" binding:"required"`
	}
	var params Params
	if err := c.ShouldBindUri(&params); err != nil {
		NewResult(c).Fail("参数错误")
	} else {
		var domain model.Domain
		err := util.GetDB().Get(&domain, "SELECT * FROM domain WHERE id=$1", params.DomainID)
		if err != nil {
			NewResult(c).Fail("无权限获取该域信息或域不存在")
		} else {
			var permission int
			userID, _ := c.Get("userID")
			err := util.GetDB().Get(&permission,
				`SELECT permission FROM role WHERE id = (
					SELECT role_id FROM domain_user WHERE user_id=$1 AND domain_id=$2
				)`,
				userID,
				params.DomainID,
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
	}
}

func addDomainRoute(r *gin.Engine) {
	api := r.Group("/domain")
	api.GET("/:id", getDomainByID)
	api.GET("/list", getDomainsByUserID)
}
