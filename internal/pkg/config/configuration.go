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
	Env    string `yaml:"env"`
	Server struct {
		Port string `yaml:"port" json:"port"`
	} `yaml:"server" json:"server"`

	DBCfg DBConfig `yaml:"db"`
}

// DBConfig define the database connection infomation
type DBConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Address  string `yaml:"address"`
	DBName   string `yaml:"DBName"`
	Env      string `yaml:"env"`
}

// NewConfiguration create and return a Configuration object
func NewConfiguration(fileName string) *Configuration {
	flag.Parse()
	cfg := Configuration{}

	if len(fileName) == 0 {
		fileName = "config.yml"
	}
	rootDirPath := os.Getenv("YAMCHA_CONFIG")
	if rootDirPath == "" {
		rootDirPath = "./configs"
	}
	configPath := filepath.Join(rootDirPath, fileName)
	_, err := os.Stat(configPath)
	if err != nil {
		log.Fatalf("[CONFIG] file error: %s", err.Error())
	}

	// config exists
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
func (conf *Configuration) BasicSetting() {
	// set default env
	if len(conf.Env) == 0 {
		conf.Env = "release"
	}

	// set default port
	if len(conf.Server.Port) == 0 {
		conf.Server.Port = ":18180"
	}
	// for heroku
	herokuPort := os.Getenv("PORT")
	if len(herokuPort) != 0 {
		conf.Server.Port = ":" + herokuPort
	}
}
