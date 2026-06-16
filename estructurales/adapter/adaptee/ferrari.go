package adaptee

type FerrariModerna struct{}

func (f *FerrariModerna) PresionarFreno() string {
	return "Ferrari: freno presionado"
}

func (f *FerrariModerna) PonerEnNeutral() string {
	return "Ferrari: caja en neutral"
}

func (f *FerrariModerna) ApretarBotonStart() string {
	return "Ferrari: botón start del volante presionado, motor encendido"
}
