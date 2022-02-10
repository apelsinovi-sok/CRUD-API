package usecase

import "test/interanl/core/interfaces"

type useCase struct {
	userRepository interfaces.UserRepository
}

func New(userRepository interfaces.UserRepository) interfaces.UseCase {
	return &useCase{userRepository: userRepository}
}
