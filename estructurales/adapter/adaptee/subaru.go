package adaptee

type SubaruImpreza struct{}

func (s *SubaruImpreza) PresionarPedalFreno() string {
	return "Subaru: pedal de freno presionado"
}

func (s *SubaruImpreza) InsertarLlave() string {
	return "Subaru: llave insertada"
}

func (s *SubaruImpreza) GirarLlave() string {
	return "Subaru: llave girada, motor encendido"
}

func (s *SubaruImpreza) SoltarFreno() string {
	return "Subaru: pedal de freno soltado"
}
