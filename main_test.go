package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

// func TestGetAuthors(t *testing.T) {

// 	authors = append(authors, Author{ID: "1", FirstName: "Arthur", LastName: "Kipling"})
// 	authors = append(authors, Author{ID: "2", FirstName: "Ruskin", LastName: "Bond"})
// 	fmt.Println("THIS IS PRINTING", authors)

// 	req, err := http.NewRequest("GET", "/authors", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(getAuthors)
// 	handler.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}
// 	// Check the response body is what we expect.
// 	expected := `[{"id":"1","firstName":"Arthur","lastName":"Kipling"},{"id":"2","firstName":"Ruskin","lastName":"Bond"}]`
// 	if rr.Body.String() != expected {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			rr.Body.String(), expected)
// 	}
// }

func TestGetAuthorByID(t *testing.T) {

	req, err := http.NewRequest("GET", "/authors/1", nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(req)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(getAuthorById)
	handler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"id":"1","firstName":"Arthur","lastName":"Kipling"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

// func TestCreateAuthor(t *testing.T) {

// 	var jsonStr = []byte(`{"id":"3","firstName":"Sidney","lastName":"Sheldon"}`)

// 	req, err := http.NewRequest("POST", "/authors", bytes.NewBuffer(jsonStr))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(createAuthor)
// 	handler.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}
// 	expected := `{"id":"3","firstName":"Sidney","lastName":"Sheldon"}`
// 	if rr.Body.String() != expected {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			rr.Body.String(), expected)
// 	}
// }

// func TestUpdateAuthor(t *testing.T) {

// 	var jsonStr = []byte(`{"id":"1","firstName":"Rudyard","last_name":"Kipling"}`)

// 	req, err := http.NewRequest("PUT", "/authors", bytes.NewBuffer(jsonStr))
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	req.Header.Set("Content-Type", "application/json")
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(updateAuthor)
// 	handler.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}
// 	expected := `{"id":"1","firstName":"Rudyard","last_name":"Kipling"},{"id":"2","firstName":"Ruskin","lastName":"Bond"},{"id":"3","firstName":"Sidney","lastName":"Sheldon"}`
// 	if rr.Body.String() != expected {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			rr.Body.String(), expected)
// 	}
// }

// func TestDeleteEntry(t *testing.T) {
// 	req, err := http.NewRequest("DELETE", "/authors", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	q := req.URL.Query()
// 	q.Add("id", "2")
// 	req.URL.RawQuery = q.Encode()
// 	rr := httptest.NewRecorder()
// 	handler := http.HandlerFunc(deleteAuthor)
// 	handler.ServeHTTP(rr, req)
// 	if status := rr.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v want %v",
// 			status, http.StatusOK)
// 	}
// 	expected := `{"id":"1","firstName":"Rudyard","last_name":"Kipling"},{"id":"3","firstName":"Sidney","lastName":"Sheldon"}`
// 	if rr.Body.String() != expected {
// 		t.Errorf("handler returned unexpected body: got %v want %v",
// 			rr.Body.String(), expected)
// 	}
// }
