package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const VERSION string = "v1"

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/"+VERSION+"/users", GetUser).Methods("GET")
	r.HandleFunc("/"+VERSION+"/orders", PostOrders).Methods("POST")
	r.HandleFunc("/"+VERSION+"/orders/shipping_fee", PostShippingFee).Methods("POST")
	http.Handle("/", r)

	log.Println("Webサーバーを開始します。")
	if err := http.ListenAndServe(":18080", nil); err != nil {
		log.Fatal("Webサーバーの起動に失敗しました。：", err)
	}
}
