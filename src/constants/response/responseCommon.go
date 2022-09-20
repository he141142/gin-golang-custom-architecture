package response

import (
	"github.com/gin-gonic/gin"
)

func SystemResponse(code int, object any, context *gin.Context) {
	context.JSON(code, object)
}

func BaseResponseData[T any](object *T) map[string]interface {
} {
	return map[string]interface{}{
		"data": object,
	}
}
