package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func AuthHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求的路径
		path := c.Request.URL.Path

		// 如果请求的是登录、注册页面或静态资源，直接放行
		if path == "/login" ||
			path == "/register" ||
			path == "/v1/user/login" ||
			path == "/v1/user/register" ||
			path == "/v1/user/logout" ||
			strings.HasPrefix(path, "/css/") ||
			strings.HasPrefix(path, "/js/") || strings.HasPrefix(path, "/image/") ||
			strings.HasPrefix(path, "/plugins/") {
			c.Next()
			return
		}

		// 否则，进行认证检查
		session := sessions.Default(c)
		userID := session.Get("userID")

		// 如果userID为空，表示用户未登录，重定向到登录页面
		if userID == nil {
			c.Redirect(http.StatusFound, "/login")
			c.Abort() // 终止请求处理
			return
		}

		// // 用户已登录，将 userID 存入上下文，以便后续处理函数使用
		// c.Set("userID", userID)
		c.Next()
	}
}
