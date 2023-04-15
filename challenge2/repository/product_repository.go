package repository

import (
	"chap3-challenge2/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (pr *ProductRepository) Add(product model.Product) (model.Product, error) {
	tx := pr.db.Create(&product)
	return product, tx.Error
}

func (pr *ProductRepository) Get() ([]model.Product, error) {
	product := make([]model.Product, 0)

	tx := pr.db.Find(&product)
	return product, tx.Error
}

func (pr *ProductRepository) GetByUserId(userId string) ([]model.Product, error) {
	product := make([]model.Product, 0)

	tx := pr.db.Where("user_id = ?", userId).Find(&product)
	return product, tx.Error
}

func (pr *ProductRepository) GetOne(id string) (model.Product, error) {
	product := model.Product{}

	tx := pr.db.First(&product, "id = ?", id)

	return product, tx.Error
}

func (pr *ProductRepository) UpdateOne(updateProduct model.Product, id string) (model.Product, error) {

	tx := pr.db.
		Clauses(clause.Returning{
			Columns: []clause.Column{
				{Name: "id"},
				{Name: "user_id"},
				{Name: "created_at"},
				{Name: "updated_at"},
			},
		},
		).
		Where("id = ?", id).
		Updates(&updateProduct)
	return updateProduct, tx.Error
}

func (pr *ProductRepository) DeleteOne(deleteProduct model.Product, id string) error {
	tx := pr.db.Delete(&deleteProduct, "id = ?", id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}
