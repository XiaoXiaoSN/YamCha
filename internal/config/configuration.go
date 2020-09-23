package config

import (
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// Configuration is config
type Configuration struct {
	Env    string        `yaml:"env"`
	Server Server        `yaml:"server"`
	DBCfg  DBConfig      `yaml:"db"`
	BotCfg LineBotConfig `yaml:"line_bot"`
}

// Server ...
type Server struct {
	Port string `yaml:"port"`
}

// DBConfig define the database connection infomation
type DBConfig struct {
	ConnectDSN string `yaml:"dsn" description:"priority use the column"`

	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Address  string `yaml:"address"`
	DBName   string `yaml:"DBName"`
	Env      string `yaml:"env"`
}

// LineBotConfig ...
type LineBotConfig struct {
	ChannelSecret string `yaml:"channel_secret"`
	ChannelToken  string `yaml:"channel_token"`
}

// NewConfiguration create and return a Configuration object
func NewConfiguration() *Configuration {
	log.Warn("!!!!!!! start to init config !!!!!!!")

	flag.Parse()
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

	// check config exists
	file, err := ioutil.ReadFile(filepath.Clean(configPath))
	if err != nil {
		log.Fatalf("[CONFIG] read file error: %s", err.Error())
	}

	err = yaml.Unmarshal(file, &cfg)
	if err != nil {
		log.Fatalf("[CONFIG] yaml unmarshal error: %v", err)
	}

	cfg.BasicSetting()
	return &cfg
}

// BasicSetting set default value
func (cfg *Configuration) BasicSetting() {
	// set default env
	if len(cfg.Env) == 0 {
		cfg.Env = "release"
	}

	// set default port
	if len(cfg.Server.Port) == 0 {
		cfg.Server.Port = ":18180"
	}

	// for heroku
	herokuPort := os.Getenv("PORT")
	if len(herokuPort) != 0 {
		cfg.Server.Port = ":" + herokuPort
	}

	// for db connect
	if dsn := os.Getenv("MYSQL_DSN"); dsn != "" {
		cfg.DBCfg.ConnectDSN = dsn
	}

	// for line bot
	lineChannelSecret := os.Getenv("LINECORP_PLATFORM_CHANNEL_CHANNELSECRET")
	lineChannelToken := os.Getenv("LINECORP_PLATFORM_CHANNEL_CHANNELTOKEN")
	cfg.BotCfg = LineBotConfig{
		ChannelSecret: lineChannelSecret,
		ChannelToken:  lineChannelToken,
	}
}
