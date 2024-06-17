package route

import (
	"backend/model"
	"backend/util"
	"encoding/json"
	"fmt"
	"math"
	"time"

	"strconv"

	"github.com/gin-gonic/gin"
)

const (
	PROBLEM_PAGE_SIZE = 20
)

func getProblems(c *gin.Context) {
	domainID, err1 := getQueryDomainID(c)
	curPageNumStr := c.DefaultQuery("page", "1")
	searchKeyword := c.DefaultQuery("keyword", "")
	diff := c.DefaultQuery("diff", "0")
	flag := c.DefaultQuery("flag", "false")
	fmt.Println("flag", flag)
	curPage, err2 := strconv.Atoi(curPageNumStr)
	if err1 != nil || err2 != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	if curPage < 1 {
		curPage = 1
	}
	problems := make([]model.Problem, 0)
	// 动态sql
	params := []interface{}{domainID, PROBLEM_PAGE_SIZE, (curPage - 1) * PROBLEM_PAGE_SIZE}
	extra_where := ""
	if flag != "true" {
		extra_where += " AND public = true"
	}
	if diff != "0" {
		params = append(params, diff)
		extra_where += fmt.Sprintf(" AND diff=$%d", len(params))
	}
	if searchKeyword != "" {
		params = append(params, "%"+searchKeyword+"%")
		extra_where += fmt.Sprintf(" AND title LIKE $%d", len(params))
	}
	sql := fmt.Sprintf(`
		SELECT id,title,diff,ac_num,submit_num
		FROM problem 
		WHERE domain_id=$1 AND is_deleted = false  %s
		ORDER BY id DESC LIMIT $2 OFFSET $3`, extra_where,
	)
	util.GetDB().Select(&problems, sql, params...)
	retProblems := make([]map[string]interface{}, len(problems))
	for i, problem := range problems {
		retProblems[i] = map[string]interface{}{
			"id":        problem.ID,
			"title":     problem.Title,
			"diff":      problem.Diff,
			"acNum":     problem.ACNum,
			"submitNum": problem.SubmitNum,
		}
	}
	var count int
	extra_where = ""
	params = []interface{}{domainID}
	if flag != "true" {
		extra_where += " AND public = true"
	}
	if diff != "0" {
		params = append(params, diff)
		extra_where += fmt.Sprintf(" AND diff=$%d", len(params))
	}
	if searchKeyword != "" {
		params = append(params, "%"+searchKeyword+"%")
		extra_where += fmt.Sprintf(" AND title LIKE $%d", len(params))
	}
	countSql := fmt.Sprintf(`
		SELECT COUNT(*)
		FROM problem 
		WHERE is_deleted = false AND public=true AND domain_id=$1 %s
		`, extra_where,
	)
	util.GetDB().Get(&count, countSql, params...)
	pageNum := math.Ceil(float64(count) / float64(PROBLEM_PAGE_SIZE))
	NewResult(c).Success("", map[string]interface{}{"problems": retProblems, "pageNum": pageNum})
}

func getProblemByID(c *gin.Context) {
	type Params struct {
		ProblemID string `uri:"id" binding:"required"`
	}
	var params Params
	if err := c.ShouldBindUri(&params); err != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	var problem model.Problem
	err := util.GetDB().Get(&problem, "SELECT * FROM problem WHERE id = $1", params.ProblemID)
	if err != nil {
		fmt.Println(err)
		NewResult(c).Fail("不存在该题目")
		return
	}
	var testCases []model.TestCase
	var valid bool
	err = json.Unmarshal([]byte(problem.TestCases), &testCases)
	if err != nil {
		testCases = make([]model.TestCase, 0)
		valid = false
	} else {
		valid = len(testCases) > 0
	}
	isEdit := c.Query("edit")
	var retCases = []model.TestCase{}
	if isEdit == "true" {
		retCases = testCases
	} else {
		for _, c := range testCases {
			if c.IsSample {
				retCases = append(retCases, c)
			}
		}
	}
	NewResult(c).Success("", map[string]interface{}{
		"problem": map[string]interface{}{
			"id":           problem.ID,
			"title":        problem.Title,
			"desc":         problem.Desc,
			"outputFormat": problem.OutFmt,
			"inputFormat":  problem.InFmt,
			"other":        problem.Other,
			"memoryLimit":  problem.MemoryLimit / 1024,
			"timeLimit":    problem.TimeLimit,
			"diff":         problem.Diff,
			"createTime":   problem.CreateTime,
			"pub":          problem.Public,
			"testCases":    retCases,
			"judgeType":    problem.JudgeType,
			"specialCode":  problem.SpecialCode,
			"valid":        valid,
			"acNum":        problem.ACNum,
			"submitNum":    problem.SubmitNum,
		},
	})
}

func upsertProblem(c *gin.Context) {
	domainID, err1 := getQueryDomainID(c)
	type ReqData struct {
		ProblemID   int              `json:"problemID"`
		JudgeType   int              `json:"judgeType"`
		Title       string           `json:"title" binding:"required"`
		Desc        string           `json:"desc" binding:"required"`
		InFmt       string           `json:"inputFormat" binding:"required"`
		OutFmt      string           `json:"outputFormat" binding:"required"`
		Other       string           `json:"other" binding:"required"`
		Public      bool             `json:"pub"`
		MemoryLimit int              `json:"memoryLimit" binding:"required"`
		TimeLimit   int              `json:"timeLimit" binding:"required"`
		Diff        int              `json:"diff"`
		TestCases   []model.TestCase `json:"testCases"`
		SpecialCode string           `json:"specialCode"`
	}
	reqData := ReqData{}
	err2 := c.ShouldBindJSON(&reqData)
	if err1 != nil || err2 != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	// 兆转千
	reqData.MemoryLimit *= 1024
	testCasesBytes, _ := json.Marshal(reqData.TestCases)
	testCasesStr := ""
	testCasesStr = string(testCasesBytes)
	if reqData.ProblemID == 0 {
		userID, _ := c.Get("userID")
		util.GetDB().MustExec(
			`INSERT INTO problem(
						title,description,in_fmt,
						out_fmt,other,memory_limit,
						time_limit,diff,domain_id,test_cases,judge_type,special_code,
						public,creator_id,create_time,update_time,is_deleted,
						submit_num,ac_num
					) VALUES($1,$2,$3,
						$4,$5,$6,
						$7,$8,$9,$10,$11,$12,
						$13,$14,$15,$15,false,
						0,0
					)`,
			reqData.Title, reqData.Desc, reqData.InFmt,
			reqData.OutFmt, reqData.Other, reqData.MemoryLimit,
			reqData.TimeLimit, reqData.Diff, domainID, testCasesStr, reqData.JudgeType, reqData.SpecialCode,
			reqData.Public, userID, time.Now(),
		)
	} else {
		util.GetDB().MustExec(
			`UPDATE problem SET
					title=$1,description=$2,in_fmt=$3,
					out_fmt=$4,other=$5,memory_limit=$6,test_cases=$7,judge_type=$8,special_code=$9,
					time_limit=$10,diff=$11,public=$12,update_time=$13
					WHERE id=$14`,
			reqData.Title, reqData.Desc, reqData.InFmt,
			reqData.OutFmt, reqData.Other, reqData.MemoryLimit, testCasesStr, reqData.JudgeType, reqData.SpecialCode,
			reqData.TimeLimit, reqData.Diff, reqData.Public, time.Now(),
			reqData.ProblemID,
		)
	}
	NewResult(c).Success("", nil)
}

func removeProblem(c *gin.Context) {
	problemID := c.Param("id")
	util.GetDB().Exec("UPDATE problem SET is_deleted=true WHERE id=$1", problemID)
	NewResult(c).Success("", nil)
}
func addProblemRoute(r *gin.Engine) {
	api := r.Group("/problem")
	api.GET("/:id", getProblemByID)
	api.GET("/list", getProblems)
	api.POST("", upsertProblem)
	api.DELETE("/:id", removeProblem)
}
