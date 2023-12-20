package database

import (
	"context"
	"database/sql"
	"errors"
	"guild-be/src/custom"
	"guild-be/src/models"
)

const (
	INSERT_USER             string = "INSERT INTO Users (Name, Surname, Username, BattleTag) VALUES ($1, $2, $3, $4)"
	INSERT_PG               string = "INSERT INTO Personaggi (Name, UserID, UserUsername, Class, TierSetPieces, Rank) VALUES ($1, $2, $3, $4, $5, $6)"
	SELECT_USER_ID          string = "SELECT ID FROM Users WHERE Username = $1"
	GET_ALL                 string = "SELECT Users.ID, Users.Name, Users.Surname, Users.Username, Users.BattleTag, Personaggi.ID, Personaggi.Name, Personaggi.Class, Personaggi.TierSetPieces,Personaggi.Rank FROM Users INNER JOIN Personaggi ON Users.ID = Personaggi.UserID;"
	GET_ALL_PG_FOREACH_USER string = "SELECT Users.ID, Users.Name, Users.Surname, Users.Username, Users.BattleTag, Personaggi.Name, Personaggi.Class, Personaggi.TierSetPieces, Personaggi.Rank FROM Users INNER JOIN Personaggi ON Users.ID = Personaggi.UserID WHERE Users.Name = $1;"
)

type DBService struct {
	DB *sql.DB
}
func (db *DBService) InsertUser(ctx context.Context, user models.User) error {
	_, err := db.DB.QueryContext(ctx, INSERT_USER, user.Name, user.Surname, user.Username, user.BattleTag)
	if err != nil {
		return err
	}
	return nil
}

func (db *DBService) InsertPg(ctx context.Context, pg models.Personaggio) error {
	var userID int
	err := db.DB.QueryRowContext(ctx, SELECT_USER_ID, pg.UserUsername).Scan(&userID)
	if err != nil {
		return err
	}
	if userID == 0 {
		return errors.New("user not found")
	}
	_, err = db.DB.ExecContext(
		ctx,
		INSERT_PG,
		pg.Name,
		userID,
		pg.UserUsername,
		pg.Class,
		pg.TierSetPieces,
		pg.Rank,
	)
	if err != nil {
		return err
	}
	return nil
}

func (db *DBService) GetAll(ctx context.Context) (*[]models.User, error) {
	var users []models.User
	var user models.User
	var personaggio models.Personaggio

	rows, err := db.DB.QueryContext(ctx, GET_ALL)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Surname,
			&user.Username,
			&user.BattleTag,
			&personaggio.ID,
			&personaggio.Name,
			&personaggio.Class,
			&personaggio.TierSetPieces,
			&personaggio.Rank,
		); err != nil {
			return nil, err
		}
		user.Pg = personaggio
		users = append(users, user)
	}
	return &users, nil
}

func (db DBService) GetAllPgForUser(ctx context.Context, name string) (*custom.ExampleListOfPGOfAUser, error) {
	var exampleBodyUser custom.ExampleListOfPGOfAUser
	var personaggi  []custom.ExampleBodyPg
	var personaggio custom.ExampleBodyPg

	rows, err := db.DB.QueryContext(ctx, GET_ALL_PG_FOREACH_USER, name)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		if err := rows.Scan(
			&personaggio.Name,
			&personaggio.Class,
			&personaggio.TierSetPieces,
			&personaggio.Rank,
		); err != nil {
			return nil, err
		}
		personaggi = append(personaggi, personaggio)
		exampleBodyUser.PgList = personaggi

	}
	return &exampleBodyUser, nil
}
