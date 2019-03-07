package handle

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

func ContractsHandle(context *gin.Context) {
	form, err := context.MultipartForm()//多文件上传
	if err != nil {
		log.Error(err)
		context.JSON(500, gin.H{"message": "param error"})
		return
	}
	log.Debug(form.Value, form.File)
	log.Debug(form.Value["account"], len(form.File["file"]))
	if len(form.File["file"]) > 0 {
		f := form.File["file"][0]
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
		if err := context.SaveUploadedFile(f, fmt.Sprintf("/tmp/file/%s", f.Filename)); err != nil {
			log.Error("save file error:", err)
		}
	}

	context.JSON(200, gin.H{"message": "post resp"})
}
