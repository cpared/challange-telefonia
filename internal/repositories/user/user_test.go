package repository

import (
	"testing"
)

func TestGetUser(t *testing.T) {
    repo := NewUserRepository()

    tests := []struct {
        phoneNumber string
        expected    User
        found       bool
    }{
        {
            phoneNumber: "+12025550119",
            expected: User{
                Address:     "123 Main St, Anytown, USA",
                Name:        "Alice Johnson",
                PhoneNumber: "+12025550119",
                Friends:     []string{"+12025550120", "+12025550121"},
            },
            found: true,
        },
        {
            phoneNumber: "+12025550129",
            expected:    User{},
            found:       false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.phoneNumber, func(t *testing.T) {
            user, found := repo.GetUser(tt.phoneNumber)
            if found != tt.found {
                t.Errorf("expected found %v, got %v", tt.found, found)
            }
            if user.Address != tt.expected.Address {
                t.Errorf("expected Address %v, got %v", tt.expected.Address, user.Address)
            }
			if user.Name != tt.expected.Name {
                t.Errorf("expected Name %v, got %v", tt.expected.Name, user.Name)
            }
			if user.PhoneNumber != tt.expected.PhoneNumber {
                t.Errorf("expected PhoneNumber %v, got %v", tt.expected.PhoneNumber, user.PhoneNumber)
            }
        })
    }
}