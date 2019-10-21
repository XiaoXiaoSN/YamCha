package database

import (
	"fmt"
	"time"

	"github.com/cenk/backoff"

	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// Config define db connection needs
type Config struct {
	Username string
	Password string
	Address  string
	DBName   string
	Env      string
}

// NewDatabases init and return write and read DB objects
func NewDatabases(cfg Config) (*gorm.DB, error) {
	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = time.Duration(180) * time.Second
	connectionString := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", cfg.Username, cfg.Password, cfg.Address, cfg.DBName)

	log.Debugf("db: database connection string: %s", connectionString)

	var database *gorm.DB
	var err error
	err = backoff.Retry(func() error {
		database, err = gorm.Open("mysql", connectionString)
		if err != nil {
			log.Errorf("db: mysql open failed: %v", err)
			return err
		}
		err = database.DB().Ping()
		if err != nil {
			log.Errorf("db: mysql ping error: %v", err)
			return err
		}
		return nil
	}, bo)

	if err != nil {
		log.Panicf("db: mysql connect err: %s", err.Error())
	}

	database.LogMode(true)

	log.Infof("database ping success")
	database.DB().SetMaxIdleConns(150)
	database.DB().SetMaxOpenConns(300)
	database.DB().SetConnMaxLifetime(14400 * time.Second)

	if cfg.Env != "prod" && cfg.Env != "production" {
		database = database.LogMode(true)
	}

	return database, nil
}
