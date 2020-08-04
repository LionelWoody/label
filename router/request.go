package router

type AnnotationSubmitReq struct {
	VideoName string `json:"videoName"`
	TrackId string `json:"trackId"`
	AnnotationList []string `json:"annotationList"`
}

