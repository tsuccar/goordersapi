package application

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redis/go-redis/v9"
)

type App struct {
	router http.Handler
	rdb    *redis.Client
}

func New() *App {
	app := &App{
		router: loadRoutes(), //*chi.mux object implements http.Handler interface
		rdb: redis.NewClient(&redis.Options{
			Addr:     "goordersapi-redis-1:6379", //make sure to use docker networks hostname
			Password: "",                         // no password set
			DB:       0,                          // use default DB
		}), //the client manages the connection state internally
	}
	return app
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":8081",
		Handler: a.router,
	}

	err := a.rdb.Ping(ctx).Err()
	if err != nil {
		return fmt.Errorf("failed to connect to redis: %w", err)
	}

	fmt.Println("Starting server")
	err = server.ListenAndServe() //blocking

	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	return nil
}

//	err, open := <-ch //blocks our code's execution until it recieves a value or the channel is closed in which
//case the value will be nil
