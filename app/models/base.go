package models

import (
    "fmt"
    "log"
    "os"
    "time"

    "github.com/denissemo/go-todo-api/app/utils"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
)

type Model struct {
    ID        uint      `gorm:"primary_key;autoIncrement;unique" json:"id"`
    CreatedAt time.Time `gorm:"not null;type:timestamp with time zone;default:now()" json:"created_at"`
    UpdatedAt time.Time `gorm:"not null;type:timestamp with time zone;default:now()" json:"updated_at"`
}

var db *gorm.DB

func init() {
    utils.LoadEnv()

    user := os.Getenv("DB_USER")
    pass := os.Getenv("DB_PASSWORD")
    dbname := os.Getenv("DB_NAME")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")

    dsn := fmt.Sprintf(
        "user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
        user, pass, dbname, dbHost, dbPort,
    )

    conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Print(err)
    }

    db = conn
}

func GetDB() *gorm.DB {
    return db
}

func AutoMigrate() {
    err := db.Debug().AutoMigrate(&User{})
    if err != nil {
        log.Printf("Migration Error: %s", err)
    }
}
