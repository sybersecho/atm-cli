package model

type Atm struct {
	Users      map[string]*User
	LoggedUser *User
}

type User struct {
	Name    string
	Balance int
}

func (a *Atm) GetUser(name string) *User {
	if u, ok := a.Users[name]; ok {
		return u
	}

	newUser := &User{
		Name:    name,
		Balance: 0,
	}
	a.Users[name] = newUser

	return newUser
}

func CreateNewAtm() *Atm {
	return &Atm{
		Users:      make(map[string]*User),
		LoggedUser: nil,
	}
}
