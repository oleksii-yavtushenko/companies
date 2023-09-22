package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	} `yaml:"database"`

	Server struct {
		Port string `yaml:"port"`
	} `yaml:"server"`
}

func LoadConfig() Config {
	f, err := os.Open("/app/configs/config.yaml")
	if err != nil {
		log.Fatal(fmt.Sprintf("error loading config, %v", err))
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal(fmt.Sprintf("error loading config, %v", err))
	}
	return cfg
}
