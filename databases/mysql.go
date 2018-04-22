package databases

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/faruqisan/gorep/configs"

	_ "github.com/go-sql-driver/mysql"
)

func getMYSQL(conf configs.DatabaseConfig) (*sql.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@/%s", conf.Username, conf.Password, conf.DBName)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return db, nil
}
