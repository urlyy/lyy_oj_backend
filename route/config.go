package route

import (
	"backend/model"
	"backend/util"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func getConfig(c *gin.Context) {
	var config model.Config
	util.GetDB().Get(&config, "SELECT * FROM config")
	NewResult(c).Success("success", map[string]interface{}{
		"config": config,
	})
}
func updateConfig(c *gin.Context) {
	reqData := model.Config{}
	if err := c.ShouldBindJSON(&reqData); err != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	fmt.Println(pq.Array(reqData.AddressList))
	util.GetDB().MustExec("UPDATE config SET address_list=$1,compilers=$2,recommend=$3,announce=$4",
		pq.Array(reqData.AddressList), pq.Array(reqData.Compilers), reqData.Recommend, reqData.Announce,
	)
	NewResult(c).Success("success", map[string]interface{}{
		"config": reqData,
	})
}
func addConfigRoute(r *gin.Engine) {
	r.GET("/config", getConfig)
	r.POST("/config", updateConfig)
}
