package models

import (
    "regexp"
    "strings"

    "github.com/denissemo/go-todo-api/app/utils"
    "golang.org/x/crypto/bcrypt"
    "gorm.io/gorm"
)

type User struct {
    Model
    Username string `gorm:"not null;type:varchar(255)" json:"username"`
    Email    string `gorm:"not null;type:varchar(255)" json:"email"`
    Password string `gorm:"not null;type:varchar(255)" json:"password"`
    Name     string `gorm:"not null;type:varchar(255)" json:"name"`
}

func (user *User) Validate() (err utils.ErrorMessage, ok bool) {
    if user.Email == "" || user.Username == "" || user.Password == "" || user.Name == "" {
        return utils.ErrorMessage{Code: 400, Message: "MissedRequiredParams"}, false
    }

    // Save email and username to lowercase.
    user.Email = strings.ToLower(user.Email)
    user.Username = strings.ToLower(user.Username)

    re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
    if !re.MatchString(user.Email) {
        return utils.ErrorMessage{Code: 400, Message: "InvalidEmailFormat"}, false
    }

    // Check duplicate email or username.
    existedUser := &User{}
    e := GetDB().Table("users").Where("email = ? OR username = ?", user.Email, user.Username).First(existedUser).Error
    if e != gorm.ErrRecordNotFound {
        if existedUser.Email == user.Email {
            return utils.ErrorMessage{Code: 409, Message: "EmailAlreadyExist"}, false
        } else if existedUser.Username == user.Username {
            return utils.ErrorMessage{Code: 409, Message: "UsernameAlreadyExist"}, false
        } else {
            return utils.ErrorMessage{Code: 409, Message: "UnknownError"}, false
        }
    }

    return utils.ErrorMessage{}, true
}

func (user *User) Create() {
    GetDB().Create(user)
    user.SetPassword(user.Password)
}

func (user *User) SetPassword(password string) {
    passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    user.Password = string(passwordHash)
    GetDB().Table("users").Save(user)
}
