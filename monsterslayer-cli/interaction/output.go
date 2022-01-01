package interaction

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/common-nighthawk/go-figure"
)

type RoundData struct {
	Action              string
	PlayerAttackingDmg  int
	PlayerHealValue     int
	MonsterAttackingDmg int
	PlayerHealth        int
	MonsterHealth       int
}

func PrintGreeting() {
	asciiFigure := figure.NewFigure("MONSTER SLAYER", "", true)
	asciiFigure.Print()
	fmt.Println("Game is starting now")
	fmt.Println("Goodluck!")
}

func ShowAvaliableActions(isSpecialAttackAvailable bool) {
	fmt.Println("Please choose your action")
	fmt.Println("--------------------------")
	fmt.Println("1) Attack Monster")
	fmt.Println("2) Heal")

	if isSpecialAttackAvailable {
		fmt.Println("3) Use special attack")
	}

}

func RoundInfo(roundData *RoundData) {
	if roundData.Action == "ATTACK" {
		fmt.Printf("Player attacked monster for %v damage.\n", roundData.PlayerAttackingDmg)
	} else if roundData.Action == "SPECIAL_ATTACK" {
		fmt.Printf("Player performed special attack on monster for %v damage.\n", roundData.PlayerAttackingDmg)
	} else {
		fmt.Printf("Player healed for %v.\n", roundData.PlayerHealValue)
	}

	fmt.Printf("Monster attacked player for %v damage.\n", roundData.MonsterAttackingDmg)
	fmt.Printf("Player Health: %v\n", roundData.PlayerHealth)
	fmt.Printf("Monster Health: %v\n", roundData.MonsterHealth)
}

func GameWinner(winner string) {
	fmt.Println(("--------------------------------"))
	asciiFigure := figure.NewColorFigure("GAME OVER", "", "red", true)
	asciiFigure.Print()
	fmt.Println("----------------------------------")
	fmt.Printf("%v won\n", winner)
}

func LogFile(gameRounds *[]RoundData) {
	exPath, err := os.Executable()

	if err != nil {
		fmt.Println("Cannot create Log file. Exiting!")
		return
	}

	exPath = filepath.Dir(exPath)

	file, err := os.Create(exPath + "/gameLog.txt")

	if err != nil {
		fmt.Println("Saving log file failed. Exiting app!")
		return
	}

	for index, value := range *gameRounds {
		rawLogData := map[string]string{
			"Round":                 fmt.Sprint(index + 1),
			"Action":                value.Action,
			"Player Attack Damage":  fmt.Sprint(value.PlayerAttackingDmg),
			"Player Heal Value":     fmt.Sprint(value.PlayerHealValue),
			"Monster Attack Damage": fmt.Sprint(value.MonsterAttackingDmg),
			"Player Health":         fmt.Sprint(value.PlayerHealth),
			"Monster Health":        fmt.Sprint(value.MonsterHealth),
		}
		stringLogData := fmt.Sprintln(rawLogData)

		_, err := file.WriteString(stringLogData)

		if err != nil {
			fmt.Println("Writing current Log to file failed!")
			continue
		}
	}

	file.Close()
	fmt.Println("Wrote data to log!")
}
