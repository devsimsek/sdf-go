package handlers

import (
	"SDF/core"
	"fmt"
	"net/http"
)

func init() {
	core.RegisterHandle("home", homeHandler, "GET")
	core.RegisterHandle("", homeHandler, "GET")
}

func homeHandler(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, core.LoadView("views/home.html", core.PageData{
		PageTitle: "Home",
	}))
	core.CheckError(err)
}
