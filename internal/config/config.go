package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"os"
	"time"
)

type Config struct {
	Env         string `yaml:"env" env-default:"local" env-required:"true"`
	StoragePath string `yaml:"storage_path" storage-required:"true"`
	HTTPServer  `yaml:"http_server" https-required:"true"`
}

type HTTPServer struct {
	Address     string        `yaml:"address" env-default:"localhost:8080"`
	Timeout     time.Duration `yaml:"timeout" env-default:"4s"`
	IdleTimeout time.Duration `yaml:"idle_timeout" env-default:"61s"`
}

func MustLoad() *Config {
	// TODO: добавить переменную окружения
	//configPath := os.Getenv("CONFIG_PATH_US")
	configPath, err := os.ReadFile("./config/local.yaml")
	if err != nil {
		log.Fatal("CONFIG_PATH environment variable is required")
	}

	if _, err := os.Stat(string(configPath)); os.IsNotExist(err) {
		log.Fatalf("Configuration file %s does not exist", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(string(configPath), &cfg); err != nil {
		log.Fatalf("Error reading configuration: %v", err)
	}

	return &cfg
}
