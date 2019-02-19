package dao

import (
	"database/sql"
	"dp/creational/factory/db_factory/factory/database_factory"
	"dp/creational/factory/db_factory/model"
	"dp/creational/factory/db_factory/util"
	"log"
	"time"
)

func NewProductDAO(dbAdapter database_factory.IDBAdapter, db *sql.DB) (productDAO, error) {

	return productDAO{
		iDBAdapter: dbAdapter,
		db:         db,
	}, nil
}

type productDAO struct {
	iDBAdapter database_factory.IDBAdapter
	db         *sql.DB
}

func (dao *productDAO) GetAllProducts() (productList []model.Product, err error) {
	defer func() { err = util.GetFullErr(err) }()
	db := dao.db

	query := `Select categoryid, name from "SwitchUpp_Core".Category`
	var stmt *sql.Stmt
	maxIntents := 10
	currentIntents := 0

	for currentIntents < maxIntents {
		stmt, err = db.Prepare(query)
		if err != nil {
			time.Sleep(time.Second)
			currentIntents++
			log.Print("*************************:", currentIntents)
			continue
		}
		currentIntents = maxIntents
	}
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var productItem model.Product
	for rows.Next() {
		err = rows.Scan(&productItem.Id, &productItem.Name)
		if err != nil {
			return nil, err
		}
		productList = append(productList, productItem)
	}

	return productList, nil
}
