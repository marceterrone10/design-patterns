abstract-factory/
├── main.go                  # Cliente: usa la fábrica sin conocer clases concretas
├── product/                 # Productos abstractos (lo que se fabrica)
│   ├── mouse.go             #   IMouse + Mouse
│   └── keyboard.go          #   IKeyboard + Keyboard
└── factory/                 # Fábricas (cómo se fabrica)
    ├── factory.go           #   IPerifericosFactory + GetPerifericosFactory
    ├── logitech.go          #   Fábrica concreta Logitech + sus productos
    └── razer.go             #   Fábrica concreta Razer + sus productos