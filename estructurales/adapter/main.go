package main

import (
	"fmt"

	"design-patterns/estructurales/adapter/adaptee"
	"design-patterns/estructurales/adapter/adapter"
	"design-patterns/estructurales/adapter/target"
)

func main() {
	vehiculos := []target.IVehicle{
		adapter.NewSubaruAdapter(&adaptee.SubaruImpreza{}),
		adapter.NewFerrariAdapter(&adaptee.FerrariModerna{}),
	}

	for _, vehiculo := range vehiculos {
		fmt.Println(vehiculo.Avanzar())
		fmt.Println()
	}
}
