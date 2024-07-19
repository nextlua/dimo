package config

import (
	"fmt"
	"github.com/codingconcepts/env"
	"github.com/joho/godotenv"
	"log"
	"os"
)

type EnvVars struct {
	Service Service
}

type Service struct {
	Environment string `env:"SERVICE_ENVIRONMENT" default:"dev"`
	Address     string `env:"HTTP_SERVER_ADDRESS" default:":8001"`
	ClientID    string `env:"CLIENT_ID" required:"true"`
	RedirectURL string `env:"REDIRECT_URL" required:"true"`
}

func LoadEnvVars() (*EnvVars, error) {
	err := godotenv.Load(".env")
	if err != nil && !os.IsNotExist(err) {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}

	s := Service{}
	if err = env.Set(&s); err != nil {
		return nil, fmt.Errorf("loading service environment variables failed, %s", err.Error())
	}

	ev := &EnvVars{
		Service: s,
	}

	return ev, nil
}
