package userinternal

import (
	"database/sql"
)

type user_database struct {
	Conn *sql.DB
}
