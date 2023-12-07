package internal

import (
	"gopkg.in/yaml.v3"
	"os"
)

type Config struct {
	Database struct {
		Host     string `yaml:"dbHost"`
		Port     string `yaml:"dbPort"`
		User     string `yaml:"dbUser"`
		Name     string `yaml:"dbName"`
		Password string `yaml:"dbPassword"`
	} `yaml:"dbSettings"`
	Server struct {
		Host             string `yaml:"bindAddress"`
		Port             int    `yaml:"port"`
		ResourceTimeout  string `yaml:"resourceTimeout"`
		AnalyticsTimeout string `yaml:"analyticsTimeout"`
	} `yaml:"programSettings"`
}

func GetConfig(configPath string) (*Config, error) {
	config := &Config{}
	file, err := os.Open(configPath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	yamlDecoder := yaml.NewDecoder(file)
	if err := yamlDecoder.Decode(&config); err != nil {
		return nil, err
	}

	return config, nil
}
