package factory

import (
	"dp/factory/factory1/constants"
	"dp/factory/factory1/factory/database_factory"
	"errors"
)

func GetDBAdapter(dbType string) (database_factory.IDBAdapter, error) {
	switch dbType {
	case constants.Postgresql:
		return database_factory.PostgresAdapter{}, nil
	case constants.Mysql:
		return database_factory.MysqlAdapter{}, nil
	default:
		return nil, errors.New("data base type not recognized")
	}
}
