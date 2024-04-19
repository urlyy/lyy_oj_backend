package route

import (
	"backend/model"
	"backend/util"
	"encoding/json"
	"math"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func sendNotification(c *gin.Context) {
	domainID, err := getPathInt(c, "id")
	if err != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	title := c.PostForm("title")
	content := c.PostForm("content")
	sendTime := time.Now()
	var id int
	err = util.GetDB().QueryRow(`
		INSERT INTO notification(
			domain_id, title,content, create_time, is_deleted
		) VALUES($1,$2,$3,$4,$5)
		RETURNING id`,
		domainID, title, content, sendTime, false,
	).Scan(&id)
	if err != nil {
		NewResult(c).Fail("服务端发送通知异常")
		return
	}
	msg := map[string]interface{}{
		"title":    title,
		"content":  content,
		"sendTime": sendTime,
	}
	str, err := json.Marshal(msg)
	if err != nil {
		NewResult(c).Fail("服务端发送通知异常")
		return
	}
	util.SSEBroadcast(strconv.Itoa(domainID), string(str))
	NewResult(c).Success("发送成功", map[string]interface{}{
		"id":         id,
		"title":      title,
		"content":    content,
		"createTime": sendTime,
	})
}

func removeNotification(c *gin.Context) {
	domainID, err1 := getPathInt(c, "id")
	notificationID, err2 := getPathInt(c, "nid")
	if err1 != nil || err2 != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	util.GetDB().MustExec(`
		UPDATE notification 
		SET is_deleted = true WHERE id = $1 AND domain_id = $2
	`, notificationID, domainID)
	NewResult(c).Success("", nil)
}

const (
	NOTIFICATION_PAGE = 10
)

func getNotificationList(c *gin.Context) {
	curPage, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	domainID, err1 := getPathInt(c, "id")
	if err != nil || err1 != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	if curPage < 1 {
		curPage = 1
	}
	notifications := make([]model.Notification, 0)
	util.GetDB().Select(&notifications, `
		SELECT id, title, content, create_time
		FROM notification
		WHERE domain_id =$1 AND is_deleted = false
		ORDER BY create_time DESC
		LIMIT $2 OFFSET $3
	`, domainID, NOTIFICATION_PAGE, (curPage-1)*NOTIFICATION_PAGE)
	var count int
	util.GetDB().Get(&count, `
		SELECT COUNT(*)
		FROM notification
		WHERE domain_id =$1 AND is_deleted = false
	`, domainID, NOTIFICATION_PAGE, (curPage-1)*NOTIFICATION_PAGE)
	pageNum := math.Ceil(float64(count) / float64(NOTIFICATION_PAGE))
	resList := make([]map[string]interface{}, len(notifications))
	for i, tmp := range notifications {
		resList[i] = map[string]interface{}{
			"id":         tmp.ID,
			"title":      tmp.Title,
			"content":    tmp.Content,
			"createTime": tmp.CreateTime,
		}
	}
	NewResult(c).Success("", map[string]interface{}{
		"list":    resList,
		"pageNum": pageNum,
	})
}

func addNotificationRoute(r *gin.Engine) {
	api := r.Group("/notify/:id")
	api.POST("", sendNotification)
	api.GET("/list", getNotificationList)
	api.DELETE("/:nid", removeNotification)
}
