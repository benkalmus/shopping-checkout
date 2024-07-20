# Kata for a shopping checkout system

Implement the code for a checkout system that handles pricing schemes such as "pineapples cost 50, three pineapples cost 130."

## Requirements
Implement the code for a supermarket checkout that calculates the total price of a number of items. In a normal supermarket, things are identified using Stock Keeping Units, or SKUs. 

In our store, we'll use individual letters of the alphabet (A, B, C, and so on). Our goods are priced individually. 
In addition, some items are multi-priced: buy `n` of them, and they'll cost you `y` pence. 

For example, item A might cost 50 individually, but this week we have a special offer: buy three As and they'll cost you 130. 

In fact the prices are:

| SKU | Unit Price | Special Price |
|-----|------------|---------------|
| A   | 50         | 3 for 130     |
| B   | 30         | 2 for 45      |
| C   | 20         |               |
| D   | 15         |               |
 
The checkout accepts items in any order, so that if we scan a B, an A, and another B, we'll recognize the two Bs and price them at 45 (for a total price so far of 95). 

The implementation should consider if the pricing model may change frequently.

Please implement your solution to implement the interface:
 
```go
type ICheckout interface {
    Scan(SKU string)(err error)
    GetTotalPrice()(totalPrice int, err error)
}
```
_________
# Solution

Initialised project with empty .go files, [changelog](CHANGELOG.md), [version](vsn.mk) and [Makefile](Makefile). 
The [Makefile](Makefile) serves ease of use for the purpose of testing, running and building the application by running a single command, without worrying about arguments and options.
It also makes the project easier to change and configure in the future, if for example a DB backend is added running in a Docker container. 

See `make help` for usage. 


# Build & Run

The requirements do not specify an Input or Output, therefore the solution doesn't provide it. For this reason, running the app doesn't do anything. 

```sh
make build
make run
```

# Test

Run tests with:

```sh
make test
```

# Development

See `package shopping` in [shopping.go](shopping.go) and [shopping_test.go](shopping_test.go) 

Usage and examples are provided on all public API exposed to `ShoppingCheckout{}`. 

## TODO & Improvements

- Ability to `Clear` the shopping cart or `Remove` individual items
- Concurrency safe access, currently multiple go routines cannot access the same ShoppingCheckout safely. 
- Currently, the app doesn't produce anything. Perhaps in the future read shopping list from a file?
    - Input file .e.g csv, SQLite

- Data generator, to produce random SKUs with Prices and Discounts and write to file or DB