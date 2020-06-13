package database

import (
	"os"
	"testing"
)

var configFilepath string = "./database.yml"

type Test struct {
	ID     int64
	Coluna string
}

func TestMigration(t *testing.T) {
	setup()
	dbConnection := CreateConnection(configFilepath)
	defer dbConnection.Close()
	if err := dbConnection.Database.Insert(&Test{Coluna: "123"}); err != nil {
		t.Errorf("Could not insert into database: %s", err.Error())
	}
	count, err := dbConnection.Database.Model(&Test{ID: 1}).SelectAndCount()
	if err != nil || count != 1 {
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
