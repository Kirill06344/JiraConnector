package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"time"
)

type Config struct {
	DB      DBSettings      `yaml:"DBSettings"`
	Program ProgramSettings `yaml:"ProgramSettings"`
}

type DBSettings struct {
	Host     string `yaml:"dbHost"`
	Port     string `yaml:"dbPort"`
	User     string `yaml:"dbUser"`
	Name     string `yaml:"dbName"`
	Password string `yaml:"dbPassword"`
}

type ProgramSettings struct {
	Host              string        `yaml:"bindAddress"`
	Port              uint          `yaml:"port"`
	JiraUrl           string        `yaml:"jiraUrl"`
	ThreadCount       uint          `yaml:"threadCount"`
	IssueInOneRequest uint          `yaml:"issueInOneRequest"`
	MaxTimeSleep      time.Duration `yaml:"maxTimeSleep"`
	MinTimeSleep      time.Duration `yaml:"minTimeSleep"`
}

func Load() (*Config, error) {
	config := &Config{}
	file, err := os.Open("./config/config.yaml")
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
