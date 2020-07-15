package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/ekozlova94/internal/handler"
	"github.com/ekozlova94/internal/middleware"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "access-time.db")
	if err != nil {
		log.Fatalf("database problem: %s", err.Error())
	}
	if err = db.Ping(); err != nil {
		log.Fatalf("no ping to database: %s", err.Error())
	}
	//noinspection GoUnhandledErrorResult
	defer db.Close()

	mux := http.NewServeMux()

	mux.Handle("/set-access-time", http.HandlerFunc(handler.SetTime))
	mux.Handle("/get-first-access-time", http.HandlerFunc(handler.GetFirstTime))
	mux.Handle("/check-access-time", http.HandlerFunc(handler.GetLastTime))
	err = http.ListenAndServe("localhost:8000", middleware.Db(db, middleware.Ip(middleware.Save(mux))))
	if err != nil {
		log.Fatalf("can not start: %s", err.Error())
	}
}
