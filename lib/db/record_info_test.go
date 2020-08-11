package db

import (
	`testing`
	
	`label-backend/config`
)

func TestUpdateRecordInfoEndTimeByTrackId(t *testing.T) {
	config.InitConf()
	config.Conf.DataBaseConfig.DB = "store_business_engine"
	config.Conf.DataBaseConfig.HostPort = "localhost:8084"
	config.Conf.DataBaseConfig.UserPassword = "root:password"
	InitDB()
	InitDBTable()
	
	// InsertRecordInfo(&model.RecordInfo{
	// 	TrackId: "001",
	// 	StartTime: 0,
	// 	EndTime: 1,
	// 	Videoname: "xxx",
	// })
	type args struct {
		endTime int
		trackId string
		videoName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{name: "updateRecordInfoEndTimeByTrackId", args: args{endTime: 1000,trackId: "001",videoName: "xxx"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := UpdateRecordInfoEndTimeByTrackId(tt.args.endTime, tt.args.trackId, tt.args.videoName); (err != nil) != tt.wantErr {
				t.Errorf("UpdateRecordInfoEndTimeByTrackId() error = %v, wantErr %v", err, tt.wantErr)
			}else{
				t.Logf("error err =%+v",err)
			}
		})
	}
}