package commands

import (
	"atm-cli/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithdrawCommand_Execute(t *testing.T) {
	atm := model.CreateNewAtm()
	loginArgs := []string{"login", "alice"}
	loginCommand := LoginCommand{Atm: atm}
	loginCommand.Execute(loginArgs)

	assert.Equal(t, loginArgs[1], loginCommand.Atm.LoggedUser.Name)

	depositArgs := []string{"deposit", "20"}
	depositCommand := DepositCommand{Atm: atm}
	t.Log("deposit 20")
	err := depositCommand.Execute(depositArgs)

	t.Log("withdraw 10, expected 10")
	args := []string{"withdraw", "10"}
	withdrawCommand := WithdrawCommand{Atm: atm}
	err = withdrawCommand.Execute(args)
	assert.NoErrorf(t, err, "success withdraw")

	assert.Equal(t, 10, withdrawCommand.Atm.LoggedUser.Balance)
}

func TestWithdrawCommand_Execute_withEmptyBalance(t *testing.T) {
	atm := model.CreateNewAtm()
	loginArgs := []string{"login", "alice"}
	loginCommand := LoginCommand{Atm: atm}
	loginCommand.Execute(loginArgs)

	assert.Equal(t, loginArgs[1], loginCommand.Atm.LoggedUser.Name)

	t.Log("withdraw 10, expected 10")
	args := []string{"withdraw", "10"}
	withdrawCommand := WithdrawCommand{Atm: atm}
	err := withdrawCommand.Execute(args)
	assert.Errorf(t, err, "empty balance")

	assert.Equal(t, 0, withdrawCommand.Atm.LoggedUser.Balance)
}
