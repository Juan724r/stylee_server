package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
)

const (
	host     = "localhost"
	port     = 5432
	dbname   = "postgres"
	user     = "admin"
	password = "admin"
)

type PostgresConn struct {
	conn *sql.DB
}

func DbConn() *PostgresConn {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	maxRetries := 3
	var conn *sql.DB
	var err error

	for i := 0; i < maxRetries; i++ {
		conn, err = sql.Open("postgres", connStr)
		if err != nil {
			log.Printf("Error connecting to database: %v\n", err)
			continue
		}
		if err = conn.Ping(); err != nil {
			log.Printf("Error pinging database: %v\n", err)
			conn.Close()
			continue
		}
		break
	}
	if err != nil {
		log.Fatalf("Не удалось подключиться к базе данных после %d попыток: %v\n", maxRetries, err)
	}

	log.Println("Соединение с базой данных успешно установлено.")

	// Create migration instance
	m, err := migrate.New(
		"file://db/migrations",
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", user, password, host, port, dbname))
	if err != nil {
		log.Fatalf("Error creating migration instance: %v", err)
	}
	defer m.Close()

	// Run migrations
	err = m.Up()
	if err != nil {
		log.Fatalf("Error running migrations: %v", err)
	}

	return &PostgresConn{conn: conn}
}
