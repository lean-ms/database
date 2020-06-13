package database

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"

	"github.com/go-pg/pg/v10"
	"gopkg.in/yaml.v2"
)

func loadDatabaseOptions(configPath, environment string) map[string]string {
	filepath := path.Join(configPath, "database.yml")
	data, _ := ioutil.ReadFile(filepath)
	var databaseOptions map[string]map[string]string
	yaml.Unmarshal(data, &databaseOptions)
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
