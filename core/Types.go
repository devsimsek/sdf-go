package core

import (
	"encoding/gob"
	"net/http"
)

type PageData struct {
	PageTitle string
	PageBody  map[string]interface{}
	ViewData  map[string]interface{}
}

type UserData struct {
	Username      string
	AuthLevel     interface{}
	Authenticated bool
}

type RegHandler struct {
	Path     string
	Method   string
	Function func(w http.ResponseWriter, r *http.Request)
}

func init() {
	gob.Register(PageData{})
	gob.Register(RegHandler{})
}
