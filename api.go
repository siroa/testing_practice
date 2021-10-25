package main

import (
	"encoding/json"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

type User struct {
	UserID int64  `json:"user_id"`
	Name   string `json:"name"`
	Grade  string `json:"grade"`
}

type UsersArray []*User

type Users struct {
	Users UsersArray `json:"users"`
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	responseUser1 := User{UserID: 101, Name: "John Doe", Grade: "gold"}
	responseUser2 := User{UserID: 201, Name: "Randy Barrett", Grade: "silver"}
	responseUser3 := User{UserID: 301, Name: "Andrew Cary", Grade: "nonmember"}
	var users UsersArray
	users = append(users, &responseUser1)
	users = append(users, &responseUser2)
	users = append(users, &responseUser3)
	responseUsers := Users{Users: users}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseUsers)
}

type Purchase struct {
	UserID int64 `json:"user_id" validate:"required"`
	Price  uint  `json:"price" validate:"required"`
	Target bool  `json:"target" validate:"required"`
}

type Receipt struct {
	OrderID     int64 `json:"order_id"`
	ShippingFee int   `json:"shipping_fee"`
	SumPrice    int   `json:"sum_price"`
}

func PostOrders(w http.ResponseWriter, r *http.Request) {
	//order id をランダム生成
	rand.Seed(time.Now().UnixNano())

	purchase := Purchase{}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	if err := json.Unmarshal(body, &purchase); err != nil {
		w.WriteHeader(400)
		return
	}

	receipt := Receipt{}
	receipt.OrderID = int64(rand.Intn(99) + 100)
	receipt.ShippingFee = 310
	if purchase.Price == 0 {
		w.WriteHeader(400)
		return
	}
	status := CalcOrder(&receipt, &purchase)
	w.WriteHeader(status)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(receipt)
}

type ShippingFeeElem struct {
	Size   string `json:"size" validate:"required"`
	Option string `json:"option" validate:"required"`
	Region string `json:"region" validate:"required"`
}

type ShippingFee struct {
	ShippingFee uint `json:"shipping_fee"`
}

func PostShippingFee(w http.ResponseWriter, r *http.Request) {
	shippingFeeElem := ShippingFeeElem{}
	shippingFee := ShippingFee{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	if err := json.Unmarshal(body, &shippingFeeElem); err != nil {
		w.WriteHeader(400)
		return
	}

	status := ReturnSizeCost(&shippingFeeElem, &shippingFee)
	if status != 200 {
		w.WriteHeader(status)
		return
	}
	status = ReturnOptionCost(&shippingFeeElem, &shippingFee)
	if status != 200 {
		w.WriteHeader(status)
		return
	}
	status = ReturnRegionCost(&shippingFeeElem, &shippingFee)
	if status != 200 {
		w.WriteHeader(status)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shippingFee)
}
