package http

import (
	"log"
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type Config struct {
	Port              int           `mapstructure:"port"`
	AllowOrigins      string        `mapstructure:"allow_origins"`
	MaxConcurrentConn int           `mapstructure:"max_concurrent_conn"`
	ReadTimeout       time.Duration `mapstructure:"read_timeout"`
	WriteTimeout      time.Duration `mapstructure:"write_timeout"`
}

func (c Config) Verify() {
	err := validation.ValidateStruct(&c,
		validation.Field(&c.Port, validation.Required),
		validation.Field(&c.ReadTimeout, validation.Required),
		validation.Field(&c.WriteTimeout, validation.Required),
		validation.Field(&c.MaxConcurrentConn, validation.Required))
	if err != nil {
		log.Fatalf("invalid http config %v", err)
	}
}
