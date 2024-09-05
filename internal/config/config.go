package config

import (
	"flag"
	"os"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	GRPCServer `yaml:"grpc" env-required:"true"`
	Database   `yaml:"database" env-required:"true"`
}
type GRPCServer struct {
	Address     string        `yaml:"address" env:"ADDRESS" env-default:":4000"`
	Timeout     time.Duration `yaml:"timeout" env:"TIMEOUT" env-default:"4s" `
	IdleTimeout time.Duration `yaml:"idle_timeout" env:"IDLE_TIMEOUT" env-default:"60s"`
}

type Database struct {
	Port     string `yaml:"port" env:"PORT" env-default:"5432"`
	Host     string `yaml:"host" env:"HOST" env-default:"localhost"`
	Name     string `yaml:"name" env:"NAME" env-default:"postgres"`
	DBName   string `yaml:"db_name" env:"DB_NAME" env-default:"users"`
	User     string `yaml:"user" env:"USER" env-default:"user"`
	Password string `yaml:"password" env:"PASSWORD" env-required:"true"`
}

func MustLoad() *Config {
	configPath := fetchConfigFlag()
	if configPath == "" {
		panic("config path to empty")
	}
	return MustLoadByPath(configPath)
}

func MustLoadByPath(configPath string) *Config {
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		panic("config file does not exist: " + configPath)
	}
	var cfg Config
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		panic("failed to read config: " + err.Error())
	}
	return &cfg
}

func fetchConfigFlag() string {
	var res string
	flag.StringVar(&res, "config", "config/development.yaml", "path to config file")
	flag.Parse()
	if res == "" {
		res = os.Getenv("CONFIG_PATH")
	}
	return res
}
