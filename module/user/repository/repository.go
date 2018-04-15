package repository

import (
	models "github.com/faruqisan/gorep/module/user"
)

type UserRepository interface {
	Get() ([]*models.User, error)
	Find(id int64) (*models.User, error)
	Store(u *models.User) error
	Update(u *models.User) error
	Delete(id int64) error
}
