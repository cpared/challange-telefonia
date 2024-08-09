package services

type Response struct {
	User                     User
	Calls                    []Call
	TotalInternationaSeconds int
	TotalNationaSeconds      int
	TotalFriendSeconds       int
	Total                    float32
}

type User struct {
	Address     string
	Name        string
	PhoneNumber string
	Friends     []string
}

type Call struct {
	PhoneNumber string
	Duration    int
	TimeStamp   string
	Amount      float32
}
