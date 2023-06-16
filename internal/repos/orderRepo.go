package repos

import (
	"github.com/yosa12978/mockPizza/internal/domain"
	"gorm.io/gorm"
)

type OrderRepo interface {
	FindById(id uint) (domain.Order, error)
	Create(order domain.Order) (domain.Order, error)
	Delete(id uint) (domain.Order, error)
	Update(order domain.Order) (domain.Order, error)
}

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) OrderRepo {
	return &orderRepo{db: db}
}

func (repo *orderRepo) FindById(id uint) (domain.Order, error) {
	var order domain.Order
	err := repo.db.First(&order, "id = ?", id).Error
	return order, err
}

func (repo *orderRepo) Create(order domain.Order) (domain.Order, error) {
	return order, repo.db.Create(&order).Error
}

func (repo *orderRepo) Delete(id uint) (domain.Order, error) {
	var order domain.Order
	err := repo.db.First(&order, "id = ?", id).Error
	if err != nil {
		return order, err
	}
	return order, repo.db.Delete(&order).Error
}

func (repo *orderRepo) Update(order domain.Order) (domain.Order, error) {
	err := repo.db.Save(&order).Error
	return order, err
}
