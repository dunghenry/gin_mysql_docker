package configs

import (
	"database/sql"
	"log"
)

func GetDataBase() *sql.DB {
	db, err := sql.Open("mysql", "root:@/mydb")
	// db, err := sql.Open("mysql", "root:root@tcp(db:3306)/gin")

	if err != nil {
		log.Fatal(err)
	}
	return db
}
