package app

type Config struct {
	Port string `toml:"port"`
}

func newConfig() *Config {
	return &Config{
		Port: ":8000",
	}
}
