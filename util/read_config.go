package util

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type SMTP struct {
	Server   string `toml:"server"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	From     string `toml:"from"`
	Port     int    `toml:"port"`
}

type Database struct {
	Host           string `toml:"host"`
	Port           int    `toml:"port"`
	User           string `toml:"user"`
	Password       string `toml:"password"`
	DBname         string `toml:"dbname"`
	MaxConnections int    `toml:"max_connections"`
}

type Server struct {
	Port int `toml:"port"`
}

type JWT struct {
	Secret string `toml:"secret"`
	Expire int    `toml:"expire"`
}

type MQ struct {
	URL string `toml:"url"`
}

type MyConfig struct {
	SMTP     SMTP     `toml:"smtp"`
	Database Database `toml:"database"`
	JWT      JWT      `toml:"jwt"`
	Server   Server   `toml:"server"`
	MQ       MQ       `toml:"mq"`
}

func initConfig() *MyConfig {
	config, err := os.ReadFile("config.toml")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	cfg = new(MyConfig)
	//反序列化
	err = toml.Unmarshal(config, cfg)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return cfg
}
