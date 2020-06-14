package database

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/lean-ms/database/config"
	"github.com/lean-ms/database/models"
)

// CreateDatabase exposes method to create database
// based on yaml database config file
func CreateDatabase(configFilepath string) error {
	databaseName, db := getConnectionWithoutDatabase(configFilepath)
	defer db.Close()
	_, err := db.Exec("CREATE DATABASE " + databaseName)
	return err
}

func DropDatabase(configFilepath string) error {
	databaseName, db := getConnectionWithoutDatabase(configFilepath)
	defer db.Close()
	_, err := db.Exec("DROP DATABASE " + databaseName)
	return err
}

func CreateTable(configFilepath string, model interface{}, opts *models.CreateTableOptions) error {
	db := getConnectionWithDefaultOptions(configFilepath)
	defer db.Close()
	return db.CreateTable(model, (*orm.CreateTableOptions)(opts))
}

func DropTable(configFilepath string, model interface{}, opts *models.DropTableOptions) error {
	db := getConnectionWithDefaultOptions(configFilepath)
	defer db.Close()
	return db.DropTable(model, (*orm.DropTableOptions)(opts))
}

func getConnectionWithDefaultOptions(configFilepath string) *pg.DB {
	dbOptions := config.LoadDatabaseConfig(configFilepath)
	return pg.Connect(dbOptions)
}

func getConnectionWithoutDatabase(configFilepath string) (string, *pg.DB) {
	dbOptions := config.LoadDatabaseConfig(configFilepath)
	name := dbOptions.Database
	dbOptions.Database = ""
	db := pg.Connect(dbOptions)
	return name, db
}

// CreateConnection opens connection with a given yml in configFilepath dir
func CreateConnection(configFilepath string) *models.DbConnection {
	dbOptions := config.LoadDatabaseConfig(configFilepath)
	db := pg.Connect(dbOptions)
	return &models.DbConnection{Database: db}
}
