package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cart struct {
	_id        primitive.ObjectID `bson:"_id,omitempty"`
	Products   []Product          `bson:"products"`
	Voucher    Voucher            `bson:"voucher"`
	TotalPrice float64            `bson:"total_price"`
}

type Product struct {
	Name     string  `bson:"name"`
	Price    float64 `bson:"price"`
	Image    string  `bson:"image"`
	Quantity int     `bson:"quantity"`
}

type Voucher struct {
	Code     string  `bson:"code"`
	Discount float64 `bson:"discount"`
}
