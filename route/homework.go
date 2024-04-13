package route

import (
	"backend/model"
	"backend/util"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

const (
	HOMEWORK_PAGE_SIZE = 7
)

func upsertHomework(c *gin.Context) {
	domainID, err1 := getQueryDomainID(c)

	userID, _ := c.Get("userID")
	type ReqData struct {
		Title      string `json:"title"  binding:"required"`
		HomeworkID int    `json:"homeworkID"`
		Desc       string `json:"desc"`
		Public     *bool  `json:"pub"`
		StartTime  string `json:"start"   binding:"required"`
		EndTime    string `json:"end"   binding:"required"`
		ProblemIDs []int  `json:"problemIDs" binding:"required"`
	}
	reqData := ReqData{}
	err2 := c.ShouldBindJSON(&reqData)
	if err1 != nil || err2 != nil {
		fmt.Println(err1, err2)
		NewResult(c).Fail("参数错误")
		return
	}
	start, err1 := parseTime(reqData.StartTime)
	end, err2 := parseTime(reqData.EndTime)
	if err1 != nil || err2 != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	now := time.Now()
	if reqData.HomeworkID == 0 {
		util.GetDB().MustExec(`
		INSERT INTO homework(
			title,description,public,
			creator_id,domain_id,start_time,
			end_time,create_time,update_time,problem_ids,is_deleted
		) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,false)`,
			reqData.Title, reqData.Desc, reqData.Public, userID,
			domainID, start, end, now, now, pq.Array(reqData.ProblemIDs),
		)
	} else {
		fmt.Println(reqData.ProblemIDs)
		util.GetDB().MustExec(`
		UPDATE homework
		SET title=$1,description=$2,public=$3,
			start_time=$4,end_time=$5,update_time=$6,problem_ids=$7
		WHERE id=$8
		`,
			reqData.Title, reqData.Desc, reqData.Public,
			start, end, now, pq.Array(reqData.ProblemIDs), reqData.HomeworkID,
		)
	}
	NewResult(c).Success("", nil)
}

func getHomeworks(c *gin.Context) {
	domainID, err1 := getQueryDomainID(c)
	curPageStr := c.DefaultQuery("page", "1")
	flag := c.DefaultQuery("flag", "false")
	curPage, err2 := strconv.Atoi(curPageStr)
	if err1 != nil || err2 != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	if curPage < 1 {
		curPage = 1
	}
	homeworks := make([]model.Homework, 0)
	extra_where := ""
	if flag != "true" {
		extra_where = " AND public=true"
	}
	sql := fmt.Sprintf(`
		SELECT id,title,start_time,end_time
		FROM homework 
		WHERE domain_id=$1 AND is_deleted = false %s
		ORDER BY id DESC
		LIMIT $2 OFFSET $3`, extra_where,
	)
	util.GetDB().Select(&homeworks, sql, domainID, HOMEWORK_PAGE_SIZE, (curPage-1)*HOMEWORK_PAGE_SIZE)
	ret_homeworks := make([]map[string]interface{}, len(homeworks))
	for i, problem := range homeworks {
		ret_homeworks[i] = map[string]interface{}{
			"id":        problem.ID,
			"title":     problem.Title,
			"startTime": problem.StartTime,
			"endTime":   problem.EndTime,
		}
	}
	var count int
	countSql := fmt.Sprintf(`
		SELECT COUNT(*)
		FROM homework 
		WHERE domain_id=$1 AND is_deleted = false %s
	`, extra_where)
	util.GetDB().Get(&count, countSql, domainID)
	pageNum := math.Ceil(float64(count) / float64(HOMEWORK_PAGE_SIZE))
	NewResult(c).Success("", map[string]interface{}{"homeworks": ret_homeworks, "pageNum": pageNum})
}

func getHomeworkByID(c *gin.Context) {
	type Params struct {
		HomeworkID string `uri:"id" binding:"required"`
	}
	var params Params
	if err := c.ShouldBindUri(&params); err != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	var homework model.Homework
	err := util.GetDB().Get(&homework, "SELECT * FROM homework WHERE id = $1", params.HomeworkID)
	if err != nil {
		fmt.Println(err)
		NewResult(c).Fail("不存在该作业")
	} else {
		pids := []int64(homework.ProblemIDs)
		NewResult(c).Success("", map[string]interface{}{
			"homework": map[string]interface{}{
				"id":         homework.ID,
				"title":      homework.Title,
				"desc":       homework.Desc,
				"creatorID":  homework.CreatorID,
				"startTime":  homework.StartTime,
				"endTime":    homework.EndTime,
				"problemIDs": pids,
				"public":     homework.Public,
			},
		})
	}
}

func addHomeworkProblems(c *gin.Context) {
	homeworkID, _ := getPathInt(c, "id")
	type ReqData struct {
		ProblemIDs []int `json:"problemIDs" binding:"required"`
	}
	reqData := ReqData{}
	err := c.ShouldBindJSON(&reqData)
	if err != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	util.GetDB().MustExec(`
		UPDATE problem SET problem_ids=$1
		WHERE id=$2
	`, reqData.ProblemIDs, homeworkID)
	NewResult(c).Success("", nil)
}

func removeHomework(c *gin.Context) {
	contestID := c.Param("id")
	util.GetDB().Exec("UPDATE homework SET is_deleted=true WHERE id=$1", contestID)
	NewResult(c).Success("", nil)
}

func addHomeworkRoute(r *gin.Engine) {
	api := r.Group("/homework")
	api.GET("/:id", getHomeworkByID)
	api.GET("/list", getHomeworks)
	api.POST("", upsertHomework)
	api.POST("/:id/problem", addHomeworkProblems)
	api.DELETE("/:id", removeHomework)
}
