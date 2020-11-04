package routes

import (
    "github.com/denissemo/go-todo-api/app/api/controllers"
    "github.com/gorilla/mux"
)

func UserRoutes(router *mux.Router) {
    router.HandleFunc("/users/{id:[0-9]+}", controllers.UserById).Methods("GET")
}
