package api

import (
	"encoding/json"
	"net/http"

	"github.com/faruqisan/gorep/modules/post/controller"
)

type PostApi interface {
	Get(w http.ResponseWriter, r *http.Request)
	GetUserPosts(w http.ResponseWriter, r *http.Request)
}

type postApi struct {
	PostController controller.PostController
}

func NewPostApi(pc controller.PostController) PostApi {
	return &postApi{pc}
}

func (pa *postApi) Get(w http.ResponseWriter, r *http.Request) {
	p, _ := json.Marshal(pa.PostController.Get())
	w.Header().Set("Content-Type", "application/json")
	w.Write(p)
}

func (pa *postApi) GetUserPosts(w http.ResponseWriter, r *http.Request) {
	up, _ := json.Marshal(pa.PostController.GetUserPosts())
	w.Header().Set("Content-Type", "application/json")
	w.Write(up)
}
