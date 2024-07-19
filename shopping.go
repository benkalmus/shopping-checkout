package shopping

type ShoppingCheckout struct {
	// map that stores string with a count of items scanned
	Shopping map[string]int

	// a mapping for SKUs to a price
	SKUToPriceMap map[string]int
}

// ShoppingCheckout Public API
// =======================================================

// ShoppingCheckout constructor
func NewShoppingCheckout() *ShoppingCheckout {
	return &ShoppingCheckout{
		Shopping:      make(map[string]int),
		SKUToPriceMap: make(map[string]int),
	}
}

// Configures string to Price mapping for Shopping Checkout
// This is used to calculate final price of shopping
// Merges item Price map with existing map, note this will override string collisions
// Example:
//	itemPrices := map[string]int{
//		"A": 50, 
//		"B": 30, 
//		"C": 20, 
//		"D": 15
//	}
// 	s.SetSKUToPriceMapping(itemPrices)
func (s *ShoppingCheckout)SetSKUToPriceMapping(itemPriceMap map[string]int) {
	for item, price := range itemPriceMap {
		s.SKUToPriceMap[item] = price
	}
}


// ShoppingCheckout Private Func
// =======================================================

