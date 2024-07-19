package shopping 

import (
	"testing"
)

func TestCheckout(t *testing.T) {

	t.Run("scanning one item returns correct price", func(t *testing.T) {
		shoppingCheckout := NewShoppingCheckout()
		item := "A"
		price := 50
		// shopping checkout requires a map of SKU to prices, this will be required to calculate price
		itemPriceMap := map[string]int{item: price}
		
		shoppingCheckout.SetSKUToPriceMapping(itemPriceMap)

		err := shoppingCheckout.Scan(item)

		if err != nil {
			t.Fatalf("Unexpected error scanning item: %v", err)
		}

		totalPrice, err := shoppingCheckout.GetTotalPrice()
	
		if totalPrice != price {
			t.Fatalf("Expected total price %d got %d", price, totalPrice)
		}
	})

}