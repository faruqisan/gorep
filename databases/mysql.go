package databases

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func GetDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@/modellingkol")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return db, nil
}
