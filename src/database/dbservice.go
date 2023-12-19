package database

import (
	"context"
	"database/sql"
	"guild-be/src/models"
)

const (
	INSERT_USER string = "INSERT INTO Users (Name, Surname, Username, BattleTag) VALUES ($1, $2, $3, $4) RETURNING id"
	INSERT_PG string = "INSERT INTO Personaggi (Name, UserID, Class, TierSetPieces, Rank) VALUES ($1, $2, $3, $4, $5)"
	GET_ALL string = "SELECT Users.ID, Users.Name, Users.Surname, Users.Username, Users.BattleTag, Personaggi.ID, Personaggi.Name, Personaggi.Class, Personaggi.TierSetPieces,Personaggi.Rank FROM Users INNER JOIN Personaggi ON Users.ID = Personaggi.UserID;"
	GET_ALL_PG_FOREACH_USER string = "SELECT Users.ID, Users.Name, Users.Surname, Users.Username, Users.BattleTag, Personaggi.Name, Personaggi.Class, Personaggi.TierSetPieces, Personaggi.Rank FROM Users INNER JOIN Personaggi ON Users.ID = Personaggi.UserID WHERE Users.Name = $1;"
	UPDATE string = "UPDATE Users SET name = $1, surname = $2 WHERE id = $3"
	DELETE string = "DELETE FROM Users WHERE id = $1"
)

type DBService struct {
	DB *sql.DB
}

func (db *DBService) DoTrx(ctx context.Context, user models.User) error {
	Tx, err := db.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	if lastInsert, err := InserUser(Tx, user); err != nil {
		Tx.Rollback()
		return err
	} else {
		if err := InsertPg(Tx, user, lastInsert); err != nil {
			return err
		}
	}
	if err := Tx.Commit(); err != nil {
		return err
	}
	return nil
}

func InserUser(tx *sql.Tx, user models.User) (int, error) {
	var lastInsertedId int = 0
	if x, err := tx.Prepare(INSERT_USER); err != nil {
		return 0, err
	} else {
		if err := x.QueryRow(user.Name, user.Surname, user.Username, user.BattleTag).Scan(&lastInsertedId); err != nil {
			return 0, err
		}
	}
	return lastInsertedId, nil
}

func InsertPg(tx *sql.Tx, user models.User, lastInsertedId int) error {
	_, err := tx.Exec(
		INSERT_PG,
		user.Pg.Name,
		lastInsertedId,
		user.Pg.Class,
		user.Pg.TierSetPieces,
		user.Pg.Rank,
	)
	if err != nil {
		return err
	}
	return nil
}

func (db *DBService) GetAll() (*[]models.User, error) {
	var users []models.User
	var user models.User
	var personaggio models.Personaggio

	rows, err := db.DB.Query(GET_ALL)
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

func (db DBService) GetAllPgForUser(name string) (*[]models.User, error) {
	var users []models.User
	var user models.User
	var personaggio models.Personaggio

	rows, err := db.DB.Query(GET_ALL_PG_FOREACH_USER, name)
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