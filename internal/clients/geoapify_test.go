package clients

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockGeoapifyClient struct {
	mock.Mock
}

func (m *MockGeoapifyClient) SearchCEPCoordinates(cep string) (*CEPCoordinates, error) {
	args := m.Called(cep)
	return args.Get(0).(*CEPCoordinates), args.Error(1)
}

func TestGeoapifyClient_Success(t *testing.T) {
	mockClient := new(MockGeoapifyClient)
	mockClient.On("SearchCEPCoordinates", "78590-000").Return(&CEPCoordinates{Lon: -55.51, Lat: -11.85}, nil)

	loc, err := mockClient.SearchCEPCoordinates("78590-000")

	assert.NoError(t, err)
	assert.Equal(t, -11.85, loc.Lat)
	assert.Equal(t, -55.51, loc.Lon)
}

func TestGeoapifyClient_StatusOk(t *testing.T) {
	mockData := &GeoapifyResponse{
		Results: []struct {
			Lon float64 `json:"lon"`
			Lat float64 `json:"lat"`
		}{
			{
				Lon: -55.51,
				Lat: -11.85,
			},
		},
	}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(mockData)
		if err != nil {
			t.Fatalf("Error on try to send mock response: %v", err)
		}
	}))
	defer mockServer.Close()

	os.Setenv("GEOAPIFY_BASE_URL", mockServer.URL)
	os.Setenv("GEOAPIFY_API_KEY", "**secret**")
	defer os.Unsetenv("GEOAPIFY_BASE_URL")
	defer os.Unsetenv("GEOAPIFY_API_KEY")

	coordinates, err := SearchCEPCoordinates("78590-000")

	assert.NoError(t, err)
	assert.Equal(t, -11.85, coordinates.Lat)
	assert.Equal(t, -55.51, coordinates.Lon)
}

func TestGeoapifyClient_StatusInternalServerError(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal server", http.StatusInternalServerError)
	}))
	defer mockServer.Close()

	os.Setenv("GEOAPIFY_BASE_URL", mockServer.URL)
	os.Setenv("GEOAPIFY_API_KEY", "**secret**")
	defer os.Unsetenv("GEOAPIFY_BASE_URL")
	defer os.Unsetenv("GEOAPIFY_API_KEY")

	coordinates, err := SearchCEPCoordinates("78590-000")

	assert.Nil(t, coordinates)
	assert.Error(t, err)
}

func TestGeoapifyClient_NothingToShow(t *testing.T) {
	mockData := &GeoapifyResponse{
		Results: []struct {
			Lon float64 `json:"lon"`
			Lat float64 `json:"lat"`
		}{},
	}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(mockData)
		if err != nil {
			t.Fatalf("Error on try to send mock response: %v", err)
		}
	}))
	defer mockServer.Close()

	os.Setenv("GEOAPIFY_BASE_URL", mockServer.URL)
	os.Setenv("GEOAPIFY_API_KEY", "**secret**")
	defer os.Unsetenv("GEOAPIFY_BASE_URL")
	defer os.Unsetenv("GEOAPIFY_API_KEY")

	coordinates, err := SearchCEPCoordinates("78590-000")

	assert.Nil(t, coordinates)
	assert.Error(t, err)
}

func TestGeoapifyClient_DecodeFails(t *testing.T) {
	mockData := struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	}{
		Lon: -55.51,
		Lat: -11.85,
	}

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		err := json.NewEncoder(w).Encode(mockData)
		if err != nil {
			t.Fatalf("Error on try to send mock response: %v", err)
		}
	}))
	defer mockServer.Close()

	os.Setenv("GEOAPIFY_BASE_URL", mockServer.URL)
	os.Setenv("GEOAPIFY_API_KEY", "**secret**")
	defer os.Unsetenv("GEOAPIFY_BASE_URL")
	defer os.Unsetenv("GEOAPIFY_API_KEY")

	coordinates, err := SearchCEPCoordinates("78590-000")

	assert.Nil(t, coordinates)
	assert.Error(t, err)
}
