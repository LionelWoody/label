package router

import (
	`testing`
	
	`label-backend/config`
	`label-backend/lib/db`
	`label-backend/model`
)

func TestGetTrackListByDataList(t *testing.T) {
	config.InitConf()
	type args struct {
		videoName string
	}
	tests := []struct {
		name    string
		args    args

	}{
		{name: "trackList", args:args{"nanling_guangzhou_bydz-ch01003-ch01003_20200525123500-fid_track_body-s252055-1590381300-1590381600"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := GetTrackListByDataList(tt.args.videoName)
			for _, item :=  range got{
				t.Logf("item : %v",item)
			}
			t.Logf("len : %v",len(got))
		})
	}
}

func TestGetTrackListByDb(t *testing.T) {
	config.InitConf()
	db.InitDB()
	type args struct {
		videoName string
	}
	tests := []struct {
		name    string
		args    args

	}{
		{name: "nanling_guangzhou_bydz-ch01003-ch01003_20200525123500-fid_track_body-s252055-1590381300-1590381600",
			args: args{ "nanling_guangzhou_bydz-ch01003-ch01003_20200525123500-fid_track_body-s252055-1590381300-1590381600"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := GetTrackListByDb(tt.args.videoName)
			for _ , item := range got{
				t.Logf("item : %v",item)
			}
		})
	}
}

func Test_chekoutMergeList(t *testing.T) {
	config.InitConf()
	
	config.Conf.DataBaseConfig.DB = "store_business_engine"
	config.Conf.DataBaseConfig.HostPort = "localhost:8084"
	config.Conf.DataBaseConfig.UserPassword = "root:password"
	db.InitDB()
	trackList :=[]*model.TrackInfo{
		&model.TrackInfo{
		TrackId: "1",
		LabelTrackId: "1",
		Videoname: "1",
	},
	&model.TrackInfo{
			TrackId: "2",
			LabelTrackId: "1",
			Videoname: "1",
		},
		&model.TrackInfo{
			TrackId: "3",
			LabelTrackId: "1",
			Videoname: "1",
		},
	}
	for _, track := range trackList {
		db.InsertTrackInfo(track)
	}
	type args struct {
		videoName string
		mergeList []string
	}
	tests := []struct {
		name        string
		args        args
	}{
		{name: "Test_chekoutMergeList", args: args{videoName: "1", mergeList: []string{"1","2","3"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mergeId , list, err := chekoutMergeList(tt.args.videoName, tt.args.mergeList)
			t.Logf("chekoutMergeList error err =%+v",err)
			t.Logf("mergeId : %v",mergeId)
			t.Logf("list : %v",list)
		})
	}
}