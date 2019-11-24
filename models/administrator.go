package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	. "github.com/wubingwei/veigar/databases"
	"github.com/wubingwei/veigar/dbmodel"
)

type administrator struct{}

func (a *administrator) Create(admin *dbmodel.Administrator) error {
	return GetDB().Create(admin).Error
}

func (a *administrator) UpdateEmail(id int, email string) error {
	return GetDB().Model(&dbmodel.Administrator{}).Where("id = ?", id).Update("email", email).Error
}

func (a *administrator) Find(id int) (*dbmodel.Administrator, error) {
	admin := new(dbmodel.Administrator)
	err := GetDB().First(admin, id).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, errors.New("cannot find this administrator, please make sure it created")
		}
		return nil, err
	}
	return admin, nil
}

func (a *administrator) Delete(id int) error {
	return GetDB().Where("id = ?", id).Delete(&dbmodel.Administrator{}).Error
}
