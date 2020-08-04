package middleware

import (
	`github.com/gin-gonic/gin`
	`github.com/gofrs/uuid`
)

func GenRequsetId() string {
	u1 := uuid.Must(uuid.NewV4())
	return u1.String()
}

func TrackLogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId :=GenRequsetId()
		c.Set("requestId", requestId)
		c.Next()
	}
}

