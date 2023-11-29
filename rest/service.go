package rest

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	DB *sql.DB
}
type IRest interface {
	GetAll(c *gin.Context)
}

func (r *Rest) GetAll(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
