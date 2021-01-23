package database

import (
	"fmt"
	"time"

	pkgConfig "yamcha/internal/config"

	"github.com/cenk/backoff"
	log "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
		dial := mysql.Open(cfg.ConnectDSN)

		database, err = gorm.Open(dial, &gorm.Config{})
		if err != nil {
			log.Errorf("db: mysql open failed: %v", err)
			return err
		}
		db, err := database.DB()
		if err != nil {
			log.Errorf("db: get db error: %v", err)
			return err
		}
		err = db.Ping()
		if err != nil {
			log.Errorf("db: ping error: %v", err)
			return err
		}
		return nil
	}, bo)

	if err != nil {
		log.Panicf("db: mysql connect err: %s", err.Error())
	}

	database.Debug()
	log.Infof("database ping success")

	return database, nil
}
