package main

import (
	"ChessWeb/backend/api"
	"log"
	"net/http"
)

func main() {

	fs := http.FileServer(http.Dir("./frontend"))
	http.Handle("/", fs)
	http.HandleFunc("/vMoveCheck", api.VMoveCheck)

	log.Println("Server starting on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
