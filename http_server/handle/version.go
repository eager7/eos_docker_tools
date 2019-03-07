package handle

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const version="v1.0.0"

func Version(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": version})
}

