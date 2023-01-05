package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcOrder(t *testing.T) {

	t.Run("success CalcOrder1", func(t *testing.T) {
		o := NewOrderRepositoryStub()
		expectSumPrice := 3210
		expectShippingFee := 300
		o.Purchase = Purchase{UserID: 101, Price: 3000, Contained: true}

		status := o.CalcOrder()
		if status != 200 {
			t.Error("panic error")
		}

		assert.Equal(t, expectSumPrice, o.Receipt.SumPrice)
		assert.Equal(t, expectShippingFee, o.Receipt.ShippingFee)
	})

	t.Run("success CalcOrder2", func(t *testing.T) {
		o := NewOrderRepositoryStub()
		expectSumPrice := 4650
		expectShippingFee := 0
		o.Purchase = Purchase{UserID: 101, Price: 5000, Contained: true}

		status := o.CalcOrder()
		if status != 200 {
			t.Error("panic error")
		}

		assert.Equal(t, expectSumPrice, o.Receipt.SumPrice)
		assert.Equal(t, expectShippingFee, o.Receipt.ShippingFee)
	})

	t.Run("success CalcOrder3", func(t *testing.T) {
		o := NewOrderRepositoryStub()
		expectSumPrice := 8670
		expectShippingFee := 300
		o.Purchase = Purchase{UserID: 101, Price: 9000, Contained: false}

		status := o.CalcOrder()
		if status != 200 {
			t.Error("panic error")
		}

		assert.Equal(t, expectSumPrice, o.Receipt.SumPrice)
		assert.Equal(t, expectShippingFee, o.Receipt.ShippingFee)
	})

	t.Run("success CalcOrder4", func(t *testing.T) {
		o := NewOrderRepositoryStub()
		expectSumPrice := 9900
		expectShippingFee := 0
		o.Purchase = Purchase{UserID: 101, Price: 11000, Contained: true}

		status := o.CalcOrder()
		if status != 200 {
			t.Error("panic error")
		}

		assert.Equal(t, expectSumPrice, o.Receipt.SumPrice)
		assert.Equal(t, expectShippingFee, o.Receipt.ShippingFee)
	})

	t.Run("success CalcOrder5", func(t *testing.T) {
		o := NewOrderRepositoryStub()
		expectSumPrice := 9300
		expectShippingFee := 300
		o.Purchase = Purchase{UserID: 101, Price: 10000, Contained: false}

		status := o.CalcOrder()
		if status != 200 {
			t.Error("panic error")
		}

		assert.Equal(t, expectSumPrice, o.Receipt.SumPrice)
		assert.Equal(t, expectShippingFee, o.Receipt.ShippingFee)
	})

	t.Run("success CalcOrder6", func(t *testing.T) {
		o := NewOrderRepositoryStub()
		expectSumPrice := 400
		expectShippingFee := 300
		o.Purchase = Purchase{UserID: 201, Price: 100, Contained: false}

		status := o.CalcOrder()
		if status != 200 {
			t.Error("panic error")
		}

		assert.Equal(t, expectSumPrice, o.Receipt.SumPrice)
		assert.Equal(t, expectShippingFee, o.Receipt.ShippingFee)
	})

	t.Run("success CalcOrder7", func(t *testing.T) {
		o := NewOrderRepositoryStub()
		expectSumPrice := 9030
		expectShippingFee := 300
		o.Purchase = Purchase{UserID: 201, Price: 9000, Contained: true}

		status := o.CalcOrder()
		if status != 200 {
			t.Error("panic error")
		}

		assert.Equal(t, expectSumPrice, o.Receipt.SumPrice)
		assert.Equal(t, expectShippingFee, o.Receipt.ShippingFee)
	})

	t.Run("success CalcOrder8", func(t *testing.T) {
		o := NewOrderRepositoryStub()
		expectSumPrice := 9500
		expectShippingFee := 0
		o.Purchase = Purchase{UserID: 201, Price: 10000, Contained: true}

		status := o.CalcOrder()
		if status != 200 {
			t.Error("panic error")
		}

		assert.Equal(t, expectSumPrice, o.Receipt.SumPrice)
		assert.Equal(t, expectShippingFee, o.Receipt.ShippingFee)
	})

	t.Run("success CalcOrder9", func(t *testing.T) {
		o := NewOrderRepositoryStub()
		expectSumPrice := 9800
		expectShippingFee := 300
		o.Purchase = Purchase{UserID: 201, Price: 10000, Contained: false}

		status := o.CalcOrder()
		if status != 200 {
			t.Error("panic error")
		}

		assert.Equal(t, expectSumPrice, o.Receipt.SumPrice)
		assert.Equal(t, expectShippingFee, o.Receipt.ShippingFee)
	})

	t.Run("success CalcOrder10", func(t *testing.T) {
		o := NewOrderRepositoryStub()
		expectSumPrice := 301
		expectShippingFee := 300
		o.Purchase = Purchase{UserID: 301, Price: 1, Contained: false}

		status := o.CalcOrder()
		if status != 200 {
			t.Error("panic error")
		}

		assert.Equal(t, expectSumPrice, o.Receipt.SumPrice)
		assert.Equal(t, expectShippingFee, o.Receipt.ShippingFee)
	})

	t.Run("success CalcOrder11", func(t *testing.T) {
		o := NewOrderRepositoryStub()
		expectSumPrice := 5300
		expectShippingFee := 300
		o.Purchase = Purchase{UserID: 301, Price: 5000, Contained: false}

		status := o.CalcOrder()
		if status != 200 {
			t.Error("panic error")
		}

		assert.Equal(t, expectSumPrice, o.Receipt.SumPrice)
		assert.Equal(t, expectShippingFee, o.Receipt.ShippingFee)
	})

	t.Run("success CalcOrder12", func(t *testing.T) {
		o := NewOrderRepositoryStub()
		expectSumPrice := 48800
		expectShippingFee := 300
		o.Purchase = Purchase{UserID: 301, Price: 50000, Contained: false}

		status := o.CalcOrder()
		if status != 200 {
			t.Error("panic error")
		}

		assert.Equal(t, expectSumPrice, o.Receipt.SumPrice)
		assert.Equal(t, expectShippingFee, o.Receipt.ShippingFee)
	})
}
