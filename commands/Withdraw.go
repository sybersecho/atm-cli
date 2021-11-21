package commands

import (
	"atm-cli/model"
	"fmt"
	"strconv"
)

type WithdrawCommand struct {
	Atm *model.Atm
}

func (w *WithdrawCommand) Execute(args []string) error {
	atm := w.Atm
	if len(args) == 1 {
		println("invalid withdraw argument")
		return fmt.Errorf("invalid withdraw argument")
	}

	if atm.LoggedUser == nil {
		println("please login first")
		return fmt.Errorf("please login first")
	}

	amount, _ := strconv.Atoi(args[1])
	balance := atm.LoggedUser.Balance
	if amount > balance {
		println(fmt.Sprintf("insufficient balance. Your balance is $%v", balance))
		return fmt.Errorf("insufficient balance. Your balance is $%v", balance)
	}
	atm.LoggedUser.Balance = balance - amount
	atm.Users[atm.LoggedUser.Name] = atm.LoggedUser

	println(fmt.Sprintf("your balance is $%v", atm.LoggedUser.Balance))

	return nil
}
