package repo

import (
	"gitlab.ezrpro.in/godemo/internal/model"
	"gorm.io/gorm"
)

type mysqlProjectRepo struct {
	db *gorm.DB
}

// GetById implements IProjectRepo.
func (p *mysqlProjectRepo) GetById(id int) (res model.Project, err error) {

	err = p.db.Find(&res, id).Error
	return
}

func NewMysqlProjectRepo(db *gorm.DB) IProjectRepo {
	return &mysqlProjectRepo{
		db: db,
	}
}
