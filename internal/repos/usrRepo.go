package repos

import "gorm.io/gorm"

type UsrRepo interface {
}

type usrRepo struct {
	db *gorm.DB
}

func NewUsrRepo(db *gorm.DB) UsrRepo {
	return &usrRepo{db: db}
}
