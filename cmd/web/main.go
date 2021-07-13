package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"neerajsidhaye.com/snippetbox/config"
)

type application struct {

	ErrorLog *log.Logger
	InfoLog *log.Logger
}

func main() {

	// directly passing value from command line without having config struct.
	// addr := flag.String("addr", ":3000", "http server port")
	
	cmdline := new(config.Cmdline)
	flag.StringVar(&cmdline.Addr, "addr", ":4040", "Http Server Port")
	flag.StringVar(&cmdline.StaticDir, "staticDir", "./ui/static", "path to static resources")
	flag.Parse()

	// creating loggers
	f, err := os.OpenFile("./tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
	if err!=nil {
		log.Fatal(err)
	}
	defer f.Close()

	infoLog := log.New(f, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr,"ERROR\t",log.Ldate|log.Ltime|log.Lshortfile)

	//initialize application struct
	app := &application{
		ErrorLog: errorLog,
		InfoLog: infoLog,
	}

	// custom http.Server struct. Telling to use http address from cmd line flag
	// use custom error log and use the handler defined above. 
	// Rest all values of http.server struct will be set to default as per Go library.
	srv := &http.Server{
		Addr: cmdline.Addr,
		ErrorLog: errorLog,
		Handler: app.routes(),
	}

	infoLog.Printf("starting server on port %s ", cmdline.Addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
