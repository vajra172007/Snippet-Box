package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello from SnippetBox"))
}
func snippetView(w http.ResponseWriter, r *http.Request){
	id , err := strconv.Atoi(r.PathValue("id"))
	if err !=nil || id < 1{
		http.NotFound(w, r)
		return
	}
	msg := fmt.Sprintf("Display a specific Snippet with ID %d...",id)
	w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func main(){
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}",home)
	mux.HandleFunc("/snippet/view/{id}",snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	err := http.ListenAndServe(":400",mux)
	log.Fatal(err)
}