package repository

import (
	"gorm.io/gorm"
	"test/interanl/core/entity"
	"test/interanl/core/interfaces"
)

type pgUserRepository struct {
	DB gorm.DB
}

func NewPgUserRepository(
	db *gorm.DB,
) interfaces.UserRepository {
	return &pgUserRepository{
		DB: *db,
	}
}

func (r *pgUserRepository) CreateUser(user *entity.User) error {
	_ = r.DB.Create(user)
	return nil
}

func (r *pgUserRepository) GetUser(id string) (entity.User, error) {
	user := entity.User{}

	DB := r.DB.Where("id = ?", id).First(&user)
	if DB.Error != nil {
		return user, DB.Error
	}

	return user, nil
}

func (r *pgUserRepository) UpdateUser(id string, newUser entity.User) error {
	oldUser, err := r.GetUser(id)
	if err != nil {
		return err
	}
	DB := r.DB.Model(&oldUser).Updates(newUser)
	if DB.Error != nil {
		return DB.Error
	}
	return nil
}
