singleton/
├── main.go                  # Cliente: obtiene siempre la misma conexión a DB
└── singleton/               # Singleton concreto
    └── database.go          #   Database + GetInstance (sync.Mutex)
