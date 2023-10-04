package main

import (
    "fmt"
)

type Pizza interface {
    GetPrice() int
}

type VeggiePizza struct{}

func (p *VeggiePizza) GetPrice() int {
    return 15
}

type ToppingDecorator struct {
    pizza Pizza
}

func (t *ToppingDecorator) GetPrice() int {
    pizzaPrice := t.pizza.GetPrice()
    return pizzaPrice + t.GetToppingPrice()
}

func (t *ToppingDecorator) GetToppingPrice() int {
    return 0
}

type CheeseTopping struct {
    *ToppingDecorator
}

func (c *CheeseTopping) GetToppingPrice() int {
    return 10
}

type TomatoTopping struct {
    *ToppingDecorator
}

func (t *TomatoTopping) GetToppingPrice() int {
    return 7
}

func main() {
    pizza := &VeggiePizza{}

    pizzaWithCheese := &CheeseTopping{
        &ToppingDecorator{
            pizza: pizza,
        },
    }

    pizzaWithTomatoAndCheese := &TomatoTopping{
        &ToppingDecorator{
            pizza: pizzaWithCheese,
        },
    }

    price := pizzaWithTomatoAndCheese.GetPrice()

    fmt.Printf("Price of VeggiePizza with tomato and cheese toppings is %d\n", price)
}