package services

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	repository "challange/internal/repositories/user"
)

func TestCalculateInvoice(t *testing.T) {}

func TestIsInternational(t *testing.T) {
	service := NewUserService(nil, nil)
	tests := []struct {
		phoneNumber string
		call        string
		expected    bool
	}{
		{
			phoneNumber: "+13025550119",
			call:     "+12025550119",
			expected: true,
		},
		{
			phoneNumber: "+12025550119",
			call:     "+12025550129",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.call, func(t *testing.T) {
			if service.isInternational(tt.phoneNumber, tt.call) != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, service.isInternational(tt.phoneNumber, tt.call))
			}
		})
	}
}

func TestIsFriend(t *testing.T) {
	service := NewUserService(nil, nil)
	tests := []struct {
		friends []string
		friend  string
		expected bool
	}{
		{
			friends: []string{"+13025550119", "+12025550119"},
			friend:  "+12025550119",
			expected: true,
		},
		{
			friends: []string{"+13025550119", "+12025550119"},
			friend:  "+12025550129",
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.friend, func(t *testing.T) {
			if service.isFriend(tt.friends, tt.friend) != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, service.isFriend(tt.friends, tt.friend))
			}
		})
	}
}

func TestCalculateFare(t *testing.T) {
	service := NewUserService(nil, nil)
	tests := []struct {
		fare     float32
		free     bool
		expected float32
	}{
		{
			fare:     10,
			free:     true,
			expected: 0,
		},
		{
			fare:     10,
			free:     false,
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run("calculate fare", func(t *testing.T) {
			if service.calculateFare(tt.fare, tt.free) != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, service.calculateFare(tt.fare, tt.free))
			}
		})
	}
}

func TestCalculateCallsDuration(t *testing.T) {
	users := map[string]repository.User{
		"+1234567890": {
			PhoneNumber: "+1234567890",
			Friends:     []string{"+0987654321"},
		},
	}

	calls := map[string][]repository.Call{
		"national": {
			{
				Origin:      "+1234567890",
				Destination: "+1237654321",
				Duration:    300,
				TimeStamp:   time.Date(2024, 8, 1, 12, 34, 56, 0, time.UTC),
			},
		},
		"international": {
			{
				Origin:      "+1234567890",
				Destination: "+441234567890",
				Duration:    200,
				TimeStamp:   time.Date(2024, 8, 1, 15, 0, 0, 0, time.UTC),
			},
		},
		"friend_discount": {
			{
				Origin:      "+1234567890",
				Destination: "+0987654321",
				Duration:    100,
				TimeStamp:   time.Date(2024, 8, 1, 10, 0, 0, 0, time.UTC),
			},
		},
		"no_calls": {
			{
				Origin:      "+1234567890",
				Destination: "+1111111111",
				Duration:    300,
				TimeStamp:   time.Date(2023, 8, 3, 10, 0, 0, 0, time.UTC),
			},
		},
	}

	userRepo := &mockUserRepository{users: users}
	callRepo := &mockCallRepository{calls: map[string][]repository.Call{
		"+1234567890": calls["national"],
	}}

	service := NewUserService(userRepo, callRepo)

	dateBefore := time.Date(2024, 7, 31, 0, 0, 0, 0, time.UTC)
	dateAfter := time.Date(2024, 8, 2, 0, 0, 0, 0, time.UTC)

	tests := []struct {
		name               string
		calls              []repository.Call
		expectedTotal      float32
		expectedInternational int
		expectedNational   int
		expectedFriends    int
	}{
		{
			name:               "NationalCall",
			calls:              calls["national"],
			expectedTotal:      2.5,
			expectedInternational: 0,
			expectedNational:   300,
			expectedFriends:    0,
		},
		{
			name:               "InternationalCall",
			calls:              calls["international"],
			expectedTotal:      150.0,
			expectedInternational: 200,
			expectedNational:   0,
			expectedFriends:    0,
		},
		{
			name:               "FriendCallWithDiscount",
			calls:              calls["friend_discount"],
			expectedTotal:      0.0,
			expectedInternational: 100,
			expectedNational:   0,
			expectedFriends:    100,
		},
		{
			name:               "NoCallsInDateRange",
			calls:              calls["no_calls"],
			expectedTotal:      0.0,
			expectedInternational: 0,
			expectedNational:   0,
			expectedFriends:    0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			callRepo.calls["+1234567890"] = tt.calls

			total, international, national, friends, _ := service.calculateCallsDuration(users["+1234567890"], tt.calls, dateBefore, dateAfter)

			assert.Equal(t, tt.expectedTotal, total)
			assert.Equal(t, tt.expectedInternational, international)
			assert.Equal(t, tt.expectedNational, national)
			assert.Equal(t, tt.expectedFriends, friends)
		})
	}
}