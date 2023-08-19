package repository

import "app-golang-unit-test/entity"

type CategoryRepository interface {
	/// buat kontrak fungsi FindById dimana parameternya id dan balikannya berupa pointer kategori entity
	FindById(id string) *entity.Category
}