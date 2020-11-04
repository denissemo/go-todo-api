package controllers

import (
    "encoding/json"
    "net/http"

    "github.com/denissemo/go-todo-api/app/middleware"
    "github.com/denissemo/go-todo-api/app/models"
    "github.com/denissemo/go-todo-api/app/utils"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

func Login(w http.ResponseWriter, r *http.Request) {
    type loginBody struct {
        Login    string `json:"login"`
        Password string `json:"password"`
    }

    body := &loginBody{}
    if err := json.NewDecoder(r.Body).Decode(body); err != nil {
        utils.Respond(w, r, utils.ErrorMessage{Code: 400, Message: "InvalidBody"})
        return
    }

    user := &models.User{}
    if err := models.GetDB().Table("users").Where(
        "email = ? OR username = ?", body.Login, body.Login,
    ).First(user).Error; err != nil {
        if err == gorm.ErrRecordNotFound {
            utils.Respond(w, r, utils.ErrorMessage{Code: 403, Message: "InvalidCredentials"})
            return
        }
        utils.Respond(w, r, utils.ErrorMessage{Code: 409, Message: "UnknownError"})
        return
    }

    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
    if err == bcrypt.ErrMismatchedHashAndPassword {
        utils.Respond(w, r, utils.ErrorMessage{Code: 403, Message: "InvalidCredentials"})
        return
    }

    tokenType := middleware.AuthToken{
        UserId: user.ID,
        Email:  user.Email,
    }
    token := tokenType.Sign()

    user.Password = ""
    response := make(map[string]interface{})
    response["user"] = user
    response["token"] = "JWT " + token

    utils.Respond(w, r, response)
    return
}

func SignUp(w http.ResponseWriter, r *http.Request) {
    user := &models.User{}

    if err := json.NewDecoder(r.Body).Decode(user); err != nil {
        utils.Respond(w, r, utils.ErrorMessage{Code: 400, Message: "InvalidBody"})
        return
    }

    if err, ok := user.Validate(); !ok {
        utils.Respond(w, r, err)
        return
    }

    user.Create()
    user.Password = "" // Don`t send password in response.

    utils.Respond(w, r, user, 201)
    return
}
