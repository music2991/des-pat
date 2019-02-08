package database_factory

import (
	"database/sql"
	"dp/factory/factory1/config"
	"dp/factory/factory1/constants"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlAdapter struct{}

func (MysqlAdapter) GetConnection() (*sql.DB, error) {
	// DNS = username:password@protocol(address)/dbname?param=value
	format := "%s:%s@tcp(%s)/%s"
	server, err := getMysqlConfig()
	if err != nil {
		return nil, err
	}
	dataSourceName := fmt.Sprintf(format, server.User, server.Password, server.Host, server.DbName)

	db, err := sql.Open(constants.Mysql, dataSourceName)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func getMysqlConfig() (serverConfig, error) {
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
