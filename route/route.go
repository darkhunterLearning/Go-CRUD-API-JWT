package route

import (
	"fmt"
	"log"
	"net/http"

	"github.com/darkhunterLearning/Go-CRUD-API-JWT/api"
	"github.com/gorilla/mux"
)

func Init_Route() {
	router := mux.NewRouter()

	router.HandleFunc("/signup/", api.SignUp).Methods("POST")
	router.HandleFunc("/login/", api.LogIn).Methods("POST")
	fmt.Println("Listening to port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
