package config

import "os"

type Config struct {
    AuthServiceURL   string
    UploadServiceURL string
    GhibliServiceURL string
    OrderServiceURL  string
    JWTSecret        string
}

func LoadConfigFromEnv() *Config {
    return &Config{
        AuthServiceURL:   os.Getenv("AUTH_SERVICE_URL"),
        UploadServiceURL: os.Getenv("UPLOAD_SERVICE_URL"),
        GhibliServiceURL: os.Getenv("GHIBLI_SERVICE_URL"),
        OrderServiceURL:  os.Getenv("ORDER_SERVICE_URL"),
        JWTSecret:        os.Getenv("JWT_SECRET"), // same as in Auth Service
    }
}
