package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/psenna/isup/api"
	"github.com/stretchr/testify/assert"
)

func TestHomeRoute(t *testing.T) {
	router := api.GetApiServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"message\":\"Everything is up in your life?\",\"system\":\"Is up is up!\"}", w.Body.String())
}
