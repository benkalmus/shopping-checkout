package shopping

type SKU string

type ShoppingCheckout struct {
	// map that stores SKU with a count of items scanned
	Shopping map[SKU]int

	// a mapping for SKUs to a price
	SKUToPriceMap map[SKU]int
}

func NewShoppingCheckout() *ShoppingCheckout {
	return &ShoppingCheckout{
		Shopping:      make(map[SKU]int),
		SKUToPriceMap: make(map[SKU]int),
	}
}
