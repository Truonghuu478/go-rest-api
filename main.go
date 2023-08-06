package main

import (
	"context"
	"fmt"
	"go-restAPI/config"
	"go-restAPI/models"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func main() {
	// Get the MongoDB client from the config package
	client := config.Client

	// Get the collection
	collection := client.Database("demo").Collection("cart")
	fmt.Println(collection)
	// config http

	router := gin.Default()
	resp, err := collection.Find(context.Background(), bson.M{})
	router.GET("/product", func(ctx *gin.Context) {
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Close(context.Background())
		// Duyệt qua từng document và đưa vào slice
		var results []models.Cart
		for resp.Next(context.Background()) {
			var item models.Cart
			err := resp.Decode(&item)
			if err != nil {
				log.Fatal(err)
			}
			results = append(results, item)
		}

		// Kiểm tra lỗi
		if err := resp.Err(); err != nil {
			log.Fatal(err)
		}
		// In kết quả

		ctx.JSON(200, gin.H{
			"data":   results,
			"status": 200,
		})
	})
	router.Run("localhost:6000")

}
