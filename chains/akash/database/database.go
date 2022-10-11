package database

import (
	"fmt"

	db "github.com/forbole/juno/v3/database"

	basedatabase "github.com/forbole/bdjuno/v3/chains/base/database"
)

type Db struct {
	*basedatabase.Db
}

// Builder allows to create a new Db instance implementing the db.Builder type
func Builder(ctx *db.Context) (db.Database, error) {
	db, err := basedatabase.Builder(ctx)
	if err != nil {
		return nil, err
	}

	baseDb, ok := db.(*basedatabase.Db)
	if !ok {
		return nil, fmt.Errorf("invalid database type: %T", db)
	}

	return &Db{
		baseDb,
	}, nil
}

// Cast allows to cast the given db to a Db instance
func Cast(db db.Database) *Db {
	bdDatabase, ok := db.(*Db)
	if !ok {
		panic(fmt.Errorf("given database instance is not a Db"))
	}
	return bdDatabase
}
