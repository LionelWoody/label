package db

import (
	`label-backend/model`
)

func UpdateRecordInfoEndTimeByTrackId(endTime   int  ,trackId  , videoName string)error{
	return  DB.Model(&model.RecordInfo{}).Where("track_id =?",trackId).Updates(map[string]interface{}{"end_time" :endTime}).Error
}

func InsertRecordInfo(item *model.RecordInfo) error{
	return DB.Model(&model.RecordInfo{}).Save(item).Error
}