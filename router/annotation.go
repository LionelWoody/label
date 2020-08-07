package router

import (
	`fmt`
	`io/ioutil`
	`os`
	`path/filepath`
	`sort`
	`strconv`
	`strings`
	`time`
	
	`github.com/bitly/go-simplejson`
	`github.com/gin-gonic/gin`
	`github.com/sirupsen/logrus`
	
	`label-backend/config`
	`label-backend/lib/db`
	`label-backend/model`
)

func GetTimeFormat(timestamp int64) string {
	cstSh, _ := time.LoadLocation("Asia/Shanghai")
	return time.Unix(timestamp, 0).In(cstSh).Format("2006-01-02 15:04:05")
}

func AnnotationQalist(c *gin.Context){
	videoName := c.Query("videoName")
	logrus.Infof("videoName : %s",videoName)
	ret , err := db.GetTrackInfoListoByVideoName(videoName, false)
	if err != nil{
		logrus.Errorf("[GetTrackListByDb] error err =%+v",err)
		Response(c, "1000",err.Error(),nil)
		return
	}
	
	mp := make(map[string][]*TrackListResp)
	for _, track := range ret{
		if track.LabelTrackId == ""{
			continue
		}
		state := 0
		if track.LabelTrackId != ""{
			state  = 1
		}
		mp[track.LabelTrackId] = append(mp[track.LabelTrackId] , &TrackListResp{
			StartTime: GetTimeFormat(int64(track.StartTime)/1000),
			EndTime:  GetTimeFormat(int64(track.EndTime)/1000),
			State:state,
			Dest: track.Dest,
			TrackDuration: int(track.EndTime)/1000 - int(track.StartTime)/1000,
			TrackId: track.TrackId,
		})
	}
	
	trackRespList := make([]*TrackQA, 0)
	for  trackId, list := range mp{
		for i := 0; i < len(list);i++{
			if list[i].TrackId == trackId{
				list[i], list[0]  = list[0],list[i]
			}
		}
		trackRespList = append(trackRespList,&TrackQA{
			TrackId:trackId,
			LabelList: list,
		})
	}
	sort.Slice(trackRespList, func(i, j int) bool {
		return trackRespList[i].TrackId < trackRespList[j].TrackId
	})
	data := TrackQAList{
		List:  trackRespList,
	}
	Response(c, "0","",data)
}

func Annotationlist(c *gin.Context){
		videoName := c.Query("videoName")
		isRelabel := c.Query("isRelabel")
		isRelabelInt, _ := strconv.Atoi(isRelabel)
		logrus.Infof("videoName : %s",videoName)
		if !db.IsTrackInfoVideoNameHave(videoName) {
			list  ,err := GetTrackListByDataList(videoName)
			if err != nil{
				logrus.Errorf("[GetTrackListByDataList] error err =%+v",err)
				Response(c, "1000",err.Error(),nil)
				return
			}
			for _, item := range list {
				db.InsertTrackInfo(item)
			}
		}
		ret , err := GetTrackListByDb(videoName, isRelabelInt == 1)
		if err != nil{
			logrus.Errorf("[GetTrackListByDb] error err =%+v",err)
			Response(c, "1000",err.Error(),nil)
			return
		}
		data := TrackList{
			List:  ret,
		}
		Response(c, "0","",data)
}

func GetTrackListByDb(videoName string, isRelabel bool)([]*TrackListResp, error){
	ret := make([]*TrackListResp, 0)
	list1, err := db.GetTrackInfoListoByVideoName(videoName, isRelabel)
	if err != nil{
		return nil, err
	}
	list:=make([]*model.TrackInfo,0)
	orderList  := make([]*model.TrackInfo,0)
	
	for _, item := range list1{
		if item.LabelTrackId != ""{
			orderList = append(orderList, item)
		}else{
			list = append(list, item)
		}
	}
	list = append(list, orderList...)
	for _, item := range list{
		state := 0
		if item.LabelTrackId != ""{
				state = 1
		}
		trackIDList := strings.Split(item.Candidates, ",")
		if err != nil{
			return nil ,err
		}
		
		filterTrackIdList := make([]string, 0)
		for _, item := range trackIDList{
			if item == ""{
				continue
			}
			filterTrackIdList = append(filterTrackIdList, item)
		}
		newTrack := &TrackListResp{
			TrackId: item.TrackId,
			StartTime: GetTimeFormat(int64(item.StartTime)/1000),
			EndTime:  GetTimeFormat(int64(item.EndTime)/1000),
			State:state,
			RealedList: filterTrackIdList,
			Dest: item.Dest,
			TrackDuration: int(item.EndTime)/1000 - int(item.StartTime)/1000,
		}
		if state == 1{
			newTrack.MarkedByOrigin = item.LabelTrackId
		}
		ret = append(ret, newTrack)
		
	}
	return ret, nil
}

func GetTrackListByDataList(videoName string)([]*model.TrackInfo, error){
	ret := make([]*model.TrackInfo, 0)
	videoNamePath := ""
	filepath.Walk(config.Conf.DataPath, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir(){
			if strings.Contains(info.Name(),videoName){
				videoNamePath = path
			}
		}
		return nil
	})
	logrus.Infof("[GetTrackListByDataList] videoName : %s videoNamePath: %s ",videoName,videoNamePath)
	byte, _ := ioutil.ReadFile(videoNamePath)
	js, err  := simplejson.NewJson(byte)
	if  err != nil {
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
		endTime  , _ := js.Get(track).Get("end_time").Int()
		desc  , _ := js.Get(track).Get("dets").Int()
		trackInfo.StartTime = uint64(startTime)
		trackInfo.EndTime = uint64(endTime)
		trackInfo.Dest = desc
		logrus.Infof("endTime : %v",endTime)
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
	labelTrackId := req.TrackId
	if len(req.MergeList) > 0 && req.IsRelabel == 0{
		mergeId, ok, err := chekoutMergeList(req.VideoName, req.MergeList)
		if err != nil {
			Response(c, "2000", err.Error(), nil)
		}
		if ok {
			labelTrackId = mergeId
			annotationList = append(annotationList, req.MergeList...)
		}
	}
	for _, item := range annotationList{
		db.UpdateTrackInfo(req.VideoName, item, labelTrackId)
	}
	db.UpdateTrackInfo(req.VideoName, req.TrackId, labelTrackId)
	Response(c, "0", "successful", nil)
}
func AnnotationDelete(c *gin.Context){
	var req AnnotationDeleteReq
	
	err := c.BindJSON(&req)
	if err != nil{
		c.JSON(400, err.Error())
		return
	}
	err = db.SetTrackInfoNil(req.VideoName, req.TrackId)
	if err != nil{
		Response(c, "1000", err.Error(), nil)
		return
	}
	Response(c, "0", "successful", nil)
}

func AnnotationRelabel(c *gin.Context){
	var req AnnotationDeleteReq
	
	err := c.BindJSON(&req)
	if err != nil{
		c.JSON(400, err.Error())
		return
	}
	err = db.SetTrackInfoReLabel(req.VideoName, req.TrackId)
	if err != nil{
		Response(c, "1000", err.Error(), nil)
		return
	}
	Response(c, "0", "successful", nil)
}


func chekoutMergeList(videoName string , mergeList []string)(mergeId string, ok bool, err error){
	for _, item := range mergeList{
		trackInfo , err := db.GetTrackInfoByVideoNameAndTrackId(videoName, item)
		if err != nil{
			return "", false, err
		}
		if mergeId == ""{
			mergeId = trackInfo.LabelTrackId
		}
		if trackInfo.LabelTrackId != mergeId{
			ok = false
			return"", false, err
		}
	}
	if mergeId != "" {
		ok = true
	}
	return
}
