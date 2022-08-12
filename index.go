package main

import (
	"fmt"
	// "html/template"

	// "html/template"
	"./views"

	"net/http"

	"github.com/gorilla/mux"
)

// func handlerFunc(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html")
// 	if r.URL.Path == "/" {
// 		fmt.Fprint(w, "<h1>Selamat datang di skillplus</h1>")
// 	} else if r.URL.Path == "/aboutus" {
// 		fmt.Fprint(w, "<h1>ini adalah web tutorial seputar informasi teknologi</h1>")
// 	} else {
// 		fmt.Fprint(w, "<h1>Halaman yang dicari tidak ditemukan</h1>")
// 		w.WriteHeader(http.StatusNotFound)
// 	}
// }
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
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	data := User{
		Name:  "John Doe",
		Dog:   Dog{Name: "Blackie"},
		Slice: []string{"a", "b"},
	}
	if err := homeTpl.Template.Execute(w, data); err != nil {
		panic(err)
	}
}

func aboutUsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// fmt.Fprint(w, "<h1>Ini Adalah testing</h1>")
	if err := aboutusTpl.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>Halaman Tidak Ketemu</h1>")
}

func main() {
	// http.HandleFunc("/", handlerFunc)
	// http.HandleFunc("/aboutus", handlerFunc)
	// http.ListenAndServe(":3000", nil)
	// var err error
	// homeTpl, err = template.ParseFiles("views/index.html", "views/layouts/footer.html")
	// if err != nil {
	// 	panic(err)
	// }

	homeTpl = views.NewView("views/index.html")
	aboutusTpl = views.NewView("views/aboutus.html")
	// aboutusTpl, err = template.ParseFiles("views/aboutus.html", "views/layouts/footer.html")
	// if err != nil {
	// 	panic(err)
	// }
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	r.HandleFunc("/", homeHandler)
	r.HandleFunc("/aboutus", aboutUsHandler)
	http.ListenAndServe("localhost:3000", r)
}
