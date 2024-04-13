package route

import (
	"backend/model"
	"backend/util"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	SUBMISSION_PAGE_SIZE   = 15
	USER_RECORDS_PAGE_SIZE = 15
	RANK_PAGE_SIZE         = 40
)

func getSubmissions(c *gin.Context) {
	cuePageStr := c.DefaultQuery("page", "1")
	username := c.DefaultQuery("username", "")
	problemTitle := c.DefaultQuery("problemTitle", "")
	compiler := c.DefaultQuery("compiler", "")
	statusStr := c.DefaultQuery("status", "")
	curPage, err1 := strconv.Atoi(cuePageStr)
	domainID, err2 := getQueryDomainID(c)
	// myID, _ := c.Get("userID")
	if err1 != nil || err2 != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	if curPage < 1 {
		curPage = 1
	}
	// 动态sql
	params := []interface{}{domainID, SUBMISSION_PAGE_SIZE, (curPage - 1) * SUBMISSION_PAGE_SIZE}
	countParams := []interface{}{domainID}
	where1, userWhere, problemWhere := "", "", ""
	countWhere1, countUserWhere, countProblemWhere := "", "", ""
	if statusStr != "" {
		status, err1 := strconv.Atoi(statusStr)
		if err1 != nil {
			NewResult(c).Fail("参数错误")
			return
		}
		params = append(params, status)
		where1 += fmt.Sprintf(" AND status=$%d", len(params))
		countParams = append(countParams, status)
		countWhere1 += fmt.Sprintf(" AND status=$%d", len(countParams))
	}
	if compiler != "" {
		params = append(params, compiler)
		where1 += fmt.Sprintf(" AND lang=$%d", len(params))
		countParams = append(countParams, compiler)
		countWhere1 += fmt.Sprintf(" AND lang=$%d", len(countParams))
	}
	if username != "" {
		params = append(params, "%"+username+"%")
		userWhere += fmt.Sprintf(" AND username LIKE $%d", len(params))
		countParams = append(countParams, "%"+username+"%")
		countUserWhere += fmt.Sprintf(" AND username LIKE $%d", len(countParams))
	}
	if problemTitle != "" {
		params = append(params, "%"+problemTitle+"%")
		problemWhere += fmt.Sprintf(" AND title LIKE $%d", len(params))
		countParams = append(countParams, "%"+problemTitle+"%")
		countProblemWhere += fmt.Sprintf(" AND title LIKE $%d", len(countParams))
	}
	type Result struct {
		ID           int       `db:"id" json:"id"`
		ProblemTitle string    `db:"problem_title" json:"problemTitle"`
		UserID       int       `db:"user_id" json:"userID"`
		ProblemID    int       `db:"problem_id" json:"problemID"`
		Username     string    `db:"username" json:"username"`
		Status       int       `db:"status" json:"status"`
		MaxTime      int64     `db:"max_time" json:"maxTime"`
		MaxMemory    int64     `db:"max_memory" json:"maxMemory"`
		PassPercent  float64   `db:"pass_percent" json:"passPercent"`
		SubmitTime   time.Time `db:"submit_time" json:"submitTime"`
		Lang         string    `db:"lang" json:"lang"`
		FromType     string    `db:"from_type" json:"fromType"`
		FromID       int       `db:"from_id" json:"fromID"`
		FromName     string    `db:"from_name" json:"fromName"`
	}
	sql := fmt.Sprintf(`
		SELECT s.id,s.problem_id,s1.title AS problem_title,s.submit_time,s.status,s.lang,s.from_type,s.from_id,s.max_time,s.max_memory,s.pass_percent,s2.username
		FROM (
			SELECT * 
			FROM submission 
			WHERE domain_id=$1 %s
			ORDER BY id DESC
			LIMIT $2 OFFSET $3
		)s INNER JOIN (
			SELECT id,title 
			FROM problem 
			WHERE domain_id=$1 %s
		)s1 ON s.problem_id=s1.id
		INNER JOIN (
			SELECT username,id
			FROM "user" 
			WHERE 1=1 %s
		)s2 ON s.user_id=s2.id`, where1, problemWhere, userWhere,
	)
	resList := make([]Result, 0)
	err := util.GetDB().Select(&resList, sql, params...)
	if err != nil {
		fmt.Println(err)
		NewResult(c).Fail("数据库错误")
		return
	}
	for i, sub := range resList {
		fromName := ""
		if sub.FromType == "contest" {
			var contest model.Contest
			util.GetDB().Get(&contest, "SELECT * FROM contest WHERE id=$1", sub.FromID)
			fromName = contest.Title
		} else if sub.FromType == "homework" {
			var homework model.Homework
			util.GetDB().Get(&homework, "SELECT * FROM homework WHERE id=$1", sub.FromID)
			fromName = homework.Title
		}
		resList[i].FromName = fromName
	}
	var count int
	countSql := fmt.Sprintf(`
		SELECT s.id,s.problem_id,s1.title AS problem_title,s.submit_time,s.status,s.lang,s.from_type,s.from_id,s.max_time,s.max_memory,s.pass_percent,s2.username
		FROM (
			SELECT * 
			FROM submission 
			WHERE domain_id=$1 %s
		)s INNER JOIN (
			SELECT id,title 
			FROM problem 
			WHERE domain_id=$1 %s
		)s1 ON s.problem_id=s1.id
		INNER JOIN (
			SELECT username,id
			FROM "user" 
			WHERE 1=1 %s
		)s2 ON s.user_id=s2.id
		ORDER BY s.id DESC
		`, countWhere1, countProblemWhere, countUserWhere,
	)
	util.GetDB().Get(&count, countSql, countParams...)
	pageNum := math.Ceil(float64(count) / float64(SUBMISSION_PAGE_SIZE))
	NewResult(c).Success("", map[string]interface{}{
		"submissions": resList,
		"pageNum":     pageNum,
	})
}

func getRank(c *gin.Context) {
	domainID, err1 := getQueryDomainID(c)
	curPageStr := c.DefaultQuery("page", "1")
	username := c.DefaultQuery("username", "")
	curPage, err2 := strconv.Atoi(curPageStr)
	if err1 != nil || err2 != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	type Data struct {
		UserID   int    `json:"userID" db:"user_id"`
		Username string `json:"username"`
		TotalAC  int    `json:"totalAC" db:"total_ac"`
		Submit   int    `json:"submit" db:"submit"`
		AC       int    `json:"ac" db:"ac"`
	}
	extraWhere := ""
	params := []interface{}{domainID, ACCEPT, (curPage - 1) * RANK_PAGE_SIZE, RANK_PAGE_SIZE}
	if username != "" {
		params = append(params, "%"+username+"%")
		extraWhere += fmt.Sprintf(" AND username LIKE $%d", len(params))
	}
	sql := fmt.Sprintf(`
	SELECT s.user_id,u.username,s.submit,s.total_ac,s1.ac FROM(
		SELECT id,username
		FROM "user"
		WHERE id IN (
			SELECT user_id 
			FROM domain_user
			WHERE domain_id=$1 AND is_deleted=false
		) %s
	)u 
	INNER JOIN(
		SELECT user_id,COUNT(*) AS submit,COUNT(CASE WHEN status=0 THEN 1 ELSE NULL END) AS total_ac
		FROM submission 
		WHERE domain_id=$1
		GROUP BY user_id
	)s ON u.id=s.user_id
	INNER JOIN (
		SELECT DISTINCT ON (user_id) user_id,COUNT(*) AS ac
		FROM submission
		WHERE domain_id = $1 AND status = $2
		GROUP BY user_id,problem_id
		ORDER BY user_id,problem_id
	)s1 ON s.user_id=s1.user_id
	ORDER BY s1.ac DESC,s.submit,s
	OFFSET $3 LIMIT $4
	`, extraWhere)
	rankData := make([]Data, 0)
	err := util.GetDB().Select(&rankData, sql, params...)
	if err != nil {
		fmt.Println(err)
		NewResult(c).Fail("数据库错误")
		return
	}
	extraWhere = ""
	params = []interface{}{domainID}
	if username != "" {
		params = append(params, "%"+username+"%")
		extraWhere += fmt.Sprintf(" AND username LIKE $%d", len(params))
	}
	countSql := fmt.Sprintf(`
		SELECT COUNT(*)
		FROM "user"
		WHERE id IN (
			SELECT user_id 
			FROM domain_user
			WHERE domain_id=$1 AND is_deleted=false
		) %s`, extraWhere)
	var count int
	util.GetDB().Get(&count, countSql, params...)
	pageNum := math.Ceil(float64(count) / float64(RANK_PAGE_SIZE))
	NewResult(c).Success("", map[string]interface{}{
		"rankData": rankData,
		"pageNum":  pageNum,
		"pageSize": RANK_PAGE_SIZE,
	})
}

// 获得比赛、作业的排名
func getSubmissions4Rank(c *gin.Context) {
	typeID := c.DefaultQuery("id", "0")
	fromType := c.DefaultQuery("type", "")
	domainID, err2 := getQueryDomainID(c)
	// myID, _ := c.Get("userID")
	if err2 != nil || (fromType != "contest" && fromType != "homework") || typeID == "0" {
		NewResult(c).Fail("参数错误")
		return
	}
	submissions := make([]model.Submission, 0)
	params := []interface{}{domainID}
	extra_where := ""
	if fromType != "" {
		params = append(params, fromType, typeID)
		extra_where = fmt.Sprintf("AND from_type=$%d AND from_id=$%d", len(params)-1, len(params))
	}
	sql := fmt.Sprintf(`
		SELECT t2.user_id,t2.submit_time,t2.status,t2.id,t2.problem_id,t2.pass_percent FROM 
		(
			SELECT user_id FROM domain_user
			WHERE domain_id=$1
		)t1 
		INNER JOIN (
			SELECT * FROM submission
			WHERE domain_id=$1 %s
		)t2
		ON t2.user_id=t1.user_id
		ORDER BY t2.submit_time
	`, extra_where)
	util.GetDB().Select(&submissions, sql, params...)
	retSubmissions := make([]map[string]interface{}, len(submissions))
	for i, submission := range submissions {
		retSubmissions[i] = map[string]interface{}{
			"userID":       submission.UserID,
			"problemID":    submission.ProblemID,
			"submissionID": submission.ID,
			"status":       submission.Status,
			"submitTime":   submission.SubmitTime,
			"passPercent":  submission.PassPercent,
		}
	}
	NewResult(c).Success("", map[string]interface{}{
		"submissions": retSubmissions,
	})
}

func getSubmissionByID(c *gin.Context) {
	sid, err := getPathInt(c, "sid")
	if err != nil {
		fmt.Println(err)
		NewResult(c).Fail("参数错误")
		return
	}
	var submission model.Submission
	util.GetDB().Get(&submission, "SELECT * FROM submission WHERE id=$1", sid)
	NewResult(c).Success("", map[string]interface{}{
		"submission": map[string]interface{}{
			"id":   submission.ID,
			"code": submission.Code,
			"log":  submission.Log,
		},
	})
}

func getUserSubmissionPie(c *gin.Context) {
	userID, err := getPathInt(c, "uid")
	domainID, err2 := getQueryDomainID(c)
	if err != nil || err2 != nil {
		fmt.Println(err)
		NewResult(c).Fail("参数错误")
		return
	}
	type Data struct {
		Status int `db:"status" json:"status"`
		Count  int `db:"count" json:"count"`
	}
	pieData := make([]Data, 0)
	util.GetDB().Select(&pieData, `
	SELECT COUNT(*),status FROM submission
	WHERE user_id=$1 AND domain_id=$2 AND status!=$3
	GROUP BY status
	`, userID, domainID, JUDGING)

	resData := map[int]int{
		ACCEPT:              0,
		COMPILE_ERROR:       0,
		WRONG_ANSWER:        0,
		RUNTIME_ERROR:       0,
		TIME_LIMIT_EXCEED:   0,
		MEMORY_LIMIT_EXCEED: 0,
	}
	for _, datum := range pieData {
		resData[datum.Status] = datum.Count
	}
	var passProblemNum Data
	util.GetDB().Get(&passProblemNum, `
	SELECT COUNT(DISTINCT problem_id) 
	FROM submission
	WHERE user_id=$1 AND domain_id=$2 AND status=$3	
	`, userID, domainID, ACCEPT)
	NewResult(c).Success("", map[string]interface{}{
		"pie":            resData,
		"passProblemNum": passProblemNum.Count,
	})
}

func getUserSubmissionRecords(c *gin.Context) {
	userID, err1 := getPathInt(c, "uid")
	domainID, err2 := getQueryDomainID(c)
	curPageStr := c.DefaultQuery("page", "1")
	curPage, err3 := strconv.Atoi(curPageStr)
	filterTitle := c.DefaultQuery("title", "")
	if err1 != nil || err2 != nil || err3 != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	if curPage < 1 {
		curPage = 1
	}
	type Record struct {
		ProblemID    int       `db:"problem_id" json:"problemID"`
		ProblemTitle string    `db:"title" json:"problemTitle"`
		Count        int       `db:"count" json:"count"`
		Recent       time.Time `db:"recent" json:"recent"`
	}
	records := make([]Record, 0)
	params := []interface{}{userID, domainID, (curPage - 1) * SUBMISSION_PAGE_SIZE, SUBMISSION_PAGE_SIZE}
	extraWhere := ""
	if filterTitle != "" {
		params = append(params, "%"+filterTitle+"%")
		extraWhere += fmt.Sprintf("AND title LIKE $%d ", len(params))
	}
	sql := fmt.Sprintf(`
	SELECT s.problem_id, p.title, s.count, s.recent  
	FROM (  
		SELECT problem_id, COUNT(*), MAX(submit_time) AS recent
		FROM submission  
		WHERE user_id = $1 AND domain_id = $2
		GROUP BY problem_id  
		OFFSET $3 LIMIT $4
	)s  
	INNER JOIN problem p ON s.problem_id = p.id 
	WHERE 1=1 %s
	ORDER BY s.recent DESC`, extraWhere)
	err := util.GetDB().Select(&records, sql, params...)
	if err != nil {
		fmt.Println(err)
		NewResult(c).Fail("获取失败")
		return
	}
	params = []interface{}{userID, domainID}
	extraWhere = ""
	if filterTitle != "" {
		params = append(params, "%"+filterTitle+"%")
		extraWhere += fmt.Sprintf(" AND title LIKE $%d ", len(params))
	}
	var count int
	countSql := fmt.Sprintf(`
	SELECT COUNT(*)
	FROM (  
		SELECT problem_id, COUNT(*), MAX(submit_time) AS recent
		FROM submission  
		WHERE user_id = $1 AND domain_id = $2
		GROUP BY problem_id 
	)s  
	INNER JOIN problem p ON s.problem_id = p.id 
	WHERE 1=1 %s
	`, extraWhere)
	util.GetDB().Get(&count, countSql, params...)
	pageNum := math.Ceil(float64(count) / float64(USER_RECORDS_PAGE_SIZE))
	NewResult(c).Success("", map[string]interface{}{
		"records": records,
		"pageNum": pageNum,
	})

}

func getUserProblemSubmissions(c *gin.Context) {
	userID, err1 := getPathInt(c, "uid")
	problemID, err2 := getPathInt(c, "pid")
	domainID, err3 := getQueryDomainID(c)
	if err1 != nil || err2 != nil || err3 != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	submissions := make([]model.Submission, 0)
	util.GetDB().Select(&submissions, `
	SELECT id,status,submit_time,lang,max_memory,max_time FROM submission
	WHERE user_id=$1 AND problem_id=$2 AND domain_id=$3 
	ORDER BY submit_time DESC
	`, userID, problemID, domainID)
	res := make([]map[string]interface{}, len(submissions))
	for idx, submission := range submissions {
		res[idx] = map[string]interface{}{
			"id":         submission.ID,
			"status":     submission.Status,
			"submitTime": submission.SubmitTime,
			"maxMemory":  submission.MaxMemory,
			"maxTime":    submission.MaxTime,
			"lang":       submission.Lang,
		}
	}
	NewResult(c).Success("", map[string]interface{}{
		"submissions": res,
	})
}
func addSubmissionRoute(r *gin.Engine) {
	api := r.Group("/submission")
	api.GET("/list", getSubmissions)
	api.GET("/special/rank", getSubmissions4Rank)
	api.GET("/rank", getRank)
	api.GET("/:sid", getSubmissionByID)
	api.GET("/pie/:uid", getUserSubmissionPie)
	api.GET("/records/:uid", getUserSubmissionRecords)
	api.GET("/record/:uid/:pid", getUserProblemSubmissions)
}
