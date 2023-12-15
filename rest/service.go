package rest

import (
	"apocalypse/database"
	"apocalypse/models"
	"context"
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
)

type Rest struct {
	DB *sql.DB
}

func (r *Rest) GetAll(c *gin.Context) {
	list, err := database.GetAll(r.DB, c)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, list)
}

func (r *Rest) PostOne(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	person := new(models.User)
	if err := c.BindJSON(person); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := CustomValidatorGin(person); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := database.DoTrx(r.DB, ctx, *person); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{
		"message": "created",
		"status":  "201",
		"body":    person,
	})
}
