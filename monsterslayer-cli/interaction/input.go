package interaction

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetPlayerChoice(isSpecialAttackAvailable bool) string {

	for {
		playerChoice, _ := getPlayerInput()

		if playerChoice == "1" {
			return "ATTACK"
		} else if playerChoice == "2" {
			return "HEAL"
		} else if playerChoice == "3" && isSpecialAttackAvailable {
			return "SPECIAL_ATTACK"
		} else {
			fmt.Println("Getting user input failed. Please try again.")
		}
	}
}

func getPlayerInput() (string, error) {
	fmt.Print("Enter your choice: ")

	userInput, err := reader.ReadString('\n')

	if err != nil {
		return "", err
	}

	userInput = strings.Replace(userInput, "\n", "", -1)

	return userInput, nil
}
