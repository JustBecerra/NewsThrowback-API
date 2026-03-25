package config

import "os"

type Config struct {
	DatabaseURL             string
	ChroniclingAmericaURL   string
	Port                    string
}

func Load() Config {
	return Config{
		DatabaseURL:           os.Getenv("DATABASE_URL"),
		ChroniclingAmericaURL: os.Getenv("CHRONICLING_AMERICA_API_URL"),
		Port:                  os.Getenv("PORT"),
	}
}
