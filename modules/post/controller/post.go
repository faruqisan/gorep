package controller

import (
	postmodel "github.com/faruqisan/gorep/modules/post/model"
	postrepository "github.com/faruqisan/gorep/modules/post/repository"
	usermodel "github.com/faruqisan/gorep/modules/user/model"
	userrepository "github.com/faruqisan/gorep/modules/user/repository"
)

type PostController interface {
	Get() []postmodel.Post
	GetUserPosts() []postmodel.UserPosts
}

type postController struct {
	postRepo postrepository.PostRepository
	userRepo userrepository.UserRepository
}

func NewPostController(pr postrepository.PostRepository, ur userrepository.UserRepository) PostController {
	return &postController{
		postRepo: pr,
		userRepo: ur,
	}
}

func (pc *postController) Get() []postmodel.Post {
	posts := pc.postRepo.Get()
	return posts
}

func (pc *postController) GetUserPosts() []postmodel.UserPosts {
	users := pc.userRepo.Get()
	posts := pc.postRepo.Get()

	return mapUserPosts(users, posts)
}

func mapUserPosts(users []usermodel.User, posts []postmodel.Post) (res []postmodel.UserPosts) {

	for _, user := range users {

		userPosts := postmodel.UserPosts{}
		userPosts.UserID = user.ID

		for _, post := range posts {

			if user.ID == post.UserID {
				userPosts.Posts = append(userPosts.Posts, post)
			}

		}

		res = append(res, userPosts)

	}

	return res
}
