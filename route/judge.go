package route

import (
	"backend/model"
	pb "backend/proto/judge"
	"backend/util"
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const (
	ACCEPT              = 0
	COMPILE_ERROR       = 1
	WRONG_ANSWER        = 2
	RUNTIME_ERROR       = 3
	TIME_LIMIT_EXCEED   = 4
	MEMORY_LIMIT_EXCEED = 5
	JUDGING             = 6
)

type JudgeResult struct {
	Status       int     `json:"status"`
	MaxTime      uint64  `json:"maxTime"`
	MaxMemory    uint64  `json:"maxMemory"`
	SubmissionID int     `json:"submissionID"`
	Output       string  `json:"output"`
	PassPercent  float64 `json:"passPercent"`
	Log          string  `json:"log"`
}

func transReply(reply *pb.JudgeReply, submissionID int) *JudgeResult {
	res := JudgeResult{SubmissionID: submissionID, Output: ""}
	var passNum int = 0
	if reply.Compile.Status == ACCEPT {
		ac := true
		for _, run := range reply.Run {
			if run.Status != ACCEPT {
				ac = false
				res.Status = int(run.Status)
				break
			} else {
				passNum += 1
				if res.MaxTime < run.Time {
					res.MaxTime = run.Time
				}
				if res.MaxMemory < run.Memory {
					res.MaxMemory = run.Memory
				}
			}
		}
		if ac {
			res.Output = reply.LastOutput
		}
	} else {
		res.Status = COMPILE_ERROR
		res.Log = reply.Compile.Log
	}
	passPercent := float64(passNum) / float64(len(reply.Run))
	res.PassPercent = math.Round(passPercent*100) / 100
	return &res
}

func getInOut(problemID int) ([]string, []string, int, int, int, string, error) {
	var problem model.Problem
	util.GetDB().Get(&problem, "SELECT * FROM problem WHERE id=$1", problemID)
	var testCases []model.TestCase
	err := json.Unmarshal([]byte(problem.TestCases), &testCases)
	if err != nil {
		return nil, nil, 0, 0, 0, "", err
	}
	inputList := []string{}
	expectList := []string{}
	for _, c := range testCases {
		inputList = append(inputList, c.Input)
		expectList = append(expectList, c.Expect)
	}
	return inputList, expectList, problem.TimeLimit, problem.MemoryLimit, problem.JudgeType, problem.SpecialCode, nil
}

func submitTest(c *gin.Context) {
	problemID, err1 := getPathInt(c, "pid")
	// domainID, err2 := getQueryDomainID(c)
	type ReqData struct {
		Code      string `json:"code"`
		TestInput string `json:"testInput"`
		Compiler  string `json:"compiler"`
	}
	reqData := ReqData{}
	err3 := c.ShouldBindJSON(&reqData)
	if err1 != nil || err3 != nil {
		fmt.Println(err1, err3)
		NewResult(c).Fail("参数错误")
		return
	}
	dummySubmissionID, err := util.GenUUID()
	if err != nil {
		fmt.Println(err, 1)
		NewResult(c).Fail("服务端错误")
		return
	}
	var problem model.Problem
	util.GetDB().Get(&problem, "SELECT * FROM problem WHERE id=$1", problemID)
	var config model.Config
	util.GetDB().Get(&config, "SELECT * FROM config")
	inputList := []string{reqData.TestInput}
	reply, err := util.Judge(dummySubmissionID, config.AddressList, reqData.Code, inputList, inputList, reqData.Compiler, uint64(problem.TimeLimit), uint64(problem.MemoryLimit), problem.JudgeType == 1, problem.SpecialCode)
	fmt.Println("out ", time.Now())
	if err != nil {
		fmt.Println(err, "2")
		NewResult(c).Fail("服务端错误")
		return
	}
	// 因为肯定会答案错误，调整方便他
	if reply.Run[0].Status == WRONG_ANSWER {
		reply.Run[0].Status = ACCEPT
	}
	// fmt.Println(reply.Run)
	res := transReply(reply, 0)
	NewResult(c).Success("", map[string]interface{}{
		"result": res,
	})
}

func submitJudge(c *gin.Context) {
	problemID, err1 := getPathInt(c, "pid")
	domainID, err2 := getQueryDomainID(c)
	userID, _ := c.Get("userID")
	type ReqData struct {
		Code     string `json:"code"`
		Compiler string `json:"compiler"`
		Type     string `json:"type"`
		FromID   int    `json:"fromID"`
	}
	reqData := ReqData{}
	err3 := c.ShouldBindJSON(&reqData)
	if err1 != nil || err2 != nil || err3 != nil || (reqData.Type != "contest" && reqData.Type != "homework" && reqData.Type != "problem") {
		NewResult(c).Fail("参数错误")
		fmt.Println(err1, err2, err3)
		return
	}
	now := time.Now()
	if reqData.Type == "contest" {
		var contest model.Contest
		util.GetDB().Get(&contest, "select * from contest where id = $1", reqData.FromID)
		if contest.EndTime.Before(now) {
			NewResult(c).Fail("比赛已结束")
			return
		}
	} else if reqData.Type == "homework" {
		var homework model.Homework
		util.GetDB().Get(&homework, "select * from homework where id = $1", reqData.FromID)
		if homework.EndTime.Before(now) {
			NewResult(c).Fail("作业已结束")
			return
		}
	}
	var submissionID int
	util.GetDB().QueryRow(`
		INSERT INTO 
		submission(problem_id,domain_id,from_type,user_id,submit_time,last_judge_time,lang,code,from_id,status,max_time,max_memory,pass_percent,log)
		VALUES($1,$2,$3,$4,$5,$5,$6,$7,$8,$9,$10,$10,$11,$12)
		RETURNING id
		`, problemID, domainID, reqData.Type, userID, now, reqData.Compiler, reqData.Code, reqData.FromID, JUDGING, 0, 0.0, "",
	).Scan(&submissionID)
	util.GetDB().MustExec("UPDATE problem SET submit_num=submit_num+1 WHERE id=$1", problemID)
	var config model.Config
	util.GetDB().Get(&config, "SELECT * FROM config")
	inputList, expectList, timeLimit, memoryLimit, judgeType, specialCode, err := getInOut(problemID)
	if err != nil {
		NewResult(c).Fail("服务端错误")
		return
	}
	reply, err := util.Judge(strconv.Itoa(submissionID), config.AddressList, reqData.Code, inputList, expectList, reqData.Compiler, uint64(timeLimit), uint64(memoryLimit), judgeType == 1, specialCode)
	if err != nil {
		NewResult(c).Fail("服务端错误")
		return
	}
	res := transReply(reply, submissionID)
	util.GetDB().MustExec(`UPDATE submission SET status=$1,max_time=$2,max_memory=$3,pass_percent=$4,log=$5 WHERE id=$6`, res.Status, res.MaxTime, res.MaxMemory, res.PassPercent, res.Log, submissionID)
	if res.Status == ACCEPT {
		util.GetDB().MustExec("UPDATE problem SET ac_num=ac_num+1 WHERE id=$1", problemID)
	}
	NewResult(c).Success("", map[string]interface{}{
		"result": res,
	})
}

func getCompilers(c *gin.Context) {
	var config model.Config
	util.GetDB().Get(&config, "SELECT * FROM config")
	compilers := make([][]string, 0)
	err := json.Unmarshal([]byte(config.Compilers), &compilers)
	if err != nil {
		NewResult(c).Fail("服务端错误")
		return
	}
	NewResult(c).Success("", map[string]interface{}{
		"compilers": compilers,
	})
}

func rejudge(c *gin.Context) {
	submissionID, err1 := getPathInt(c, "sid")
	if err1 != nil {
		NewResult(c).Fail("参数错误")
		fmt.Println(err1)
		return
	}
	// domainID, err2 := getQueryDomainID(c)
	// userID, _ := c.Get("userID")
	var config model.Config
	util.GetDB().Get(&config, "SELECT * FROM config")
	var submission model.Submission
	util.GetDB().Get(&submission, "SELECT * FROM submission WHERE id=$1", submissionID)
	inputList, expectList, timeLimit, memoryLimit, judgeType, specialCode, err := getInOut(submission.ProblemID)
	if err != nil {
		NewResult(c).Fail("服务端错误")
		return
	}
	reply, err := util.Judge(strconv.Itoa(submissionID), config.AddressList, submission.Code, inputList, expectList, submission.Lang, uint64(timeLimit), uint64(memoryLimit), judgeType == 1, specialCode)
	if err != nil {
		NewResult(c).Fail("服务端错误")
		return
	}
	res := transReply(reply, submissionID)
	util.GetDB().MustExec(`UPDATE submission SET status=$1,max_time=$2,max_memory=$3,pass_percent=$4,log=$5,last_judge_time=$6 WHERE id=$7`, res.Status, res.MaxTime, res.MaxMemory, res.PassPercent, res.Log, time.Now(), submissionID)
	if res.Status == ACCEPT && submission.Status != ACCEPT {
		util.GetDB().MustExec("UPDATE problem SET ac_num=ac_num+1 WHERE id=$1", submission.ProblemID)
	}
	if res.Status != ACCEPT && submission.Status == ACCEPT {
		util.GetDB().MustExec("UPDATE problem SET ac_num=ac_num-1 WHERE id=$1", submission.ProblemID)
	}
	NewResult(c).Success("", map[string]interface{}{
		"result": res,
	})
}

func addJudgeRoute(r *gin.Engine) {
	api := r.Group("/judge")
	api.POST("/:pid/test", submitTest)
	api.POST("/:pid", submitJudge)
	api.POST("/re/:sid", rejudge)
	api.GET("/compiler", getCompilers)
}
