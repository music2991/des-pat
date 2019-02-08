package database_factory

import (
	"database/sql"
	"dp/creational/factorynal/factory/factory1/config"
	"dp/creational/factorynal/factory/factory1/constants"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
)

type PostgresAdapter struct{}

func (PostgresAdapter) GetConnection() (*sql.DB, error) {
	format := "user=%s password=%s host=%s dbname=%s sslmode=disable"
	server, err := getPostgresConfig()
	if err != nil {
		return nil, err
	}
	dataSourceName := fmt.Sprintf(format, server.User, server.Password, server.Host, server.DbName)

	db, err := sql.Open(constants.Postgresql, dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func getPostgresConfig() (serverConfig, error) {
	switch config.Environment {
	case constants.Production:
		return serverConfig{
			Host:     "host",
			DbName:   "db_name",
			User:     "user",
			Password: "pwd",
		}, nil
	case constants.Development:
		return serverConfig{
			Host:     "host",
			DbName:   "db_name",
			User:     "user",
			Password: "pwd",
		}, nil
	case constants.Local:
		return serverConfig{
			Host:     "host",
			DbName:   "db_name",
			User:     "user",
			Password: "pwd",
		}, nil
	default:
		return serverConfig{}, errors.New("database environment not recognized")
	}
}
