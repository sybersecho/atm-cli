package commands

import (
	"atm-cli/model"
	"fmt"
	"strconv"
)

type TransferCommand struct {
	Atm *model.Atm
}

func (t *TransferCommand) Execute(args []string) error {
	atm := t.Atm
	if len(args) <= 2 {
		println("invalid transfer argument")
		return fmt.Errorf("invalid transfer argument")
	}

	if atm.LoggedUser == nil {
		println("please login first")
		return fmt.Errorf("please login first")
	}

	username := args[1]
	amountToTransfer, _ := strconv.Atoi(args[2])
	user := atm.GetUser(username)
	user.Balance += amountToTransfer
	atm.LoggedUser.Balance -= amountToTransfer

	println(fmt.Sprintf("your balance is $%v", atm.LoggedUser.Balance))

	return nil
}
