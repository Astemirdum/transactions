package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
	"gopkg.in/yaml.v3"
)

type Server struct {
	Host        string        `yaml:"host"`
	Port        string        `yaml:"port"`
	ReadTimeout time.Duration `yaml:"readTimeout"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yaml:"user"`
	Password string `yaml:"password"`
	NameDB   string `yaml:"dbname"`
}

type JS struct {
	URL            string        `yaml:"url"`
	MaxReconnects  int           `yaml:"maxReconnects"`
	ReconnectWait  time.Duration `yaml:"reconnectWait"`
	ConnectTimeout time.Duration `yaml:"connectTimeout"`
	WorkerCount    int           `yaml:"workerCount"`
}

type AuthClient struct {
	Host    string        `yaml:"host"`
	Port    string        `yaml:"port"`
	Timeout time.Duration `yaml:"timeout"`
}

type Config struct {
	Server     Server     `yaml:"server"`
	Database   DB         `yaml:"db"`
	JS         JS         `yaml:"js"`
	AuthClient AuthClient `yaml:"authClient"`
	InitCash   int
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
		cfg = new(Config)
		cfg.JS = newJSConf(
			withConnectTimeout(20*time.Second),
			withMaxReconnect(nats.DefaultMaxReconnect))
		if err = yaml.NewDecoder(file).Decode(&cfg); err != nil {
			log.Fatal(err)
		}
		cfg.Database.Password = os.Getenv("DB_PASSWORD_BALANCE")
		cfg.InitCash = 1000
		envCash := os.Getenv("INIT_CASH")
		if envCash != "" {
			if d, err := strconv.Atoi(envCash); err == nil {
				cfg.InitCash = d
			}
		}
		_ = printConfig(cfg) //nolint:errcheck
	})

	return cfg
}

func printConfig(cfg *Config) error {
	jscfg, err := json.MarshalIndent(cfg, "", "	")
	fmt.Println(string(jscfg))
	return err
}
