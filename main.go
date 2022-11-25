package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"net/http/httputil"
)

// properties with datatypes and how they are present json object
type Author struct {
	ID        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// slice authors of type Author
// slice like array but length doesn't have to be specified in this
var authors []Author

// to get all authors
// * because r can hold value of http.Request only
// content-type is set to application/json
// encode and return json object
func getAuthors(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}

// to get authors by id
// the url contains id, params will hold the id as params[id]
// range returns index,element
func getAuthorById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	reqDump, _ := httputil.DumpRequestOut(r, true)
	fmt.Println(reqDump)
	fmt.Println(params["id"])
	var found bool
	found = false
	for _, item := range authors {
		fmt.Println("Individual", item.ID)
		fmt.Println("Parameter", params)
		if item.ID == params["id"] {
			found = true
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	if !found {
		json.NewEncoder(w).Encode(&Author{
			ID:        "id not found",
			FirstName: "",
			LastName:  "",
		})
	}
}

//to insert author

func createAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var author Author
	_ = json.NewDecoder(r.Body).Decode(&author)
	authors = append(authors, author)
	json.NewEncoder(w).Encode(&author)
}

// to update author
// creates new row with new details and deletes old row
func updateAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range authors {
		if item.ID == params["id"] {
			authors = append(authors[:index], authors[index+1:]...)
			var author Author
			_ = json.NewDecoder(r.Body).Decode(&author)
			author.ID = params["id"]
			authors = append(authors, author)
			json.NewEncoder(w).Encode(&author)
			return
		}
	}
	json.NewEncoder(w).Encode(authors)
}

// to delete an entry
func deleteAuthor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range authors {
		if item.ID == params["id"] {
			authors = append(authors[:index], authors[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(authors)
}
func main() {
	router := mux.NewRouter()

	authors = append(authors, Author{ID: "1", FirstName: "Arthur", LastName: "Kipling"})
	authors = append(authors, Author{ID: "2", FirstName: "Ruskin", LastName: "Bond"})

	router.HandleFunc("/authors", getAuthors).Methods("GET")
	router.HandleFunc("/authors", createAuthor).Methods("POST")
	router.HandleFunc("/authors/{id}", getAuthorById).Methods("GET")
	router.HandleFunc("/authors/{id}", updateAuthor).Methods("PUT")
	router.HandleFunc("/authors/{id}", deleteAuthor).Methods("DELETE")

	http.ListenAndServe(":3000", router)
}
