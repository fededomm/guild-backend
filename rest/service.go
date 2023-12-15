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
	var users []models.User
	var user models.User
	var personaggio models.Personaggio

	rows, err := database.GetAll(r.DB)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	for rows.Next() {
		if err := rows.Scan(
			&user.ID, 
			&user.Name, 
			&user.Surname, 
			&user.Username, 
			&user.BattleTag, 
			&user.Pg,
			&personaggio.ID,
			&personaggio.UserID, 			
			&personaggio.Name, 
			&personaggio.Class, 
			&personaggio.TierSetPieces, 
			&personaggio.Rank,
		); err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		user.Pg = personaggio
		users = append(users, user)
		
	}
	c.JSON(200, users)
}

func (r *Rest) PostOne(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	person := new(models.User)
	if err := c.BindJSON(person); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := CustomValidatorGin(person, c); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if err := database.DoTrx(r.DB, ctx, *person, c); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{
		"message": "created",
		"status":  "201",
		"body":    person,
	})
}
