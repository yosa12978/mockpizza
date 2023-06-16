package repos

import (
	"github.com/yosa12978/mockPizza/internal/domain"
	"gorm.io/gorm"
)

type UsrRepo interface {
	FindById(id uint) (domain.Usr, error)
	FindByUsername(username string) (domain.Usr, error)
	Create(usr domain.Usr) (domain.Usr, error)
	Delete(id uint) (domain.Usr, error)
	Update(usr domain.Usr) (domain.Usr, error)
}

type usrRepo struct {
	db *gorm.DB
}

func NewUsrRepo(db *gorm.DB) UsrRepo {
	return &usrRepo{db: db}
}

func (repo *usrRepo) FindById(id uint) (domain.Usr, error) {
	var usr domain.Usr
	err := repo.db.First(&usr, "id = ?", id).Error
	return usr, err
}

func (repo *usrRepo) FindByUsername(username string) (domain.Usr, error) {
	var usr domain.Usr
	err := repo.db.First(&usr, "username = ?", username).Error
	return usr, err
}

func (repo *usrRepo) Create(usr domain.Usr) (domain.Usr, error) {
	return usr, repo.db.Create(&usr).Error
}

func (repo *usrRepo) Delete(id uint) (domain.Usr, error) {
	var usr domain.Usr
	err := repo.db.First(&usr, "id = ?", id).Error
	if err != nil {
		return usr, err
	}
	return usr, repo.db.Delete(&usr).Error
}

func (repo *usrRepo) Update(usr domain.Usr) (domain.Usr, error) {
	err := repo.db.Save(&usr).Error
	return usr, err
}
