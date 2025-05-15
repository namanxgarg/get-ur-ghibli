package config

import "os"

type Config struct {
    DBHost    string
    DBPort    string
    DBUser    string
    DBPass    string
    DBName    string
    JWTSecret string
}

func LoadConfigFromEnv() *Config {
    return &Config{
        DBHost:    os.Getenv("DB_HOST"),
        DBPort:    os.Getenv("DB_PORT"),
        DBUser:    os.Getenv("DB_USER"),
        DBPass:    os.Getenv("DB_PASS"),
        DBName:    os.Getenv("DB_NAME"),
        JWTSecret: os.Getenv("JWT_SECRET"),
    }
}
