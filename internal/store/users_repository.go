package store

import "github.com/tmrrwnxtsn/const-payments-api/internal/model"

const usersTable = "users"

// UsersRepository представляет таблицу сущности "Пользователь" в базе данных.
type UsersRepository interface {
	// Create создаёт пользователя в базе данных.
	Create(user model.User) (uint64, error)
}

type usersRepository struct {
	store Store
}

func NewUsersRepository(store Store) UsersRepository {
	return &usersRepository{store: store}
}

func (u *usersRepository) Create(user model.User) (uint64, error) {
	panic("implement me")
}
