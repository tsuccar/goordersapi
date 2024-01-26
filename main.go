package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {

	//Port
	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}
	//Chi Router
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/hello", basicHandler)
	router.Post("/hello", basicHandler)
	server := &http.Server{
		Addr:    ":" + port,
		Handler: router,
	}



	// Server Listening
	pid := os.Getpid()
	log.Printf("Current process ID: %d\n", pid)
	log.Printf("Serving on port: %s\n", port)
		// err:= server.ListenAndServe()
	// if err != nil {
	// 	fmt.Println("failed to listen to server",err)
	// }
	log.Fatalf("Error Occured : %s \n", server.ListenAndServe().Error())

}

func basicHandler(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			if r.URL.Path == "/foo"{
				w.Write([]byte("you're in : /foo , Method : Get"))
			}
		}
		if r.Method == http.MethodPost {
			w.Write([]byte("you're in : / , Method : Post"))
		 }
		w.Write([]byte("Hello, world! 你好"))
	}
