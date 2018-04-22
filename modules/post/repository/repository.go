package repository

import (
	"github.com/faruqisan/gorep/modules/post/model"
)

type PostRepository interface {
	Get() []model.Post
}
