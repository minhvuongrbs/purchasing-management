package database

import (
	"log"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Config struct {
	User                 string        `mapstructure:"user"`
	Passwd               string        `mapstructure:"passwd" json:"-"`
	AllowNativePasswords bool          `mapstructure:"allow_native_passwords"`
	Address              string        `mapstructure:"addr"`
	DatabaseName         string        `mapstructure:"db_name"`
	MaxOpenConn          int           `mapstructure:"max_open_conn"`
	MaxIdleConn          int           `mapstructure:"max_idle_conn"`
	ConnMaxLifeTime      time.Duration `mapstructure:"conn_max_life_time"`
}

func (c Config) Verify() {
	err := validation.ValidateStruct(&c,
		validation.Field(&c.User, validation.Required),
		validation.Field(&c.Address, validation.Required),
		validation.Field(&c.DatabaseName, validation.Required),
		validation.Field(&c.MaxOpenConn, validation.Required),
		validation.Field(&c.MaxIdleConn, validation.Required),
		validation.Field(&c.ConnMaxLifeTime, validation.Required))
	if err != nil {
		log.Fatalf("invalid database config, %v", err)
	}
}

func (c *Config) SetDefault() {}
