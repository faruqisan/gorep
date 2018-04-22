package main

import (
	"log"
	"net/http"

	"github.com/faruqisan/gorep/databases"
	"github.com/faruqisan/gorep/modules/post"
)

func main() {

	db, err := databases.GetDB()
	if err != nil {
		log.Panicln(err)
	}
	defer db.Close()

	post.InitPost(db)

	log.Fatal(http.ListenAndServe(":3000", nil))

}
