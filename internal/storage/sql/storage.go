package sqlstorage

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cronnoss/tk-api/internal/storage/models"
	_ "github.com/jackc/pgx/stdlib" // nolint: revive
	"github.com/jmoiron/sqlx"
)

type Storage struct {
	dsn string
	db  *sqlx.DB
}

type ShowSQL struct {
	ID   sql.NullInt64
	Name sql.NullString
}

func New(dsn string) *Storage {
	return &Storage{dsn: dsn}
}

func (s *Storage) Connect(ctx context.Context) error {
	db, err := sqlx.Open("pgx", s.dsn)
	if err != nil {
		return fmt.Errorf("failed to load driver: %w", err)
	}
	s.db = db
	err = s.db.PingContext(ctx)
	if err != nil {
		return fmt.Errorf("failed to connect to db: %w", err)
	}
	return nil
}

func (s *Storage) Close(ctx context.Context) error {
	s.db.Close()
	ctx.Done()
	return nil
}

func stringNull(s string) sql.NullString {
	if len(s) == 0 {
		return sql.NullString{}
	}
	return sql.NullString{String: s, Valid: true}
}

// GetShows returns shows.
func (s *Storage) GetShows(ctx context.Context) ([]models.Show, error) {
	var shows []models.Show
	query := `SELECT id, name FROM shows`
	if err := s.db.SelectContext(ctx, &shows, query); err != nil {
		return nil, fmt.Errorf("failed to get shows: %w", err)
	}
	return shows, nil
}

// CreateShows creates shows.
func (s *Storage) CreateShows(ctx context.Context, shows []models.Show) ([]models.Show, error) {
	insertedShows := make([]models.Show, 0, len(shows))
	for _, show := range shows {
		var newShow models.Show
		err := s.db.GetContext(ctx, &newShow,
			`INSERT INTO shows (name) VALUES ($1)
            ON CONFLICT (name) DO UPDATE SET updated_at = now()
            RETURNING *`,
			show.Name)
		if err != nil {
			return insertedShows, nil // nolint: nilerr
		}
		insertedShows = append(insertedShows, newShow)
	}
	return insertedShows, nil
}

// CreateShow creates a show.
func (s *Storage) CreateShow(ctx context.Context, show models.Show) (models.Show, error) {
	var insertedShow models.Show
	err := s.db.GetContext(ctx, &insertedShow,
		`INSERT INTO shows (name) VALUES ($1) 
		ON CONFLICT (name) DO UPDATE SET updated_at = now()
		RETURNING *`,
		show.Name)
	if err != nil {
		return insertedShow, nil // nolint: nilerr
	}

	return insertedShow, nil
}

// GetEvents returns events.
func (s *Storage) GetEvents(ctx context.Context) ([]models.Event, error) {
	var events []models.Event
	query := `SELECT id, show_id, date FROM events`
	if err := s.db.SelectContext(ctx, &events, query); err != nil {
		return nil, fmt.Errorf("failed to get events: %w", err)
	}
	return events, nil
}

// CreateEvents creates events.
func (s *Storage) CreateEvents(ctx context.Context, events []models.Event) ([]models.Event, error) {
	insertedEvents := make([]models.Event, 0, len(events))
	for _, event := range events {
		var newEvent models.Event
		err := s.db.GetContext(ctx, &newEvent,
			`INSERT INTO events (show_id, date) VALUES ($1, $2)
			ON CONFLICT (show_id, date) DO UPDATE SET updated_at = now()
			RETURNING *`,
			event.ShowID, event.Date)
		if err != nil {
			return insertedEvents, nil // nolint: nilerr
		}
		insertedEvents = append(insertedEvents, newEvent)
	}
	return insertedEvents, nil
}

// CreateEvent creates a event.
func (s *Storage) CreateEvent(ctx context.Context, event models.Event) (models.Event, error) {
	var insertedEvent models.Event
	err := s.db.GetContext(ctx, &insertedEvent,
		`INSERT INTO events (id, show_id, date) VALUES ($1, $2, $3)
	   	ON CONFLICT (id) DO UPDATE SET updated_at = now()
		RETURNING *`,
		event.ID, event.ShowID, event.Date)
	if err != nil {
		return insertedEvent, nil // nolint: nilerr
	}

	return insertedEvent, nil
}

// GetPlaces returns places.
func (s *Storage) GetPlaces(ctx context.Context) ([]models.Place, error) {
	var places []models.Place
	query := `SELECT id, x, y, width, height, is_available FROM places`
	if err := s.db.SelectContext(ctx, &places, query); err != nil {
		return nil, fmt.Errorf("failed to get places: %w", err)
	}
	return places, nil
}

// CreatePlaces creates places.
func (s *Storage) CreatePlaces(ctx context.Context, places []models.Place) ([]models.Place, error) {
	insertedPlaces := make([]models.Place, 0, len(places))
	for _, place := range places {
		var newPlace models.Place
		err := s.db.GetContext(ctx, &newPlace,
			`INSERT INTO places (x, y, width, height, is_available) VALUES ($1, $2, $3, $4, $5)
			ON CONFLICT (x, y, width, height) DO UPDATE SET updated_at = now()
			RETURNING *`,
			place.X, place.Y, place.Width, place.Height, place.IsAvailable)
		if err != nil {
			return insertedPlaces, nil // nolint: nilerr
		}
		insertedPlaces = append(insertedPlaces, newPlace)
	}
	return insertedPlaces, nil
}

// CreatePlace creates a place.
func (s *Storage) CreatePlace(ctx context.Context, place models.Place) (models.Place, error) {
	var insertedPlace models.Place
	err := s.db.GetContext(ctx, &insertedPlace,
		`INSERT INTO places (id, x, y, width, height, is_available) VALUES ($1, $2, $3, $4, $5, $6)
		ON CONFLICT (id) DO UPDATE SET updated_at = now()
	   	RETURNING *`,
		place.ID, place.X, place.Y, place.Width, place.Height, place.IsAvailable)
	if err != nil {
		return insertedPlace, nil // nolint: nilerr
	}

	return insertedPlace, nil
}
