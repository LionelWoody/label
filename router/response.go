package router

import (
	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, code string, msg string, data interface{}) {
	requestId := c.MustGet("requestId")
	c.JSON(200, map[string]interface{}{
		"data":       data,
		"error_no":   code,
		"error_msg":  msg,
		"request_id": requestId,
	})
}

func AbortResponse(c *gin.Context, code string, msg string, data interface{}) {
	requestId := c.MustGet("requestId")
	c.AbortWithStatusJSON(200, map[string]interface{}{
		"data":       data,
		"error_no":   code,
		"error_msg":  msg,
		"request_id": requestId,
	})
}

type TrackListResp struct {
	TrackId string `json:"trackId"`	
	StartTime int `json:"start_time"`
	State  int `json:"state"`
	RealedList []string `json:"relatedList"`
}
