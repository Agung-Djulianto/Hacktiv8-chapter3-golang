package repository

import (
	"chap3-challenge2/model"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{
		db: db,
	}
}

func (ur *UserRepository) Save(user model.User) (model.User, error) {
	newUser := model.User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
		Role:     user.Role,
	}
	tx := ur.db.Create(&newUser)
	if tx.Error != nil {
		return model.User{}, tx.Error
	}
	return newUser, nil
}

func (ur *UserRepository) GetByEmail(email string) (model.User, error) {
	user := model.User{}
	tx := ur.db.First(&user, "email = ?", email)
	return user, tx.Error
}
