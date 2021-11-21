package commands

import (
	"atm-cli/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLogoutCommand_ExecuteSuccess(t *testing.T) {
	atm := model.CreateNewAtm()
	loginArgs := []string{"login", "alice"}
	loginCommand := LoginCommand{Atm: atm}
	loginCommand.Execute(loginArgs)

	assert.Equal(t, loginArgs[1], loginCommand.Atm.LoggedUser.Name)

	args := []string{"logout"}
	command := LogoutCommand{Atm: atm}
	command.Execute(args)

	user := command.Atm.LoggedUser
	assert.Nil(t, user)
}
func TestLogoutCommand_ExecuteLogoutWhenNoUserLogin(t *testing.T) {
	atm := model.CreateNewAtm()

	args := []string{"logout"}
	command := LogoutCommand{Atm: atm}
	command.Execute(args)

	user := command.Atm.LoggedUser
	assert.Nil(t, user)
}
