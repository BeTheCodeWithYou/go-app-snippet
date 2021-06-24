package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type Config struct {

	Addr string
	StaticDir string
}

func main() {

	// directly passing value from command line without having config struct.
	// addr := flag.String("addr", ":3000", "http server port")
	
	cfg := new(Config)
	flag.StringVar(&cfg.Addr, "addr", ":4040", "Http Server Port")
	flag.StringVar(&cfg.StaticDir, "staticDir", "./ui/static", "path to static resources")
	flag.Parse()

	// creating loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr,"ERROR\t",log.Ldate|log.Ltime|log.Lshortfile)


	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileserver := http.FileServer(http.Dir(cfg.StaticDir))

	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	// custom http.Server struct. Telling to use http address from cmd line flag
	// use custom error log and use the handler defined above. 
	// Rest all values of http.server struct will be set to default as per Go library.
	srv := &http.Server{
		Addr: cfg.Addr,
		ErrorLog: errorLog,
		Handler: mux,
	}

	infoLog.Printf("starting server on port %s ", cfg.Addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
