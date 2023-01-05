package domain

type UserRepositoryStub struct {
	user  User
	users Users
}

func (u UserRepositoryStub) Users() (*Users, error) {
	return &u.users, nil
}

// テストデータ作成
func NewUserRepositoryStub() UserRepositoryStub {
	responseUser1 := User{UserID: 101, Name: "John Doe", Grade: "gold"}
	responseUser2 := User{UserID: 201, Name: "Randy Barrett", Grade: "silver"}
	responseUser3 := User{UserID: 301, Name: "Andrew Cary", Grade: "nonmember"}
	var users UsersArray
	users = append(users, &responseUser1)
	users = append(users, &responseUser2)
	users = append(users, &responseUser3)
	responseUsers := Users{Users: users}

	return UserRepositoryStub{user: User{}, users: responseUsers}
}
