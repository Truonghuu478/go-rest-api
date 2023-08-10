package controllers

import (
	"context"
	"go-restAPI/configs"
	"go-restAPI/models"
	"go-restAPI/responses"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
)

var validate = validator.New()
var collection = configs.Client.Database("demo").Collection("cart")

func create() {

}

func getListProduct() gin.HandlerFunc {
	return func(c *gin.Context) {
		var cart models.Cart
		result, err := collection.Find(context.Background(), bson.M{})
		if err != nil {
			cartRes := responses.CartResponse{
				Status:  http.StatusInternalServerError,
				Message: "error",
				Data:    map[string]interface{}{"data": err.Error()},
			}
			c.JSON(http.StatusInternalServerError, cartRes)
			return
		}

		c.JSON(http.StatusOK, responses.CartResponse{
			Status:  http.StatusOK,
			Message: "success",
			Data:    map[string]interface{}{"data": cart},
		})

	}
}
