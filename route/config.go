package route

import (
	"backend/model"
	"backend/util"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

type ConfigData struct {
	AddressList []string   ` json:"addressList"`
	Compilers   [][]string ` json:"compilers"`
	Recommend   string     ` json:"recommend"`
	Announce    string     `json:"announce"`
}

func getConfig(c *gin.Context) {
	var config model.Config
	util.GetDB().Get(&config, "SELECT * FROM config")
	var compilers [][]string
	err := json.Unmarshal([]byte(config.Compilers), &compilers)
	if err != nil {
		NewResult(c).Fail("服务端异常")
		return
	}
	NewResult(c).Success("success", map[string]interface{}{
		"config": ConfigData{
			AddressList: config.AddressList,
			Compilers:   compilers,
			Recommend:   config.Recommend,
			Announce:    config.Announce,
		},
	})
}
func updateConfig(c *gin.Context) {
	var reqData ConfigData
	if err := c.ShouldBindJSON(&reqData); err != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	compilers, err := json.Marshal(reqData.Compilers)
	if err != nil {
		NewResult(c).Fail("服务端异常")
		return
	}
	util.GetDB().MustExec("UPDATE config SET address_list=$1,compilers=$2,recommend=$3,announce=$4",
		pq.Array(reqData.AddressList), compilers, reqData.Recommend, reqData.Announce,
	)
	NewResult(c).Success("success", map[string]interface{}{
		"config": reqData,
	})
}
func addConfigRoute(r *gin.Engine) {
	r.GET("/config", getConfig)
	r.POST("/config", updateConfig)
}
