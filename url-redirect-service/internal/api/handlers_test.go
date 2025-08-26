package api_test

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/iton0/duss/url-redirect-service/internal/api"
	"github.com/iton0/duss/url-redirect-service/internal/core/services"
)

type MockRedirectService struct {
	ReturnURL string
	ReturnErr error
}

func (m *MockRedirectService) GetOriginalURL(ctx context.Context, shortKey string) (string, error) {
	return m.ReturnURL, m.ReturnErr
}

func TestHandleRedirect(t *testing.T) {
	gin.SetMode(gin.TestMode)

	testCases := []struct {
		name                string
		shortKey            string
		mockReturnURL       string
		mockReturnErr       error
		expectedStatusCode  int
		expectedRedirectURL string
	}{
		{
			name:                "Success - Valid Redirect",
			shortKey:            "testkey",
			mockReturnURL:       "https://example.com/long/url",
			mockReturnErr:       nil,
			expectedStatusCode:  http.StatusMovedPermanently,
			expectedRedirectURL: "https://example.com/long/url",
		},
		{
			name:                "Not Found Error",
			shortKey:            "nonexistent",
			mockReturnURL:       "",
			mockReturnErr:       services.ErrURLNotFound,
			expectedStatusCode:  http.StatusNotFound,
			expectedRedirectURL: "",
		},
		{
			name:                "Internal Server Error",
			shortKey:            "badkey",
			mockReturnURL:       "",
			mockReturnErr:       errors.New("database connection failed"),
			expectedStatusCode:  http.StatusInternalServerError,
			expectedRedirectURL: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			mockService := &MockRedirectService{
				ReturnURL: tc.mockReturnURL,
				ReturnErr: tc.mockReturnErr,
			}

			redirectHandler := api.NewRedirectHandler(mockService)

			w := httptest.NewRecorder()
			req, _ := http.NewRequest(http.MethodGet, "/"+tc.shortKey, nil)

			c, _ := gin.CreateTestContext(w)
			c.Request = req

			c.Params = gin.Params{
				{Key: "shortKey", Value: tc.shortKey},
			}

			redirectHandler.HandleRedirect(c)

			if w.Code != tc.expectedStatusCode {
				t.Errorf("expected status code %d, but got %d", tc.expectedStatusCode, w.Code)
			}

			if tc.expectedStatusCode == http.StatusMovedPermanently {
				locationHeader := w.Header().Get("Location")
				if locationHeader != tc.expectedRedirectURL {
					t.Errorf("expected redirect URL %s, but got %s", tc.expectedRedirectURL, locationHeader)
				}
			}
		})
	}
}
