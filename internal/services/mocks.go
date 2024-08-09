package services

import repository "challange/internal/repositories/user"

type mockUserRepository struct {
	users map[string]repository.User
}

func (m *mockUserRepository) GetUser(phoneNumber string) (repository.User, bool) {
	user, found := m.users[phoneNumber]
	return user, found
}

type mockCallRepository struct {
	calls map[string][]repository.Call
}

func (m *mockCallRepository) GetCalls(phoneNumber string) []repository.Call {
	return m.calls[phoneNumber]
}