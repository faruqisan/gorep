package repository

import (
	"github.com/faruqisan/gorep/modules/user/model"
)

type UserRepository interface {
	Get() []model.User
}
