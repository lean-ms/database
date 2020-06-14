package database

import (
	"os"
	"testing"

	"github.com/lean-ms/database/models"
)

var configFilepath string = "./database.yml"

func TestDatabaseSetup(t *testing.T) {
	model := &models.TestModel{Coluna: "123"}
	CreateDatabase(configFilepath)
	CreateTable(configFilepath, model, &models.CreateTableOptions{IfNotExists: true})
	dbConnection := CreateConnection(configFilepath)
	dbConnection.Database.Insert(model)
	defer dbConnection.Close()
	DropTable(configFilepath, model, &models.DropTableOptions{IfExists: true})
	DropDatabase(configFilepath)
}

func TestMigration(t *testing.T) {
	setup()
	dbConnection := CreateConnection(configFilepath)
	defer dbConnection.Close()
	if err := dbConnection.Database.Insert(&models.TestModel{Coluna: "123"}); err != nil {
		t.Errorf("Could not insert into database: %s", err.Error())
	}
	testModel := &models.TestModel{ID: 1}
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
	CreateTable(configFilepath, &models.TestModel{}, nil)
}

func tearDown() {
	DropDatabase(configFilepath)
}
