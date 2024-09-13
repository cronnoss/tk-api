package model

type ShowResponse struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type ShowListResponse struct {
	Response []ShowResponse `json:"response"`
}

type EventResponse struct {
	ID     int64  `json:"id"`
	ShowID int64  `json:"showId"`
	Date   string `json:"date"`
}

type EventListResponse struct {
	Response []EventResponse `json:"response"`
}

type PlaceResponse struct {
	ID          int64   `json:"id"`
	X           float64 `json:"x"`
	Y           float64 `json:"y"`
	Width       float64 `json:"width"`
	Height      float64 `json:"height"`
	IsAvailable bool    `json:"is_available"` // nolint64: tagliatelle
}

type PlaceListResponse struct {
	Response []PlaceResponse `json:"response"`
}
