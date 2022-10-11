package database

import (
	db "github.com/forbole/juno/v3/database"

	basedatabase "github.com/forbole/bdjuno/v3/chains/base/database"
)

// Builder allows to create a new Db instance implementing the db.Builder type
func Builder(ctx *db.Context) (db.Database, error) {
	return basedatabase.Builder(ctx)
}
