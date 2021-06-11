package user

import (
	"backend/entity"

	"gorm.io/gorm"
)

type Repository interface {
	RShowAllUser() ([]entity.User, error)
	RRegisterUser(user entity.User) (entity.User, error)
	RFindUserByID(ID string) (entity.User, error)
	RDeleteUserByID(ID string) (string, error)
	RUpdateUserByID(ID string, dataUpdate map[string]interface{}) (entity.User, error)
	RFindUserByEmail(email string) (entity.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) RShowAllUser() ([]entity.User, error) {
	var Users []entity.User

	err := r.db.Find(&Users).Error
	if err != nil {
		return Users, err
	}

	return Users, nil
}

func (r *repository) RRegisterUser(user entity.User) (entity.User, error) {

	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) RFindUserByID(ID string) (entity.User, error) {
	var user entity.User

	err := r.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) RFindUserByEmail(email string) (entity.User, error) {
	var user entity.User

	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) RDeleteUserByID(ID string) (string, error) {
	if err := r.db.Where("id = ?", ID).Delete(&entity.User{}).Error; err != nil {
		return "error", err
	}

	return "success", nil
}

func (r *repository) RUpdateUserByID(ID string, dataUpdate map[string]interface{}) (entity.User, error) {

	var user entity.User

	if err := r.db.Model(&user).Where("id = ?", ID).Updates(dataUpdate).Error; err != nil {
		return user, err
	}

	if err := r.db.Where("id = ?", ID).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
