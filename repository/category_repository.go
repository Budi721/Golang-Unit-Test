package repository

import "github.com/Budi721/Golang-Unit-Test/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
