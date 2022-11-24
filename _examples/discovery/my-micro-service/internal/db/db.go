package db

import (
	"context"
	"fmt"
	"github.com/god-jay/gools/tools/conf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Conf struct {
	Host     string `yaml:"Host"`
	Port     int    `yaml:"Port"`
	Database string `yaml:"Database"`
	Username string `yaml:"Username"`
	Password string `yaml:"Password"`
}

type Client struct {
	*gorm.DB
}

func NewClient(configFile string) (*Client, error) {
	var config Conf
	conf.MustResolveYaml(configFile, &config)

	db, dbErr := gorm.Open(mysql.Open(config.dsn()))

	if dbErr != nil {
		return nil, dbErr
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(200)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return &Client{db}, nil
}

func (conf *Conf) dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.Database)
}

func (c *Client) User(ctx context.Context) *User {
	return &User{c.WithContext(ctx)}
}
