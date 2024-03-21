package route

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Ctx *gin.Context
}

type ResultCont struct {
	Success bool                   `json:"success"` // 自增
	Msg     string                 `json:"msg"`     //
	Data    map[string]interface{} `json:"data"`
}

func NewResult(ctx *gin.Context) *Result {
	return &Result{Ctx: ctx}
}

// 返回成功
func (r *Result) Success(msg string, data map[string]interface{}) {
	if data == nil {
		data = gin.H{}
	}
	res := ResultCont{}
	res.Success = true
	res.Msg = ""
	res.Data = data
	r.Ctx.JSON(http.StatusOK, res)
}

// 返回失败
func (r *Result) Fail(msg string) {
	res := ResultCont{}
	res.Success = false
	res.Msg = msg
	res.Data = gin.H{}
	// r.Ctx.JSON(http.StatusBadRequest, res)
	r.Ctx.JSON(http.StatusOK, res)
	r.Ctx.Abort()
}
