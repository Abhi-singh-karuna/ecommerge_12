package models

import (
	"time"
)

//
type Product struct {
	ID          *string   `json:"id,omitempty" bson:"_id,omitempty"`
	Name        *string   `json:"product_name"`
	Price       *int      `json:"price,omitempty" bson:"price,omitempty"`
	Quantity    *int      `json:"Quantity"`
	Description *string   `json:"Description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type ProductCart struct {
	Name     *string `json:"product_name"`
	Price    *int    `json:"price,omitempty" bson:"price,omitempty"`
	Quantity *int    `json:"Quantity"`
}

type Total struct {
	Price int `json:"price,omitempty" bson:"price,omitempty"`
}

type Add struct {
	ID        *string `json:"id,omitempty" bson:"_id,omitempty"`
	Full_Name *string `json:"full_name" bson:"full_name"`
	Address_a *string `json:"address_a" bson:"address_a"`
	Address_b *string `json:"address_b" bson:"address_b"`
	City      *string `json:"city" bson:"city"`
	State     *string `json:"state" bson:"state"`
	PinCode   *int    `json:"pincode" bson:"pincode"`
}

type OrderDeatails struct {
	Name     *string `json:"product_name"`
	Quantity *int    `json:"Quantity"`
	Price    int     `json:"price,omitempty" bson:"price,omitempty"`
}
