package userService

import (
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user Users) (Users, error)
	GetAllUsers() ([]Users, error)
	UpdateUserByID(id uint, user Users) (Users, error)
	DeleteUserByID(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}
func (r *userRepository) CreateUser(user Users) (Users, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return Users{}, result.Error
	}
	return user, nil
}
func (r *userRepository) GetAllUsers() ([]Users, error) {
	var users []Users
	err := r.db.Find(&users).Error
	return users, err
}
func (r *userRepository) UpdateUserByID(id uint, user Users) (Users, error) {
	var thatUser Users
	result := r.db.First(&thatUser, id)
	if result.Error != nil {
		return Users{}, result.Error
	}
	thatUser.Email = user.Email
	thatUser.Password = user.Password
	result = r.db.Save(&thatUser)
	if result.Error != nil {
		return Users{}, result.Error
	}
	return thatUser, nil
}
func (r *userRepository) DeleteUserByID(id uint) error {
	var user Users
	result := r.db.First(&user, id)
	if result.Error != nil {
		return result.Error
	}
	result = r.db.Delete(&user)
	return result.Error
}
