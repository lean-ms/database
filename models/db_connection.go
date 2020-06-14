package models

import "github.com/go-pg/pg"

// DbConnection encapsulates database implementation
// TODO: create common interface for model operations (e.g.: select)
type DbConnection struct {
	Database *pg.DB
}

// Close closes connection
func (db DbConnection) Close() error {
	return db.Database.Close()
}
