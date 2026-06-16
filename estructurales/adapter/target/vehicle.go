package target

import "fmt"

type IVehicle interface {
	Avanzar() string
}

type Vehicle struct {
	Nombre string
}

func Avanzando(auto string) string {
	return fmt.Sprintf("%s avanzando", auto)
}

func PrintAvanzando(auto string) {
	fmt.Println(Avanzando(auto))
}

func PrintAvance(secuencia, auto string) {
	fmt.Println(secuencia)
	PrintAvanzando(auto)
}
