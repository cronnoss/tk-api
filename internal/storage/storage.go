package storage

import (
	"context"
	"fmt"
	"os"

	"github.com/cronnoss/tk-api/internal/storage/models"
	sqlstorage "github.com/cronnoss/tk-api/internal/storage/sql"
)

type Conf struct {
	DB  string `toml:"db"`
	DSN string `toml:"dsn"`
}

type Storage interface {
	Connect(ctx context.Context) error
	Close(ctx context.Context) error
	GetShows(ctx context.Context) ([]models.Show, error)
	CreateShows(ctx context.Context, shows []models.Show) ([]models.Show, error)
	CreateShow(ctx context.Context, shows models.Show) (models.Show, error)
	GetEvents(ctx context.Context) ([]models.Event, error)
	CreateEvents(ctx context.Context, events []models.Event) ([]models.Event, error)
	CreateEvent(ctx context.Context, event models.Event) (models.Event, error)
	// GetPlaces(ctx context.Context) ([]models.Place, error)
	// CreatePlaces(ctx context.Context, places []models.Place) ([]models.Place, error)
	// CreatePlace(ctx context.Context, place models.Place) (models.Place, error)
}

func NewStorage(conf Conf) Storage {
	switch conf.DB {
	// case "in_memory":
	// 	return memorystorage.New()
	case "sql":
		return sqlstorage.New(conf.DSN)
	}

	fmt.Fprintln(os.Stderr, "wrong DB")
	os.Exit(1)
	return nil
}
