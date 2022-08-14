package controllers

import (
	"fmt"
	"net/http"

	"../views"
)

func NewUser() *User {
	return &User{
		NewView: views.NewView("bootstrap", "./views/newuser.html"),
	}
}

type User struct {
	NewView *views.View
}

type SignupForm struct {
	Email string `schema:"email"`
	Pswrd string `schema:"pswrd"`
}

func (u *User) New(w http.ResponseWriter, r *http.Request) {
	err := u.NewView.Template.ExecuteTemplate(w, u.NewView.Layout, nil)
	if err != nil {
		panic(err)
	}
}

func (u *User) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}

	fmt.Fprintln(w, form)
	fmt.Fprintln(w, r.PostForm["email"])
	fmt.Fprintln(w, r.PostForm["pswrd"])
}
