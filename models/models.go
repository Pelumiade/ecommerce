package models

import ()

type User struct {
	ID              primitive.objectID `json:"_id" bson:"_id"`
	First_name      *string            `json:"first_name"`
	Last_name       *string            `json:"last_name"`
	Password        *string            `json:"password"`
	Email           *string            `json:"email"`
	Phone           *string            `json:"phone"`
	Token           *string            `json:"token"`
	Refresh_Token   *string            `json:"refresh_token"`
	Created_At      time.Time          `json:"created_at"`
	Updated_At      time.Time          `json:"uppdated_at"`
	User_ID         string             `json:"user_id"`
	UserCart        []ProductUser      `json:"usercart" bson:"usercart"`
	Address_Details []Address          `json:"address" bson:"address"`
	Order_Status    []Order            `json:"orders" bson:"orders"`
}

type Product struct {
	Product_ID   primitive.bjectID `bson:"_id"`
	Product_Name *string           `json:"product_name"`
	Price        *uint64           `json:"price"`
	Rating       *uint8            `json:"rating"`
	Image        *string           `json:"image"`
}

type ProductUser struct {
	Product_ID   primitive.ObjectID `bson:"_id"`
	Product_Name *string
	Price        int
	Rating       uint
	Image        *string
}

type Address struct {
	Address_id *string
	House      *string
	Street     *string
	City       *string
	Pincode    *string
}

type Order struct {
	Order_ID       primitive.ObjectID
	Order_Cart     []ProductUser
	Ordered_At     time.Time
	Price          int
	Discount       *int
	Payment_Method Payment
}

type Payment struct {
	Digital bool
	COD     bool
}
