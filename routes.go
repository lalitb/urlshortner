package main

import (
    "net/http"
    "github.com/gorilla/mux"
)

type route struct {
    name    string
    method  string
    pattern string
    HandlerFunc http.HandlerFunc
}

type routes []route

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    for _, route := range routes {
        router.Methods(route.method).
              .Path(route.pattern).
              .Name(route.name).
              .Handler(route.HandlerFunc)
    }
    return router
}

var routes = Routes {
    Route{
        "shorten",
        "POST",
        "/shorten",
        hortenURL
    },
    Route{
        "original",
        "POST",
        "/original",
        originalURL
    },
}
