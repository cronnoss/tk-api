package internalhttp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"time"

	"github.com/cronnoss/tk-api/internal/common/srv"
	"github.com/cronnoss/tk-api/internal/model"
	"github.com/cronnoss/tk-api/internal/server"
	"github.com/cronnoss/tk-api/internal/storage/models"
	"github.com/gorilla/mux"
)

type ctxKeyID int

const (
	KeyLoggerID ctxKeyID = iota
)

type Server struct {
	srv  http.Server
	app  server.Application
	log  Logger
	host string
	port string
}

type Logger interface {
	Fatalf(format string, a ...interface{})
	Errorf(format string, a ...interface{})
	Warningf(format string, a ...interface{})
	Infof(format string, a ...interface{})
	Debugf(format string, a ...interface{})
}

func NewServer(log Logger, app server.Application, host, port string) *Server {
	return &Server{log: log, app: app, host: host, port: port}
}

func (s *Server) helperDecode(stream io.ReadCloser, w http.ResponseWriter, data interface{}) error {
	decoder := json.NewDecoder(stream)
	if err := decoder.Decode(&data); err != nil {
		s.log.Errorf("Can't decode json:%v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("{\"error\": \"Can't decode json:%v\"}\n", err)))
		return err
	}
	return nil
}

func (s *Server) GetShows(w http.ResponseWriter, r *http.Request) {
	// Step 1: Make a GET request to the remote API
	remoteURL := "https://leadbook.ru/test-task-api/shows"
	req, err := http.NewRequestWithContext(r.Context(), http.MethodGet, remoteURL, nil)
	if err != nil {
		srv.RespondWithError(fmt.Errorf("failed to create request: %w", err), w, r)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		srv.RespondWithError(fmt.Errorf("failed to do request: %w", err), w, r)
		return
	}
	defer resp.Body.Close()

	// Step 2: Decode the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		srv.RespondWithError(fmt.Errorf("failed to read response body: %w", err), w, r)
		return
	}

	var showListResponse model.ShowListResponse
	if err := json.Unmarshal(body, &showListResponse); err != nil {
		srv.RespondWithError(fmt.Errorf("failed to decode response: %w", err), w, r)
		return
	}

	// Step 3: Iterate over shows and store them in the local service
	for _, show := range showListResponse.Response {
		_, err := s.app.CreateShow(r.Context(), models.Show{
			ID:   show.ID,
			Name: show.Name,
		})
		if err != nil {
			srv.RespondWithError(fmt.Errorf("failed to create show: %w", err), w, r)
			return
		}
	}

	srv.RespondOK(showListResponse.Response, w, r)
}

func (s *Server) GetEvents(w http.ResponseWriter, r *http.Request) {
	// Step 1: Make a GET request to the remote API
	vars := mux.Vars(r)
	id := vars["id"]
	remoteURL := "https://leadbook.ru/test-task-api/shows/" + id + "/events"
	req, err := http.NewRequestWithContext(r.Context(), http.MethodGet, remoteURL, nil)
	if err != nil {
		srv.RespondWithError(fmt.Errorf("failed to create request: %w", err), w, r)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		srv.RespondWithError(fmt.Errorf("failed to do request: %w", err), w, r)
		return
	}
	defer resp.Body.Close()

	// Step 2: Decode the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		srv.RespondWithError(fmt.Errorf("failed to read response body: %w", err), w, r)
		return
	}

	var eventListResponse model.EventListResponse
	if err := json.Unmarshal(body, &eventListResponse); err != nil {
		srv.RespondWithError(fmt.Errorf("failed to decode response: %w", err), w, r)
		return
	}

	// Step 3: Iterate over events and store them in the local service
	for _, event := range eventListResponse.Response {
		_, err := s.app.CreateEvent(r.Context(), models.Event{
			ID:     event.ID,
			ShowID: event.ShowID,
			Date:   event.Date,
		})
		if err != nil {
			srv.RespondWithError(fmt.Errorf("failed to create event: %w", err), w, r)
			return
		}
	}

	srv.RespondOK(eventListResponse.Response, w, r)
}

func (s *Server) GetPlaces(w http.ResponseWriter, r *http.Request) {
	// Step 1: Make a GET request to the remote API
	vars := mux.Vars(r)
	id := vars["id"]
	remoteURL := "https://leadbook.ru/test-task-api/events/" + id + "/places"
	req, err := http.NewRequestWithContext(r.Context(), http.MethodGet, remoteURL, nil)
	if err != nil {
		srv.RespondWithError(fmt.Errorf("failed to create request: %w", err), w, r)
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		srv.RespondWithError(fmt.Errorf("failed to do request: %w", err), w, r)
		return
	}
	defer resp.Body.Close()

	// Step 2: Decode the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		srv.RespondWithError(fmt.Errorf("failed to read response body: %w", err), w, r)
		return
	}

	var placeListResponse model.PlaceListResponse
	if err := json.Unmarshal(body, &placeListResponse); err != nil {
		srv.RespondWithError(fmt.Errorf("failed to decode response: %w", err), w, r)
		return
	}

	// Step 3: Iterate over places and store them in the local service
	for _, place := range placeListResponse.Response {
		_, err := s.app.CreatePlace(r.Context(), models.Place{
			ID:          place.ID,
			X:           place.X,
			Y:           place.Y,
			Width:       place.Width,
			Height:      place.Height,
			IsAvailable: place.IsAvailable,
		})
		if err != nil {
			srv.RespondWithError(fmt.Errorf("failed to create place: %w", err), w, r)
			return
		}
	}

	srv.RespondOK(placeListResponse.Response, w, r)
}

func (s *Server) Start(ctx context.Context) error {
	addr := net.JoinHostPort(s.host, s.port)
	midLogger := NewMiddlewareLogger()

	router := mux.NewRouter()

	router.Handle("/healthz", midLogger.setCommonHeadersMiddleware(
		midLogger.loggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK healthz\n"))
		}))))

	router.Handle("/readiness", midLogger.setCommonHeadersMiddleware(
		midLogger.loggingMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("OK readiness\n"))
		}))))

	router.Handle("/shows", midLogger.setCommonHeadersMiddleware(
		midLogger.loggingMiddleware(http.HandlerFunc(s.GetShows))))
	router.Handle("/shows/{id:[0-9]+}/events", midLogger.setCommonHeadersMiddleware(
		midLogger.loggingMiddleware(http.HandlerFunc(s.GetEvents))))
	router.Handle("/events/{id:[0-9]+}/places", midLogger.setCommonHeadersMiddleware(
		midLogger.loggingMiddleware(http.HandlerFunc(s.GetPlaces))))

	s.srv = http.Server{
		Addr:              addr,
		Handler:           router,
		ReadHeaderTimeout: 2 * time.Second,
		BaseContext: func(_ net.Listener) context.Context {
			bCtx := context.WithValue(ctx, KeyLoggerID, s.log)
			return bCtx
		},
	}

	s.log.Infof("http server started on %s:%s\n", s.host, s.port)
	return s.srv.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	if err := s.srv.Shutdown(ctx); err != nil {
		return err
	}
	s.log.Infof("http server shutdown\n")
	return nil
}
