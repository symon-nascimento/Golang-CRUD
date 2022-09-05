package main

import (
	"log"
	"net/http"
    "github.com/gorilla/mux"
	"github.com/symon-nascimento/api/db"
	"github.com/symon-nascimento/api/handler"
)


func main() {
    DB := db.Init()
    h := handler.New(DB)
    router := mux.NewRouter()

    router.HandleFunc("/hello",handler.Hello )
    router.HandleFunc("/headers", handler.Headers)
    router.HandleFunc("/print", handler.WorkerPrint)
    router.HandleFunc("/persist", handler.WorkerPersist)
    router.HandleFunc("/sync", handler.WorkerSync)
    router.HandleFunc("/student", h.AddStudent).Methods(http.MethodPost)
    router.HandleFunc("/student/{id}", h.RmStudent).Methods(http.MethodDelete)
    router.HandleFunc("/student/{id}", h.UpStudent).Methods(http.MethodPut)
    router.HandleFunc("/student/{id}", h.FindOneStudent).Methods(http.MethodGet)
    
    log.Println("API Running 8090")
    http.ListenAndServe(":8090", router)
    
}