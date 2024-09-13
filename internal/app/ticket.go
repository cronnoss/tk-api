package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/cronnoss/tk-api/internal/logger"
	"github.com/cronnoss/tk-api/internal/server"
	"github.com/cronnoss/tk-api/internal/storage"
	"github.com/cronnoss/tk-api/internal/storage/models"
	"golang.org/x/sync/errgroup"
)

type TicketConf struct {
	Logger  logger.Conf  `toml:"logger"`
	Storage storage.Conf `toml:"storage"`
	HTTP    struct {
		Host string `toml:"host"`
		Port string `toml:"port"`
	} `toml:"http-server"`
}

type Ticket struct {
	conf    TicketConf
	log     server.Logger
	storage Storage
}

type Storage interface {
	Connect(ctx context.Context) error
	Close(ctx context.Context) error
	GetShows(ctx context.Context) ([]models.Show, error)
	CreateShows(ctx context.Context, shows []models.Show) ([]models.Show, error)
	CreateShow(ctx context.Context, shows models.Show) (models.Show, error)
	// GetEvents(ctx context.Context) ([]models.Event, error)
	// CreateEvents(ctx context.Context, events []models.Event) ([]models.Event, error)
	// CreateEvent(ctx context.Context, event models.Event) ([]models.Event, error)
	// GetPlaces(ctx context.Context) ([]models.Place, error)
	// CreatePlaces(ctx context.Context, places []models.Place) ([]models.Place, error)
	// CreatePlace(ctx context.Context, place models.Place) (models.Place, error)
}

type Server interface {
	Start(context.Context) error
	Stop(context.Context) error
}

func (t *Ticket) GetShows(ctx context.Context) ([]models.Show, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return t.storage.GetShows(ctx)
}

func (t *Ticket) CreateShows(ctx context.Context, shows []models.Show) ([]models.Show, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return t.storage.CreateShows(ctx, shows)
}

func (t *Ticket) CreateShow(ctx context.Context, show models.Show) (models.Show, error) {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()
	return t.storage.CreateShow(ctx, show)
}

func NewTicket(log server.Logger, conf TicketConf, storage Storage) (*Ticket, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := storage.Connect(ctx); err != nil {
		server.Exitfail(fmt.Sprintf("Can't connect to storage:%v", err))
	}

	return &Ticket{log: log, conf: conf, storage: storage}, nil
}

func (t Ticket) Run(httpsrv Server) {
	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	g, ctxEG := errgroup.WithContext(ctx)

	func1 := func() error {
		return httpsrv.Start(ctxEG)
	}

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := httpsrv.Stop(ctx); err != nil {
			if !errors.Is(err, http.ErrServerClosed) &&
				!errors.Is(err, context.Canceled) {
				t.log.Errorf("failed to stop HTTP-server:%v\n", err)
			}
		}

		if err := t.storage.Close(ctx); err != nil {
			t.log.Errorf("failed to close db:%v\n", err)
		}
	}()

	g.Go(func1)

	if err := g.Wait(); err != nil {
		if !errors.Is(err, http.ErrServerClosed) &&
			!errors.Is(err, context.Canceled) {
			t.log.Errorf("%v\n", err)
		}
	}
}
