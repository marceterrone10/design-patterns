singleton/
├── main.go                  # Cliente: obtiene siempre la misma instancia
└── singleton/               # Singleton concreto
    └── logger.go            #   Logger + GetInstance (sync.Once)
