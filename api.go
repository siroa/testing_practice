package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
)

var (
	PropErr    = "MISSING_REQUEST_PROPERTY"
	PropMsg    = "Please set the correct value."
	SizeMsg    = "Please choose from large, medium and small."
	OptionMsg  = "Please choose from regular, designation and haste."
	RegionMsg  = "Please choose from the prefectures."
	NotFond    = "MISSING_REQUEST_USERID"
	NotFondMsg = "Non-existent user ID."
)

type ErrorMessage struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}

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
	UserID    int64 `json:"user_id" validate:"required"`
	Price     uint  `json:"price" validate:"required,max=1000000,min=0"`
	Contained bool  `json:"contained" validate:"required"`
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
		eMessage := ErrorMessage{Type: PropErr, Message: PropMsg}
		w.WriteHeader(400)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(eMessage)
		return
	}
	if err := json.Unmarshal(body, &purchase); err != nil {
		eMessage := ErrorMessage{Type: PropErr, Message: PropMsg}
		w.WriteHeader(400)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(eMessage)
		return
	}

	validate := validator.New()
	if err := validate.Struct(purchase); err != nil {
		errs := err.(validator.ValidationErrors)
		for _, errrs := range errs {
			fmt.Printf("Field=%s, Tag=%s\n", errrs.Field(), errrs.Tag())
		}
		eMessage := ErrorMessage{Type: PropErr, Message: PropMsg}
		w.WriteHeader(400)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(eMessage)
		return
	}

	receipt := Receipt{}
	receipt.OrderID = int64(rand.Intn(99) + 100)
	if purchase.Price == 0 {
		eMessage := ErrorMessage{Type: PropErr, Message: PropMsg}
		w.WriteHeader(400)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(eMessage)
		return
	}
	status := CalcOrder(&receipt, &purchase)
	if status == 404 {
		eMessage := ErrorMessage{Type: NotFond, Message: NotFondMsg}
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(eMessage)
		return
	}

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

// エラーメッセージ入れ替え
func PostShippingFee(w http.ResponseWriter, r *http.Request) {
	shippingFeeElem := ShippingFeeElem{}
	shippingFee := ShippingFee{}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		eMessage := ErrorMessage{Type: PropErr, Message: PropMsg}
		w.WriteHeader(400)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(eMessage)
		return
	}
	if err := json.Unmarshal(body, &shippingFeeElem); err != nil {
		eMessage := ErrorMessage{Type: PropErr, Message: PropMsg}
		w.WriteHeader(400)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(eMessage)
		return
	}

	status := ReturnSizeCost(&shippingFeeElem, &shippingFee)
	if status != 200 {
		eMessage := ErrorMessage{Type: PropErr, Message: SizeMsg}
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(eMessage)
		return
	}
	status = ReturnOptionCost(&shippingFeeElem, &shippingFee)
	if status != 200 {
		eMessage := ErrorMessage{Type: PropErr, Message: RegionMsg}
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(eMessage)
		return
	}
	status = ReturnRegionCost(&shippingFeeElem, &shippingFee)
	if status != 200 {
		eMessage := ErrorMessage{Type: PropErr, Message: OptionMsg}
		w.WriteHeader(status)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(eMessage)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(shippingFee)
}
