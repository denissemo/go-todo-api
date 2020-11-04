package middleware

import (
    "log"
    "net/http"
)

var RequestLogger = func(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        method := r.Method
        uri := r.URL.String()
        log.Printf("<-- [%s] %s\n", method, uri)

        next.ServeHTTP(w, r)
    })
}
