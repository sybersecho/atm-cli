package commands

import (
	"atm-cli/model"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestTransferCommand_Execute_WithoutLogin(t *testing.T) {
	atm := model.CreateNewAtm()

	args := []string{"transfer", "bob", "10"}
	transferCommand := TransferCommand{Atm: atm}
	err := transferCommand.Execute(args)
	assert.Errorf(t, err, "user not login")

	user := transferCommand.Atm.LoggedUser
	assert.Nil(t, user)
}

func TestTransferCommand_Execute(t *testing.T) {
	atm := model.CreateNewAtm()
	loginArgs := []string{"login", "alice"}
	loginCommand := LoginCommand{Atm: atm}
	loginCommand.Execute(loginArgs)

	assert.Equal(t, loginArgs[1], loginCommand.Atm.LoggedUser.Name)

	depositArgs := []string{"deposit", "20"}
	depositCommand := DepositCommand{Atm: atm}

	err := depositCommand.Execute(depositArgs)

	assert.NoErrorf(t, err, "deposit 10")
	expected, _ := strconv.Atoi(depositArgs[1])
	assert.Equal(t, expected, loginCommand.Atm.LoggedUser.Balance)

	args := []string{"transfer", "bob", "10"}
	transferCommand := TransferCommand{Atm: atm}
	err = transferCommand.Execute(args)
	assert.NoErrorf(t, err, "success transfer")

	assert.Equal(t, 10, transferCommand.Atm.LoggedUser.Balance)
}
