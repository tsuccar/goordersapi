package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

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
	// this is to close our redis connection after reached the end of the error channel
	// processing. it's in annoymous wrap b/c defer doen'st work with errors.

	defer func() {
		if err := a.rdb.Close(); err != nil {
			fmt.Println("failed to close redis", err)
		}
	}()

	fmt.Println("Starting server")

	ch := make(chan error, 1)

	go func() {
		err = server.ListenAndServe() //blocking
		if err != nil {
			ch <- fmt.Errorf("failed to start server: %w", err)
		}
		close(ch)
	}()

	//	err, open := <-ch //This actually blocks our code's execution until
	//1) it recieves a value or
	//2) the channel is closed in which case the value will be nil

	// Also there is another channel used by Context
	// ctx.Done - reurns a channel that's closed when work done on behalf of this context should be canceled.
	// Done may return nil if this context can never be cancled. successive calls return the asme value
	// The secon par of our context is - ctx.Done(). This method return a channel inside, which is how
	// it was signalled if the channel was cancelled.
	// We want use it in conjuction with error to determine if we wil close()

	// because we are now listening to Two blocking Channels, we will use "select" switch statement.
	//The first case that has it's value to be read, will have it's case to be resolved and the code will be
	// able to contiue.

	// btw this select statment is why we are using a buffer for our error channel. in the event that
	// this channel is called first, then we won't read from this channel again and
	// therefore, we don't want for our servers go routine to be deadlocked
	select {
	case err = <-ch:
		return err
	case <-ctx.Done():
		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}

	//handling of ctx.Done - To call our server's shutdown method. we need to pass another context
	// we could pass the same context we have alrady been using, However, this would prohibit any graceful
	// termination as the context has already been canceled as this point so there'll be no time for
	//any requests in flight to be resolved.so we need to create a brand new context using 'context.background' function.
	// this function alone can't do the job, our job could run indefinatly. So will use context.timeout function
	// we could give 10 seconds for any flight request to resovle.
	fmt.Print("This won't get printed of course, unless something happens")
	return nil
}
