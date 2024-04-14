package route

import (
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

const (
	PAGE_SIZE = 25
)

func init() {
	r = gin.Default()
	addMiddleware(r)
	addUserRoute(r)
	addDomainRoute(r)
	addProblemRoute(r)
	addHomeworkRoute(r)
	addContestRoute(r)
	addDiscussionRoute(r)
	addJudgeRoute(r)
	addSubmissionRoute(r)
	addConfigRoute(r)
	addSSERoute(r)
}

func GetRouter() *gin.Engine {
	return r
}
