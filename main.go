package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Notes struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func main() {
	fmt.Println("Server Listening on http://localhost:2222")
	if err := http.ListenAndServe(":2222", nil); err != nil {
		log.Fatal(err)
	}
}
