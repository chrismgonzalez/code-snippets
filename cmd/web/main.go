package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

//application struct to hold app-wide dependencies

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	//command line flag for localhost addr with :4000 as the default
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	//logs
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	//http.Server struct
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
