package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"label-backend/config"
	"label-backend/middleware"
)

func Init() {
	r := gin.Default()
	Route(r)
	r.Run(config.Conf.HostPort)
}

func Route(r *gin.Engine) {
	r.Use(gin.Recovery())
	r.Use(middleware.TrackLogMiddleware())
	r.StaticFS("/data", http.Dir(config.Conf.DataPath))
	g := r.Group("/v1")
	{
		g.GET("/annotation/list", Annotationlist)
		g.GET("/annotation/qa/list", AnnotationQalist)
		g.POST("/annotation/submit", AnnotationSubmit)
		g.POST("/annotation/delete", AnnotationDelete)
		g.POST("/annotation/rejectGroup", AnnotationRelabel)
		g.POST("/annotation/recordTime", RecordTime)
	}
}
