package database

import (
	"context"

	"github.com/NaphatBuranathaworn/go-microservice/internal"
)

type Service interface {
    Add(ctx context.Context, doc *internal.Document) (string, error)
    Get(ctx context.Context, filters ...internal.Filter) ([]internal.Document, error)
}