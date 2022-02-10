package usecase

import (
	"test/interanl/core/entity"
)

func (usecase *useCase) CreateUser(user *entity.User) error {
	err := usecase.userRepository.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}
