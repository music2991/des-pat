package main

import (
	"dp/creational/factory/db_factory/constants"
	"dp/creational/factory/db_factory/dao"
	"dp/creational/factory/db_factory/model"
	"encoding/json"
	"fmt"
	"sync"
)

func main() {
	executeOne()
	//executeInGoroutines()
}

func executeOne() {
	var productList []model.Product

	productDAO, err := dao.NewProductDAO(constants.Postgresql)
	if err != nil {
		fmt.Println(err, ":A")
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
	var totalCicles = 10
	var waitGroup sync.WaitGroup

	waitGroup.Add(totalCicles)
	for cicle < totalCicles {
		go func(currentCicle int) {
			defer waitGroup.Done()
			var productList []model.Product

			productDAO, err := dao.NewProductDAO(constants.Postgresql)
			if err != nil {
				fmt.Println(currentCicle, "*** error")
				return
			}

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
