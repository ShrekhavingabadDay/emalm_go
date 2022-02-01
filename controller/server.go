package controller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var router *mux.Router

func initHandlers() {
	router.HandleFunc("/api/all_users", GetAllUsers).Methods("GET")
	router.HandleFunc("/api/images/all", GetAllImages).Methods("GET")
	router.HandleFunc("/api/comments/all", GetAllComments).Methods("GET")
	router.HandleFunc("/api/upload/image", UploadImage).Methods("POST")
	router.HandleFunc("/api/upload/comments/{content_type}/{uuid}", PostComment).Methods("POST")
}

func Start() {
	router = mux.NewRouter()

	initHandlers()
	fmt.Printf("Router initialized and serving on 6969\n")
	log.Fatal(http.ListenAndServe(":6969", router))
}
