package config

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-pg/pg"
	"gopkg.in/yaml.v2"
)

func LoadDatabaseConfig(configFilepath string) *pg.Options {
	environment := getEnvironment()
	envOptions := loadDatabaseOptions(configFilepath, environment)
	address := fmt.Sprintf("%s:%s", envOptions["host"], envOptions["port"])
	return &pg.Options{
		Addr:     address,
		User:     envOptions["user"],
		Password: envOptions["password"],
		Database: envOptions["database"],
	}
}

func getEnvironment() string {
	env := os.Getenv("LEANMS_ENV")
	if len(env) > 0 {
		return env
	}
	return "development"
}

func loadDatabaseOptions(configFile, environment string) map[string]string {
	expandedData := os.ExpandEnv(readConfigFile(configFile))
	var databaseOptions map[string]map[string]string
	if err := yaml.Unmarshal([]byte(expandedData), &databaseOptions); err != nil {
		log.Fatal("Could not deserialize database config file: " + configFile)
	}
	return databaseOptions[environment]
}

func readConfigFile(configFile string) string {
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		log.Fatal("Could not find database config file: " + configFile)
	}
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal("Could not read database config file: " + configFile)
	}
	return string(data)
}
