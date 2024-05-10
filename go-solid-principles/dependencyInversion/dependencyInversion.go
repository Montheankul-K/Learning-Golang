package main

import "fmt"

type (
	Database interface {
		Insert()
	}
	PostgresDb struct{}
	MockDb     struct{}
)

func NewPostgresDatabase() Database {
	return &PostgresDb{}
}

func NewMockDatabase() Database {
	return &MockDb{}
}

func (p *PostgresDb) Insert() {
	fmt.Println("Postgres insert")
}

func (m *MockDb) Insert() {
	fmt.Println("Mock insert")
}

func InsertPlayerItem(db Database) {
	db.Insert()
}

func main() {
	postgresDb := NewPostgresDatabase()
	InsertPlayerItem(postgresDb)

	mockDb := NewMockDatabase()
	InsertPlayerItem(mockDb)
}
