package perso

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func (c *Character) StartTrainingFight() {
	reader := bufio.NewReader(os.Stdin)
	wave := 1

	fmt.Println(Yellow + "\n  ==========================================  ")
	fmt.Println("       SYS.NET // ARENE D'ENTRAÎNEMENT INFINIE")
	fmt.Println("  ==========================================  " + Reset)

	for {
		enemy := GenerateBot(wave)
		fmt.Printf(Red+"\n [ENTRÉE EN ARÈNE] VAGUE %d : Un %s s'approche !\n\n"+Reset, wave, enemy.Name)

		// Boucle de combat contre un ennemi
		for !c.IsDead() && !enemy.IsDead() {
			fmt.Printf("--- [ VAGUE %d - TOUR ] ---------------------------\n", wave)
			fmt.Printf(Green+" %s : %d/%d HP  │   %s : %d/%d HP\n"+Reset, c.Name, c.CurrentHP, c.MaxHP, enemy.Name, enemy.CurrentHP, enemy.MaxHP)
			fmt.Println("--------------------------------------------------")
			fmt.Println(Red + "1. Attaquer (Compétences)" + Reset)
			fmt.Println(Yellow + "2. Inventaire (Utiliser un objet)" + Reset)
			fmt.Print("Choisissez votre action (1-2) : ")

			actionInput, _ := reader.ReadString('\n')
			action := strings.TrimSpace(actionInput)

			turnExecuted := false

			switch action {
			case "1":
				fmt.Println(Red + "\n--- Compétences ---" + Reset)
				for i, s := range c.Skills {
					fmt.Printf("%d. %s (%d dmg)\n", i+1, s.Name, s.Damage)
				}
				fmt.Print("Choix : ")

				skillInput, _ := reader.ReadString('\n')
				skillChoice := strings.TrimSpace(skillInput)

				var skillIndex int
				_, err := fmt.Sscanf(skillChoice, "%d", &skillIndex)

				if err == nil && skillIndex >= 1 && skillIndex <= len(c.Skills) {
					selectedSkill := c.Skills[skillIndex-1]
					enemy.CurrentHP -= selectedSkill.Damage
					if enemy.CurrentHP < 0 {
						enemy.CurrentHP = 0
					}
					fmt.Printf(Red+"\n [%s] inflige %d dégâts à %s !\n"+Reset, selectedSkill.Name, selectedSkill.Damage, enemy.Name)
					turnExecuted = true
				} else {
					fmt.Println(Red + "\nChoix invalide." + Reset)
				}

			case "2":
				if len(c.Inventory) == 0 {
					fmt.Println(Red + "\nInventaire vide." + Reset)
				} else {
					fmt.Println(Yellow + "\n--- Inventaire ---" + Reset)
					for i, item := range c.Inventory {
						fmt.Printf("%d. %s\n", i+1, item)
					}
					fmt.Print("Choix (0 pour annuler) : ")

					itemInput, _ := reader.ReadString('\n')
					itemChoice := strings.TrimSpace(itemInput)

					var itemIndex int
					_, err := fmt.Sscanf(itemChoice, "%d", &itemIndex)

					if err == nil && itemIndex >= 1 && itemIndex <= len(c.Inventory) {
						item := c.Inventory[itemIndex-1]
						if strings.EqualFold(item, "Stimpack") {
							c.Heal(itemIndex - 1)
							turnExecuted = true
						} else {
							fmt.Printf(Red+"\nImpossible d'utiliser %s en combat.\n"+Reset, item)
						}
					}
				}

			default:
				fmt.Println(Red + "\nAction invalide." + Reset)
			}

			// Riposte de l'ennemi
			if turnExecuted && !enemy.IsDead() {
				time.Sleep(500 * time.Millisecond)
				c.CurrentHP -= enemy.Attack
				if c.CurrentHP < 0 {
					c.CurrentHP = 0
				}
				fmt.Printf(Red+" %s attaque et inflige %d dégâts !\n\n"+Reset, enemy.Name, enemy.Attack)

				if c.CurrentHP <= 0 {
					c.Death() // Déclenche la gestion de mort
					return
				}
			}
		}

		// Ennemi Vaincu
		if enemy.IsDead() {
			fmt.Printf(Yellow+"\n [CIBLE DÉTRUITE] %s éliminé !\n"+Reset, enemy.Name)
			c.Money += enemy.RewardED
			fmt.Printf(Yellow+" +%d $ED gagnés (Solde: %d $ED)\n"+Reset, enemy.RewardED, c.Money)
			c.AddXP(enemy.RewardXP)

			// Demande si le joueur veut continuer les vagues
			fmt.Println(Green + "\nVoulez-vous continuer vers la vague suivante ?")
			fmt.Println("1. Continuer le combat (Vague suivante)")
			fmt.Println("2. Quitter l'arène" + Reset)
			fmt.Print("Choix (1-2) : ")

			nextInput, _ := reader.ReadString('\n')
			nextChoice := strings.TrimSpace(nextInput)

			if nextChoice == "1" {
				wave++
			} else {
				fmt.Println("\n Retrait de l'arène d'entraînement. Retour au menu.")
				return
			}
		}
	}
}
