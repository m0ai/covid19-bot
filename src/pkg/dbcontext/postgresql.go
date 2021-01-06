package dbcontext

import (
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"scrapper/internal/entity"
)

type contextKey int

// DB represents a DB connection that can be used to run SQL queries.
type DB struct {
	db *gorm.DB
}

func New(db *gorm.DB) *DB {
	return &DB{db}
}

const (
	txKey contextKey = iota
)

// DB returns the gorm.DB wrapped by this object.
func (db *DB) DB() *gorm.DB {
	return db.db
}

// With returns a Builder that can be used to build and execute SQL queries.
// With will return the transaction if it is found in the given context
func (db *DB) With(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(txKey).(*gorm.DB); ok {
		return tx
	}
	return db.db.WithContext(ctx)
}

// Will be remove :D
func DbInitConfig() *gorm.DB {
	dsn := fmt.Sprint(""+
		"host=", os.Getenv("POSTGRES_HOST"), " ",
		"dbname=", os.Getenv("POSTGRES_DB"), " ",
		"user=", os.Getenv("POSTGRES_USER"), " ",
		"password=", os.Getenv("POSTGRES_PASSWORD"), " ",
		"port=5432 sslmode=disable TimeZone=Asia/Seoul connect_timeout=15")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connect database")
	}

	err = db.AutoMigrate(&entity.Covid19InfoEntity{})
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
