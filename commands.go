package database

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/lean-ms/database/config"
)

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
