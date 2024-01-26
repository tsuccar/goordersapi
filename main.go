package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	if port == "" {
		log.Fatal("PORT environment variable is not set")
	}
  
	server := &http.Server{
		Addr:    ":" + port,
		Handler: http.HandlerFunc(basicHandler),
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
