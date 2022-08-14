package main

import (
	"fmt"

	"net/http"

	"./controllers"
	"./views"

	"github.com/gorilla/mux"
)

type Dog struct {
	Name string
}

type User struct {
	Name  string
	Dog   Dog
	Slice []string
}

var (
	// homeTpl    *template.Template
	// aboutusTpl *template.Template
	homeTpl    *views.View
	aboutusTpl *views.View
	signupTpl  *views.View
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	data := User{
		Name:  "John Doe",
		Dog:   Dog{Name: "Blackie"},
		Slice: []string{"a", "b"},
	}
	if err := homeTpl.Template.ExecuteTemplate(w, homeTpl.Layout, data); err != nil {
		panic(err)
	}
}

func aboutUsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := aboutusTpl.Template.ExecuteTemplate(w, aboutusTpl.Layout, nil); err != nil {
		panic(err)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Halaman Tidak Ketemu</h1>")
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	err := signupTpl.Template.ExecuteTemplate(w, signupTpl.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func main() {
	homeTpl = views.NewView("bootstrap", "views/index.html")
	aboutusTpl = views.NewView("bootstrap", "views/aboutus.html")
	signupTpl = views.NewView("bootstrap", "views/newuser.html")

	cUser := controllers.NewUser()

	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/aboutus", aboutUsHandler).Methods("GET")
	r.HandleFunc("/signup", cUser.New).Methods("GET")

	r.HandleFunc("/signup", cUser.Create).Methods("POST")

	http.ListenAndServe("localhost:3000", r)
}
