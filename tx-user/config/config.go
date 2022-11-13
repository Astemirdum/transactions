package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"gopkg.in/yaml.v3"
)

type Server struct {
	Host    string        `yaml:"host"`
	Port    string        `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type BalanceClient struct {
	Host    string        `yaml:"host"`
	Port    string        `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"user"`
	Password string `yaml:"password"`
	NameDB   string `yaml:"dbname"`
}

type Redis struct {
	Host       string        `yaml:"host"`
	Port       string        `yaml:"port"`
	Password   string        `yaml:"password"`
	SessionTTL time.Duration `yaml:"sessionTTL"` //nolint:tagliatelle
}

type Config struct {
	Server        Server        `yaml:"server"`
	Database      DB            `yaml:"db"`
	Redis         Redis         `yaml:"redis"`
	BalanceClient BalanceClient `yaml:"balanceClient"`
}

var (
	once sync.Once
	cfg  *Config
)

func GetConfigYML(configYML string) *Config {
	once.Do(func() {
		file, err := os.Open(configYML)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		if err = yaml.NewDecoder(file).Decode(&cfg); err != nil {
			log.Fatal(err)
		}
		cfg.Database.Password = os.Getenv("DB_PASSWORD_USER")
		_ = printConfig(cfg) //nolint:errcheck
	})

	return cfg
}

func printConfig(cfg *Config) error {
	jscfg, err := json.MarshalIndent(cfg, "", "	")
	fmt.Println(string(jscfg))
	return err
}
