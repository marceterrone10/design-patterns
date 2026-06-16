package main

import (
	"fmt"
	"log"

	"design-patterns/creacionales/singleton/singleton"
)

func main() {
	db1, err := singleton.GetInstance()
	if err != nil {
		log.Fatal(err)
	}

	db2, err := singleton.GetInstance()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("¿Misma instancia? %t\n", db1 == db2)

	conn := db1.Conn()

	_, err = conn.Exec("CREATE TABLE IF NOT EXISTS logs (id INTEGER PRIMARY KEY, message TEXT)")
	if err != nil {
		log.Fatal(err)
	}

	_, err = conn.Exec("INSERT INTO logs (message) VALUES (?)", "Primera entrada")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db2.Conn().Exec("INSERT INTO logs (message) VALUES (?)", "Segunda entrada")
	if err != nil {
		log.Fatal(err)
	}

	rows, err := conn.Query("SELECT message FROM logs ORDER BY id")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("Mensajes en la base de datos:")
	for rows.Next() {
		var message string
		if err := rows.Scan(&message); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("  - %s\n", message)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
