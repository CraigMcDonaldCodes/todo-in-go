package main

import (
	"net/http"

	"github.com/craigmcdonaldcodes/todo-in-go/handlers"
	"github.com/craigmcdonaldcodes/todo-in-go/utils"
)

func main() {

	utils.DisplayBanner()

	mux := http.NewServeMux()

	mux.HandleFunc("/", handlers.Home)
	mux.HandleFunc("/add", handlers.Add)
	mux.HandleFunc("/delete", handlers.Delete)
	mux.HandleFunc("/purge", handlers.Purge)
	mux.HandleFunc("/login", handlers.Login)

	staticFiles := http.FileServer(http.Dir("./pub"))
	mux.Handle("/pub/", http.StripPrefix("/pub", staticFiles))

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
