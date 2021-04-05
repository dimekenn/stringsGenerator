package server

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"main/config"
	"main/handler"
	"net/http"
	"time"
)
//server
func StartHTTPServer(ctx context.Context, errCh chan <- error)  {
	fmt.Println("listening")

	router := mux.NewRouter()

	cfg, err := config.LoadConfiguration("./config.json")
	if err != nil{
		errCh <- err
	}

	handler2 := handler.NewHandler(cfg)

	router.HandleFunc("/generator", handler2.GeneratorHandler).Methods("POST")

	srv := &http.Server{
		Addr: cfg.Port,
		Handler: router,
		WriteTimeout: 15*time.Second,
		ReadTimeout: 15*time.Second,
	}

	err = srv.ListenAndServe()
	if err != nil{
		errCh <- err
	}
}

