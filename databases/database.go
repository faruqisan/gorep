package databases

import (
	"database/sql"
	"os"

	"github.com/faruqisan/gorep/configs"
)

func GetDatabase(conf configs.Config) (db *sql.DB, err error) {

	dbConfig := conf.DBConf

	switch os.Getenv("DB") {
	case "PGQL":
		db, err = getPGQL(dbConfig)
		break
	case "MY":
		db, err = getMYSQL(dbConfig)
		break
	default:
		db, err = getPGQL(dbConfig)
		break
	}

	return db, nil
}
