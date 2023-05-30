package databases

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

// This function will make a connection to the database only once.
func InitDatabase() {
	var err error

	connStr,ok := os.LookupEnv("POSTGRES_CONNECT_URL")

	if !ok{
		fmt.Println("Not found POSTGRES_CONNECT_URL")
	}

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	
	if err = db.Ping(); err != nil {
		panic(err)
	}
	
    insertStmt := `CREATE TABLE public.albums (
    name text NOT NULL,
    id bigint NOT NULL,
    price money,
    description text,
    PRIMARY KEY (id)
    )`
    _, e := db.Exec(insertStmt)

	if e != nil {
		panic(err)
	}

	fmt.Println("The database is connected")
}
