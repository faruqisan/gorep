package post

import (
	"database/sql"

	postcontroller "github.com/faruqisan/gorep/modules/post/controller"
	postrepository "github.com/faruqisan/gorep/modules/post/repository"
	userrepository "github.com/faruqisan/gorep/modules/user/repository"
	"github.com/faruqisan/gorep/routes"
)

func InitPost(db *sql.DB) {
	postRepository := postrepository.NewPostQueries(db)
	userRepository := userrepository.NewUserQueries(db)
	postController := postcontroller.NewPostController(postRepository, userRepository)
	routes.PostRoutes(postController)
}
