package singleton

import (
	"database/sql"
	"sync"

	_ "modernc.org/sqlite"
)

type Database struct {
	conn *sql.DB
}

var (
	instance *Database
	mu       sync.Mutex
)

func GetInstance() (*Database, error) {
	mu.Lock()         // bloquea el acceso a la instancia
	defer mu.Unlock() // desbloquea siempre cuando la funcion getInstance termina, automaticamente

	if instance == nil {
		conn, err := sql.Open("sqlite", ":memory:")
		if err != nil {
			return nil, err
		}

		if err := conn.Ping(); err != nil {
			conn.Close()
			return nil, err
		}

		instance = &Database{conn: conn}
	}

	return instance, nil
}

func (db *Database) Conn() *sql.DB {
	return db.conn
}

func (db *Database) Close() error {
	return db.conn.Close()
}
