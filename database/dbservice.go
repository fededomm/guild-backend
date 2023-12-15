package database

import (
	"apocalypse/models"
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"
)

const (
	INSERT_USER string = "INSERT INTO Users (Name, Surname, Username, BattleTag) VALUES ($1, $2, $3, $4) RETURNING id"
	INSERT_PG   string = "INSERT INTO Personaggi (Name, UserID, Class, TierSetPieces, Rank) VALUES ($1, $2, $3, $4, $5)"
	GET_ALL     string = "SELECT Users.ID, Users.Name, Users.Surname, Users.Username, Users.BattleTag, Personaggi.ID, Personaggi.Name, Personaggi.Class, Personaggi.TierSetPieces, Personaggi.Rank FROM Users INNER JOIN Personaggi ON Users.ID = Personaggi.UserID;"
	GET_BY_ID   string = "SELECT * FROM Users WHERE id = $1"
	UPDATE      string = "UPDATE Users SET name = $1, surname = $2 WHERE id = $3"
	DELETE      string = "DELETE FROM Users WHERE id = $1"
)

func DoTrx(db *sql.DB, ctx context.Context, user models.User, c *gin.Context) error {
	Tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		c.JSON(500, gin.H{"step":"begin transaction","error": err.Error()})
		return err
	}
	if lastInsert, err := InserUser(Tx, user, c); err != nil {
		c.JSON(500, gin.H{"step":"insert on table users","error": err.Error()})
		Tx.Rollback()
		return err
	} else {
		if err := InsertPg(Tx, user, lastInsert); err != nil {
			
			c.JSON(500, gin.H{"step":"insert on table personaggi","error": err.Error()})
			return err
		}
	}
	if err := Tx.Commit(); err != nil {
		c.JSON(500, gin.H{"step":"commit transaction","error": err.Error()})
		return err
	}
	return nil
}

func InserUser(tx *sql.Tx, user models.User, c *gin.Context) (int,error) {
	var lastInsertedId int = 0
	if x, err := tx.Prepare(INSERT_USER); err != nil {
		c.JSON(500, gin.H{"step":"prepare insert on table users","error": err.Error()})
		return 0,err
	} else {
		if err := x.QueryRow(user.Name,user.Surname,user.Username,user.BattleTag,).Scan(&lastInsertedId); err != nil {
			c.JSON(500, gin.H{"step":"insert on table users","error": err.Error()})
			return 0,err
		}
	}
	return lastInsertedId,nil
}

func InsertPg(tx *sql.Tx, person models.User, lastInsertedId int) error {
	_, err := tx.Exec(
		INSERT_PG,
		person.Pg.Name,
		lastInsertedId,
		person.Pg.Class,
		person.Pg.TierSetPieces,
		person.Pg.Rank,
	)
	if err != nil {
		return err
	}
	return nil
}

func GetAll(db *sql.DB) (*sql.Rows, error) {
	row, err := db.Query(GET_ALL)
	if err != nil {
		return nil, err
	}
	return row, nil
}

