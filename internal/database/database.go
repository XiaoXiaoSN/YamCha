package database

import (
	"fmt"
	"time"

	pkgConfig "yamcha/internal/config"

	"github.com/cenk/backoff"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// NewDatabases init and return write and read DB objects
func NewDatabases(cfg pkgConfig.DBConfig) (*gorm.DB, error) {
	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = time.Duration(180) * time.Second

	if cfg.ConnectDSN == "" {
		cfg.ConnectDSN = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", cfg.Username, cfg.Password, cfg.Address, cfg.DBName)
	}

	var database *gorm.DB
	var err error
	err = backoff.Retry(func() error {
		database, err = gorm.Open("mysql", cfg.ConnectDSN)
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
