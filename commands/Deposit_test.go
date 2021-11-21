package commands

import (
	"atm-cli/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDepositCommand_Execute(t *testing.T) {
	atm := model.CreateNewAtm()
	loginArgs := []string{"login", "alice"}
	loginCommand := LoginCommand{Atm: atm}
	loginCommand.Execute(loginArgs)

	assert.Equal(t, loginArgs[1], loginCommand.Atm.LoggedUser.Name)

	args:=[]string{ "deposit", "10"}
	depositCommand := DepositCommand{Atm: atm}
	err := depositCommand.Execute(args)
	assert.NoErrorf(t, err, "success deposit")

	assert.Equal(t, 10, depositCommand.Atm.LoggedUser.Balance)
}

func TestDepositCommand_ExecuteWithoutLogin(t *testing.T) {
	atm := model.CreateNewAtm()

	args:=[]string{ "deposit", "10"}
	depositCommand := DepositCommand{Atm: atm}
	err := depositCommand.Execute(args)
	assert.Errorf(t, err, "no user login")

	user := depositCommand.Atm.LoggedUser
	assert.Nil(t, user)
}
