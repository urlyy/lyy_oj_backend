package route

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func getDomainID(c *gin.Context) (int, error) {
	domainIDStr := c.Query("d")
	domainID, err := strconv.Atoi(domainIDStr)
	if err != nil {
		return 0, err
	} else {
		return domainID, nil
	}
}

func parseTime(timeStr string) (time.Time, error) {
	const layout = "2006-01-02 15:04"
	// 使用 time.Parse 解析时间字符串
	t, err := time.Parse(layout, timeStr)
	if err != nil {
		return time.Time{}, err
	} else {
		return t, nil
	}
}
