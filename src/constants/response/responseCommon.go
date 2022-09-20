package response

import (
	"github.com/gin-gonic/gin"
)

func SystemResponse(code int, object any, context *gin.Context) {
	context.JSON(code, object)
}


