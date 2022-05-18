package db

import (
	"github.com/zmicro-team/zmicro/core/config"
	"github.com/zmicro-team/zmicro/core/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	DriverName string
	DataSource string
	MaxIdle    int
	MaxOpen    int
}

func Setup() {
	c := Config{}

	if err := config.Scan("mysql", &c); err != nil {
		log.Fatal(err)
	}

	// TODO: gorm logger
	db, err := gorm.Open(mysql.Open(c.DataSource), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		//return err
	}
	if c.MaxIdle > 0 {
		sqlDB.SetMaxIdleConns(c.MaxIdle)
	}

	if c.MaxOpen > 0 {
		sqlDB.SetMaxOpenConns(c.MaxOpen)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatal(err)
	}

	//runtime.SetDB(db)
}

func Open(c *Config) (db *gorm.DB, err error) {
	db, err = gorm.Open(mysql.Open(c.DataSource), &gorm.Config{})
	if err != nil {
		return
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	if c.MaxIdle > 0 {
		sqlDB.SetMaxIdleConns(c.MaxIdle)
	}

	if c.MaxOpen > 0 {
		sqlDB.SetMaxOpenConns(c.MaxOpen)
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	return
}
