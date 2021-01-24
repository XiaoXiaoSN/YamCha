package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/jinzhu/configor"
	log "github.com/sirupsen/logrus"
)

var globalConfig = &Configuration{}

// Configuration is config
type Configuration struct {
	Env    string        `yaml:"env" env:"ENVIRONMENT" default:"release"`
	Server Server        `yaml:"server"`
	DBCfg  DBConfig      `yaml:"db"`
	BotCfg LineBotConfig `yaml:"line_bot"`
	Sentry SentryConfig  `yaml:"sentry"`
}

// Server ...
type Server struct {
	// heroku will bind env PORT as the exporting service
	Port int `yaml:"port" env:"PORT" default:"18080"`
}

// DBConfig define the database connection infomation
type DBConfig struct {
	ConnectDSN string `yaml:"dsn" env:"MYSQL_DSN"`

	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Address  string `yaml:"address"`
	DBName   string `yaml:"DBName"`
	Env      string `yaml:"env"`
}

// LineBotConfig ...
type LineBotConfig struct {
	ChannelSecret string `yaml:"channel_secret" env:"LINECORP_PLATFORM_CHANNEL_CHANNELSECRET"`
	ChannelToken  string `yaml:"channel_token" env:"LINECORP_PLATFORM_CHANNEL_CHANNELTOKEN"`
}

// NewConfiguration create and return a Configuration object
func NewConfiguration() *Configuration {
	cfg := Configuration{}

	var fileName, rootDirPath string
	if fileName = os.Getenv("YAMCHA_CONFIG_NAME"); fileName == "" {
		fileName = "config.yml"
	}
	if rootDirPath = os.Getenv("YAMCHA_CONFIG"); rootDirPath == "" {
		rootDirPath = "./configs"
	}
	configPath := filepath.Join(rootDirPath, fileName)
	_, err := os.Stat(configPath)
	if err != nil {
		log.Fatalf("[CONFIG] file error: %s", err.Error())
	}

	// Enable debug mode or set env `CONFIGOR_DEBUG_MODE` to true when running your application
	err = configor.New(&configor.Config{Debug: false}).Load(&cfg, configPath)
	if err != nil {
		log.Fatalf("[CONFIG] configor error: %s", err.Error())
	}

	// register sentry setting
	cfg.Sentry.init()

	// fmt.Printf("%+v\n\n", prettyPrint(cfg))
	globalConfig = &cfg
	return &cfg
}

// Config get the global config
func Config() Configuration {
	return *globalConfig
}

func prettyPrint(data interface{}) string {
	jsonByte, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("%s\n", jsonByte)
}
