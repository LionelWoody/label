package router

import (
	`fmt`
	`io/ioutil`
	`path/filepath`
	`strings`
	
	`github.com/bitly/go-simplejson`
	`github.com/gin-gonic/gin`
	`github.com/sirupsen/logrus`
	
	`label-backend/config`
	`label-backend/lib/db`
	`label-backend/model`
)

func Annotationlist(c *gin.Context){
		videoName := c.Query("videoName")
		logrus.Infof("videoName : %s",videoName)
		if !db.IsTrackInfoVideoNameHave(videoName) {
			list  ,err := GetTrackListByDataList(videoName)
			if err != nil{
				Response(c, "1000",err.Error(),nil)
				return
			}
			for _, item := range list {
				db.InsertTrackInfo(item)
			}
		}
		ret , err := GetTrackListByDb(videoName)
		if err != nil{
			Response(c, "1000",err.Error(),nil)
			return
		}
		Response(c, "0","",ret)
}

func GetTrackListByDb(videoName string)([]*TrackListResp, error){
	ret := make([]*TrackListResp, 0)
	list, err := db.GetTrackInfoListoByVideoName(videoName)
	if err != nil{
		return nil, err
	}
	for _, item := range list{
		state := 0
		if item.LabelTrackId != ""{
				state = 1
		}
		trackIDList := strings.Split(item.Candidates, ",")
		if err != nil{
			return nil ,err
		}
		ret = append(ret, &TrackListResp{
			TrackId: item.TrackId,
			StartTime:  int(item.StartTime),
			State:state,
			RealedList: trackIDList,
		})
	}
	return ret, nil
}

func GetTrackListByDataList(videoName string)([]*model.TrackInfo, error){
	ret := make([]*model.TrackInfo, 0)
	videoNamePath := filepath.Join(config.Conf.DataPath, videoName+".json")
	byte, _ := ioutil.ReadFile(videoNamePath)
	js, err  := simplejson.NewJson(byte)
	if  err != nil{
		return nil, err
	}
	mp , err := js.Map()
	if err != nil{
		return nil, err
	}
	for track, _ := range mp{
		trackInfo := &model.TrackInfo{
			TrackId:  track,
			Videoname:  videoName,
		}
		startTime , _ := js.Get(track).Get("start_time").Int()
		trackInfo.StartTime = uint64(startTime)
		list, _ := js.Get(track).Get("candidates").Array()
		str := ""
		for i, item := range list{
			str += fmt.Sprintf("%v",item)
			if i != len(list) - 1 {
				str += ","
			}
		}
		trackInfo.Candidates = str
		logrus.Infof("Candidates = %v",trackInfo.Candidates)
		ret = append(ret, trackInfo)
	}
	return ret, nil
}

func AnnotationSubmit(c *gin.Context){
	var req AnnotationSubmitReq
	
	err := c.BindJSON(&req)
	if err != nil{
		c.JSON(400, err.Error())
		return
	}
	annotationList := req.AnnotationList
	for _, item := range annotationList{
		db.UpdateTrackInfo(req.VideoName, item, req.TrackId)
	}
	db.UpdateTrackInfo(req.VideoName, req.TrackId, req.TrackId)
	Response(c, "0", "successful", nil)
}
