package memory

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/cronnoss/tk-api/internal/storage/models"
)

type mapShow map[int64]*models.Show

type mapEvent map[int64]*models.Event

type mapPlace map[int64]*models.Place

type Storage struct {
	dataShow  mapShow
	dataEvent mapEvent
	dataPlace mapPlace
	mu        sync.RWMutex
}

var GenID int64

func getNewIDSafe() int64 {
	return atomic.AddInt64(&GenID, 1)
}

func New() *Storage {
	return &Storage{
		dataShow:  make(mapShow),
		dataEvent: make(mapEvent),
		dataPlace: make(mapPlace),
		mu:        sync.RWMutex{},
	}
}

func (s *Storage) Connect(_ context.Context) error {
	return nil
}

func (s *Storage) Close(_ context.Context) error {
	return nil
}

// GetShows returns shows.
func (s *Storage) GetShows(_ context.Context) ([]models.Show, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	sliceS := []models.Show{}
	for _, v := range s.dataShow {
		sliceS = append(sliceS, *v)
	}
	return sliceS, nil
}

// CreateShows creates shows.
func (s *Storage) CreateShows(_ context.Context, shows []models.Show) ([]models.Show, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range shows {
		shows[i].ID = getNewIDSafe()
		s.dataShow[shows[i].ID] = &shows[i]
	}
	return shows, nil
}

// CreateShow creates a show.
func (s *Storage) CreateShow(_ context.Context, show models.Show) (models.Show, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	show.ID = getNewIDSafe()
	s.dataShow[show.ID] = &show
	return show, nil
}

// GetEvents returns events.
func (s *Storage) GetEvents(_ context.Context) ([]models.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	sliceE := []models.Event{}
	for _, v := range s.dataEvent {
		sliceE = append(sliceE, *v)
	}
	return sliceE, nil
}

// CreateEvents creates events.
func (s *Storage) CreateEvents(_ context.Context, events []models.Event) ([]models.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range events {
		events[i].ID = getNewIDSafe()
		s.dataEvent[events[i].ID] = &events[i]
	}
	return events, nil
}

// CreateEvent creates an event.
func (s *Storage) CreateEvent(_ context.Context, event models.Event) (models.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	event.ID = getNewIDSafe()
	s.dataEvent[event.ID] = &event
	return event, nil
}

// GetPlaces returns places.
func (s *Storage) GetPlaces(_ context.Context) ([]models.Place, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	sliceP := []models.Place{}
	for _, v := range s.dataPlace {
		sliceP = append(sliceP, *v)
	}
	return sliceP, nil
}

// CreatePlaces creates places.
func (s *Storage) CreatePlaces(_ context.Context, places []models.Place) ([]models.Place, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i := range places {
		places[i].ID = getNewIDSafe()
		s.dataPlace[places[i].ID] = &places[i]
	}
	return places, nil
}

// CreatePlace creates a place.
func (s *Storage) CreatePlace(_ context.Context, place models.Place) (models.Place, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	place.ID = getNewIDSafe()
	s.dataPlace[place.ID] = &place
	return place, nil
}
