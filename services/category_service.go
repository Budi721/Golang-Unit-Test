package services

import (
	"errors"
	"github.com/Budi721/Golang-Unit-Test/entity"
	"github.com/Budi721/Golang-Unit-Test/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if category == nil {
		return category, errors.New("Category not found")
	} else {
		return category, nil
	}
}