package models

import (
	"fmt"
	api_responses "github.com/frootlo/api-responses"
	"github.com/frootlo/dbhelper"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type ProductRate struct {
	Id            int64   `json:"id"`
	ProductPack   string  `json:"product_pack"`
	Price         int64 `json:"price"`
	ProductTypeId int64   `json:"product_type_id"`
}

type ProductType struct {
	Id          int64         `json:"id"`
	Name        string        `json:"name"`
	ProductRate []ProductRate `json:"product_rate"`
	ProductId   int64         `json:"product_id"`
}

type ProductImage struct {
	ImageUrl  string `json:"image_url"`
	ProductId int64  `json:"product_id"`
}

type Product struct {
	Id            int64          `json:"id"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	ProductImages []ProductImage `json:"product_images"`
	ProductTypes  []ProductType  `json:"product_types"`
	VisibleToRecommendation bool `json:"visible_to_recommendation"`
}

func GetAllProducts() ([]*Product, *api_responses.CustomError) {
	products := []*Product{}
	db, err := dbhelper.GetDBConnection()
	if err != nil {
		return nil, &api_responses.CustomError{Status: 0, Message: api_responses.InvalidDatabaseErr, Error: err.Error()}
	}
	defer db.Close()

	if errArr := db.Preload("ProductImages").Find(&products).GetErrors(); len(errArr) > 0 {
		fmt.Println(fmt.Sprintf("Failed to get products. Error: %v", errArr))
		return nil, &api_responses.CustomError{Status: 0, Message: api_responses.ServerErr, Error: errArr[0].Error()}
	}
	for _, p := range products {
		productType := []ProductType{}
		if errArr := db.Preload("ProductRate").Where("product_id=?", p.Id).Find(&productType).GetErrors(); len(errArr) > 0 {
			fmt.Println(fmt.Sprintf("Failed to get product types and rates for product <%v>. Error: %v", p.Id, errArr))
			return nil, &api_responses.CustomError{Status: 0, Message: api_responses.ServerErr, Error: errArr[0].Error()}
		}
		p.ProductTypes = productType
	}

	return products, nil
}

func (p *Product) GetProduct(produuctId int64) *api_responses.CustomError {
	db, err := dbhelper.GetDBConnection()
	if err != nil {
		return &api_responses.CustomError{Status: 0, Message: api_responses.InvalidDatabaseErr, Error: err.Error()}
	}
	defer db.Close()

	if errArr := db.Preload("ProductImages").Where("id=?", produuctId).Find(&p).GetErrors(); len(errArr) > 0 {
		fmt.Println(fmt.Sprintf("Failed to get product <%v>, error: %v", produuctId, errArr))
		for _, err := range errArr {
			if err == gorm.ErrRecordNotFound {
				return &api_responses.CustomError{Status: 0, Message: api_responses.ProductNotExistErr, Error: errArr[0].Error()}
			}
		}
		return &api_responses.CustomError{Status: 0, Message: api_responses.ServerErr, Error: errArr[0].Error()}
	}

	productType := []ProductType{}
	if errArr := db.Preload("ProductRate").Where("product_id=?", produuctId).Find(&productType).GetErrors(); len(errArr) > 0 {
		fmt.Println(fmt.Sprintf("Failed to get product types and rates, product id <%v>, error: %v", produuctId, errArr))
		return &api_responses.CustomError{Status: 0, Message: api_responses.ServerErr, Error: errArr[0].Error()}
	}
	p.ProductTypes = productType

	return nil
}
