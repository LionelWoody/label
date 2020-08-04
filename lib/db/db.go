package db

import (
	"fmt"
	"time"
	
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	
	`label-backend/config`
	`label-backend/model`
)

var (
	DB  *gorm.DB
)

func InitDB() {
	var err error
	dsn := fmt.Sprintf("%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",config.Conf.DataBaseConfig.UserPassword, config.Conf.DataBaseConfig.HostPort,config.Conf.DataBaseConfig.DB)
	logrus.Infof("dsn : %s", dsn)
	DB , err = gorm.Open("mysql", dsn)
	if err != nil {
		panic("Init Databae error err : " + err.Error() +"dsn : "+dsn)
	}
	
	if err := DB.DB().Ping(); err != nil {
		panic("Connect to databae error err : " + err.Error())
	}
	DB.LogMode(true)
	DB.DB().SetConnMaxLifetime(60 * time.Second)
	DB.DB().SetMaxIdleConns(40)
	DB.DB().SetMaxOpenConns(300)
}

func InitDBTable(){
	err := DB.AutoMigrate(&model.TrackInfo{}).Error
	if err != nil {
		panic("table init error : " + err.Error())
	}
	
	logrus.Infof("start create userInfo")
	err = DB.AutoMigrate(&model.TrackInfo{}).Error
	if err != nil {
		panic("table userinfo init error : " + err.Error())
	}
}

func InsertTrackInfo(track *model.TrackInfo) error{
	return DB.Model(&model.TrackInfo{}).Save(track).Error
}

func GetTrackInfoByVideoName(videoName string) (*model.TrackInfo,error){
	var trackInfo model.TrackInfo
	err := DB.Model(&model.TrackInfo{}).Where("videoname = ?",videoName).Find(&trackInfo).Error
	if err != nil{
		return nil, err
	}
	return &trackInfo, nil
}

func UpdateTrackInfo(videoName , trackId,labelTrackId string) (*model.TrackInfo,error){
	var trackInfo model.TrackInfo
	err := DB.Model(&model.TrackInfo{}).Where("videoname = ? ",videoName).Where("track_id =?",trackId).Update("label_track_id",labelTrackId).Error
	if err != nil{
		return nil, err
	}
	return &trackInfo, nil
}

func IsTrackInfoVideoNameHave(videoName string)bool{
	_, err := GetTrackInfoByVideoName(videoName)
	if err != nil{
		return false
	}
	return true
}

func GetTrackInfoListoByVideoName(videoName string) ([]*model.TrackInfo,error){
	var trackInfo []*model.TrackInfo
	err := DB.Model(&model.TrackInfo{}).Where("videoname = ?",videoName).Find(&trackInfo).Error
	if err != nil{
		return nil, err
	}
	return trackInfo, nil
}
