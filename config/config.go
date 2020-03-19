package config

import (
	"github.com/go-ini/ini"
	"log"
)

type Configuration struct {
	PageSize  int
	JwtSecret string

	// Image
	ImageSavePath string
	ImageMaxSize  int
	ImageAllowExt []string

	// Database
	DbType     string
	DbHost     string
	DbPort     string
	DbUser     string
	DbName     string
	DbPassword string
}

var Config = &Configuration{}
var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("config/config.ini")
	if err != nil {
		log.Fatalf("config.Setup, fail to parse 'config/config.ini': %v", err)
	}

	mapTo("config", Config)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("config.mapTo %s err: %v", section, err)
	}
}
