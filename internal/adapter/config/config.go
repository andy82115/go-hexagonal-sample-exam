package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type (
	Container struct {
		App   *App
		Token *Token
		DB    *DB
		HTTP  *HTTP
		AWS   *AWS
	}

	App struct {
		Name string
		Env  string
	}

	Token struct {
		Duration string
		IsSaveSecretAtAws bool
	}

	DB struct {
		Connection string
		Host       string
		Port       string
		User       string
		Password   string
		Name       string
	}

	HTTP struct {
		Env            string
		URL            string
		Port           string
		AllowedOrigins string
	}

	AWS struct {
		AwasUrl string
	}
)

func New() (*Container, error) {
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			return nil, err
		}
	}

	app := &App{
		Name: os.Getenv("APP_NAME"),
		Env:  os.Getenv("APP_ENV"),
	}

	isSaveSecretAtAws, _ := strconv.ParseBool(os.Getenv("TOKEN_SAVE_SECRET_AT_AWS"))
	token := &Token{
		Duration: os.Getenv("TOKEN_DURATION"),
		IsSaveSecretAtAws: isSaveSecretAtAws,
	}

	db := &DB{
		Connection: os.Getenv("DB_CONNECTION"),
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		User:       os.Getenv("DB_USER"),
		Password:   os.Getenv("DB_PASSWORD"),
		Name:       os.Getenv("DB_NAME"),
	}

	http := &HTTP{
		Env:            os.Getenv("APP_ENV"),
		URL:            os.Getenv("HTTP_URL"),
		Port:           os.Getenv("HTTP_PORT"),
		AllowedOrigins: os.Getenv("HTTP_ALLOWED_ORIGINS"),
	}

	aws := &AWS{
		AwasUrl: os.Getenv("AWS_URL"),
	}

	return &Container{
		app,
		token,
		db,
		http,
		aws,
	}, nil
}
