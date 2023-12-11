package rest

import "github.com/gin-gonic/gin"

type IRest interface {
	GetAll(c *gin.Context)
	PostOne(c *gin.Context)
}
