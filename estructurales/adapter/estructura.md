adapter/
├── main.go                  # Cliente: usa Avanzar() sin conocer la secuencia de cada auto
├── target/                  # Interfaz objetivo (lo que el cliente espera)
│   └── vehicle.go           #   IVehicle + Avanzar()
├── adaptee/                 # Clases incompatibles (cada auto con su propia API)
│   ├── subaru.go            #   SubaruImpreza: freno → llave → girar → soltar
│   └── ferrari.go           #   FerrariModerna: freno → neutral → botón start
└── adapter/                 # Adaptadores (traducen adaptee → target)
    ├── subaru_adapter.go    #   Envuelve SubaruImpreza e implementa IVehicle
    └── ferrari_adapter.go   #   Envuelve FerrariModerna e implementa IVehicle
