package router

import (
	`testing`
	
	`label-backend/config`
	`label-backend/lib/db`
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