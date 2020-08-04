package router

import (
	`github.com/gin-gonic/gin`
	
	`label-backend/config`
	`label-backend/middleware`
)

func Init(){
	r := gin.Default()
	Route(r)
	r.Run(config.Conf.HostPort)
}

func  Route(r *gin.Engine){
	r.Use(gin.Recovery())
	r.Use(middleware.TrackLogMiddleware())
	g:= r.Group("/v1")
	{
		g.GET("/annotation/list" , Annotationlist)
		g.POST("/annotation/submit", AnnotationSubmit)
	}
}