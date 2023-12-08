package users


type User struct {
    Username string
    Password string
}

type UserService struct {
}

func (UserService) createUser(newUser User) {
}
