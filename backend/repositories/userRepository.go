package repositories

import (
	"book_list/entities"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user entities.User) (entities.User, error)
	UpdateUser(user entities.User) (entities.User, error)
	DeleteUserById(id int) (string, error)
	FindAllUsers() ([]entities.User, error)
	FindUserById(id int) (entities.User, error)
	FindUserByEmail(email string) (entities.User, error)
}

type userRepository struct {
	userDb *gorm.DB
}

func NewUserRepository(userDb *gorm.DB) *userRepository {
	return &userRepository{userDb}
}

func (r *userRepository) InsertUser(user entities.User) (entities.User, error) {
	if err := r.userDb.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) UpdateUserById(user entities.User) (entities.User, error) {
	if err := r.userDb.Save(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *userRepository) DeleteUserById(id int) (string, error) {
	if err := r.userDb.Where("user_id", id).Delete(&entities.User{}).Error; err != nil {
		return "Failed to delete user", err
	}

	return "Successfully delete user", nil
}

func (r *userRepository) FindAllUsers() ([]entities.User, error) {
	var usersData []entities.User
	if err := r.userDb.Find(&usersData).Error; err != nil {
		return usersData, err
	}

	return usersData, nil
}

func (r *userRepository) FindUserById(id int) (entities.User, error) {
	var userData entities.User
	if err := r.userDb.Where("user_id ?", id).Find(&userData).Error; err != nil {
		return userData, err
	}

	return userData, nil
}

func (r *userRepository) FindUserByEmail(userEmail string) (entities.User, error) {
	var userData entities.User
	if err := r.userDb.Where("email = ?", userEmail).Find(&userData).Error; err != nil {
		return userData, err
	}

	return userData, nil
}
