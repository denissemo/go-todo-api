package main

import (
    "io"
    "log"
    "net/http"
    "os"

    "github.com/denissemo/go-todo-api/app/api/routes"
    "github.com/denissemo/go-todo-api/app/middleware"
    "github.com/denissemo/go-todo-api/app/models"
    "github.com/denissemo/go-todo-api/app/utils"
    "github.com/gorilla/mux"
)

func init() {
    utils.LoadEnv()
}

func main() {
    var router *mux.Router
    router = mux.NewRouter()
    router = router.PathPrefix("/api").Subrouter()
    router.Use(middleware.RequestLogger)
    router.Use(middleware.JwtAuthentication)

    routes.AuthRoutes(router)
    routes.UserRoutes(router)

    // For check server living
    router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
        _, _ = io.WriteString(w, "PONG")
    })

    port := os.Getenv("PORT")
    if port == "" {
        // Set default port
        port = "3000"
    }

    models.AutoMigrate() // Sync database schema and migrate.

    log.Printf("INFO: Server started on http://localhost:%s", port)

    if err := http.ListenAndServe(":" + port, router); err != nil {
        log.Fatal(err)
    }
}
