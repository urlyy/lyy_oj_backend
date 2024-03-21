package route

import (
	"backend/model"
	"backend/util"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func login(c *gin.Context) {
	type Input struct {
		Password string `json:"password" binding:"required"`
		Email    string `json:"email" `
		TrueID   string `json:"trueID" `
	}
	var input Input
	c.ShouldBindJSON(&input)
	users := []model.User{}
	if input.Email != "" {
		util.GetDB().Select(&users, `SELECT * FROM "user" WHERE email=$1`, input.Email)
	} else if input.TrueID != "" {
		util.GetDB().Select(&users, `SELECT * FROM "user" WHERE true_id=$1`, input.TrueID)
	}
	if len(users) == 0 {
		NewResult(c).Fail("登录失败，请重新输入登录信息")
	} else if len(users) > 1 {
		NewResult(c).Fail("存在相同学号/工号用户,请使用邮箱登录")
	} else {
		user := users[0]
		// TODO 校验密码
		tokenn, _ := util.GenToken(int64(user.ID))
		util.GetDB().Exec(`UPDATE "user" SET session_token=$1,last_login=$2 WHERE id=$2`, tokenn, time.Now(), user.ID)
		domain_ids := []int{}
		util.GetDB().Select(&domain_ids, `SELECT domain_id FROM domain_user WHERE user_id=$1`, user.ID)
		data := map[string]interface{}{
			"token": tokenn,
			"user": map[string]interface{}{
				"id":       user.ID,
				"trueID":   user.TrueID,
				"username": user.Username,
				"email":    user.Email,
				"school":   user.School,
				"gender":   user.Gender,
			},
		}
		//看看是否直接进入域
		if len(domain_ids) == 1 {
			data["domainID"] = domain_ids[0]
		}
		NewResult(c).Success("", data)
	}
}

func sendForgetPasswordCaptcha(c *gin.Context) {
	email := c.Query("email")
	var user model.User
	err := util.GetDB().Get(&user, `SELECT * FROM "user" WHERE email=$1`, email)
	if err != nil {
		if err == sql.ErrNoRows {
			NewResult(c).Fail("邮箱未注册")
		} else {
			NewResult(c).Fail("数据库错误")
		}
	} else {
		captchaLength := 6
		content := fmt.Sprintf("验证码是:%s", util.GenCaptcha(captchaLength))
		util.SendEmail(email, "忘记密码服务-验证码", content)
		NewResult(c).Success("", nil)
	}
}

func forgetPassword(c *gin.Context) {

	type Input struct {
		Email       string `json:"email"  binding:"required"`
		NewPassword string `json:"newPassword"  binding:"required"`
		Captcha     string `json:"captcha"  binding:"required"`
	}
	var input Input
	c.ShouldBindJSON(&input)
	if strings.EqualFold(input.Captcha, "1234") {
		NewResult(c).Fail("验证码错误")
	} else {
		util.GetDB().MustExec(`UPDATE "use" SET password=$1 WHERE email=$2`, input.NewPassword, input.Email)
		NewResult(c).Success("", nil)
	}
}

func changePassword(c *gin.Context) {
	userID, _ := c.Get("userID")
	type ReqData struct {
		OldPassword string `json:"oldPassword"  binding:"required"`
		NewPassword string `json:"newPassword"  binding:"required"`
	}
	var reqData ReqData
	if err := c.ShouldBindJSON(&reqData); err != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	var user model.User
	err := util.GetDB().Get(&user, `SELECT * FROM "user" WHERE id = $1 AND password=$2`, userID, reqData.OldPassword)
	if err != nil {
		NewResult(c).Fail("密码错误")
		return
	}
	util.GetDB().MustExec(`UPDATE "user" SET password = $1 WHERE id = $2`, reqData.NewPassword, userID)
	NewResult(c).Success("", nil)
}

func changeEmail(c *gin.Context) {
	userID, _ := c.Get("userID")
	type ReqData struct {
		Password string `json:"password"  binding:"required"`
		NewEmail string `json:"newEmail"  binding:"required"`
	}
	var reqData ReqData
	if err := c.ShouldBindJSON(&reqData); err != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	var user model.User
	err := util.GetDB().Get(&user, `SELECT * FROM "user" WHERE id = $1 AND password=$2`, userID, reqData.Password)
	if err != nil {
		NewResult(c).Fail("密码错误")
		return
	}
	util.GetDB().MustExec(`UPDATE "user" SET email = $1 WHERE id = $2`, reqData.NewEmail, userID)
	NewResult(c).Success("", nil)
}

func getUserProfile(c *gin.Context) {
	type Params struct {
		UserID int `uri:"id" binding:"required"`
	}
	var params Params
	if err := c.ShouldBindUri(&params); err != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	var user model.User
	err := util.GetDB().Get(&user, `SELECT * FROM "user" WHERE id=$1 AND is_deleted=false`, params.UserID)
	if err != nil {
		NewResult(c).Fail("用户不存在")
		return
	}
	NewResult(c).Success("", map[string]interface{}{
		"user": map[string]interface{}{
			"username": user.Username,
			// "email":     user.Email,
			"school":    user.School,
			"gender":    user.Gender,
			"lastLogin": user.LastLogin,
		},
		"submitRecords": map[string]interface{}{},
	})
}

func addUserRoute(r *gin.Engine) {
	api := r.Group("/user")
	api.POST("/login", login)
	api.POST("/forget-pass/captcha", sendForgetPasswordCaptcha)
	api.POST("/forget-pass", forgetPassword)
	api.POST("/pass", changePassword)
	api.POST("/email", changeEmail)
	api.GET("/:id/profile", getUserProfile)
}
