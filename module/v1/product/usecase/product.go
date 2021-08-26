package usecase

import (
	"errors"
	"go-simple/config"
	"go-simple/model"
	"go-simple/module/v1/product/repo"
	"strings"
)

// get product list
func ProductList(conf config.Configuration) (products []model.ProductView, err error) {
	db := conf.MysqlDB
	return repo.GetProductList(db)
}

// get product detail
func ProductDetail(conf config.Configuration, Id int) (users model.ProductView, err error) {
	db := conf.MysqlDB
	return repo.GetProductDetail(db, Id)
}

// create product
func CreateProduct(conf config.Configuration, prd *model.Product) (prod model.Product, err error) {

	tx, err := conf.MysqlDB.Begin()
	if err != nil {
		return prod, err
	}

	if strings.TrimSpace(prd.Name) == "" || strings.TrimSpace(prd.SKU) == "" {
		return prod, errors.New("data can not null")
	}

	_, err = repo.CreateProduct(tx, prd)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return *prd, nil
}

// update product
func UpdateProduct(conf config.Configuration, prd *model.Product) (user model.Product, err error) {

	tx, err := conf.MysqlDB.Begin()
	if err != nil {
		return user, err
	}

	_, err = repo.UpadateProduct(tx, prd)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return *prd, nil
}

// delete product
func DeleteProduct(conf config.Configuration, prd *model.Product) (product model.Product, err error) {
	var (
		payload = model.Product{}
	)
	tx, err := conf.MysqlDB.Begin()
	if err != nil {
		return product, err
	}

	payload.Id = prd.Id
	payload.Deleted = 1

	_, err = repo.DeleteProduct(tx, &payload)
	if err != nil {
		tx.Rollback()
		return
	}

	tx.Commit()

	return *prd, nil
}
