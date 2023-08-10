package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Cart struct {
	_id        primitive.ObjectID `bson:"_id,omitempty"`
	Products   []Product          `bson:"products,omitempty" validate:"required"`
	Voucher    Voucher            `bson:"voucher,omitempty" validate:"required"`
	TotalPrice float64            `bson:"total_price,omitempty" validate:"required"`
}

type Product struct {
	Name     string  `bson:"name,omitempty" validate:"required"`
	Price    float64 `bson:"price,omitempty" validate:"required"`
	Image    string  `bson:"image,omitempty" validate:"required"`
	Quantity int     `bson:"quantity,omitempty" validate:"required"`
}

type Voucher struct {
	Code     string  `bson:"code,omitempty" validate:"required"`
	Discount float64 `bson:"discount,omitempty" validate:"required"`
}
