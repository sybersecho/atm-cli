package commands

import (
	"atm-cli/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSuccessLogin(t *testing.T) {
	// init atm
	atm := model.CreateNewAtm()
	args := make([]string, 2)
	args[0] = "login"
	args[1] = "alice"

	doLogin := LoginCommand{Atm: atm}
	doLogin.Execute(args)

	if doLogin.Atm.LoggedUser == nil {
		t.Errorf("user failed to logged in")
	}

	assert.Equal(t, args[1], doLogin.Atm.LoggedUser.Name, "expected alice")
	assert.Equal(t, 0, doLogin.Atm.LoggedUser.Balance, "expected 0")
}

func TestLoginWithMissingArgument(t *testing.T) {
	// init atm
	atm := model.CreateNewAtm()
	args := make([]string, 2)
	args[0] = "login"

	doLogin := LoginCommand{Atm: atm}
	doLogin.Execute(args)

	assert.Equal(t, "", doLogin.Atm.LoggedUser.Name, "expected empty")
	assert.Equal(t, 0, doLogin.Atm.LoggedUser.Balance, "expected 0")
}
