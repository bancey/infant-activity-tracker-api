package routes

import (
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	recorder := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/v1/ping/", nil)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 200, recorder.Code)
	assert.Equal(t, "pong", recorder.Body.String())
}
