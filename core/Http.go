package core

import (
	"errors"
	"net/http"
	"reflect"
	"runtime"
)

/**
 * Http Library
 */

var handles = map[interface{}]interface{}{}

// RegisterHandle creates new handle on specific path.
func RegisterHandle(path string, handler func(http.ResponseWriter, *http.Request), method string) {
	Console("Creating new "+method+" route "+path+" on "+runtime.FuncForPC(reflect.ValueOf(handler).Pointer()).Name()+" handler", "Router")
	handles[path+"_"+method] = RegHandler{
		Path:   path,
		Method: method,
		Function: func(w http.ResponseWriter, r *http.Request) {
			handler(w, r)
		},
	}
}

// RegisterStandaloneHandle Register standalone handle.
func RegisterStandaloneHandle(path string, handler func(http.ResponseWriter, *http.Request)) {
	http.HandleFunc(path, func(writer http.ResponseWriter, request *http.Request) {
		Console(request.Method+" request from "+request.RemoteAddr+" to "+request.RequestURI, "Request")
		handler(writer, request)
	})
}

// RegisterStaticHandle creates new static server on path
func RegisterStaticHandle(path string, localPath string) {
	if FileExists(localPath) {
		Console("Registered static server on "+path+" for local "+localPath+" path", "Router")
		fs := http.FileServer(http.Dir("./" + localPath))
		http.Handle(path, http.StripPrefix(path, fs))
	} else {
		CheckForFatal(errors.New("Path " + localPath + " does not exists."))
	}
}

func router() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		Console(request.Method+" request from "+request.RemoteAddr+" to "+request.RequestURI, "Request")
		clientReq := handles["/"+request.URL.Path[1:]+"_"+request.Method]
		if clientReq != nil {
			if clientReq.(RegHandler).Method == request.Method {
				clientReq.(RegHandler).Function(writer, request)
			} else {
				http.Error(writer, "Invalid request method.", 405)
			}
		} else {
			http.Error(writer, "Path not found.", 404)
		}
	})
}

func Serve(port string) {
	Console("Starting serving routes...", "Router")
	router()
	Console("Serving application in "+port, "Info")
	server := http.ListenAndServe(":"+port, nil)
	CheckForFatal(server)
}
