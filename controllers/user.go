package controllers

import (
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

func (u *User) New(w http.ResponseWriter, r *http.Request) {
	err := u.NewView.Template.ExecuteTemplate(w, u.NewView.Layout, nil)
	if err != nil {
		panic(err)
	}
}
