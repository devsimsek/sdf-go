package core

import (
	"errors"
	"net/http"
	"reflect"
	"runtime"
	"strings"
)

/**
 * Http Library
 * Copyright (C)devsimsek
 * This file contains functions for routing purposes
 * and answering client requests.
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

func matchPath(path, routePath string) bool {
	parts := strings.Split(routePath, "/")
	pathParts := strings.Split(path, "/")
	if path == routePath {
		return true
	}

	if len(parts) != len(pathParts) {
		return false
	}

	for i, part := range parts {
		if part != pathParts[i] && !strings.HasPrefix(part, "{") {
			return false
		}
	}

	return true
}

func router() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Console(r.Method+" request from "+r.RemoteAddr+" to "+r.RequestURI, "Router")
		path := r.URL.Path[1:]
		method := r.Method
		pathFound := false

		for _, v := range handles {
			handler, ok := v.(RegHandler)
			if !ok {
				continue
			}

			if matchPath(path, handler.Path) && (method == handler.Method || handler.Method == "ANY") {
				handler.Function(w, r)
				pathFound = true
				break
			}
		}

		if !pathFound {
			http.Error(w, "Path not found.", 404)
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
