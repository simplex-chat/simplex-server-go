package db

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var pool *sql.DB

type NewSimplex struct {
	Recipient_id  string
	Sender_id     string
	Recipient_key []byte
}

func Open() {
	connStr := "postgres://localhost:5432/postgres?sslmode=disable"
	var err error
	if pool, err = sql.Open("postgres", connStr); err != nil {
		log.Fatal(err)
	}
	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(5)
	pool.SetMaxOpenConns(5)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := Ping(ctx); err != nil {
		log.Fatal(err)
	}
}

func Close() {
	pool.Close()
}

func Ping(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	return pool.PingContext(ctx)
}

func CreateConnection(ctx context.Context, simplex NewSimplex) sql.Result {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	conn, err := pool.Conn(ctx)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer conn.Close()
	tx, err := conn.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Println(err)
	}
	conn.ExecContext(ctx, `INSERT INTO ids (id) VALUES ($1)`, simplex.Recipient_id)
	conn.ExecContext(ctx, `INSERT INTO ids (id) VALUES ($1)`, simplex.Sender_id)
	result, err := conn.ExecContext(ctx, `INSERT INTO connections
		(recipient_id, sender_id, recipient_key) VALUES ($1, $2, $3);`,
		simplex.Recipient_id, simplex.Sender_id, simplex.Recipient_key)

	tx.Commit()

	return result
}
