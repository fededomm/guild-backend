package database

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"guild-be/src/custom"
	"guild-be/src/models"
)

const (
	INSERT_USER             string = "INSERT INTO Users (Name, Surname, Username, BattleTag) VALUES ($1, $2, $3, $4)"
	INSERT_PG               string = "INSERT INTO Personaggi (Name, UserUsername, Class, TierSetPieces, Rank) VALUES ($1, $2, $3, $4, $5)"
	SELECT_USER_ID          string = "SELECT Username FROM Users WHERE Username = $1"
	GET_ALL                 string = `SELECT Users.ID, Users.Name, Users.Surname, Users.Username, Users.BattleTag FROM Users`
	GET_ALL_PG_FOREACH_USER string = `SELECT Users.Name, Users.Username, Users.BattleTag, json_agg(Personaggi) as PgList
	FROM Users
	LEFT JOIN Personaggi ON Users.Username = Personaggi.UserUsername
	WHERE Users.Username = $1
	GROUP BY users.name, users.username, users.battletag;`
)

type DBService struct {
	DB     *sql.DB
}

func (db *DBService) InsertUser(ctx context.Context, user models.User) error {
	_, err := db.DB.QueryContext(ctx, INSERT_USER, user.Name, user.Surname, user.Username, user.BattleTag)
	if err != nil {
		return err
	}
	return nil
}

func (db *DBService) InsertPg(ctx context.Context, pg models.Personaggio) error {
	var userName string
	err := db.DB.QueryRowContext(ctx, SELECT_USER_ID, pg.UserUsername).Scan(&userName)
	if err != nil && err != sql.ErrNoRows {

		return err
	}
	if sql.ErrNoRows == err {

		return errors.New("userID not found")
	}
	_, err = db.DB.ExecContext(
		ctx,
		INSERT_PG,
		pg.Name,
		userName,
		pg.Class,
		pg.TierSetPieces,
		pg.Rank,
	)
	if err != nil {

		return err
	}
	return nil
}

func (db *DBService) GetAll(ctx context.Context) (*[]custom.ExampleBodyUser, error) {
	var users []custom.ExampleBodyUser
	var user custom.ExampleBodyUser
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
		); err != nil {

			return nil, err
		}
		users = append(users, user)
	}
	return &users, nil
}

func (db *DBService) GetAllPgForUser(ctx context.Context, username string) (*custom.ExampleListOfPGOfAUser, error) {
	rows, err := db.DB.QueryContext(ctx, GET_ALL_PG_FOREACH_USER, username)
	if err != nil {

		return &custom.ExampleListOfPGOfAUser{}, err
	}
	defer rows.Close()

	var user custom.ExampleListOfPGOfAUser
	for rows.Next() {
		var pg []byte
		err := rows.Scan(
			&user.Name,
			&user.Username,
			&user.BattleTag,
			&pg,
		)
		if err != nil {

			return &custom.ExampleListOfPGOfAUser{}, err
		}
		var pgList []custom.ExampleBodyPg
		err = json.Unmarshal(pg, &pgList)
		if err != nil {

			return &custom.ExampleListOfPGOfAUser{}, err
		}
		user.PgList = pgList
	}

	if err := rows.Err(); err != nil {

		return &custom.ExampleListOfPGOfAUser{}, err
	}

	return &user, nil
}

func (db *DBService) DeleteUserAndPg(ctx context.Context, username string) error {
	tx, err := db.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	err = DeletePg(tx, ctx, models.Personaggio{UserUsername: username})
	if err != nil {
		tx.Rollback()
		return err
	}
	err = DeleteUser(tx, ctx, models.User{Username: username})
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func DeletePg(Tx *sql.Tx, ctx context.Context, pg models.Personaggio) error {
	_, err := Tx.ExecContext(
		ctx,
		"DELETE FROM Personaggi WHERE Name = $1 AND UserUsername = $2",
		pg.Name,
		pg.UserUsername,
	)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(Tx *sql.Tx, ctx context.Context, user models.User) error {
	_, err := Tx.ExecContext(
		ctx,
		"DELETE FROM Users WHERE Username = $1",
		user.Username,
	)
	if err != nil {
		return err
	}
	return nil
}