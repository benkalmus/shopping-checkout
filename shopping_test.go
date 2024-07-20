package shopping

import (
	"fmt"
	"testing"
)

func TestCheckout(t *testing.T) {

	t.Run("scanning one item returns correct price", func(t *testing.T) {
		shoppingCheckout := NewShoppingCheckout()
		item := "A"
		price := 50
		// shopping checkout requires a map of string to prices, this will be required to calculate price
		itemPriceMap := map[string]int{item: price}

		shoppingCheckout.SetSKUToPriceMapping(itemPriceMap)

		err := shoppingCheckout.Scan(item)

		if err != nil {
			t.Fatalf("Unexpected error scanning item: %v", err)
		}

		totalPrice, err := shoppingCheckout.GetTotalPrice()

		if err != nil {
			t.Fatalf("Unexpected error getting total price: %v", err)
		}

		if totalPrice != price {
			t.Fatalf("Expected total price %d got %d", price, totalPrice)
		}
	})
	//things to test
	// negative item prices
	// unknown item
	// discount
	// discount for item not found in SKUPriceMap
	//

	t.Run("setting a negative price on an item returns error", func(t *testing.T) {
		shoppingCheckout := NewShoppingCheckout()
		item := "A"
		price := -50
		itemPriceMap := map[string]int{item: price}

		err := shoppingCheckout.SetSKUToPriceMapping(itemPriceMap)

		if err == nil {
			t.Fatalf("Expected error returned on negative price")
		}

	})

}

func TestCheckoutWithoutDiscount(t *testing.T) {
	//setup test
	itemPriceMap := map[string]int{
		"A": 50,
		"B": 30,
		"C": 20,
		"D": 15,
	}

	cases := []struct {
		items []string
		total int
	}{
		{[]string{"A", "B"}, 80},
		{[]string{"A", "A"}, 100},
		{[]string{"A", "B", "C", "D"}, 115},
		{[]string{"A", "A", "A", "C"}, 170},
		{[]string{"A", "A", "C", "C"}, 140},
	}

	for _, testData := range cases {
		t.Run(fmt.Sprintf("%v returns total price %d", testData.items, testData.total), func(t *testing.T) {
			shopping := NewShoppingCheckout()
			shopping.SetSKUToPriceMapping(itemPriceMap)
			for _, item := range testData.items {
				err := shopping.Scan(item)
				if err != nil {
					t.Fatalf("Unexpected error scanning item: %v", err)
				}
			}
			gotTotal, err := shopping.GetTotalPrice()
			if err != nil {
				t.Fatalf("Unexpected error getting total price: %v", err)
			}
			if gotTotal != testData.total {
				t.Fatalf("got %q, want %q", gotTotal, testData.total)
			}
		})
	}

}
