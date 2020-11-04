package middleware

import (
    "context"
    "net/http"
    "os"
    "strings"

    "github.com/denissemo/go-todo-api/app/models"
    "github.com/denissemo/go-todo-api/app/utils"
    "github.com/dgrijalva/jwt-go"
)

type AuthToken struct {
    jwt.StandardClaims
    UserId uint
    Email  string
}

func (tk *AuthToken) Sign() string {
    token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
    signed, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
    return signed
}

var JwtAuthentication = func(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        notAuthUrls := []string{"/api/auth/login", "/api/auth/sign-up", "/api/ping"}
        path := r.URL.Path

        // Ignore some urls.
        for _, url := range notAuthUrls {
            if url == path {
                next.ServeHTTP(w, r)
                return
            }
        }

        headerToken := strings.Split(r.Header.Get("Authorization"), " ")
        if len(headerToken) != 2 {
            utils.Respond(w, r, utils.ErrorMessage{Code: 403, Message: "MissedAuthToken"})
            return
        }

        tokenPart := headerToken[1]
        tk := &AuthToken{}

        token, err := jwt.ParseWithClaims(tokenPart, tk, func(token *jwt.Token) (interface{}, error) {
            return []byte(os.Getenv("JWT_SECRET")), nil
        })

        if err != nil {
            utils.Respond(w, r, utils.ErrorMessage{Code: 403, Message: "InvalidAuthToken"})
            return
        }

        if !token.Valid {
            utils.Respond(w, r, utils.ErrorMessage{Code: 401, Message: "Unauthorized"})
            return
        }

        user := &models.User{}
        if err := models.GetDB().Table("users").Where("id = ?", tk.UserId).First(user).Error; err != nil {
            utils.Respond(w, r, utils.ErrorMessage{Code: 401, Message: "Unauthorized"})
            return
        }

        ctx := context.WithValue(r.Context(), "user", user)
        r = r.WithContext(ctx)
        next.ServeHTTP(w, r)
    })
}
