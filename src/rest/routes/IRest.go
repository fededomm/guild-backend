package routes

import "github.com/gin-gonic/gin"

type IRest interface {
	GetAll(c *gin.Context)
	PostOne(c *gin.Context)
	GetAllPgByUser(c *gin.Context)
}
