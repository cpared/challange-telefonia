package models

type Response struct {
	User                     User    `json:"user"`
	Calls                    []Call  `json:"calls"`
	TotalInternationaSeconds int     `json:"total_international_seconds"`
	TotalNationaSeconds      int     `json:"total_national_seconds"`
	TotalFriendSeconds       int     `json:"total_friend_seconds"`
	Total                    float32 `json:"total"`
}

type User struct {
	Address     string `json:"address"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type Call struct {
	PhoneNumber string  `json:"phone_number"`
	Duration    int     `json:"duration"`
	TimeStamp   string  `json:"timestamp"`
	Amount      float32 `json:"amount"`
}
