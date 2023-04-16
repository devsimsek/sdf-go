package handlers

import (
	"SDF/core"
	"fmt"
	"net/http"
)

func init() {
	core.RegisterHandle("p/{all}", func (w http.ResponseWriter, r *http.Request)  {
    w.Write([]byte("You are viewing " + core.ParseUrl(r)[2] + " page."))
	}, "GET")
	core.RegisterHandle("home", homeHandler, "GET")
	core.RegisterHandle("", homeHandler, "GET")
}

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, core.LoadView("views/home.html", core.PageData{
		PageTitle: "Home",
	}))
	core.CheckError(err)
}
