package rest

import (
	"apocalypse/database"
	"apocalypse/models"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type Rest struct {
	DB *sql.DB
}
type IRest interface {
	GetAll(c *gin.Context)
	PostOne(c *gin.Context)
}

var validate = validator.New()

func (r *Rest) GetAll(c *gin.Context) {

	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (r *Rest) PostOne(c *gin.Context) {

	person := new(models.Users)

	if err := c.BindJSON(person); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if validationError := validate.Struct(person); validationError != nil {
		c.JSON(400, gin.H{"error": validationError.Error()})
		return
	}

	_, err := r.DB.Exec(
		database.INSERT.String(),
		person.Name,
		person.Surname,
		person.Username,
		person.Class,
		person.BattleTag,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"message": "created",
	})
}
