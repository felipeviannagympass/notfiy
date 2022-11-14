package config

import (
	"encoding/json"
	"os"

	"github.com/kelseyhightower/envconfig"
)

type ServiceConfig struct {
	AccessToken string `json:"access_token"`
	Repos       []Repo `json:"repos"`
}

type Repo struct {
	Name string `json:"name"`
	Repo string `json:"repo"`
}

// LoadServiceConfig ...
func LoadServiceConfig(configFile string) (*ServiceConfig, error) {
	var cfg ServiceConfig

	if err := loadServiceConfigFromFile(configFile, &cfg); err != nil {
		return nil, err
	}

	if err := envconfig.Process("", &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func loadServiceConfigFromFile(configFile string, cfg *ServiceConfig) error {
	_, err := os.Stat(configFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	jsonFile, err := os.ReadFile(configFile)
	if err != nil {
		return err
	}

	return json.Unmarshal(jsonFile, &cfg)
}

func Dump(configName string, bc *ServiceConfig) error {
	fl, err := os.OpenFile(configName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	defer fl.Close()
	if err != nil {
		return err
	}
	data, err := json.MarshalIndent(bc, "", " ")
	if err != nil {
		return err
	}
	_, err = fl.Write(data)
	return err
}
