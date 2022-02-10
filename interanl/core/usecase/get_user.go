package usecase

import "test/interanl/core/entity"

func (usecase *useCase) GetUser(id string) (entity.User, error) {
	user, err := usecase.userRepository.GetUser(id)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
