package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
)

func parseForm(r *http.Request, target interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}

	dec := schema.NewDecoder()

	if err := dec.Decode(target, r.PostForm); err != nil {
		return err
	}
	return nil
}
