package route

import (
	"backend/model"
	"backend/util"
	"fmt"
	"time"

	"strconv"

	"github.com/gin-gonic/gin"
)

func getProblems(c *gin.Context) {
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
		var problems []model.Problem
		util.GetDB().Select(&problems,
			`SELECT id,title,diff
			FROM problem 
			WHERE domain_id=$1 AND is_deleted = false 
			LIMIT $2 OFFSET $3`,
			domainID, PAGE_SIZE, (pageNum-1)*PAGE_SIZE)
		ret_problems := make([]map[string]interface{}, len(problems))
		for i, problem := range problems {
			ret_problems[i] = map[string]interface{}{
				"id":    problem.ID,
				"title": problem.Title,
				"diff":  problem.Diff,
			}
		}
		NewResult(c).Success("", map[string]interface{}{"problems": ret_problems})
	}
}

func getProblemByID(c *gin.Context) {
	type Params struct {
		ProblemID string `uri:"id" binding:"required"`
	}
	var params Params
	if err := c.ShouldBindUri(&params); err != nil {
		NewResult(c).Fail("参数错误")
	} else {
		var problem model.Problem
		err := util.GetDB().Get(&problem, "SELECT * FROM problem WHERE id = $1", params.ProblemID)
		if err != nil {
			NewResult(c).Fail("不存在该题目")
		} else {
			NewResult(c).Success("", map[string]interface{}{
				"problem": map[string]interface{}{
					"id":           problem.ID,
					"title":        problem.Title,
					"desc":         problem.Desc,
					"outputFormat": problem.OutFmt,
					"inputFormat":  problem.InFmt,
					"other":        problem.Other,
					"memoryLimit":  problem.MemoryLimit,
					"timeLimit":    problem.TimeLimit,
					"diff":         problem.Diff,
					"createTime":   problem.CreateTime,
					"pub":          problem.Public,
				},
			})
		}
	}
}

func addProblem(c *gin.Context) {
	domainID, err := getDomainID(c)
	if err != nil {
		NewResult(c).Fail("域参数错误")
	} else {
		type ReqData struct {
			ProblemID   int    `json:"problemID"`
			Title       string `json:"title" binding:"required"`
			Desc        string `json:"desc" binding:"required"`
			InFmt       string `json:"inputFormat" binding:"required"`
			OutFmt      string `json:"outputFormat" binding:"required"`
			Other       string `json:"other" binding:"required"`
			Public      *bool  `json:"pub" binding:"required"`
			MemoryLimit int    `json:"memoryLimit" binding:"required"`
			TimeLimit   int    `json:"timeLimit" binding:"required"`
			Diff        *int   `json:"diff" binding:"required"`
		}
		reqData := ReqData{}
		if err := c.ShouldBindJSON(&reqData); err != nil {
			fmt.Println(err)
			NewResult(c).Fail("参数错误")
		} else {
			if reqData.ProblemID == 0 {
				userID, _ := c.Get("userID")
				util.GetDB().MustExec(
					`INSERT INTO problem(
						title,description,in_fmt,
						out_fmt,other,memory_limit,
						time_limit,diff,domain_id,
						public,create_time,is_deleted,creator_id,update_time
					) VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,false,$12,$13)`,
					reqData.Title, reqData.Desc, reqData.InFmt,
					reqData.OutFmt, reqData.Other, reqData.MemoryLimit,
					reqData.TimeLimit, *reqData.Diff, domainID,
					*reqData.Public, time.Now(), userID, time.Now(),
				)
			} else {
				util.GetDB().MustExec(
					`UPDATE problem SET
					title=$1,description=$2,in_fmt=$3,
					out_fmt=$4,other=$5,memory_limit=$6,
					time_limit=$7,diff=$8,public=$9,update_time=$10
					WHERE id=$11`,
					reqData.Title, reqData.Desc, reqData.InFmt,
					reqData.OutFmt, reqData.Other, reqData.MemoryLimit,
					reqData.TimeLimit, *reqData.Diff, *reqData.Public, time.Now(),
					reqData.ProblemID,
				)
			}
			NewResult(c).Success("", nil)
		}
	}

}

func addProblemRoute(r *gin.Engine) {
	api := r.Group("/problem")
	api.GET("/:id", getProblemByID)
	api.GET("/list", getProblems)
	api.POST("", addProblem)
}
