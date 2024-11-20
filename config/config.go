package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	AppName   string `yaml:"app_name"`
	Port      string `yaml:"port"`
	JWTSecret string `yaml:"jwt_secret"`
	Database  struct {
		Host     string `yaml:"host"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	} `yaml:"database"`
}

func LoadConfig(filepath string) Config {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatalf("Could not open config file: %v", err)
	}
	defer file.Close()

	var cfg Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		log.Fatalf("Could not decode config file: %v", err)
	}
	return cfg
}
