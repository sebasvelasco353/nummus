package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	App    AppConfig
	DB     DBConfig
	Server ServerConfig
}

type AppConfig struct {
	Env string // "development" || "production".
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

func (db DBConfig) DBAddress() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		db.User,
		db.Password,
		db.Host,
		db.Port,
		db.Name,
		db.SSLMode,
	)
}

type ServerConfig struct {
	Port           string
	JWTSecret      string
	JWTExpiryHours string
}

func Load() (*Config, error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	if env == "development" {
		if err := godotenv.Load("../.env"); err != nil {
			fmt.Printf("warning: could not load ../.env: %v\n", err)
		}
	}

	cfg := &Config{
		App: AppConfig{
			Env: env,
		},
		DB: DBConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     os.Getenv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
			SSLMode:  os.Getenv("DB_SSLMODE"),
		},
		Server: ServerConfig{
			Port:           os.Getenv("SERVER_PORT"),
			JWTSecret:      os.Getenv("SERVER_JWT_SECRET"),
			JWTExpiryHours: os.Getenv("SERVER_JWT_EXPIRY_HOURS"),
		},
	}

	// Apply defaults for optional fields with sensible values.
	if cfg.DB.Port == "" {
		cfg.DB.Port = "5432"
	}
	if cfg.DB.SSLMode == "" {
		cfg.DB.SSLMode = "disable"
	}
	if cfg.Server.Port == "" {
		cfg.Server.Port = "8080"
	}
	if cfg.Server.JWTExpiryHours == "" {
		cfg.Server.JWTExpiryHours = "0.25"
	}

	if err := cfg.validate(); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (c *Config) validate() error {
	required := []struct {
		value string
		name  string
	}{
		{c.DB.Host, "DB_HOST"},
		{c.DB.User, "DB_USER"},
		{c.DB.Password, "DB_PASSWORD"},
		{c.DB.Name, "DB_NAME"},
		{c.Server.JWTSecret, "SERVER_JWT_SECRET"},
	}

	var missing []string
	for _, r := range required {
		if strings.TrimSpace(r.value) == "" {
			missing = append(missing, r.name)
		}
	}

	if len(missing) > 0 {
		return fmt.Errorf(
			"missing required environment variables: %s\n"+
				"copy .env.example to .env and fill in the values",
			strings.Join(missing, ", "),
		)
	}

	return nil
}

func (c *Config) IsDevelopment() bool {
	return c.App.Env == "development"
}

