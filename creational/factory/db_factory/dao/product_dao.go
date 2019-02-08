package dao

import (
	"dp/creational/factorynal/factory/factory1/factory"
	"dp/creational/factorynal/factory/factory1/factory/database_factory"
	"dp/creational/factorynal/factory/factory1/model"
	"dp/creational/factorynal/factory/factory1/util"
)

func NewProductDAO(dbType string) (productDAO, error) {
	dbAdapter, err := factory.GetDBAdapter(dbType)
	if err != nil {
		return productDAO{}, err
	}

	return productDAO{
		iDBAdapter: dbAdapter,
	}, nil
}

type productDAO struct {
	iDBAdapter database_factory.IDBAdapter
}

func (dao *productDAO) GetAllProducts() (productList []model.Product, err error) {
	defer func() { err = util.GetFullErr(err) }()

	db, err := dao.iDBAdapter.GetConnection()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	query := `Select id, name, defaultvalue from Product`
	stmt, err := db.Prepare(query)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}

	var productItem model.Product
	for rows.Next() {
		err = rows.Scan(&productItem.Id, &productItem.Name, &productItem.DefaultValue)
		if err != nil {
			return nil, err
		}
		productList = append(productList, productItem)
	}

	return productList, nil
}
