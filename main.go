package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/DmitriyPrischep/Course-Project-of-Data-Base/database"
	"github.com/DmitriyPrischep/Course-Project-of-Data-Base/handlers"
)

func main() {

	database.DB.Connect()

	router := handlers.CreateRouter()

	srv := &http.Server{
		Handler:      router,
		Addr:         "0.0.0.0:5000",
		WriteTimeout: 150 * time.Second,
		ReadTimeout:  150 * time.Second,
	}
	fmt.Println("Starting server at PORT: 5000")
	log.Fatal(srv.ListenAndServe())
}
