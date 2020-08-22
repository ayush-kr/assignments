package main

import (
	cm "coffeemachine/coffeemachine"
	"fmt"
	"log"
	"time"
)

func main() {

	outlets := 4

	//Initialize the coffee machine
	cm.NewMachine(outlets)

	//Get the coffee machine instance
	coffeeMachine := cm.GetCoffeeMachine()

	//Add items in coffee machine
	coffeeMachine.AddIngredient("Water", 50, 100, 10)
	coffeeMachine.AddIngredient("CoffeePowder", 25, 50, 5)
	coffeeMachine.AddBeverage("BlackCoffee", []string{"Water", "CoffeePowder"}, []int{10, 5})

	//Start the coffee machine
	go coffeeMachine.Run()

	//Order beverages
	for i := 0; i < 4; i++ {
		coffeeMachine.OrderBeverage("BlackCoffee")
	}
	time.Sleep(10 * time.Second)
	coffeeMachine.OrderBeverage("HotCoffee")

	for i := 0; i < 3; i++ {
		coffeeMachine.OrderBeverage("BlackCoffee")
	}

	time.Sleep(10 * time.Second)
	coffeeMachine.ReFillIngredient("Water", 50)
	coffeeMachine.ReFillIngredient("CoffeePowder", 25)

	for i := 0; i < 3; i++ {
		coffeeMachine.OrderBeverage("BlackCoffee")
	}

	time.Sleep(5 * time.Second)
	coffeeMachine.AddIngredient("Milk", 50, 100, 10)
	coffeeMachine.AddBeverage("HotCoffee", []string{"Milk", "CoffeePowder"}, []int{10, 5})

	for i := 0; i < 10; i++ {
		coffeeMachine.OrderBeverage("HotCoffee")
	}

	//Press any key and enter to stop
	var a int
	fmt.Scanf("%d", &a)
	log.Println("Done")
}
