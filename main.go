package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/tsuccar/goordersapi/application"
)

func main() {

	app := application.New(application.LoadConfig())
	//creating our own interrupt that will respond to interrupt signal
	// context.Background()- used as root level cascader of interruptions, and to derive
	//a new context from,only in main (), initialization and test.
	//os.Interrupt provides the system SIGINT interrupt

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel() //cancel all derived context and any children underneath it.

	// err := app.Start(ctx)
	err := app.Start(ctx) //blocking
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
	fmt.Println("End of go routine. There was an error, otherwise app would be blocking")

}
