module github.com/denissemo/go-todo-api

// +heroku goVersion go1.14
go 1.14

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/gorilla/mux v1.8.0
	github.com/joho/godotenv v1.3.0
	golang.org/x/crypto v0.0.0-20200622213623-75b288015ac9
	gorm.io/driver/postgres v1.0.5
	gorm.io/gorm v1.20.5
)
