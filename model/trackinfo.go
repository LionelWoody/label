package model

import (
	"github.com/jinzhu/gorm"
)

type TrackInfo struct {
	gorm.Model  //column:beast_id
	TrackId string `gorm:"column:track_id; index:idx_track; type:varchar(128)  not null  default ' ' comment '轨迹id'"`
	StartTime  uint64  `gorm:"column:start_time; comment '轨迹开始时间'"`
	EndTime     uint64 `gorm:"column:end_time;  comment '轨迹结束时间'"`
	Videoname   string `gorm:"column:videoname;  comment '视频名称'"`
	Candidates  string `gorm:"column:candidates; type:text;"`
	LabelTrackId string `gorm:"column:label_track_id; type:varchar(128)  not null  default ' ' comment '标注的轨迹id'"`
}

func (TrackInfo) TableName() string {
	return "annotation_info"
}
