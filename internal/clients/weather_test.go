package clients

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCurrentTemperature_Success(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"current": {"temp_c": 25.7, "temp_f": 89.5}} `)
	}))
	defer mockServer.Close()

	os.Setenv("WEATHER_BASE_URL", mockServer.URL)
	os.Setenv("WEATHER_API_KEY", "**secret**")
	defer os.Unsetenv("WEATHER_BASE_URL")
	defer os.Unsetenv("WEATHER_API_KEY")

	currentWeather, err := GetCurrentWeather(-25, -55)

	assert.NoError(t, err)
	assert.Equal(t, 25.7, currentWeather.Temp_c)
	assert.Equal(t, 89.5, currentWeather.Temp_f)
	assert.Equal(t, 25.7+273, currentWeather.Temp_k)
}

func TestGetCurrentTemperature_InternalServerError(t *testing.T) {
	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Internal server", http.StatusInternalServerError)
	}))
	defer mockServer.Close()

	os.Setenv("WEATHER_BASE_URL", mockServer.URL)
	os.Setenv("WEATHER_API_KEY", "**secret**")
	defer os.Unsetenv("WEATHER_BASE_URL")
	defer os.Unsetenv("WEATHER_API_KEY")

	currentWeather, err := GetCurrentWeather(-25, -55)

	assert.Nil(t, currentWeather)
	assert.Error(t, err)
}
