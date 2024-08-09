package repository

type User struct {
	Address     string
	Name        string
	PhoneNumber string
	Friends    []string
}

type UserRepository struct {
	users map[string]User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		map[string]User{
			"+5491167980950": {
				Address:     "123 Main St, Buenos Aires, Argentina",
				Name:        "Alice Johnson",
				PhoneNumber: "+5491167980950",
				Friends:     []string{"+191167980952", "+5491167950940", "+5491167910920"},
			},
			"+191167980952": {
				Address:     "456 Oak St, New York, USA",
				Name:        "Bob Smith",
				PhoneNumber: "+191167980952",
				Friends:     []string{"+5491167980950", "+5491167980953", "+5491167910920"},
			},
			"+5491167910920": {
				Address:     "789 Pine St, Buenos Aires, Argentina",
				Name:        "Carol White",
				PhoneNumber: "+5491167910920",
				Friends:     []string{"+191167980952", "+5491167980950", "+5491167980953"},
			},
			"+5491167950940": {
				Address:     "101 Maple St, Buenos Aires, Argentina",
				Name:        "David Brown",
				PhoneNumber: "+5491167950940",
				Friends:     []string{"+5491167980950", "+191167970944", "+5491167980951"},
			},
			"+191167970944": {
				Address:     "202 Birch St, New York, USA",
				Name:        "Eve Davis",
				PhoneNumber: "+191167970944",
				Friends:     []string{"+5491167950940", "+5491167980954", "+5491167920944"},
			},
			"+5491167980953": {
				Address:     "303 Cedar St, Buenos Aires, Argentina",
				Name:        "Frank Miller",
				PhoneNumber: "+5491167980953",
				Friends:     []string{"+5491167980950", "+191167980952", "+5491167950940"},
			},
			"+5491167980951": {
				Address:     "404 Elm St, Buenos Aires, Argentina",
				Name:        "Grace Wilson",
				PhoneNumber: "+5491167980951",
				Friends:     []string{"+5491167950940", "+191167980953", "+5491167940999"},
			},
			"+191167980953": {
				Address:     "505 Ash St, New York, USA",
				Name:        "Hank Lee",
				PhoneNumber: "+191167980953",
				Friends:     []string{"+5491167980951", "+5491167950940", "+5491167940999"},
			},
			"+5491167940999": {
				Address:     "606 Willow St, Buenos Aires, Argentina",
				Name:        "Ivy Green",
				PhoneNumber: "+5491167940999",
				Friends:     []string{"+5491167980951", "+191167980953", "+5491167920930"},
			},
			"+5491167920930": {
				Address:     "707 Redwood St, Buenos Aires, Argentina",
				Name:        "Jack Black",
				PhoneNumber: "+5491167920930",
				Friends:     []string{"+5491167940999", "+5491167980950", "+5491167920944"},
			},
		},
	}
}

// GetUser returns a user by phone number
// If the user is not found, the second return value is false
func (u *UserRepository) GetUser(phoneNumber string) (User, bool) {
	user, found := u.users[phoneNumber]
	return user, found
}
