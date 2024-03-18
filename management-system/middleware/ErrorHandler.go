package middleware

import (
	"fmt"
	"management-system/response"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				err, ok := r.(error)
				if !ok {
					err = fmt.Errorf("%v", r)
				}
				app := response.Gin{C: c}
				app.ResponseErr(err.Error(), nil)
				c.Abort()
			}
		}()
		c.Next()
	}
}
