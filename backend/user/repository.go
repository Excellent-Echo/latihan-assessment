package user

import (
	"gorm.io/gorm"
	"latihan-assessment/entity"
)

type Repository interface {
	GetAllUsers() ([]entity.User, error)
	CreateUsers(user entity.User) (entity.User, error)
	GetUserById(ID string) (entity.User, error)
	DeleteUserById(ID string) (string, error)
	UpdateUserById(ID string, userUpdate map[string]interface{}) (entity.User, error)
	GetUserByEmail(email string) (entity.User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) GetAllUsers() ([]entity.User, error) {
	var users []entity.User

	if err := r.db.Find(&users).Error; err != nil {
		return users, nil
	}

	return users, nil
}

func (r *repository) CreateUsers(user entity.User) (entity.User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, nil
	}

	return user, nil
}

func (r *repository) GetUserById(ID string) (entity.User, error) {
	var user entity.User

	if err := r.db.Where("id = ?", ID).Find(user).Error; err != nil {
		return user, nil
	}

	return user, nil
}

func (r *repository) DeleteUserById(ID string) (string, error) {
	if err := r.db.Where("id = ?", ID).Delete(&entity.User{}).Error; err != nil {
		return "error", nil
	}

	return "success", nil
}

func (r *repository) UpdateUserById(ID string, userUpdate map[string]interface{}) (entity.User, error) {
	var userWillUpdate entity.User

	if err := r.db.Model(&userWillUpdate).Where("id = ?", ID).Updates(userUpdate); err != nil {
		return userWillUpdate, nil
	}

	if err := r.db.Where("id = ?", ID).Find(userWillUpdate).Error; err != nil {
		return userWillUpdate, nil
	}

	return userWillUpdate, nil

}

//func (r *repository) GetUserByEmail(email string) (entity.User, error) {
//
//}
