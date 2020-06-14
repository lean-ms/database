package database

import (
	"os"
	"testing"
)

type Test struct {
	ID     int64
	Coluna string
}

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

func TestMigration(t *testing.T) {
	setup()
	dbConnection := CreateConnection(configFilepath)
	defer dbConnection.Close()
	if err := dbConnection.Database.Insert(&Test{Coluna: "123"}); err != nil {
		t.Errorf("Could not insert into database: %s", err.Error())
	}
	testModel := &Test{ID: 1}
	err := dbConnection.Database.Select(testModel)
	if err != nil || testModel.Coluna != "123" {
		t.Errorf("Could not setup database. Error was: %s", err.Error())
	}
	tearDown()
}

func setup() {
	os.Setenv("LEANMS_ENV", "test")
	DropDatabase(configFilepath)
	CreateDatabase(configFilepath)
	CreateTable(configFilepath, &Test{}, nil)
}

func tearDown() {
	DropDatabase(configFilepath)
}