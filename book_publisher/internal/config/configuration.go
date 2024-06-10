package config

import (
	"encoding/json"
	"os"

	"go.uber.org/zap"
)

type Config struct {
	Port             string      `json:"port"`
	BookPublishTopic string      `json:"kafka.book_publish_topic"`
	logger           *zap.Logger `json:"omit"`
}

func NewConfig(logger *zap.Logger) *Config {
	return &Config{
		logger: logger,
	}
}

func (config *Config) LoadConfig() {
	file := config.getConfigFile()
	defer file.Close()

	newConfig := config.parseConfigFile(file)

	config.setConfigParameters(newConfig)

	config.logger.Info("Configurations loaded successfully")
}

func (config *Config) getConfigFile() (file *os.File) {
	file, err := os.Open("../../config.json")

	if err != nil {
		config.logger.Fatal("Config file could not opened")
	}

	return file
}

func (config *Config) parseConfigFile(file *os.File) *Config {
	var parseConfig *Config
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&parseConfig)

	if err != nil {
		config.logger.Fatal("Could not decode config file")
	}

	return parseConfig
}

func (config *Config) setConfigParameters(parsedConfig *Config) {
	parsedConfig.logger = config.logger
	*config = *parsedConfig
}
