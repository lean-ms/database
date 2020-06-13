package database

import (
	"os"
	"testing"

	"github.com/go-pg/pg/orm"
)

var configFilepath string = "./database.yml"

type Test struct {
	ID     int64
	Coluna string
}

func TestMigration(t *testing.T) {
	os.Setenv("LEANMS_ENV", "test")
	CreateDatabase(configFilepath)
	dbConnection := CreateConnection(configFilepath)
	defer dbConnection.Close()
	dbConnection.Database.CreateTable(&Test{}, &orm.CreateTableOptions{
		Temp: true,
	})
	dbConnection.Database.Insert(&Test{Coluna: "123"})
	count, _ := dbConnection.Database.Model(&Test{ID: 1}).SelectAndCount()
	if count != 1 {
		t.Error("Could not setup database")
	}
	DropDatabase(configFilepath)
}
