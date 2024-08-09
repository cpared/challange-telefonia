package services

import (
	"log"
	"os"
	"time"

	repository "challange/internal/repositories/user"
)

const (
	INTERNATIONAL_FARE  = 0.75
	NATIONAL_FARE       = 2.5
	FREE_AMOUNT_FRIENDS = 10
)

type UserRepository interface {
	GetUser(phoneNumber string) (repository.User, bool)
}

type CallRepository interface {
	GetCalls(phoneNumber string) []repository.Call
}

type UserService struct {
	usersRepo UserRepository
	callsRepo CallRepository
	logger    *log.Logger
}

func NewUserService(u UserRepository, c CallRepository) *UserService {
	return &UserService{
		usersRepo: u,
		callsRepo: c,
		logger:    log.New(os.Stdout, "", log.LstdFlags),
	}
}

// CalculateInvoice calculates the invoice for a user
// It returns a Response with the user information and the invoice details
// If the user is not found, it returns an empty Response
// If there is an error loading the calls, it returns an empty Response
func (u *UserService) CalculateInvoice(dateBefore, dateAfter time.Time, phoneNumber string) Response {
	user, found := u.usersRepo.GetUser(phoneNumber)
	if !found {
		u.logger.Printf("Phone: %s not found", phoneNumber)
		return Response{}
	}

	calls := u.callsRepo.GetCalls(phoneNumber)
	if len(calls) == 0 {
		u.logger.Printf("No calls found for phone: %s", phoneNumber)
		return Response{}
	}

	total, international, national, friends, userCalls := u.calculateCallsDuration(user, calls, dateBefore, dateAfter)

	return Response{
		User: User{
			Address:     user.Address,
			Name:        user.Name,
			PhoneNumber: user.PhoneNumber,
		},
		Calls:                    userCalls,
		TotalInternationaSeconds: international,
		TotalNationaSeconds:      national,
		TotalFriendSeconds:       friends,
		Total:                    total,
	}
}

// calculateCallsDuration calculates the total duration of the calls
// It returns the total duration, the total international duration, the total national duration and the total duration with friends
// It also calculates the total amount of the invoice
func (u *UserService) calculateCallsDuration(user repository.User, 
	calls []repository.Call, dateBefore, dateAfter time.Time) (float32, int, int, int, []Call) {

	international := 0
	national := 0
	friends := 0
	amount_friend_calls := 0
	userCalls := []Call{}
	total := float32(0)

	for _, call := range calls {
		if user.PhoneNumber != call.Origin || call.TimeStamp.Before(dateBefore) || call.TimeStamp.After(dateAfter) {
			continue
		}

		isFriend := u.isFriend(user.Friends, call.Destination)

		if isFriend {
			friends += call.Duration
			amount_friend_calls += 1
		}

		cost := float32(0)
		if u.isInternational(user.PhoneNumber, call.Destination) {
			international += call.Duration
			cost += u.calculateFare(INTERNATIONAL_FARE*float32(call.Duration), isFriend && amount_friend_calls <= FREE_AMOUNT_FRIENDS)
		} else {
			national += call.Duration
			cost += u.calculateFare(NATIONAL_FARE, isFriend && amount_friend_calls <= FREE_AMOUNT_FRIENDS)
		}

		userCalls = append(userCalls, Call{
			PhoneNumber: call.Destination,
			Duration:    call.Duration,
			TimeStamp:   call.TimeStamp.Format("2006-01-02 15:04:05"),
			Amount:      cost,
		})

		total += cost
	}
	return total, international, national, friends, userCalls
}

func (u *UserService) isInternational(in, out string) bool {
	return in[:3] != out[:3]
}

func (u *UserService) isFriend(friends []string, friend string) bool {
	for _, f := range friends {
		if f == friend {
			return true
		}
	}
	return false
}

func (u *UserService) calculateFare(amount float32, applyFriendDiscount bool) float32 {
	if applyFriendDiscount {
		return 0
	}
	return amount
}
