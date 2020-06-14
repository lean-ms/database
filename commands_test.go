package database

import (
	"testing"
)

var configFilepath string = "./database.yml"

func TestDatabaseSetup(t *testing.T) {
	model := &Test{Coluna: "123"}
	CreateDatabase(configFilepath)
	CreateTable(configFilepath, model, &CreateTableOptions{IfNotExists: true})
	dbConnection := CreateConnection(configFilepath)
	dbConnection.Database.Insert(model)
	defer dbConnection.Close()
	DropTable(configFilepath, model, &DropTableOptions{IfExists: true})
	DropDatabase(configFilepath)
}
