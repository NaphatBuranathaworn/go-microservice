package database

import (
	"context"

	"github.com/NaphatBuranathaworn/go-microservice/internal"
)


type databaseservice struct{}

func NewService() Service { return &databaseservice{} }

func (d *databaseservice) Add(_ context.Context, doc *internal.Document) (string, error) {
    return "", nil
}

func (d *databaseservice) Get(_ context.Context, filters ...internal.Filter) ([]internal.Document, error) {
    doc := internal.Document{
        Content: "book-database",
    }
    return []internal.Document{doc}, nil
}