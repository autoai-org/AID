package contrib

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const FILE_FIELD = "file"

func parseFormFail(context *gin.Context) {
	context.JSON(http.StatusBadRequest, gin.H{
		"message": "Can not parse form",
	})
}
