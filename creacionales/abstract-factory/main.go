package main

import (
	"fmt"

	"design-patterns/creacionales/abstract-factory/factory"
	"design-patterns/creacionales/abstract-factory/product"
)

func main() {
	logitechFactory, _ := factory.GetPerifericosFactory("logitech")
	razerFactory, _ := factory.GetPerifericosFactory("razer")

	printMouseInfo(logitechFactory.CreateMouse())
	printKeyboardInfo(logitechFactory.CreateKeyboard())
	printMouseInfo(razerFactory.CreateMouse())
	printKeyboardInfo(razerFactory.CreateKeyboard())
}

func printMouseInfo(mouse product.IMouse) {
	fmt.Printf("Mouse: %s, Price: %.2f\n", mouse.GetLogo(), mouse.GetPrice())
}

func printKeyboardInfo(keyboard product.IKeyboard) {
	fmt.Printf("Keyboard: %s, Price: %.2f\n", keyboard.GetLogo(), keyboard.GetPrice())
}
