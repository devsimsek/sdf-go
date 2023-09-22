package core

import (
	"errors"
	"net/http"
	"reflect"
	"regexp"
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

func router() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		Console(r.Method+" request from "+r.RemoteAddr+" to "+r.RequestURI, "Router")
		pathFound := false
		for _, v := range handles {
			match := strings.Replace(v.(RegHandler).Path, "{url}", "([0-9a-zA-Z]+)", -1)
			match = strings.Replace(v.(RegHandler).Path, "{id}", "([0-9]+)", -1)
			match = strings.Replace(v.(RegHandler).Path, "{all}", "(.*)", -1)
			p := regexp.MustCompile(match)
			matches := p.FindAllString("/"+r.URL.Path[1:], -1)
			possibleMatch := handles[r.URL.Path+"_"+r.Method]
			if possibleMatch.(RegHandler).Path != "" {
				possibleMatch.(RegHandler).Function(w, r)
				pathFound = true
				break
			}
			if (len(matches) > 0) && strings.Join(matches, "") != "/" && strings.Join(matches, "") != "" {
				if r.Method == v.(RegHandler).Method {
					v.(RegHandler).Function(w, r)
					pathFound = true
					break
				}
			} else {
				if r.URL.Path[1:] == v.(RegHandler).Path {
					if r.Method == v.(RegHandler).Method {
						v.(RegHandler).Function(w, r)
						pathFound = true
						break
					}
				}
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
