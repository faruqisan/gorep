package repository

import (
	"database/sql"

	models "github.com/faruqisan/gorep/module/user"
)

type userDatabase struct {
	Conn *sql.DB
}

func NewUserDatabase(Conn *sql.DB) UserRepository {
	return &userDatabase{Conn}
}

func (userDB *userDatabase) Get() ([]*models.User, error) {
	return nil, nil
}
func (userDB *userDatabase) Find(id int64) (*models.User, error) {
	return nil, nil
}
func (userDB *userDatabase) Store(u *models.User) error {
	return nil
}
func (userDB *userDatabase) Update(u *models.User) error {
	return nil
}
func (userDB *userDatabase) Delete(id int64) error {
	return nil
}
