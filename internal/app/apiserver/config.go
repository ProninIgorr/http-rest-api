package apiserver

import "github.com/ProninIgorr/http-rest-api/internal/app/store"

// Config
type Config struct {
	BindAdrr string `toml:"bind_addr"`
	LogLevel string `toml:"log_level"`
	Store    *store.Config
}

// NewConfig
func NewConfig() *Config {
	return &Config{
		BindAdrr: "8080",
		LogLevel: "debug",
		Store:    store.NewConfig(),
	}
}
