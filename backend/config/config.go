package config

import "github.com/ganeshdipdumbare/goenv"

type EnvVar struct {
	LogLevel      string `json:"log_level"`
	CodePlayerURL string `json:"code_player_url"`
	Port          string `json:"port"`
}

var envVars = &EnvVar{
	LogLevel:      "info",
	Port:          "8000",
	CodePlayerURL: "https://go.dev",
}

func init() {
	goenv.SyncEnvVar(&envVars)
}

func Get() *EnvVar {
	return envVars
}
