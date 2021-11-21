package main

import (
	"atm-cli/commands"
	"atm-cli/model"
	"bufio"
	"os"
	"strings"
)

func main() {

	atm := model.CreateNewAtm()
	scanner := bufio.NewScanner(os.Stdin)
	doExit := false
	for scanner.Scan() {

		command := scanner.Text()
		args := strings.Split(command, " ")
		switch args[0] {
		case "exit":
			println("exit command")
			doExit = true
			break
		case "login":
			{
				loginCommand := commands.LoginCommand{Atm: atm}
				loginCommand.Execute(args)
			}
		case "deposit":
			{
				depositCommand := commands.DepositCommand{Atm: atm}
				depositCommand.Execute(args)
			}
		case "withdraw":
			{
				withdrawCommand := commands.WithdrawCommand{Atm: atm}
				withdrawCommand.Execute(args)
			}
		case "transfer":
			{
				transferCommand := commands.TransferCommand{Atm: atm}
				transferCommand.Execute(args)
			}
		case "logout":
			{
				logoutCommand := commands.LogoutCommand{Atm: atm}
				logoutCommand.Execute(args)
			}
		default:
			println("command not found")
		}

		if doExit {
			break
		}
	}
}
