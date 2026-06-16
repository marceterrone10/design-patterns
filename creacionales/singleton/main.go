package main

import (
	"fmt"

	"design-patterns/creacionales/singleton/singleton"
)

func main() {
	logger1 := singleton.GetInstance()
	logger2 := singleton.GetInstance()

	logger1.Log("Primera entrada")
	logger2.Log("Segunda entrada")

	fmt.Printf("¿Misma instancia? %t\n", logger1 == logger2)
	fmt.Printf("Mensajes: %v\n", logger1.GetMessages())
}
