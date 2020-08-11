package router

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"label-backend/lib/db"
	"label-backend/model"
)

func RecordTime(c *gin.Context) {
	var req RecordTimeReq

	err := c.BindJSON(&req)
	if err != nil {
		c.JSON(400, err.Error())
		return
	}

	err = db.UpdateRecordInfoEndTimeByTrackId(int(time.Now().UnixNano()/1e6), req.MasterTrackId, req.VideoName)
	if err != nil && err != gorm.ErrRecordNotFound {
		Response(c, "4000", err.Error(), nil)
		return
	}

	err = db.InsertRecordInfo(&model.RecordInfo{
		Videoname: req.VideoName,
		TrackId:   req.SubTrackId,
		StartTime: uint64(time.Now().UnixNano()/1e6),
		EndTime:  uint64(time.Now().UnixNano()/1e6),
	})
	if err != nil {
		Response(c, "4000", err.Error(), nil)
		return
	}
	
	Response(c, "0", "successful!", nil)
}
