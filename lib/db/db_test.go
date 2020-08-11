package db

import (
	`testing`
	
	`label-backend/config`
	`label-backend/model`
)

func init(){
	config.InitConf()
	config.Conf.DataBaseConfig.DB = "store_business_engine"
	config.Conf.DataBaseConfig.HostPort = "localhost:8084"
	config.Conf.DataBaseConfig.UserPassword = "root:password"
	InitDB()
	InitDBTable()
}

///Users/guoxiaoshuang/project/label-backend/config.yaml
func TestInitDB(t *testing.T){
	InitDB()
	InitDBTable()
}

func TestDb(t *testing.T) {
	InitDB()
	track := &model.TrackInfo{TrackId: "xxx"}
	InsertTrackInfo(track)
}


func TestGetTrackInfoByVideoName(t *testing.T){
	config.InitConf()
	config.Conf.DataBaseConfig.DB = "store_business_engine"
	config.Conf.DataBaseConfig.HostPort = "localhost:8084"
	config.Conf.DataBaseConfig.UserPassword = "root:password"
	InitDB()
	InitDBTable()
	track , err  := GetTrackInfoByVideoName("xxxx")
	if err != nil{
		t.Errorf("GetTrackInfoByVideoName error err =%+v",err)
	}
	t.Logf("track : %v",track)
}
