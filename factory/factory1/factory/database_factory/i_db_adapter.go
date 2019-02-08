package database_factory

import "database/sql"

type IDBAdapter interface {
	GetConnection() (*sql.DB, error)
}
