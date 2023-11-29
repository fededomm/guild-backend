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
	"INSERT INTO <table_bame> (column1, column2) VALUES (?, ?)",
	"SELECT * FROM <table_bame>",
	"SELECT * FROM <table_bame> WHERE id = ?",
	"UPDATE <table_bame> SET column1 = ?, column2 = ? WHERE id = ?",
	"DELETE FROM <table_bame> WHERE id = ?",
}

func (q Query) String() string {
	if q < 0 || int(q) >= len(queries) {
		return "Unknown"
	}
	return queries[q]
}

