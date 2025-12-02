package category

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/mytheresa/go-hiring-challenge/models"
	"github.com/stretchr/testify/assert"
)

func TestCategoryService_GetAllCategories(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name   string
		setup  func(repo *models.MockCategoryRepository)
		expect []Category
		err    error
	}{
		{
			name: "success",
			setup: func(repo *models.MockCategoryRepository) {
				repo.EXPECT().GetAllCategories().Return([]models.Category{{Code: "C1"}}, nil)
			},
			expect: []Category{{Code: "C1"}},
			err:    nil,
		},
		{
			name: "error",
			setup: func(repo *models.MockCategoryRepository) {
				repo.EXPECT().GetAllCategories().Return(nil, errors.New("error"))
			},
			expect: nil,
			err:    errors.New("error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := models.NewMockCategoryRepository(ctrl)
			tt.setup(repo)
			svc := NewCategoryService(repo)

			result, err := svc.GetAllCategories()

			assert.Equal(t, tt.expect, result)
			assert.Equal(t, tt.err, err)
		})
	}
}

func TestCategoryService_CreateCategory(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	tests := []struct {
		name  string
		setup func(repo *models.MockCategoryRepository)
		input Category
		err   error
	}{
		{
			name: "success",
			setup: func(repo *models.MockCategoryRepository) {
				repo.EXPECT().CreateCategory(gomock.Any()).Return(nil)
			},
			input: Category{Code: "C1"},
			err:   nil,
		},
		{
			name: "error",
			setup: func(repo *models.MockCategoryRepository) {
				repo.EXPECT().CreateCategory(gomock.Any()).Return(errors.New("error"))
			},
			input: Category{Code: "C1"},
			err:   errors.New("error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := models.NewMockCategoryRepository(ctrl)
			tt.setup(repo)
			svc := NewCategoryService(repo)

			err := svc.CreateCategory(tt.input)

			assert.Equal(t, tt.err, err)
		})
	}
}
