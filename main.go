package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"example.com/m/db"
	"example.com/m/handlers"
	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	db.InitDatabase()
	l := log.New(os.Stdout, "urlshortener ", log.LstdFlags)
	sm := mux.NewRouter()

	sh := handlers.NewShorturl(l)

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/api/shorturl/{short}", sh.GetShorturl)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/api/shorturl/{original}", sh.PostShorturl)

	s := &http.Server{
		Addr:         ":8080",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Kill)
	signal.Notify(sigChan, os.Interrupt)

	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	tc, c := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
	c()
}
