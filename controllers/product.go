package controllers

import (
	api_responses "github.com/frootlo/api-responses"
	"github.com/frootlo/models"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	"strconv"
)

func GetAllProducts(ctx *gin.Context) {
	products, err := models.GetAllProducts()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
	} else {
		ctx.JSON(http.StatusOK, api_responses.Response{Status: 1, Message: "Request completed successfully", Data: products})
	}
}

func GetProduct(ctx *gin.Context) {
	productId := ctx.Params.ByName("id")
	if _, err := strconv.Atoi(productId); err != nil {
		ctx.JSON(http.StatusBadRequest, api_responses.CustomError{Status: 0, Message: api_responses.InvalidProductIdErr, Error: err.Error()})
	} else {
		product := models.Product{}
		customErr := product.GetProduct(cast.ToInt64(productId))
		if customErr != nil {
			ctx.JSON(http.StatusBadRequest, customErr)
		} else {
			ctx.JSON(http.StatusOK, api_responses.Response{Status: 1, Message: "Request completed successfully", Data: product})
		}
	}
}
