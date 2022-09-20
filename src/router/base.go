package router

import "github.com/gin-gonic/gin"

type BaseRouter struct {
	name    string
	router  *gin.RouterGroup
}