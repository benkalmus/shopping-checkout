package shopping

import "fmt"

// Type definitions
// =======================================================
type ShoppingCheckout struct {
	// map that stores string with a count of items scanned
	Shopping map[string]int

	// a mapping for SKUs to a price
	SKUToPriceMap map[string]int

	// a mapping for SKUs to a Discount
	SKUToDiscountMap map[string]Discount
}

type Discount struct {
	NumItems int
	Price    int
}

// ShoppingCheckout Public API
// =======================================================

// ShoppingCheckout constructor
func NewShoppingCheckout() *ShoppingCheckout {
	return &ShoppingCheckout{
		Shopping:         make(map[string]int),
		SKUToPriceMap:    make(map[string]int),
		SKUToDiscountMap: make(map[string]Discount),
	}
}

// Configures string to Price mapping for Shopping Checkout
// This is used to calculate final price of shopping
// Merges item Price map with existing map, note this will override string collisions
// Example:
//
//	itemPrices := map[string]int{
//		"A": 50,
//		"B": 30,
//		"C": 20,
//		"D": 15
//	}
//	s.SetSKUToPriceMapping(itemPrices)
func (s *ShoppingCheckout) SetSKUToPriceMapping(itemPriceMap map[string]int) error {
	for item, price := range itemPriceMap {
		if price < 0 {
			return fmt.Errorf("item %s's price cannot be negative %d", item, price)
		}
		s.SKUToPriceMap[item] = price
	}
	return nil
}

func (s *ShoppingCheckout) SetDiscountPriceMapping(itemDiscountMap map[string]Discount) error {
	for item, discount := range itemDiscountMap {
		regularPrice, ok := s.SKUToPriceMap[item]
		// cannot apply discount on an item that is not in SKUToPriceMap
		if !ok {
			return fmt.Errorf("item %s not recognised by shop", item)
		}

		// discount doesn't make sense if it's more than the price of an item
		nonDiscountPrice := (regularPrice * discount.NumItems)
		if discount.Price >= nonDiscountPrice {
			return fmt.Errorf("item %s's discount price cannot be more than regular price %d >= %d", item, discount.Price, nonDiscountPrice)
		}

		// store discount mapping
		s.SKUToDiscountMap[item] = discount
	}
	return nil
}

// Adds a new item to shopping checkout,
// returns error if item's price doesn't exist in SKUToPriceMap
// Example:
//
//	item := "A" 	// price for this SKU is 50 (see SetSKUToPriceMapping/1 example)
//	s.Scan(item)
func (s *ShoppingCheckout) Scan(item string) error {
	if _, ok := s.SKUToPriceMap[item]; !ok {
		return fmt.Errorf("item SKU %s not recognised by shop", item)
	}
	s.Shopping[item]++
	return nil
}

func (s *ShoppingCheckout) GetTotalPrice() (int, error) {
	totalPrice := 0
	for item, count := range s.Shopping {
		price := s.SKUToPriceMap[item] //no need to check if item exists, already checked in Scan/1

		discount, ok := s.SKUToDiscountMap[item]
		if ok {
			// apply discount on multiples of discount.NumItems
			discountPrice := discount.Price * (count / discount.NumItems)
			// use regular price on the remainder of items that don't fit into discount.NumItems
			remainderPrice := price * (count % discount.NumItems)

			// shopper savings = (price * count) - (discountPrice + remainderPrice)
			totalPrice += discountPrice + remainderPrice

			continue
		}

		totalPrice += price * count
	}
	return totalPrice, nil
}

// ShoppingCheckout Private Func
// =======================================================
