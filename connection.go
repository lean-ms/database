package database

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"github.com/go-pg/pg"
	"gopkg.in/yaml.v2"
)

func loadDatabaseOptions(configPath, environment string) map[string]string {
	filepath := path.Join(configPath, "database.yml")
	_, err := os.Stat(filepath)
	if os.IsNotExist(err) {
		log.Fatal("Could not find database config file: " + filepath)
	}
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal("Could not read database config file: " + filepath)
	}
	var databaseOptions map[string]map[string]string
	err = yaml.Unmarshal(data, &databaseOptions)
	if err != nil {
		log.Fatal("Could not deserialize database config file: " + filepath)
	}
	return databaseOptions[environment]
}

func LoadDatabaseConfig(configPath, environment string) *pg.Options {
	envOptions := loadDatabaseOptions(configPath, environment)
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

func (db DbConnection) Close() {
	db.Database.Close()
}

func CreateConnection(configPath string) *DbConnection {
	env := os.Getenv("LEANMS_ENV")
	if len(env) == 0 {
		env = "development"
	}
	dbOptions := LoadDatabaseConfig(configPath, env)
	db := pg.Connect(dbOptions)
	return &DbConnection{Database: db}
}
