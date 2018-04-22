package routes

import (
	"net/http"

	"github.com/faruqisan/gorep/modules/post/api"
	"github.com/faruqisan/gorep/modules/post/controller"
)

func PostRoutes(pc controller.PostController) {

	postAPI := api.NewPostApi(pc)

	http.HandleFunc("/posts", postAPI.Get)
	http.HandleFunc("/user-posts", postAPI.GetUserPosts)
}
