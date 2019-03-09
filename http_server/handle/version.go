package handle

import (
	"github.com/eager7/eos_docker_tools/http_server/message"
	"github.com/gin-gonic/gin"
	"net/http"
)

const version="v1.0.0"

func Version(context *gin.Context) {
	context.JSON(http.StatusOK, message.Resp{Code: message.Code_SUCCESS_CODE, Msg:  "pong",})
}

