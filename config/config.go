package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

type DbConfig struct {
	Host       string
	DbPort     string
	DbUser     string
	DbPassword string
	DbName     string
	Driver     string
}

type ApiConfig struct {
	ApiPort string
}

type TokenConfig struct {
	IssuerName    string
	SigningKey    []byte
	ExpireTime    time.Duration
	SigningMethod *jwt.SigningMethodHMAC
}

type Config struct {
	DbConfig
	ApiConfig
	TokenConfig
}

func (c *Config) Configuration() error {
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("missing env file %v", err.Error())
	}

	c.DbConfig = DbConfig{
		Host:       os.Getenv("DB_HOST"),
		DbPort:     os.Getenv("DB_PORT"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbName:     os.Getenv("DB_NAME"),
		Driver:     os.Getenv("DB_DRIVER"),
	}

	c.ApiConfig = ApiConfig{
		ApiPort: os.Getenv("API_PORT"),
	}

	tokenExpire, _ := strconv.Atoi(os.Getenv("TOKEN_EXPIRE_IN_HOUR"))

	c.TokenConfig = TokenConfig{
		IssuerName:    os.Getenv("TOKEN_ISSUER_NAME"),
		SigningKey:    []byte(os.Getenv("TOKEN_SIGNING_KEY")),
		ExpireTime:    time.Hour * time.Duration(tokenExpire),
		SigningMethod: jwt.SigningMethodHS256,
	}

	if c.Host == "" || c.DbPort == "" || c.DbUser == "" || c.DbPassword == "" || c.DbName == "" || c.Driver == "" || c.ApiPort == "" || c.IssuerName == "" || len(c.SigningKey) == 0 || c.ExpireTime < 0 || c.SigningMethod == nil {
		return fmt.Errorf("missing environment variables")
	}

	return nil
}

func NewConfig() (*Config, error) {
	config := &Config{}

	if err := config.Configuration(); err != nil {
		return nil, err
	}

	return config, nil
}
