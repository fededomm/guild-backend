package routes

import "github.com/gin-gonic/gin"

type IRest interface {
	GetAllUsers(c *gin.Context)
	PostUser(c *gin.Context)
	PostPg(c *gin.Context)
	GetAllPgByUser(c *gin.Context)
	DeletePgsAndUser(c *gin.Context)
}
