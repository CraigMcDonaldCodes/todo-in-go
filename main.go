package main

import (
	"fmt"
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
	mux.HandleFunc("/logout", handlers.Logout)

	staticFiles := http.FileServer(http.Dir("./pub"))
	mux.Handle("/pub/", http.StripPrefix("/pub", staticFiles))

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println("Failed to start web server:", err)
	}
}
