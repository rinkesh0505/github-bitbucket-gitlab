package catalog

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mytheresa/go-hiring-challenge/models"
	"github.com/stretchr/testify/assert"
)

func TestCatalogService_GetProducts(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name    string
		setup   func(repo *models.MockProductRepository)
		request ListRequest
		expect  []Product
		total   int64
		err     error
	}{
		{
			name: "success",
			setup: func(repo *models.MockProductRepository) {
				repo.EXPECT().GetProducts(gomock.Any()).Return([]models.Product{{Code: "P1"}}, int64(1), nil)
			},
			request: ListRequest{},
			expect:  []Product{{Code: "P1"}},
			total:   1,
			err:     nil,
		},
		{
			name: "error",
			setup: func(repo *models.MockProductRepository) {
				repo.EXPECT().GetProducts(gomock.Any()).Return(nil, int64(0), errors.New("error"))
			},
			request: ListRequest{},
			expect:  nil,
			total:   0,
			err:     errors.New("error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := models.NewMockProductRepository(ctrl)
			tt.setup(repo)
			svc := NewCatalogService(repo)

			result, total, err := svc.GetProducts(tt.request)

			assert.Equal(t, tt.expect, result)
			assert.Equal(t, tt.total, total)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestCatalogService_GetProductDetails(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name   string
		setup  func(repo *models.MockProductRepository)
		code   string
		expect *ProductDetails
		err    error
	}{
		{
			name: "success",
			setup: func(repo *models.MockProductRepository) {
				repo.EXPECT().GetProductByCode("P1").Return(&models.Product{Code: "P1"}, nil)
			},
			code:   "P1",
			expect: &ProductDetails{Code: "P1", Variants: []VariantDetails{}},
			err:    nil,
		},
		{
			name: "error",
			setup: func(repo *models.MockProductRepository) {
				repo.EXPECT().GetProductByCode("P1").Return(nil, errors.New("error"))
			},
			code:   "P1",
			expect: nil,
			err:    errors.New("error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := models.NewMockProductRepository(ctrl)
			tt.setup(repo)
			svc := NewCatalogService(repo)

			result, err := svc.GetProductDetails(tt.code)

			assert.Equal(t, tt.expect, result)
			assert.Equal(t, tt.err, err)
		})
	}
}
