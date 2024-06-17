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
	CONTEST_PAGE_SIZE = 7
)

func upsertContest(c *gin.Context) {
	domainID, err1 := getQueryDomainID(c)
	userID, _ := c.Get("userID")
	type ReqData struct {
		Title      string `json:"title"  binding:"required"`
		Typee      string `json:"type" binding:"required"`
		ContestID  int    `json:"contestID"`
		Desc       string `json:"desc"  binding:"required"`
		Public     bool   `json:"pub" `
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
	if reqData.ContestID == 0 {
		util.GetDB().MustExec(`
		INSERT INTO contest(
			title,description,public,type,
			creator_id,domain_id,start_time,end_time,
			create_time,update_time,is_deleted,problem_ids
		) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$9,$10,$11)`,
			reqData.Title, reqData.Desc, reqData.Public, reqData.Typee,
			userID, domainID, start, end,
			now, false, pq.Array(reqData.ProblemIDs),
		)
	} else {
		util.GetDB().MustExec(`
		UPDATE contest
		SET title=$1,description=$2,public=$3,type=$4,
			start_time=$5,end_time=$6,update_time=$7,problem_ids=$8
		WHERE id=$9
		`,
			reqData.Title, reqData.Desc, reqData.Public, reqData.Typee,
			start, end, time.Now(), pq.Array(reqData.ProblemIDs), reqData.ContestID,
		)
	}
	NewResult(c).Success("", nil)
}

func getContests(c *gin.Context) {
	domainID, err1 := getQueryDomainID(c)
	curPageStr := c.DefaultQuery("page", "1")
	curPage, err2 := strconv.Atoi(curPageStr)
	flag := c.DefaultQuery("flag", "false")
	if err1 != nil || err2 != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	if curPage < 1 {
		curPage = 1
	}
	contests := make([]model.Contest, 0)
	extraWhere := ""
	if flag != "true" {
		extraWhere = " AND public=true"
	}
	sql := fmt.Sprintf(`
		SELECT *
		FROM contest 
		WHERE domain_id=$1 AND is_deleted = false %s
		ORDER BY start_time DESC
		LIMIT $2 OFFSET $3`, extraWhere)
	util.GetDB().Select(&contests, sql, domainID, CONTEST_PAGE_SIZE, (curPage-1)*CONTEST_PAGE_SIZE)
	ret_contests := make([]map[string]interface{}, len(contests))
	for i, contest := range contests {
		ret_contests[i] = map[string]interface{}{
			"id":        contest.ID,
			"title":     contest.Title,
			"startTime": contest.StartTime,
			"endTime":   contest.EndTime,
			"type":      contest.Typee,
		}
	}
	var count int
	countSql := fmt.Sprintf(`
		SELECT COUNT(*)
		FROM contest 
		WHERE domain_id=$1 AND is_deleted = false %s
	`, extraWhere)
	util.GetDB().Get(&count, countSql, domainID)
	pageNum := math.Ceil(float64(count) / float64(CONTEST_PAGE_SIZE))
	NewResult(c).Success("", map[string]interface{}{"contests": ret_contests, "pageNum": pageNum})
}

func getContestByID(c *gin.Context) {
	type Params struct {
		ContestID string `uri:"id" binding:"required"`
	}
	var params Params
	if err := c.ShouldBindUri(&params); err != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	var contest model.Contest
	err := util.GetDB().Get(&contest, "SELECT * FROM contest WHERE id = $1", params.ContestID)
	if err != nil {
		NewResult(c).Fail("不存在该测验")
	} else {
		pids := []int64(contest.ProblemIDs)
		NewResult(c).Success("", map[string]interface{}{
			"contest": map[string]interface{}{
				"id":         contest.ID,
				"title":      contest.Title,
				"desc":       contest.Desc,
				"creatorID":  contest.CreatorID,
				"startTime":  contest.StartTime,
				"endTime":    contest.EndTime,
				"type":       contest.Typee,
				"problemIDs": pids,
				"public":     contest.Public,
			},
		})
	}
}

func addProblemProblems(c *gin.Context) {
	contestID, _ := getPathInt(c, "id")
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
		UPDATE contest SET problem_ids=$1
		WHERE id=$2
	`, reqData.ProblemIDs, contestID)
	NewResult(c).Success("", nil)
}

func removeContest(c *gin.Context) {
	contestID := c.Param("id")
	util.GetDB().Exec("UPDATE contest SET is_deleted=true WHERE id=$1", contestID)
	NewResult(c).Success("", nil)
}

func addContestRoute(r *gin.Engine) {
	api := r.Group("/contest")
	api.GET("/:id", getContestByID)
	api.GET("/list", getContests)
	api.POST("", upsertContest)
	api.POST("/:id/problem", addProblemProblems)
	api.DELETE("/:id", removeContest)
}
