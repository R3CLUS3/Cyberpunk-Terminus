// duel de legendes
package perso

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// GetLegendsCatalogue renvoie la liste des légendes affrontables
func GetLegendsCatalogue() []Monster {
	return []Monster{
		{
			Name:      Green + "Rebecca" + Reset,
			MaxHP:     250,
			CurrentHP: 250,
			Attack:    28,
			RewardED:  500,
			RewardXP:  500,
		},
		{
			Name:      Blue + "Lucy" + Reset,
			MaxHP:     220,
			CurrentHP: 220,
			Attack:    35,
			RewardED:  500,
			RewardXP:  300,
		},
		{
			Name:      Yellow + "David Martinez" + Reset,
			MaxHP:     320,
			CurrentHP: 320,
			Attack:    45,
			RewardED:  800,
			RewardXP:  500,
		},
		{
			Name:      "V",
			MaxHP:     420,
			CurrentHP: 420,
			Attack:    55,
			RewardED:  1200,
			RewardXP:  750,
		},
		{
			Name:      Red + "Adam Smasher" + Reset,
			MaxHP:     999,
			CurrentHP: 999,
			Attack:    65,
			RewardED:  2000,
			RewardXP:  1000,
		},
	}
}

// StartLegendDuel lance le menu de sélection et le combat contre une légende
func (c *Character) StartLegendDuel() {
	reader := bufio.NewReader(os.Stdin)
	legends := GetLegendsCatalogue()

	for {
		if c.IsDead() {
			return
		}

		ClearScreen()
		fmt.Println(Green + " ========================================== ")
		fmt.Println("       SYS.NET // ARENE DES LÉGENDES DE NC    ")
		fmt.Println(" ========================================== " + Reset)
		fmt.Println("Choisissez un adversaire pour un duel en 1v1 :\n")

		for i, legend := range legends {
			fmt.Printf("  %d. %-26s \n", i+1, legend.Name)
		}
		fmt.Println("  0. Retour au menu principal")
		fmt.Println("----------------------------------------------")
		fmt.Print("Votre choix (0-5) : ")

		choiceInput, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(choiceInput)

		if choice == "0" {
			return
		}

		var selectedIdx int
		_, err := fmt.Sscanf(choice, "%d", &selectedIdx)

		if err != nil || selectedIdx < 1 || selectedIdx > len(legends) {
			fmt.Println(Red + "\nChoix invalide." + Reset)
			continue
		}

		// Copie de l'adversaire sélectionné
		boss := legends[selectedIdx-1]

		ClearScreen()
		fmt.Printf(Green+"  [DUEL DÉCLENCHÉ] %s s'avance dans l'arène !\n", boss.Name)
		fmt.Println("──────────────────────────────────────────────" + Reset)

		// Réutilisation du moteur de combat de raid
		c.fightRaidEnemy(&boss)

		if boss.IsDead() {
			fmt.Printf(Yellow+"\n[EXPLOIT MONUMENTAL] Vous avez vaincu %s !\n", boss.Name)
			fmt.Printf("Récompense : +%d $ED\n"+Reset, boss.RewardED)
		}

		fmt.Print("\nAppuyez sur Entrée pour continuer...")
		reader.ReadString('\n')
	}
}
