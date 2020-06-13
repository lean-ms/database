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

func CreateTable(configFilepath string, model interface{}, opts *CreateTableOptions) {
	db := getConnectionWithDefaultOptions(configFilepath)
	defer db.Close()
	db.CreateTable(model, getCreateTableOptions(opts))
}

func DropTable(configFilepath string, model interface{}, opts *DropTableOptions) {
	db := getConnectionWithDefaultOptions(configFilepath)
	defer db.Close()
	db.DropTable(model, getDropTableOptions(opts))
}

func getDropTableOptions(opts *DropTableOptions) *orm.DropTableOptions {
	var options *orm.DropTableOptions
	if opts != nil {
		*options = orm.DropTableOptions(*opts)
	}
	return options
}

func getCreateTableOptions(opts *CreateTableOptions) *orm.CreateTableOptions {
	var options *orm.CreateTableOptions
	if opts != nil {
		*options = orm.CreateTableOptions(*opts)
	}
	return options
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
