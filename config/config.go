package config

import (
	"encoding/json"
	"io/fs"
	"os"
)

type Config struct {
	Port          int    `json:"port"`
	InterfacePath string `json:"interface-path"`
}

func Load() *Config {
	res := Config{
		Port: 16160,
		InterfacePath: "/",
	}
	data, err := os.ReadFile("config.json")
	if err != nil {
		return &res
	}
	if err := json.Unmarshal(data, &res); err != nil {
		return &res
	}
	return &res
}

func (c *Config) Reload() {
	*c = *Load();
}

func (c *Config) Save() {
	data, err := json.Marshal(c);
	if err != nil {
		return
	}
	if err := os.WriteFile("config.json", data, fs.ModePerm); err != nil {
		return
	}
}