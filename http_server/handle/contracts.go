package handle

import (
	"fmt"
	"github.com/eager7/eos_docker_tools/eos_cmd"
	"github.com/eager7/eos_docker_tools/http_server/message"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"os"
)

const account = "account"
const file = "files"

func ContractsHandle(context *gin.Context) {
	form, err := context.MultipartForm() //多文件上传
	if err != nil {
		log.Error(err)
		context.JSON(http.StatusBadRequest,
			message.Resp{Code: message.Code_ERROR_REQUEST_PARAM_INVALID, Msg: fmt.Sprintf("param error:%s", err)})
		return
	}
	log.Debug("param:", len(form.Value), "file:", len(form.File[file]))
	if len(form.Value[account]) <= 0 || len(form.File[file]) <= 0 {
		context.JSON(http.StatusBadRequest,
			message.Resp{Code: message.Code_ERROR_REQUEST_PARAM_INVALID, Msg: "can't receive account name or no files update"})
		return
	}
	contractDir := fmt.Sprintf("/tmp/%s", form.Value[account][0])
	if err := os.MkdirAll(contractDir, 0755); err != nil {
		context.JSON(http.StatusInternalServerError,
			message.Resp{Code: message.Code_ERROR_INTERNAL_SERVER, Msg: "can't receive account name or no files update"})
	}

	for _, f := range form.File[file] {
		log.Debug(f.Filename)
		if fi, err := f.Open(); err != nil {
			log.Error(err)
		} else {
			if buffer, err := ioutil.ReadAll(fi); err != nil {
				log.Error(err)
			} else {
				log.Info("context:", string(buffer))
			}
		}
		if err := context.SaveUploadedFile(f, fmt.Sprintf("%s/%s", contractDir, f.Filename)); err != nil {
			log.Error("save file error:", err)
		}
	}
	hash, err := eos_cmd.CompileCodeAndGetHash(contractDir)
	if err != nil {
		context.JSON(http.StatusInternalServerError,
			message.Resp{Code: message.Code_ERROR_INTERNAL_SERVER, Msg: err.Error()})
		return
	}

	context.JSON(http.StatusOK, message.Resp{Code: message.Code_SUCCESS_CODE, Msg: hash})
}
