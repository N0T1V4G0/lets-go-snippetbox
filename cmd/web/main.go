package main

import (
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     ":4000",
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Println("Starting app on port 4000")
	// ListenAndServe always return non-nil err
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
