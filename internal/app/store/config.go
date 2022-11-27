package store

// Config ...
type Config struct {
	DataBaseURL string `toml:"database_url"`
}

func NewConfig() *Config {
	return &Config{}
}
