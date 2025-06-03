package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var Store = newNotes()

func main() {
	router := httprouter.New()
	router.GET("/notes", HandleGetNotes)
	router.GET("/notes/:id", HandleGetNoteByID)
	router.POST("/notes", HandlePostNote)
	router.PUT("/notes/:id", HandlePutNoteByID)
	router.DELETE("/notes/:id", HandleDeleteNoteByID)

	fmt.Println("Server Listening on http://localhost:2222")
	if err := http.ListenAndServe(":2222", router); err != nil {
		log.Fatal(err)
	}
}
