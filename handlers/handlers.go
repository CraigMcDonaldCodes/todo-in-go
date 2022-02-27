package handlers

import (
	"net/http"
)

// Home is the default page
func Home(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Home todo page"))
}

// Add is the handler to add a new todo entry
func Add(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Add todo page"))
}

// Delete is the handler to delete a todo entry
func Delete(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Delete todo page"))
}

// Purge is the handler to delete all todo entries
func Purge(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte("Purge todo page"))
}
