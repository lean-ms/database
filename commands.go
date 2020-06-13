package database

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/lean-ms/database/config"
)

func CreateDatabase(configFilepath string) {
	databaseName, db := getConnectionWithoutDatabase(configFilepath)
	defer db.Close()
	db.Exec("CREATE DATABASE " + databaseName)
}

func DropDatabase(configFilepath string) {
	databaseName, db := getConnectionWithoutDatabase(configFilepath)
	defer db.Close()
	db.Exec("DROP DATABASE " + databaseName)
}

func CreateTable(configFilepath string, model interface{}, options *orm.CreateTableOptions) {
	db := getConnectionWithDefaultOptions(configFilepath)
	defer db.Close()
	db.CreateTable(model, options)
}

func DropTable(configFilepath string, model interface{}, options *orm.DropTableOptions) {
	db := getConnectionWithDefaultOptions(configFilepath)
	defer db.Close()
	db.DropTable(model, options)
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
