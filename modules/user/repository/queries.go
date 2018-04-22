package repository

import (
	"database/sql"
	"log"

	"github.com/faruqisan/gorep/modules/user/model"
)

type userQueries struct {
	db *sql.DB
}

func NewUserQueries(db *sql.DB) UserRepository {
	return &userQueries{db}
}

//Get return list of user
func (uq *userQueries) Get() []model.User {

	users := []model.User{}

	db := uq.db

	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var u model.User
		err := rows.Scan(&u.ID, &u.Name)
		if err != nil {
			log.Println(err)
		}
		users = append(users, u)
	}

	return users
}
