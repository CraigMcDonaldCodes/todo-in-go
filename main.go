package main

import (
	"net/http"

	"github.com/craigmcdonaldcodes/todo-in-go/handlers"
	"github.com/craigmcdonaldcodes/todo-in-go/utils"
)

func main() {

	utils.DisplayBanner()

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/add", handlers.Add)
	http.HandleFunc("/delete", handlers.Delete)
	http.HandleFunc("/purge", handlers.Purge)
	http.HandleFunc("/login", handlers.Login)

	http.ListenAndServe("localhost:8080", nil)
}
