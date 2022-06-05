package store

import "github.com/tmrrwnxtsn/const-payments-api/internal/model"

const userTable = "users"

// UserRepository представляет таблицу сущности "Пользователь" в базе данных.
type UserRepository interface {
	// Create создаёт пользователя в базе данных.
	Create(user model.User) (uint64, error)
}

type userRepository struct {
	store Store
}

func NewUserRepository(store Store) UserRepository {
	return &userRepository{store: store}
}

func (u *userRepository) Create(user model.User) (uint64, error) {
	panic("implement me")
}
