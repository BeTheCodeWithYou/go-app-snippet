package main

import (
	"flag"
	"log"
	"net/http"
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

	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	fileserver := http.FileServer(http.Dir(cfg.StaticDir))

	mux.Handle("/static/", http.StripPrefix("/static", fileserver))

	log.Printf("starting server on port %s ", cfg.Addr)
	err := http.ListenAndServe(cfg.Addr, mux)
	log.Fatal(err)
}
