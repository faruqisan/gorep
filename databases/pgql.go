package databases

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/faruqisan/gorep/configs"

	_ "github.com/lib/pq"
)

func getPGQL(conf configs.DatabaseConfig) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s", conf.Username, conf.Password, conf.DBName)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return db, nil
}
