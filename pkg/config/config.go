package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/thomas-armena/scrman/pkg/dir"
)

type Config struct {
	Location  string     `json:"location"`
	Arguments []Argument `json:"arguments"`
}

type Argument struct {
	Description string `json:"description"`
	Default     string `json:"default"`
}

func GetConfig(scriptName string) (*Config, error) {
	scriptDir, err := dir.GetScriptDir(scriptName)
	if err != nil {
		return nil, fmt.Errorf("unable to parse config file: %v", err)
	}
	configFile, err := os.ReadFile(scriptDir + "/config.json")
	if err != nil {
		return nil, fmt.Errorf("unable to parse config file: %v", err)
	}
	var config Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return nil, fmt.Errorf("unable to parse config file: %v", err)
	}
	return &config, nil
}
