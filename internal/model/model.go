package model

import "fmt"

type ShowResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type ShowListResponse struct {
	Response []ShowResponse `json:"response"`
}

func (s *ShowListResponse) ShowListResponseValidate() error {
	if len(s.Response) == 0 {
		return fmt.Errorf("empty response")
	}
	if s.Response[0].ID == 0 {
		return fmt.Errorf("%w: id", ErrRequired)
	}
	if s.Response[0].Name == "" {
		return fmt.Errorf("%w: name", ErrRequired)
	}
	return nil
}

type EventResponse struct {
	ID     int64  `json:"id"`
	ShowID int64  `json:"showId"`
	Date   string `json:"date"`
}

type EventListResponse struct {
	Response []EventResponse `json:"response"`
}

func (e *EventListResponse) EventListResponseValidate() error {
	if len(e.Response) == 0 {
		return fmt.Errorf("empty response")
	}
	if e.Response[0].ID == 0 {
		return fmt.Errorf("%w: id", ErrRequired)
	}
	if e.Response[0].ShowID == 0 {
		return fmt.Errorf("%w: showId", ErrRequired)
	}
	if e.Response[0].Date == "" {
		return fmt.Errorf("%w: date", ErrRequired)
	}
	return nil
}

type PlaceResponse struct {
	ID          int64   `json:"id"`
	X           float64 `json:"x"`
	Y           float64 `json:"y"`
	Width       float64 `json:"width"`
	Height      float64 `json:"height"`
	IsAvailable bool    `json:"is_available"` // nolint: tagliatelle
}

type PlaceListResponse struct {
	Response []PlaceResponse `json:"response"`
}

func (p *PlaceListResponse) PlaceListResponseValidate() error {
	if len(p.Response) == 0 {
		return fmt.Errorf("empty response")
	}
	if p.Response[0].ID == 0 {
		return fmt.Errorf("%w: id", ErrRequired)
	}
	if p.Response[0].X < 0 {
		return fmt.Errorf("%w: x", ErrNegative)
	}
	if p.Response[0].Y < 0 {
		return fmt.Errorf("%w: y", ErrNegative)
	}
	if p.Response[0].Width < 0 {
		return fmt.Errorf("%w: width", ErrNegative)
	}
	if p.Response[0].Height < 0 {
		return fmt.Errorf("%w: height", ErrNegative)
	}
	return nil
}
