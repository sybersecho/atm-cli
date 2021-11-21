package commands

import (
	"atm-cli/model"
	"fmt"
)

type LogoutCommand struct {
	Atm *model.Atm
}

func (l *LogoutCommand) Execute(args []string) {
	atm := l.Atm
	if atm.LoggedUser == nil {
		println("no user logged in")
		return
	}
	println(fmt.Sprintf("Goodbye, %v", atm.LoggedUser.Name))
	atm.LoggedUser = nil
}
