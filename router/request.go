package router

type AnnotationSubmitReq struct {
	VideoName string `json:"videoName"`
	TrackId string `json:"trackId"`
	AnnotationList []string `json:"annotationList"`
	MergeList []string `json:"mergeList"`
	IsRelabel int `json:"isRelabel"`
}

type AnnotationDeleteReq struct {
	VideoName string `json:"videoName"`
	TrackId string `json:"trackId"`
}

