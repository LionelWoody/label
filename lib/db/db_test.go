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

func TestSetTrackInfoLabelNil(t *testing.T) {
	
	config.InitConf()
	config.Conf.DataBaseConfig.DB = "store_business_engine"
	config.Conf.DataBaseConfig.HostPort = "localhost:8084"
	config.Conf.DataBaseConfig.UserPassword = "root:password"
	InitDB()
	InitDBTable()
	type args struct {
		videoName string
		trackId   string
	}
	tests := []struct {
		name    string
		args    args

	}{
		{name: "videoName", args: args{videoName: "ch01004_20191025104000",trackId: "sephora_shanghai_qbwk-ch01004-fid-track_body-123-6d59cde3-b253-4c49-bac6-0dc8bb61dd47"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			 err := SetTrackInfoLabelNil(tt.args.videoName, tt.args.trackId)
			 t.Errorf("error err =%+v",err)
		})
	}
}