package model

import (
	`github.com/jinzhu/gorm`
)

type RecordInfo struct {
	gorm.Model          //column:beast_id
	TrackId      string `gorm:"column:track_id; index:idx_track; type:varchar(128)  not null  default ' ' comment '轨迹id'"`
	StartTime    uint64 `gorm:"column:start_time; comment '标注开始时间'"`
	EndTime      uint64 `gorm:"column:end_time;  comment '标注结束时间'"`
	Videoname    string `gorm:"column:videoname;  comment '视频名称'"`
}

func (RecordInfo) TableName() string {
	return "record_info"
}


