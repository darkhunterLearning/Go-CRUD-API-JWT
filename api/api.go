package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/darkhunterLearning/Go-CRUD-API-JWT/action"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var router *mux.Router

//----------------------ROUTES-------------------------------
//create a mux router
func CreateRouter() {
	router = mux.NewRouter()
}

func ServerStart() {
	fmt.Println("Server started at http://localhost:8080")
	err := http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Access-Control-Allow-Origin", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router))
	if err != nil {
		log.Fatal(err)
	}
}

//initialize all routes
func InitializeRoute() {
	router.HandleFunc("/signup", action.SignUp).Methods("POST")
	router.HandleFunc("/signin", action.SignIn).Methods("POST")
	router.HandleFunc("/admin", action.IsAuthorized(action.AdminIndex)).Methods("GET")
	router.HandleFunc("/user", action.IsAuthorized(action.UserIndex)).Methods("GET")
	router.HandleFunc("/", action.Index).Methods("GET")
	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Request-Headers, Access-Control-Request-Method, Connection, Host, Origin, User-Agent, Referer, Cache-Control, X-header")
	})
}
