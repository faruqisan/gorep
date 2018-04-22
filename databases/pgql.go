package databases

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func GetPgQL() (*sql.DB, error) {
	connStr := "user=faruqisan password=faruqisan dbname=faruqisan"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return db, nil
}
