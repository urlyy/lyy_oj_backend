package route

import (
	"backend/model"
	"backend/util"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func addDiscussion(c *gin.Context) {
	userID, _ := c.Get("userID")
	domainID, err1 := getQueryDomainID(c)
	type ReqData struct {
		Title        string `json:"title"  binding:"required"`
		DiscussionID int    `json:"discussionID"`
		Content      string `json:"content"  binding:"required"`
	}
	reqData := ReqData{}
	err2 := c.ShouldBindJSON(&reqData)
	if err1 != nil || err2 != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	if reqData.DiscussionID == 0 {
		util.GetDB().MustExec(`
		INSERT INTO discussion(
			title,content,creator_id,domain_id,
			create_time,update_time,is_deleted
		) VALUES($1,$2,$3,$4,$5,$6,false)`,
			reqData.Title, reqData.Content, userID, domainID,
			time.Now(), time.Now(),
		)
	} else {
		util.GetDB().MustExec(`
		UPDATE discussion
		SET title=$1,content=$2,update_time=$3
		WHERE id=$4
		`, reqData.Title, reqData.Content, time.Now(), reqData.DiscussionID,
		)
	}
	NewResult(c).Success("", nil)
}

func getDiscussions(c *gin.Context) {
	domainID, err1 := getQueryDomainID(c)
	pageNumStr := c.DefaultQuery("page", "1")
	pageNum, err2 := strconv.Atoi(pageNumStr)
	if err1 != nil || err2 != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	if pageNum < 1 {
		pageNum = 1
	}
	var discussions []model.Discussion
	util.GetDB().Select(&discussions, `
		SELECT *
		FROM discussion 
		WHERE domain_id=$1 AND is_deleted = false 
		LIMIT $2 OFFSET $3`,
		domainID, PAGE_SIZE, (pageNum-1)*PAGE_SIZE)
	ret_discussions := make([]map[string]interface{}, len(discussions))
	for i, discussion := range discussions {
		ret_discussions[i] = map[string]interface{}{
			"id":         discussion.ID,
			"title":      discussion.Title,
			"createTime": discussion.CreateTime,
			"commentNum": discussion.CommentNum,
		}
		var user model.User
		err := util.GetDB().Get(&user, `SELECT * FROM "user" WHERE id=$1`, discussion.CreatorID)
		if err != nil {
			fmt.Println(err)
			NewResult(c).Fail("服务端异常")
			return
		}
		ret_discussions[i]["creatorUsername"] = user.Username
	}
	NewResult(c).Success("", map[string]interface{}{"discussions": ret_discussions})
}

func getDiscussionByID(c *gin.Context) {
	type Params struct {
		DiscussionID int `uri:"id" binding:"required"`
	}
	var params Params
	if err := c.ShouldBindUri(&params); err != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	var discussion model.Discussion
	err := util.GetDB().Get(&discussion, `SELECT * FROM discussion WHERE id = $1`, params.DiscussionID)
	if err != nil {
		NewResult(c).Fail("不存在该讨论")
		return
	}
	var creator model.User
	err = util.GetDB().Get(&creator, `SELECT * FROM "user" WHERE id = $1`, discussion.CreatorID)
	if err != nil {
		NewResult(c).Fail("服务端异常")
		return
	}
	NewResult(c).Success("", map[string]interface{}{
		"discussion": map[string]interface{}{
			"id":              discussion.ID,
			"title":           discussion.Title,
			"content":         discussion.Content,
			"createTime":      discussion.CreateTime,
			"creatorID":       discussion.CreatorID,
			"creatorUsername": creator.Username,
			"commentNum":      discussion.CommentNum,
		},
	})
}

func addDiscussionComment(c *gin.Context) {
	// domainID, err := getDomainID(c)
	// if err != nil {
	// 	NewResult(c).Fail("域参数错误")
	// 	return
	// }
	userID, _ := c.Get("userID")
	type Params struct {
		DiscussionID int `uri:"id" binding:"required"`
	}
	var params Params
	if err := c.ShouldBindUri(&params); err != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	type ReqData struct {
		Content string `json:"content"  binding:"required"`
		FloorID int    `json:"floorID"  `
		ReplyID int    `json:"replyID"  `
	}
	reqData := ReqData{}
	if err := c.ShouldBindJSON(&reqData); err != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	commentId := 0
	createTime := time.Now()
	err := util.GetDB().QueryRow(`
		INSERT INTO discussion_comment(
			discussion_id, content, creator_id, reply_id,floor_id, create_time, is_deleted
		) VALUES($1,$2,$3,$4,$5,$6,false)
		RETURNING id`,
		params.DiscussionID, reqData.Content, userID, reqData.ReplyID, reqData.FloorID, createTime,
	).Scan(&commentId)
	if err != nil {
		fmt.Println(err)
		NewResult(c).Fail("服务端异常")
		return
	}
	util.GetDB().MustExec("UPDATE discussion SET comment_num=comment_num+1 WHERE id=$1", params.DiscussionID)
	if reqData.FloorID == 0 {
		reqData.FloorID = commentId
		reqData.ReplyID = commentId
		util.GetDB().MustExec("UPDATE discussion SET reply_id=$1,floor_id=$1 WHERE id=$1", commentId)
	}
	NewResult(c).Success("", map[string]interface{}{
		"comment": map[string]interface{}{
			"id":         commentId,
			"content":    reqData.Content,
			"createTime": createTime,
			"floorID":    commentId,
			"replyID":    commentId,
		},
	})
}

func getDiscussionComments(c *gin.Context) {
	type Params struct {
		DiscussionID int `uri:"id" binding:"required"`
	}
	var params Params
	if err := c.ShouldBindUri(&params); err != nil {
		NewResult(c).Fail("参数错误")
		return
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
	var comments []model.DiscussionComment
	err = util.GetDB().Select(&comments, `SELECT * FROM discussion_comment WHERE discussion_id = $1 AND id=floor_id AND is_deleted=false OFFSET $2 LIMIT $3`, params.DiscussionID, (pageNum-1)*PAGE_SIZE, PAGE_SIZE)
	if err != nil {
		fmt.Println(err)
		NewResult(c).Fail("数据库错误")
	}
	fmt.Println(comments, params.DiscussionID, (pageNum-1)*PAGE_SIZE, PAGE_SIZE)
	ret_comments := make([]map[string]interface{}, len(comments))
	for idx, comment := range comments {
		var creator model.User
		err := util.GetDB().Get(&creator, `SELECT * FROM "user" WHERE id = $1`, comment.CreatorID)
		if err != nil {
			NewResult(c).Fail("服务端异常")
			return
		}
		ret_comments[idx] = map[string]interface{}{
			"id":              comment.ID,
			"content":         comment.Content,
			"replyID":         comment.ReplyID,
			"floorID":         comment.FloorID,
			"creatorUsername": creator.Username,
			"creatorID":       comment.CreatorID,
			"createTime":      comment.CreateTime,
			"replyUsername":   "",
		}
	}
	NewResult(c).Success("", map[string]interface{}{
		"comments": ret_comments,
	})
}

func getDiscussionReplies(c *gin.Context) {
	type Params struct {
		DiscussionID int `uri:"id" binding:"required"`
	}
	var params Params
	if err := c.ShouldBindUri(&params); err != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	floorIDStr := c.Query("floorID")
	floorID, err := strconv.Atoi(floorIDStr)
	if err != nil {
		NewResult(c).Fail("参数错误")
		return
	}
	var replies []model.DiscussionComment
	err = util.GetDB().Select(&replies, `SELECT * FROM discussion_comment WHERE discussion_id = $1 AND floor_id=$2 AND floor_id!=id AND is_deleted=false ORDER BY create_time DESC`, params.DiscussionID, floorID)
	if err != nil {
		fmt.Println(err)
		NewResult(c).Fail("数据库错误")
	}
	ret_replies := make([]map[string]interface{}, len(replies))
	for idx, reply := range replies {
		var creator model.User
		err := util.GetDB().Get(&creator, `SELECT * FROM "user" WHERE id = $1`, reply.CreatorID)
		if err != nil {
			NewResult(c).Fail("服务端异常")
			return
		}
		ret_replies[idx] = map[string]interface{}{
			"id":              reply.ID,
			"content":         reply.Content,
			"replyID":         reply.ReplyID,
			"floorID":         reply.FloorID,
			"creatorUsername": creator.Username,
			"creatorID":       reply.CreatorID,
			"createTime":      reply.CreateTime,
		}
	}
	for idx, reply := range ret_replies {
		replyID, floorID := reply["replyID"], reply["floorID"]
		if replyID == floorID {
			ret_replies[idx]["replyUsername"] = ""
		} else {
			for _, tmp := range ret_replies {
				if tmp["id"] == replyID {
					ret_replies[idx]["replyUsername"] = tmp["creatorUsername"]
					break
				}
			}
		}
	}
	NewResult(c).Success("", map[string]interface{}{
		"replies": ret_replies,
	})
}

func removeDiscussion(c *gin.Context) {
	contestID := c.Param("id")
	util.GetDB().Exec("UPDATE discussion SET is_deleted=true WHERE id=$1", contestID)
	NewResult(c).Success("", nil)
}

func removeDiscussionComment(c *gin.Context) {
	commentID := c.Param("id")
	util.GetDB().Exec("UPDATE discussion_comment SET is_deleted=true WHERE id=$1", commentID)
	NewResult(c).Success("", nil)
}

func addDiscussionRoute(r *gin.Engine) {
	api := r.Group("/discussion")
	api.GET("/:id", getDiscussionByID)
	api.GET("/list", getDiscussions)
	api.POST("", addDiscussion)
	api.GET("/:id/comment", getDiscussionComments)
	api.GET("/:id/comment/:floorID/reply", getDiscussionReplies)
	api.POST("/:id/comment", addDiscussionComment)
	api.DELETE("/:id", removeDiscussion)
	api.DELETE("/comment/:id", removeDiscussionComment)
}
