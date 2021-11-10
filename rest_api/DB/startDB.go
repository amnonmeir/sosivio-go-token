package startDB

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	//_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	var err error
	//DB, err = sql.Open("mysql", "amnon:1234@tcp(127.0.0.1:3306)/dbclient")
	DB, err = sql.Open("postgres", "postgres://amnon:1234@localhost/dbclient?sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	if err = DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Connected to DB")
}
