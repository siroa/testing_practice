package main

func CalcOrder(receipt *Receipt, purchase *Purchase) int {
	receipt.ShippingFee = 300 //配送料は一律300円とする
	if purchase.UserID == 101 {
		receipt.SumPrice = DiscountedPrice(int(purchase.Price), 0.1, 0.07, 0.03)
		if int(purchase.Price) >= 5000 && purchase.Contained {
			receipt.ShippingFee = 0
		}
		receipt.SumPrice += receipt.ShippingFee
	} else if purchase.UserID == 201 {
		receipt.SumPrice = DiscountedPrice(int(purchase.Price), 0.05, 0.03, 0.0)
		if int(purchase.Price) >= 10000 && purchase.Contained {
			receipt.ShippingFee = 0
		}
		receipt.SumPrice += receipt.ShippingFee
	} else if purchase.UserID == 301 {
		receipt.SumPrice = DiscountedPrice(int(purchase.Price), 0.03, 0.0, 0.0)
		receipt.SumPrice += receipt.ShippingFee
	} else {
		return 404
	}
	return 200
}

func DiscountedPrice(price int, a, b, c float64) int {
	sumPrice := 0
	if price >= 10000 {
		sumPrice = price - int(float64(price)*a)
	} else if price >= 5000 {
		sumPrice = price - int(float64(price)*b)
	} else {
		sumPrice = price - int(float64(price)*c)
	}
	return sumPrice
}

func ReturnSizeCost(shippingFeeElem *ShippingFeeElem, shippingFee *ShippingFee) int {
	if shippingFeeElem.Size == "small" {
		shippingFee.ShippingFee += 300
	} else if shippingFeeElem.Size == "medium" {
		shippingFee.ShippingFee += 600
	} else if shippingFeeElem.Size == "large" {
		shippingFee.ShippingFee += 800
	} else {
		return 400
	}
	return 200
}

func ReturnOptionCost(shippingFeeElem *ShippingFeeElem, shippingFee *ShippingFee) int {
	if shippingFeeElem.Option == "regular" {
		shippingFee.ShippingFee += 0
	} else if shippingFeeElem.Option == "designation" {
		shippingFee.ShippingFee += 100
	} else if shippingFeeElem.Option == "haste" {
		shippingFee.ShippingFee += 300
	} else {
		return 400
	}
	return 200
}

func ReturnRegionCost(shippingFeeElem *ShippingFeeElem, shippingFee *ShippingFee) int {
	region := shippingFeeElem.Region
	if region == "tokyo" {
		shippingFee.ShippingFee += 0
	} else if region == "kanto" || region == "toukai" || region == "kansai" {
		shippingFee.ShippingFee += 500
	} else if region == "hokkaido" || region == "sikoku" || region == "kyusyu" {
		shippingFee.ShippingFee += 1000
	} else if region == "touhoku" || region == "koushinetu" || region == "hokuriku" || region == "tyugoku" {
		shippingFee.ShippingFee += 800
	} else if region == "okinawa" {
		shippingFee.ShippingFee += 1200
	} else {
		return 400
	}
	return 200
}
