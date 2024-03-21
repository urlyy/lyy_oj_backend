package route

import (
	"backend/model"
	"backend/util"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func addHomework(c *gin.Context) {
	domainID, err := getDomainID(c)
	if err != nil {
		NewResult(c).Fail("域参数错误")
		return
	}
	userID, _ := c.Get("userID")
	type ReqData struct {
		Title     string `json:"title"  binding:"required"`
		ProblemID int    `json:"problemID"`
		Desc      string `json:"desc"  binding:"required"`
		Public    *bool  `json:"pub"  binding:"required"`
		StartTime string `json:"start"   binding:"required"`
		EndTime   string `json:"end"   binding:"required"`
	}
	reqData := ReqData{}
	if err := c.ShouldBindJSON(&reqData); err != nil {
		fmt.Println(err)
		NewResult(c).Fail("参数错误")
		return
	}
	start, err1 := parseTime(reqData.StartTime)
	end, err2 := parseTime(reqData.EndTime)
	if err1 != nil || err2 != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	if reqData.ProblemID == 0 {
		util.GetDB().MustExec(`
		INSERT INTO homework(
			title,description,public,
			creator_id,domain_id,start_time,
			end_time,create_time,update_time,is_deleted
		) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,false)`,
			reqData.Title, reqData.Desc, reqData.Public, userID,
			domainID, start, end, time.Now(), time.Now(),
		)
	} else {
		util.GetDB().MustExec(`
		UPDATE homework
		SET title=$1,description=$2,public=$3,
			start_time=$4,end_time=$5,update_time=$6
		WHERE id=$7
		`,
			reqData.Title, reqData.Desc, reqData.Public,
			start, end, time.Now(), reqData.ProblemID,
		)
	}
	NewResult(c).Success("", nil)
}

func getHomeworks(c *gin.Context) {
	domainID, err := getDomainID(c)
	if err != nil {
		NewResult(c).Fail("参数错误")
	}
	pageNumStr := c.DefaultQuery("page", "1")
	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil {
		NewResult(c).Fail("参数错误")
	} else {
		if pageNum < 1 {
			pageNum = 1
		}
		var homeworks []model.Homework
		util.GetDB().Select(&homeworks,
			`SELECT id,title,start_time,end_time
			FROM homework 
			WHERE domain_id=$1 AND is_deleted = false 
			LIMIT $2 OFFSET $3`,
			domainID, PAGE_SIZE, (pageNum-1)*PAGE_SIZE)
		ret_homeworks := make([]map[string]interface{}, len(homeworks))
		for i, problem := range homeworks {
			ret_homeworks[i] = map[string]interface{}{
				"id":        problem.ID,
				"title":     problem.Title,
				"startTime": problem.StartTime,
				"endTime":   problem.EndTime,
			}
		}
		NewResult(c).Success("", map[string]interface{}{"homeworks": ret_homeworks})
	}
}

func getHomeworkByID(c *gin.Context) {
	type Params struct {
		HomeworkID string `uri:"id" binding:"required"`
	}
	var params Params
	if err := c.ShouldBindUri(&params); err != nil {
		NewResult(c).Fail("参数错误")
	} else {
		var homework model.Homework
		err := util.GetDB().Get(&homework, "SELECT * FROM homework WHERE id = $1", params.HomeworkID)
		if err != nil {
			NewResult(c).Fail("不存在该作业")
		} else {
			// TODO 拿到作业里的问题列表
			NewResult(c).Success("", map[string]interface{}{
				"homework": map[string]interface{}{
					"id":        homework.ID,
					"title":     homework.Title,
					"desc":      homework.Desc,
					"creatorID": homework.CreatorID,
					"startTime": homework.StartTime,
					"endTime":   homework.EndTime,
				},
			})
		}
	}
}

func addHomeworkRoute(r *gin.Engine) {
	api := r.Group("/homework")
	api.GET("/:id", getHomeworkByID)
	api.GET("/list", getHomeworks)
	api.POST("", addHomework)
}
