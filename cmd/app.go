package main

import (
	"log"
	"net/http"
	"time"

	"github.com/faruqisan/gorep/databases"
	"github.com/faruqisan/gorep/modules/post"
)

func main() {

	db, err := databases.GetPgQL()
	if err != nil {
		log.Panicln(err)
	}
	defer db.Close()

	post.InitPost(db)

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	go func() {
		for t := range ticker.C {
			log.Println(t)
		}
	}()

	log.Printf("****Server Running****")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
