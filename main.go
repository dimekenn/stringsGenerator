package main

import (
	"context"
	"fmt"
	"main/server"
	"os"
	"os/signal"
	"syscall"
)

func main()  {
	ctx := context.Background()
	//httpAddr := flag.String("http", ":8080", "http listen address")


	errChan := make(chan error, 1)

	go func(){
		sigCh := make(chan os.Signal)
		signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
		errChan <- fmt.Errorf("%s", <-sigCh)
	}()

	go server.StartHTTPServer(ctx, errChan)

	fmt.Printf("Terminated: %s", <-errChan)
}

