package usecase

import "test/interanl/core/entity"

func (usecase *useCase) UpdateUser(id string, user *entity.User) error {
	err := usecase.userRepository.UpdateUser(id, *user)
	if err != nil {
		return err
	}
	return nil
}
