package controllers

import (
    "net/http"

    "github.com/denissemo/go-todo-api/app/models"
    "github.com/denissemo/go-todo-api/app/utils"
    "github.com/gorilla/mux"
)

func UserById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    user := &models.User{}
    if err := models.GetDB().Table("users").Where("id = ?", id).First(user).Error; err != nil {
        utils.Respond(w, r, utils.ErrorMessage{Code: 404, Message: "UserNotFound"})
        return
    }

    user.Password = ""
    utils.Respond(w, r, user)
    return
}
