//go:generate mockery

package server

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/cronnoss/tk-api/internal/storage/models"
)

var (
	ErrID     = errors.New("wrong ID")
	ErrUserID = errors.New("wrong UserID")
	ErrName   = errors.New("wrong Name")
)

type Logger interface {
	Fatalf(format string, a ...interface{})
	Errorf(format string, a ...interface{})
	Warningf(format string, a ...interface{})
	Infof(format string, a ...interface{})
	Debugf(format string, a ...interface{})
}

type Application interface {
	GetShows(ctx context.Context) ([]models.Show, error)
	CreateShows(ctx context.Context, shows []models.Show) ([]models.Show, error)
	CreateShow(ctx context.Context, shows models.Show) (models.Show, error)
	GetEvents(ctx context.Context) ([]models.Event, error)
	CreateEvents(ctx context.Context, events []models.Event) ([]models.Event, error)
	CreateEvent(ctx context.Context, event models.Event) (models.Event, error)
	GetPlaces(ctx context.Context) ([]models.Place, error)
	CreatePlaces(ctx context.Context, places []models.Place) ([]models.Place, error)
	CreatePlace(ctx context.Context, place models.Place) (models.Place, error)
}

func Exitfail(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}
