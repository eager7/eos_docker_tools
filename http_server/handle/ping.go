package handle

import (
	"github.com/eager7/eos_docker_tools/http_server/code"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, code.Resp{
		ErrCode: code.ErrCode_SUCCESS_CODE,
		ErrMsg:  "pong",
		Payload: nil,
	})
}
