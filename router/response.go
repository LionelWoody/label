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
	TrackId        string   `json:"trackId"`
	StartTime      string      `json:"start_time"`
	EndTime        string   `json:"end_time"`
	State          int      `json:"state"`
	MarkedByOrigin string   `json:"markedByOrigin,omitempty"`
	RealedList     []string `json:"relatedList"`
	Dest           int      `json:"dets"`
	TrackDuration  int `json:"trackDuration"`
}

type TrackList struct {
	List []*TrackListResp `json:"trackList"`
}

type TrackQA struct {
	TrackId string `json:"trackId"`
	LabelList []string `json:"LabelList"`
}

type TrackQAList struct {
	List []*TrackQA `json:"trackList"`
}
