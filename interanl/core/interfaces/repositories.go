package interfaces

import (
	"test/interanl/core/entity"
)

type UserRepository interface {
	CreateUser(user *entity.User) error
	GetUser(id string) (entity.User, error)
	UpdateUser(id string, user entity.User) error
}
