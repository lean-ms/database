package database

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/go-pg/pg"
	"gopkg.in/yaml.v2"
)

func loadDatabaseOptions(configFile, environment string) map[string]string {
	_, err := os.Stat(configFile)
	if os.IsNotExist(err) {
		log.Fatal("Could not find database config file: " + configFile)
	}
	data, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Fatal("Could not read database config file: " + configFile)
	}
	var databaseOptions map[string]map[string]string
	err = yaml.Unmarshal(data, &databaseOptions)
	if err != nil {
		log.Fatal("Could not deserialize database config file: " + configFile)
	}
	return databaseOptions[environment]
}

func LoadDatabaseConfig(configFilepath, environment string) *pg.Options {
	envOptions := loadDatabaseOptions(configFilepath, environment)
	address := fmt.Sprintf("%s:%s", envOptions["host"], envOptions["port"])
	return &pg.Options{
		Addr:     address,
		User:     envOptions["user"],
		Password: envOptions["password"],
		Database: envOptions["database"],
	}
}

type DbConnection struct {
	Database *pg.DB
}

func CreateDatabase(configFilepath string) {
	environment := os.Getenv("LEANMS_ENV")
	dbOptions := LoadDatabaseConfig(configFilepath, environment)
	name := dbOptions.Database
	dbOptions.Database = ""
	db := pg.Connect(dbOptions)
	db.Exec("CREATE DATABASE " + name)
}

func DropDatabase(configFilepath string) {
	environment := os.Getenv("LEANMS_ENV")
	dbOptions := LoadDatabaseConfig(configFilepath, environment)
	name := dbOptions.Database
	dbOptions.Database = ""
	db := pg.Connect(dbOptions)
	db.Exec("CREATE DATABASE " + name)
}

// Close closes connection
func (db DbConnection) Close() {
	db.Database.Close()
}

// CreateConnection opens connection with a given yml in configFilepath dir
func CreateConnection(configFilepath string) *DbConnection {
	env := os.Getenv("LEANMS_ENV")
	if len(env) == 0 {
		env = "development"
	}
	dbOptions := LoadDatabaseConfig(configFilepath, env)
	db := pg.Connect(dbOptions)
	return &DbConnection{Database: db}
}
