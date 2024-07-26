package config

import "os"

const (
	EnvServerHost = "SERVER_HOST"
	EnvServerPort = "SERVER_PORT"
)

type Config struct {
	Server ServerConfig `json:"server"`
}

type ServerConfig struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

func GetConfig() *Config {
	cfg := &Config{
		Server: ServerConfig{
			Host: "localhost",
			Port: "8080",
		},
	}

	host := os.Getenv(EnvServerHost)
	if host != "" {
		cfg.Server.Host = host
	}

	port := os.Getenv(EnvServerPort)
	if port != "" {
		cfg.Server.Port = port
	}

	return cfg
}
