package database

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/lean-ms/database/config"
)

type DbConnection struct {
	Database *pg.DB
}

// Close closes connection
func (db DbConnection) Close() error {
	return db.Database.Close()
}

type CreateTableOptions struct {
	Temp          bool
	IfNotExists   bool
	Varchar       int
	FKConstraints bool
}

type DropTableOptions struct {
	IfExists bool
	Cascade  bool
}

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

func CreateTable(configFilepath string, model interface{}, opts *CreateTableOptions) error {
	db := getConnectionWithDefaultOptions(configFilepath)
	defer db.Close()
	return db.CreateTable(model, (*orm.CreateTableOptions)(opts))
}

func DropTable(configFilepath string, model interface{}, opts *DropTableOptions) error {
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
func CreateConnection(configFilepath string) *DbConnection {
	dbOptions := config.LoadDatabaseConfig(configFilepath)
	db := pg.Connect(dbOptions)
	return &DbConnection{Database: db}
}
