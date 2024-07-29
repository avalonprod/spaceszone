package store

import "context"

type UsersStore struct {
}

func NewUsersStore() *UsersStore {
	return &UsersStore{}
}

func (s *UsersStore) Save() error {
	return nil
}

func (s *UsersStore) Get(ctx context.Context, id string) error {
	return nil
}

func (s *UsersStore) Update(ctx context.Context, id string) error {
	return nil
}

func (s *UsersStore) SoftDelete(ctx context.Context, id string) error {
	return nil
}
