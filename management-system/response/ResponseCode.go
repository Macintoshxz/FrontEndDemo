package response

import (
	"github.com/gin-gonic/gin"
)

const (
	SuccessCode     = 2001
	FailCode        = 4000
	CheckFailCode   = 4003
	ServerErrorCode = 5000
)

type ResponseStruct struct {
	HttpStatus int   `json:"http_status"` //http状态
	Code       int    `json:"code"`//状态码
	Data       gin.H  `json:"data"`//数据
	Msg        string `json:"msg"`//信息
}


