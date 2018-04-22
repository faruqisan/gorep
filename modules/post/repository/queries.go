package repository

import (
	"database/sql"
	"log"

	"github.com/faruqisan/gorep/modules/post/model"
)

type postQueries struct {
	db *sql.DB
}

func NewPostQueries(db *sql.DB) PostRepository {
	return &postQueries{db}
}

func (pq *postQueries) Get() []model.Post {

	posts := []model.Post{}

	db := pq.db

	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		log.Println(err)
	}

	defer rows.Close()

	for rows.Next() {
		var p model.Post
		err := rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Body)
		if err != nil {
			log.Println(err)
		}
		posts = append(posts, p)
	}

	return posts
}
