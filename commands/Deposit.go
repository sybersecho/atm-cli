package commands

import (
	"atm-cli/model"
	"errors"
	"fmt"
	"strconv"
)

type DepositCommand struct {
	Atm *model.Atm
}

func (d *DepositCommand) Execute(args []string) error {
	atm := d.Atm
	if len(args) == 1 {
		println("invalid deposit argument")
		return errors.New("invalid deposit argument")
	}

	if atm.LoggedUser == nil {
		println("please login first")
		return errors.New("please login first")
	}

	amount, _ := strconv.Atoi(args[1])
	atm.LoggedUser.Balance += amount
	atm.Users[atm.LoggedUser.Name] = atm.LoggedUser

	println(fmt.Sprintf("your balance is $%v", atm.LoggedUser.Balance))
	return nil
}
