// один из самых медленных роутеров
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func List(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "You see user list\n")
}

func Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "you try to see user %s\n", vars["id"])
}

/*
	curl -v -X PUT -H "Content-Type: application/json" -d '{"login":"evg_goncharenko"}' http://localhost:8080/users
*/

func Create(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "you try to create new user\n")
}

/*
	curl -v -X POST -H "Content-Type: application/json"  -H "X-Auth: test" -d '{"name":"Eugene Goncharenko"}' http://localhost:8080/users/evg_goncharenko
*/

func Update(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "you try to update %s\n", vars["login"])
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", List)

	r.HandleFunc("/users", List).
		Host("localhost")

	r.HandleFunc("/users", Update).
		Methods("PUT")
		/*
			можно прописать так:
			r.HandleFunc("/users", Select).
				Methods("GET")
			r.HandleFunc("/users", Insert).
				Methods("PUT")
			r.HandleFunc("/users", Update).
				Methods("POST")
		*/

	r.HandleFunc("/users/{id:[0-9]+}", Get)

	r.HandleFunc("/users/{login}", Create).
		Methods("POST").
		Headers("X-Auth", "test")

	fmt.Println("starting server at :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
