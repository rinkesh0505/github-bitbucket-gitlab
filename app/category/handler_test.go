package category

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCategoryHandler_HandleGetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name       string
		setup      func(service *MockCategoryService)
		expectCode int
		expectBody string
	}{
		{
			name: "success",
			setup: func(service *MockCategoryService) {
				service.EXPECT().GetAllCategories().Return([]Category{{Code: "C1", Name: "Category1"}}, nil)
			},
			expectCode: http.StatusOK,
			expectBody: `[{"code":"C1","name":"Category1"}]`,
		},
		{
			name: "error",
			setup: func(service *MockCategoryService) {
				service.EXPECT().GetAllCategories().Return(nil, errors.New("failed to fetch categories"))
			},
			expectCode: http.StatusInternalServerError,
			expectBody: `{"error":"failed to fetch categories"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewMockCategoryService(ctrl)
			tt.setup(service)
			handler := NewCategoryHandler(service)

			r := httptest.NewRequest(http.MethodGet, "/categories", nil)
			w := httptest.NewRecorder()

			handler.HandleGetAll(w, r)

			assert.Equal(t, tt.expectCode, w.Code)
			assert.JSONEq(t, tt.expectBody, w.Body.String())
		})
	}
}

func TestCategoryHandler_HandleCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name       string
		setup      func(service *MockCategoryService)
		body       string
		expectCode int
		expectBody string
	}{
		{
			name: "success",
			setup: func(service *MockCategoryService) {
				service.EXPECT().CreateCategory(Category{Code: "C1", Name: "Category1"}).Return(nil)
			},
			body:       `{"code":"C1","name":"Category1"}`,
			expectCode: http.StatusOK,
			expectBody: `{"code":"C1","name":"Category1"}`,
		},
		{
			name:       "invalid body",
			setup:      func(service *MockCategoryService) {},
			body:       `invalid-body`,
			expectCode: http.StatusBadRequest,
			expectBody: `{"error":"invalid request body"}`,
		},
		{
			name:       "validation error",
			setup:      func(service *MockCategoryService) {},
			body:       `{"code":"","name":""}`,
			expectCode: http.StatusBadRequest,
			expectBody: `{"error":"code is required"}`,
		},
		{
			name: "service error",
			setup: func(service *MockCategoryService) {
				service.EXPECT().CreateCategory(Category{Code: "C1", Name: "Category1"}).Return(errors.New("service error"))
			},
			body:       `{"code":"C1","name":"Category1"}`,
			expectCode: http.StatusInternalServerError,
			expectBody: `{"error":"service error"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewMockCategoryService(ctrl)
			tt.setup(service)
			handler := NewCategoryHandler(service)

			r := httptest.NewRequest(http.MethodPost, "/categories", bytes.NewBufferString(tt.body))
			w := httptest.NewRecorder()

			handler.HandleCreate(w, r)

			assert.Equal(t, tt.expectCode, w.Code)
			assert.JSONEq(t, tt.expectBody, w.Body.String())
		})
	}
}
