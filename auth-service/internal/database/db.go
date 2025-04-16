package database

import (
    "fmt"
    "github.com/jinzhu/gorm"
    _ "github.com/lib/pq"

    "github.com/example/get-ur-ghibli/auth-service/internal/config"
    "github.com/example/get-ur-ghibli/auth-service/internal/models"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPass, cfg.DBName)
    return gorm.Open("postgres", dsn)
}

func AutoMigrate(db *gorm.DB) {
    db.AutoMigrate(&models.User{})
}
