package main

import (
	"controllers"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"models"
	"net/http"
	"os"
)


func init() {
	client := models.DBClient{
		Username: "root",
		Password: "tarun123",
		DBName:   "codezorro",
	}
	models.ConnectToDB(client)
	models.MakeMigrations()
}

func serveHome(w http.ResponseWriter, r *http.Request) {
	log.Println(r.URL)
	if r.URL.Path != "/" {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	http.ServeFile(w, r, "home.html")
}



func main() {

	router := mux.NewRouter()
	//router.Use(AuthMiddleware.JwtAuthentication) //attach JWT auth middleware
	router.HandleFunc("/",serveHome)
	router.HandleFunc("/api/user/new", controllers.CreateAccount).Methods("POST")
	router.HandleFunc("/api/user/login", controllers.Authenticate).Methods("POST")
	router.HandleFunc("/api/code/new", controllers.CreateCode).Methods("POST")
	router.HandleFunc("/api/me/code", controllers.GetCodeFor).Methods("GET")
	router.HandleFunc("/api/runcode",controllers.RunCode).Methods("POST")
	router.HandleFunc("/ws",controllers.HandleSocketComs).Methods("POST")

	port := os.Getenv("PORT") //Get port from .env file, we did not specify any port so this should return an empty string when tested locally
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Println(port)

	err := http.ListenAndServe(":" + port, router) //Launch the app, visit localhost:8000/api

	if err != nil {
		fmt.Print(err)
	}
}
