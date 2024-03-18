package response

import (
	"github.com/gin-gonic/gin"
)

type Gin struct {
	C *gin.Context
}

type Response struct {
	Code int         `json:"status"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type ResponsePage struct {
	Response
	Total int `json:"total"`
}

// Response setting gin.JSON
func (g *Gin) ResponseSuccess(msg string, data interface{}) {
	g.C.JSON(200, Response{
		Code: SuccessCode,
		Msg:  msg,
		Data: data,
	})
	return
}

func (g *Gin) ResponseFail(msg string, data interface{}) {
	g.C.JSON(200, Response{
		Code: FailCode,
		Msg:  msg,
		Data: data,
	})
	return
}

func (g *Gin) ResponseErr(msg string, data interface{}) {
	g.C.JSON(200, Response{
		Code: ServerErrorCode,
		Msg:  msg,
		Data: data,
	})
	return
}

func (g *Gin) Response(httpCode, code int, msg string, data interface{}) {
	g.C.JSON(httpCode, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
	return
}

// Response setting gin.JSON
func (g *Gin) ResponsePage(msg string, data interface{}, total int) {

	g.C.JSON(200, ResponsePage{
		Response: Response{
			Code: SuccessCode,
			Msg:  msg,
			Data: data,
		},
		Total: total,
	})
	return
}
