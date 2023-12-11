package database

type Query int

const (
	INSERT Query = iota
	GET_ALL
	GET_BY_ID
	UPDATE
	DELETE
)

var queries = [...]string{
	"INSERT INTO Users (Name, Surname, Username, Class, BattleTag) VALUES (?, ?, ?, ?, ?)",
	"SELECT * FROM Users",
	"SELECT * FROM Users WHERE id = ?",
	"UPDATE Users SET name = ?, surname = ? WHERE id = ?",
	"DELETE FROM Users WHERE id = ?",
}

func (q Query) String() string {
	if q < 0 || int(q) >= len(queries) {
		return "Unknown"
	}
	return queries[q]
}
