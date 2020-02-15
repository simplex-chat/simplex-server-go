package db

import (
	"context"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func CreateConnection() sql.Result {
	connStr := "postgres://localhost:5432/postgres?sslmode=disable"
	log.Println("here 0")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("here 1")

	// c := 10
	// b := make([]byte, c)
	// _, err := rand.Read(b)
	// rand.Read()

	recipient_uri := "123"
	sender_uri := "456"
	recipient_key := []byte("abc")
	sender_key := []byte("def")

	var ctx context.Context
	tx, err := db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Println(err)
	}
	tx.Exec(`INSERT INTO uris (uri) VALUES (?)`, recipient_uri)
	tx.Exec(`INSERT INTO uris (uri) VALUES (?)`, sender_uri)
	result, _ := tx.Exec(`INSERT INTO connections
		(recipient_uri, sender_uri, recipient_key, sender_key)
		VALUES (?, ?, ?, ?);`,
		recipient_uri, sender_uri, recipient_key, sender_key)

	tx.Commit()

	log.Println("here 2")
	log.Println(err)

	return result
}
