package rest

import (
	"apocalypse/database"
	"apocalypse/models"
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type Rest struct {
	DB *sql.DB
}

func (r *Rest) GetAll(c *gin.Context) {
	var persons []models.User
	rows, err := r.DB.Query(database.GET_ALL.String())
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	for rows.Next() {
		person := new(models.User)
		if err := rows.Scan(&person.ID, &person.Name, &person.Surname, &person.Username, &person.BattleTag, &person.Pg); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		persons = append(persons, *person)
		c.JSON(200, persons)
	}

}

func (r *Rest) PostOne(c *gin.Context) {
	person := new(models.User)

	if err := c.BindJSON(person); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := CustomValidatorGin(person, c); err != nil {
		log.Err(err).Msgf("Error validating model: %s", err.Error())
		return
	}

	_, err := r.DB.Exec(
		database.INSERT.String(),
		person.Name,
		person.Surname,
		person.Username,
		person.BattleTag,
		person.Pg,
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, gin.H{
		"message": "created",
	})
}
