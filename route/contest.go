package route

import (
	"backend/model"
	"backend/util"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func addContest(c *gin.Context) {
	domainID, err := getDomainID(c)
	if err != nil {
		NewResult(c).Fail("域参数错误")
		return
	}
	userID, _ := c.Get("userID")
	type ReqData struct {
		Title        string `json:"title"  binding:"required"`
		Typee        string `json:"type" binding:"required"`
		Participants []int  `json:"participants" binding:"required"`
		ContestID    int    `json:"contestID"`
		Desc         string `json:"desc"  binding:"required"`
		Public       *bool  `json:"pub"  binding:"required"`
		StartTime    string `json:"start"   binding:"required"`
		EndTime      string `json:"end"   binding:"required"`
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
	if reqData.ContestID == 0 {
		util.GetDB().MustExec(`
		INSERT INTO contest(
			title,description,public,type,participant_num,
			creator_id,domain_id,start_time,
			end_time,create_time,update_time,is_deleted
		) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,false)`,
			reqData.Title, reqData.Desc, reqData.Public, reqData.Typee,
			len(reqData.Participants), userID, domainID, start, end, time.Now(), time.Now(),
		)
	} else {
		util.GetDB().MustExec(`
		UPDATE homework
		SET title=$1,description=$2,public=$3,type=$4,participant_num=$5,
			start_time=$6,end_time=$7,update_time=$8
		WHERE id=$9
		`,
			reqData.Title, reqData.Desc, reqData.Public, reqData.Typee,
			len(reqData.Participants), start, end, time.Now(), reqData.ContestID,
		)
	}
	NewResult(c).Success("", nil)
}

func getContests(c *gin.Context) {
	domainID, err := getDomainID(c)
	if err != nil {
		NewResult(c).Fail("参数错误")
	}
	pageNumStr := c.DefaultQuery("page", "1")
	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	if pageNum < 1 {
		pageNum = 1
	}
	var contests []model.Contest
	util.GetDB().Select(&contests, `
		SELECT id,title,start_time,end_time,type,participant_num
		FROM contest 
		WHERE domain_id=$1 AND is_deleted = false 
		LIMIT $2 OFFSET $3`,
		domainID, PAGE_SIZE, (pageNum-1)*PAGE_SIZE)
	ret_contests := make([]map[string]interface{}, len(contests))
	for i, contest := range contests {
		ret_contests[i] = map[string]interface{}{
			"id":             contest.ID,
			"title":          contest.Title,
			"startTime":      contest.StartTime,
			"endTime":        contest.EndTime,
			"type":           contest.Typee,
			"participantNum": contest.ParticipantNum,
		}
	}
	NewResult(c).Success("", map[string]interface{}{"contests": ret_contests})
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
		NewResult(c).Fail("不存在该作业")
	} else {
		// TODO 拿到作业里的问题列表
		NewResult(c).Success("", map[string]interface{}{
			"contest": map[string]interface{}{
				"id":        contest.ID,
				"title":     contest.Title,
				"desc":      contest.Desc,
				"creatorID": contest.CreatorID,
				"startTime": contest.StartTime,
				"endTime":   contest.EndTime,
				"type":      contest.Typee,
			},
		})
	}
}

func addContestRoute(r *gin.Engine) {
	api := r.Group("/contest")
	api.GET("/:id", getContestByID)
	api.GET("/list", getContests)
	api.POST("", addContest)
}
