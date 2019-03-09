package handle

import (
	"github.com/eager7/eos_docker_tools/http_server/message"
	"github.com/eager7/go/mlog"
	"github.com/gin-gonic/gin"
	"net/http"
)

var log = mlog.L

func Ping(context *gin.Context) {
	context.JSON(http.StatusOK, message.Resp{Code: message.Code_SUCCESS_CODE, Msg:  "pong",})
}
