package config

import (
    "log"
    "os"
    "github.com/joho/godotenv"
    "strconv"
)

type APIFinnhubConfig struct {
    FinnhubAPI string
}

type APIMarketstackConfig struct {
    MarketstackAPI string
    MarketstackURL string
}

type RedisConfig struct {
    Host     string
    Port     int
    DB       int
    Password string
    Username string
}

// LoadEnv loads environment variables from .env file (if present)
func LoadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Println("Warning: .env file not found, using environment variables")
    }
}

// LoadFinnhubConfig loads Finnhub API config from environment
func LoadFinnhubConfig() *APIFinnhubConfig {
    LoadEnv()
    return &APIFinnhubConfig{
        FinnhubAPI: os.Getenv("FINNHUB_API"),
    }
}

// LoadMarketstackConfig loads Marketstack API config from environment
func LoadMarketstackConfig() *APIMarketstackConfig {
    LoadEnv()
    return &APIMarketstackConfig{
        MarketstackAPI: os.Getenv("MARKETSTACK_API"),
        MarketstackURL: os.Getenv("MARKETSTACK_URL"),
    }
}

func LoadRedisConfig() *RedisConfig {
    LoadEnv()
    port, _ := strconv.Atoi(os.Getenv("REDIS_PORT"))
    db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))
    return &RedisConfig{
        Host:     os.Getenv("REDIS_HOST"),
        Port:     port,
        DB:       db,
        Password: os.Getenv("REDIS_PASSWORD"),
        Username: os.Getenv("REDIS_USERNAME"),
    }
}