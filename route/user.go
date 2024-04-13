package route

import (
	"backend/model"
	"backend/util"
	"database/sql"
	"fmt"
	"strconv"
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
	users := make([]model.User, 0)
	if input.Email != "" {
		util.GetDB().Select(&users, `SELECT * FROM "user" WHERE email=$1 AND is_deleted=false`, input.Email)
	} else if input.TrueID != "" {
		util.GetDB().Select(&users, `SELECT * FROM "user" WHERE true_id=$1 AND is_deleted=false`, input.TrueID)
	}
	if len(users) == 0 {
		NewResult(c).Fail("登录失败，请重新输入登录信息")
		return
	}
	if len(users) > 1 {
		NewResult(c).Fail("存在相同学号/工号用户,请使用邮箱登录")
		return
	}
	user := users[0]
	// TODO 校验密码
	// if user.Password != util.SM3(input.Password, user.Salt) {
	// 	NewResult(c).Fail("登录失败，请重新输入登录信息")
	// 	return
	// }
	tokenn, _ := util.GenToken(int64(user.ID))
	util.GetDB().MustExec(`UPDATE "user" SET session_token=$1,last_login=$2 WHERE id=$3`, tokenn, time.Now(), user.ID)
	domain_ids := make([]int, 0)
	util.GetDB().Select(&domain_ids, `SELECT domain_id FROM domain_user WHERE user_id=$1 AND is_deleted=false`, user.ID)
	data := map[string]interface{}{
		"token": tokenn,
		"user": map[string]interface{}{
			"id":       user.ID,
			"trueID":   user.TrueID,
			"username": user.Username,
			"email":    user.Email,
			"school":   user.School,
			"gender":   user.Gender,
			"website":  user.Website,
		},
	}
	//看看是否直接进入域
	if len(domain_ids) == 1 {
		data["domainID"] = domain_ids[0]
	}
	NewResult(c).Success("", data)
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
			"website":   user.Website,
		},
	})
}

func changeUserProfile(c *gin.Context) {
	userID, _ := c.Get("userID")
	type ReqData struct {
		Username string `json:"username"  binding:"required"`
		School   string `json:"school"`
		Gender   int    `json:"gender"`
		Website  string `json:"website"`
	}
	var reqData ReqData
	if err := c.ShouldBindJSON(&reqData); err != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	util.GetDB().MustExec(`
		UPDATE "user" SET username=$1,gender=$2,school=$3,website=$4 
		WHERE id=$5`,
		reqData.Username, reqData.Gender, reqData.School, reqData.Website, userID,
	)
	NewResult(c).Success("", nil)
}

func createUsers(c *gin.Context) {
	type User struct {
		TrueID   int    `json:"trueId"  binding:"required"`
		Gender   int    `json:"gender"  binding:"required"`
		Username string `json:"username"  binding:"required"`
		School   string `json:"school"  binding:"required"`
	}
	type ReqData struct {
		Users []User `json:"users"  binding:"required"`
	}
	var reqData ReqData
	err := c.ShouldBindJSON(&reqData)
	if err != nil {
		fmt.Println(err)
		NewResult(c).Fail("参数错误")
		return
	}
	retData := make([][]interface{}, len(reqData.Users))
	for idx, user := range reqData.Users {
		var userID int
		password := "1234"
		params := []interface{}{user.TrueID, user.Username, password, user.School, "", "salt", "token", user.Gender, false, time.Now(), ""}
		err := util.GetDB().QueryRow(`
		INSERT INTO "user"(true_id,username,password,school,
			email,salt,session_token,gender,is_deleted,last_login,website
		)VALUES($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
		RETURNING id
		`, params...).Scan(&userID)
		if err != nil {
			fmt.Println("服务端异常")
			NewResult(c).Fail("新建用户失败")
			return
		}
		email := fmt.Sprintf("%s%d@%d", user.Username, userID, userID)
		util.GetDB().Exec(`
			UPDATE "user" SET email=$1
			WHERE id=$2`, email, userID,
		)
		retData[idx] = []interface{}{user.TrueID, user.Username, email, password}
	}
	NewResult(c).Success("创建成功", map[string]interface{}{
		"users": retData,
	})
}

func searchUser(c *gin.Context) {
	pageNumStr := c.DefaultQuery("page", "1")
	keyword := c.DefaultQuery("keyword", "")
	trueID := c.DefaultQuery("trueID", "")
	school := c.DefaultQuery("school", "")
	pageNum, err := strconv.Atoi(pageNumStr)
	if err != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	users := make([]model.User, 0)
	params := []interface{}{PAGE_SIZE, (pageNum - 1) * PAGE_SIZE}
	extra_where := ""
	if trueID != "" {
		params = append(params, trueID)
		extra_where += fmt.Sprintf(" AND true_id=$%d", len(params))
	}
	if keyword != "" {
		params = append(params, "%"+keyword+"%")
		extra_where += fmt.Sprintf(" AND username LIKE $%d", len(params))
	}
	if school != "" {
		params = append(params, "%"+school+"%")
		extra_where += fmt.Sprintf(" AND school LIKE $%d", len(params))
	}
	sql := fmt.Sprintf(`
		SELECT *
		FROM "user"
		WHERE is_deleted = false %s
		LIMIT $1 OFFSET $2`, extra_where,
	)
	util.GetDB().Select(&users, sql, params...)
	retUsers := make([]map[string]interface{}, len(users))
	for i, user := range users {
		retUsers[i] = map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
			"trueID":   user.TrueID,
			"school":   user.School,
			"email":    user.Email,
			"gender":   user.Gender,
		}
	}
	NewResult(c).Success("", map[string]interface{}{
		"users": retUsers,
	})
}

func addUser2Domain(c *gin.Context) {
	domainID, err1 := getPathInt(c, "id")
	type ReqData struct {
		UserIDs []int `json:"userIDs"  binding:"required"`
	}
	var reqData ReqData
	err2 := c.ShouldBindJSON(&reqData)
	if err1 != nil || err2 != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	for _, userID := range reqData.UserIDs {
		util.GetDB().Exec(`INSERT INTO domain_user(user_id,domain_id) VALUES($1,$2)`, userID, domainID)
	}
	NewResult(c).Success("", nil)
}

func addUserRoute(r *gin.Engine) {
	api := r.Group("/user")
	api.POST("/login", login)
	api.POST("/forget-pass/captcha", sendForgetPasswordCaptcha)
	api.POST("/forget-pass", forgetPassword)
	api.POST("/pass", changePassword)
	api.POST("/email", changeEmail)
	api.GET("/:id/profile", getUserProfile)
	api.POST("/profile", changeUserProfile)
	api.POST("", createUsers)
	api.POST("/domain/:id", addUser2Domain)
	api.GET("/list", searchUser)
}
