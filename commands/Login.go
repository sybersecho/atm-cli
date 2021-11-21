package commands

import (
	"atm-cli/model"
	"fmt"
)

type LoginCommand struct {
	Atm *model.Atm
}

func (l *LoginCommand) Execute(args []string) {
	a := l.Atm
	if len(args) == 1 {
		println("invalid login argument")
		return
	}
	if a.LoggedUser != nil {
		println(fmt.Sprintf("current user is %v", a.LoggedUser.Name))
		return
	}
	user := a.GetUser(args[1])
	a.LoggedUser = user
	login(user)
}

func login(u *model.User) {
	println(fmt.Sprintf("Hello, %v!", u.Name))
	println(fmt.Sprintf("Your Balance is $%v", u.Balance))
}
