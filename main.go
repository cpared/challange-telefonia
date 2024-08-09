package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	rep "challange/internal/repositories/user"
	serv "challange/internal/services"
	"challange/models"
)

const (
	AMOUNT_OF_ARGS = 4
)

func main() {

	if len(os.Args) != AMOUNT_OF_ARGS {
		fmt.Println("No arguments provided")
		os.Exit(1)
	}

	dates := strings.Split(os.Args[2], "/")
	tbefore, err := time.Parse("2006-01-02", dates[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	tafter, err := time.Parse("2006-01-02", dates[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	// Initialize Repositories
	userRepository := rep.NewUserRepository()
	callsRepository := rep.NewCallRepository(os.Args[3])

	// Initialize Services
	userService := serv.NewUserService(userRepository, callsRepository)

	servResponse := userService.CalculateInvoice(tbefore, tafter, os.Args[1])

	response := mapToResponse(servResponse)
	jsonData, err := json.MarshalIndent(response, "", "  ")
	// jsonData, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(jsonData))
}

func mapToResponse(servicesResponse serv.Response) models.Response {
	return models.Response{
		User: models.User{
			Address:     servicesResponse.User.Address,
			Name:        servicesResponse.User.Name,
			PhoneNumber: servicesResponse.User.PhoneNumber,
		},
		Calls:                    mapToCalls(servicesResponse.Calls),
		TotalInternationaSeconds: servicesResponse.TotalInternationaSeconds,
		TotalNationaSeconds:      servicesResponse.TotalNationaSeconds,
		TotalFriendSeconds:       servicesResponse.TotalFriendSeconds,
		Total:                    servicesResponse.Total,
	}
}

func mapToCalls(calls []serv.Call) []models.Call {
	mappedCalls := make([]models.Call, len(calls))
	for i, call := range calls {
		mappedCalls[i] = models.Call{
			PhoneNumber: call.PhoneNumber,
			Duration:    call.Duration,
			TimeStamp:   call.TimeStamp,
			Amount:      call.Amount,
		}
	}
	return mappedCalls
}
