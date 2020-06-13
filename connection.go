package database

import (
	"github.com/go-pg/pg"
	"github.com/lean-ms/database/config"
)

type DbConnection struct {
	Database *pg.DB
}

// Close closes connection
func (db DbConnection) Close() error {
	return db.Database.Close()
}

// CreateConnection opens connection with a given yml in configFilepath dir
func CreateConnection(configFilepath string) *DbConnection {
	dbOptions := config.LoadDatabaseConfig(configFilepath)
	db := pg.Connect(dbOptions)
	return &DbConnection{Database: db}
}
