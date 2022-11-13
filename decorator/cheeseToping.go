package main

type CheeseTopping struct {
	IPizza
}

func (c *CheeseTopping) getPrice() int {
	pizzaPrice := c.getPrice()
	return pizzaPrice + 10
}
