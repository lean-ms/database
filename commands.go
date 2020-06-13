package database

import (
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/lean-ms/database/config"
)

type CreateTableOptions struct {
	Temp        bool
	IfNotExists bool
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
	dbOptions := config.LoadDatabaseConfig(configFilepath)
	db := pg.Connect(dbOptions)
	defer db.Close()
	options := &orm.CreateTableOptions{}
	if opts != nil {
		options.Temp = opts.Temp
		options.IfNotExists = opts.IfNotExists
	}
	db.CreateTable(model, options)
}

func DropTable(configFilepath string, model interface{}, opts *DropTableOptions) {
	dbOptions := config.LoadDatabaseConfig(configFilepath)
	db := pg.Connect(dbOptions)
	defer db.Close()
	options := &orm.DropTableOptions{}
	if opts != nil {
		options.Cascade = opts.Cascade
		options.IfExists = opts.IfExists
	}
	db.DropTable(model, options)
}

func getConnectionWithoutDatabase(configFilepath string) (string, *pg.DB) {
	dbOptions := config.LoadDatabaseConfig(configFilepath)
	name := dbOptions.Database
	dbOptions.Database = ""
	db := pg.Connect(dbOptions)
	return name, db
}
