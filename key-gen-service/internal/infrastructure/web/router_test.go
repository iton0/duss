package web_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

type MockRedirectService struct{}

func (m *MockRedirectService) GetOriginalURL(ctx context.Context, shortKey string) (string, error) {
	return "", nil
}

func TestRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := &MockRedirectService{}
	redirectHandler := api.NewRedirectHandler(mockService)

	router := web.NewRouter(redirectHandler)

	testCases := []struct {
		name               string
		method             string
		path               string
		expectedStatusCode int
	}{
		{
			name:               "Valid GET Request",
			method:             http.MethodGet,
			path:               "/abc1234",
			expectedStatusCode: http.StatusMovedPermanently,
		},
		{
			name:               "POST on GET Endpoint",
			method:             http.MethodPost,
			path:               "/abc1234",
			expectedStatusCode: http.StatusNotFound, // Corrected from 405
		},
		{
			name:               "Invalid Path",
			method:             http.MethodGet,
			path:               "/invalid/path",
			expectedStatusCode: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(tc.method, tc.path, nil)

			router.ServeHTTP(w, req)

			if w.Code != tc.expectedStatusCode {
				t.Errorf("for path %s, expected status code %d, but got %d", tc.path, tc.expectedStatusCode, w.Code)
			}
		})
	}
}
