package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	m "github.com/geokat/crud/model"
)

const (
	usersPath = "/users/"
)

func usersHandler(w http.ResponseWriter, r *http.Request) {
	var resp []byte
	var err error

	code := http.StatusOK
	id := strings.TrimPrefix(r.URL.Path, usersPath)

	// We should be using gorilla/mux or similar here.
	switch r.Method {

	case "GET":
		if id == "" {
			resp, err = m.GetUsers()
		} else {
			resp, err = m.GetUser(id)
		}
		code = http.StatusOK

	case "POST", "PUT":
		email, name := r.FormValue("email"), r.FormValue("name")
		if email == "" || name == "" {
			handleError(&w, errors.New("Need user email and name"), http.StatusBadRequest)
			return
		}

		if r.Method == "POST" {
			err = m.CreateUser(email, name)
			code = http.StatusCreated
		} else {
			if id == "" {
				handleError(&w, errors.New("No user ID provided"), http.StatusBadRequest)
				return
			}
			err = m.UpdateUser(id, email, name)
			code = http.StatusNoContent
		}

	case "DELETE":
		if id == "" {
			handleError(&w, errors.New("No user ID provided"), http.StatusBadRequest)
			return
		}
		err = m.DeleteUser(id)
		code = http.StatusNoContent
	}

	if err != nil {
		// It's an unknown error; return a 500 response
		handleError(&w, err, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(resp)
}

func main() {
	http.HandleFunc(usersPath, usersHandler)

	listen := "0.0.0.0:8080"
	fmt.Println("Starting CRUD API demo server on " + listen)
	log.Fatal(http.ListenAndServe(listen, nil))
}

//
// Helpers
//

func handleError(w *http.ResponseWriter, err error, code int) {
	log.Println(err)
	http.Error(*w, err.Error(), code)
}

/*
 * Local variables:
 * compile-command: "go build github.com/geokat/crud":
 * End:
 */
