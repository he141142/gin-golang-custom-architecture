package errorHandling

import "github.com/gin-gonic/gin"

func ValidateDto(context *gin.Context,data any) error{
	if err := context.ShouldBind(&data); err != nil {
		return err
	}
	return nil
}