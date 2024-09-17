package internalhttp

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/cronnoss/tk-api/internal/model"
	"github.com/cronnoss/tk-api/internal/server/mocks"
	"github.com/cronnoss/tk-api/internal/storage/models"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestGetShows(t *testing.T) {
	showServiceMock := mocks.NewApplication(t)
	showServiceMock.On("CreateShow", mock.Anything, mock.Anything).Return(models.Show{}, nil)

	_, err := showServiceMock.CreateShow(context.Background(), models.Show{})
	require.NoError(t, err)

	// Step 1: Mock the remote API
	mockResponse := `{
	"response": [
		{"id": 1, "name": "Show #1"},
		{"id": 2, "name": "Show #2"}
	]
}`
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte(mockResponse))
	},
	))

	// Step 2: Make a GET request to the remote API
	req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, mockServer.URL, nil)
	require.NoError(t, err)

	resp, err := http.DefaultClient.Do(req)
	require.NoError(t, err)

	defer resp.Body.Close()

	// Step 3: Read and decode the remote API response
	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	var showListResponse model.ShowListResponse
	err = json.Unmarshal(body, &showListResponse)
	require.NoError(t, err)

	require.Equal(t, 2, len(showListResponse.Response))
	require.Equal(t, int64(1), showListResponse.Response[0].ID)
	require.Equal(t, "Show #1", showListResponse.Response[0].Name)
	require.Equal(t, int64(2), showListResponse.Response[1].ID)
	require.Equal(t, "Show #2", showListResponse.Response[1].Name)
}
