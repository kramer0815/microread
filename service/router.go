package service

import (
        "github.com/gorilla/mux"
)

// NewRouter Function to build a new Router
func NewRouter() *mux.Router {

        router := mux.NewRouter().StrictSlash(true)
        for _, route := range routes {

                router.Methods(route.Method).
                        Path(route.Pattern).
                        Name(route.Name).
                        Handler(route.HandlerFunc)

        }
        return router
}