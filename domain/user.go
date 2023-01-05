package domain

type User struct {
	UserID int64  `json:"user_id"`
	Name   string `json:"name"`
	Grade  string `json:"grade"`
}

type UsersArray []*User

type Users struct {
	Users UsersArray `json:"users"`
}

type UserRepository interface {
	Users() (*Users, error)
}
