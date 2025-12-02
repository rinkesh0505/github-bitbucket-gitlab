package catalog

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCatalogHandler_HandleGet(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name       string
		setup      func(service *MockCatalogService)
		query      string
		expectCode int
		expectBody string
	}{
		{
			name: "success",
			setup: func(service *MockCatalogService) {
				service.EXPECT().GetProducts(gomock.Any()).Return([]Product{{Code: "P1", Price: 0}}, int64(1), nil)
			},
			query:      "offset=0&limit=10",
			expectCode: http.StatusOK,
			expectBody: `{"products":[{"code":"P1","price":0}],"total":1}`,
		},
		{
			name: "error",
			setup: func(service *MockCatalogService) {
				service.EXPECT().GetProducts(gomock.Any()).Return(nil, int64(0), errors.New("error"))
			},
			query:      "offset=0&limit=10",
			expectCode: http.StatusInternalServerError,
			expectBody: `{"error":"error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewMockCatalogService(ctrl)
			tt.setup(service)
			handler := NewCatalogHandler(service)

			r := httptest.NewRequest(http.MethodGet, "/catalog?"+tt.query, nil)
			w := httptest.NewRecorder()

			handler.HandleGet(w, r)

			assert.Equal(t, tt.expectCode, w.Code)
			assert.JSONEq(t, tt.expectBody, w.Body.String())
		})
	}
}

func TestCatalogHandler_HandleGetDetails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name       string
		setup      func(service *MockCatalogService)
		url        string
		expectCode int
		expectBody string
	}{
		{
			name: "success",
			setup: func(service *MockCatalogService) {
				service.EXPECT().GetProductDetails("P1").Return(&ProductDetails{Code: "P1", Price: 0, Variants: nil}, nil)
			},
			url:        "/catalog/P1",
			expectCode: http.StatusOK,
			expectBody: `{"code":"P1","price":0,"variants":null}`,
		},
		{
			name: "not found",
			setup: func(service *MockCatalogService) {
				service.EXPECT().GetProductDetails("P1").Return(nil, errors.New("not found"))
			},
			url:        "/catalog/P1",
			expectCode: http.StatusNotFound,
			expectBody: `{"error":"not found"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewMockCatalogService(ctrl)
			tt.setup(service)
			handler := NewCatalogHandler(service)

			r := httptest.NewRequest(http.MethodGet, tt.url, nil)
			w := httptest.NewRecorder()

			handler.HandleGetDetails(w, r)

			assert.Equal(t, tt.expectCode, w.Code)
			assert.JSONEq(t, tt.expectBody, w.Body.String())
		})
	}
}
