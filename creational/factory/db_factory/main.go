package main

import (
	"database/sql"
	"dp/creational/factory/db_factory/constants"
	"dp/creational/factory/db_factory/dao"
	"dp/creational/factory/db_factory/factory"
	"dp/creational/factory/db_factory/model"
	"encoding/json"
	"fmt"
	"log"
	"sync"
)

var db *sql.DB

func main() {
	//executeOne()
	go executeInGoroutines()
	executeInGoroutines()
	log.Print("===================================================")
	log.Print("|				END OF PROCESS					  |")
	log.Print("===================================================")
}

func executeOne() {
	var productList []model.Product

	dbAdapter, err := factory.GetDBAdapter(constants.Postgresql)
	if err != nil {
		panic(err)
	}

	db, err := dbAdapter.GetConnection()
	if err != nil {
		panic(err)
	}

	productDAO, err := dao.NewProductDAO(dbAdapter, db)
	if err != nil {
		return
	}

	productList, err = productDAO.GetAllProducts()
	if err != nil {
		fmt.Println(err, ":B")
		return
	}

	jsonProductList, err := json.Marshal(productList)
	if err != nil {
		fmt.Println(err, ":C")
		return
	}

	fmt.Println(string(jsonProductList))
	fmt.Println("Process terminated")
}
func executeInGoroutines() {
	var cicle int
	var totalCicles = 500
	var waitGroup sync.WaitGroup

	dbAdapter, err := factory.GetDBAdapter(constants.Postgresql)
	if err != nil {
		panic(err)
	}

	db, err := dbAdapter.GetConnection()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	productDAO, err := dao.NewProductDAO(dbAdapter, db)
	if err != nil {
		return
	}

	waitGroup.Add(totalCicles)
	for cicle < totalCicles {
		go func(currentCicle int) {
			defer waitGroup.Done()
			var productList []model.Product

			productList, err = productDAO.GetAllProducts()
			if err != nil {
				fmt.Println(currentCicle, "*** error", err)
				return
			}

			jsonProductList, err := json.Marshal(productList)
			if err != nil {
				fmt.Println(currentCicle, "*** error")
				return
			}

			fmt.Println(currentCicle, string(jsonProductList))
		}(cicle)
		cicle++
	}
	waitGroup.Wait()
	fmt.Println("Process terminated")
}
