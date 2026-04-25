package core_logger

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Level  string `envconfig:"LEVEL" required:"true"`
	Folder string `envconfig:"FOLDER" required:"true"`
}

func NewConfig() (Config, error) {
	var config Config

	//if err := godotenv.Load("../../.env"); err != nil { // Для дебага
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	if err := envconfig.Process("LOGGER", &config); err != nil {
		return Config{}, fmt.Errorf("Process envconfig: %w", err)
	}

	return config, nil
}

func NewConfigMust() Config {
	config, err := NewConfig()
	if err != nil {
		err = fmt.Errorf("Get Logger config: %w", err)
		panic(err)
	}
	return config
}
