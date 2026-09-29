package perso

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// StartRaidMilitech gère le Raid Militech sur 3 niveaux
func (c *Character) StartRaidMilitech() {
	reader := bufio.NewReader(os.Stdin)

	// Layout des 3 cartes (5x5)
	maps := [3][5][5]rune{
		{
			{'P', '.', 'S', '.', '.'},
			{'.', '#', '#', '#', '.'},
			{'.', 'S', '.', '.', '.'},
			{'#', '#', '.', '#', '.'},
			{'.', '.', '.', 'S', 'E'},
		},
		{
			{'P', '#', '.', '.', 'S'},
			{'.', '#', '.', '#', '.'},
			{'.', 'S', '.', '#', '.'},
			{'.', '#', '#', '#', '.'},
			{'.', '.', '.', 'S', 'E'},
		},
		{
			{'P', '.', '.', '.', '.'},
			{'#', '#', 'S', '#', '.'},
			{'.', '.', '.', '.', '.'},
			{'.', '#', '#', '#', '.'},
			{'.', '.', 'S', '.', 'B'},
		},
	}

	playerX, playerY := 0, 0
	currentLevel := 0
	statusMessage := " Raid débuté. Infiltrez le complexe !"

	for {
		if c.IsDead() {
			return
		}

		// Nettoyage de l'écran à chaque déplacement
		ClearScreen()

		// En-tête
		fmt.Printf(Yellow + "==========================================\n")
		fmt.Printf("      SYS.NET // RAID MILITECH - NIVEAU %d/3\n", currentLevel+1)
		fmt.Printf("==========================================\n\n" + Reset)

		// Rendu de la carte
		for y := 0; y < 5; y++ {
			fmt.Print("  ")
			for x := 0; x < 5; x++ {
				fmt.Printf("%c ", maps[currentLevel][y][x])
			}
			fmt.Println()
		}

		// Légende & Message de statut
		fmt.Println(Blue + "\n------------------------------------------" + Reset)
		fmt.Println("LÉGENDE : P = Joueur | S = Soldat | B = Boss")
		fmt.Println("          # = Mur    | E = Escalier | . = Vide")
		fmt.Println(Blue + "------------------------------------------" + Reset)

		if statusMessage != "" {
			fmt.Println(statusMessage)
			fmt.Println("------------------------------------------")
		}

		fmt.Print(Yellow + "Déplacement (Z = Haut, Q = Gauche, S = Bas, D = Droite, 0 = Fuir) : " + Reset)

		input, _ := reader.ReadString('\n')
		dir := strings.ToLower(strings.TrimSpace(input))

		if dir == "0" {
			fmt.Println(Red + "\nAbandon du raid Militech... Sale lâche !" + Reset)
			return
		}

		newX, newY := playerX, playerY

		switch dir {
		case "z":
			newY--
		case "s":
			newY++
		case "q":
			newX--
		case "d":
			newX++
		default:
			statusMessage = Red + " Touche invalide (Utilisez Z, Q, S, D)." + Reset
			continue
		}

		// Limites et Murs
		if newX < 0 || newX >= 5 || newY < 0 || newY >= 5 || maps[currentLevel][newY][newX] == '#' {
			statusMessage = "Tu peux pas traverser les murs chooms !"
			continue
		}

		targetTile := maps[currentLevel][newY][newX]

		// Déplacement du joueur
		maps[currentLevel][playerY][playerX] = '.'
		playerX, playerY = newX, newY
		statusMessage = "" // Réinitialise le message par défaut

		// Événements
		if targetTile == 'S' {
			maps[currentLevel][playerY][playerX] = 'P'
			enemy := Monster{
				Name:      fmt.Sprintf("Soldat Militech (Niveau %d)", currentLevel+1),
				MaxHP:     50 + (currentLevel * 25),
				CurrentHP: 50 + (currentLevel * 25),
				Attack:    12 + (currentLevel * 5),
				RewardED:  50 + (currentLevel * 30),
				RewardXP:  40 + (currentLevel * 20),
			}
			c.fightRaidEnemy(&enemy)
			statusMessage = fmt.Sprintf(" %s éliminé !", enemy.Name)

		} else if targetTile == 'B' {
			maps[currentLevel][playerY][playerX] = 'P'
			boss := Monster{
				Name:      "Machine de sécurité Militech",
				MaxHP:     250,
				CurrentHP: 250,
				Attack:    25,
				RewardED:  600,
				RewardXP:  400,
			}
			c.fightRaidEnemy(&boss)

			if boss.IsDead() {
				ClearScreen()
				fmt.Println(Blue + "\n==========================================")
				fmt.Println(" [VICTOIRE] Le RAID contre Militech est un succès !")
				fmt.Println("==========================================" + Reset)

				c.Inventory = append(c.Inventory, "Composant Rang S", "Composant Rang S")
				fmt.Println(Yellow + " [LOOT] +2x Composant Rang S ajoutés !" + Reset)

				hasSkill := false
				for _, s := range c.Skills {
					if s.Name == "Cybersquelette Militech" {
						hasSkill = true
						break
					}
				}

				if !hasSkill {
					c.Skills = append(c.Skills, Skill{
						Name:    "Cybersquelette Militech",
						Damage:  120,
						Uses:    8,
						IsFatal: true,
					})
					fmt.Println(Yellow + "[LOOT EXPÉRIMENTAL] Proto-Cybersquelette récupéré !" + Reset)
				}

				fmt.Print("\nAppuyez sur Entrée pour continuer...")
				reader.ReadString('\n')
				return
			}

		} else if targetTile == 'E' {
			currentLevel++
			playerX, playerY = 0, 0
			maps[currentLevel][playerY][playerX] = 'P'
			statusMessage = Green + fmt.Sprintf(" Escalier emprunté ! Arrivée au Niveau %d/3", currentLevel+1) + Reset
		} else {
			maps[currentLevel][playerY][playerX] = 'P'
		}
	}
}

// Sub-fonction pour gérer un combat individuel pendant le raid
func (c *Character) fightRaidEnemy(enemy *Monster) {
	reader := bufio.NewReader(os.Stdin)
	lastAction := " Début de l'affrontement !"

	for !c.IsDead() && !enemy.IsDead() {
		ClearScreen()
		fmt.Println("==========================================")
		fmt.Printf("   COMBAT : %s\n", enemy.Name)
		fmt.Println("==========================================")
		fmt.Printf(" Vos PV : %d/%d HP | Ennemi : %d/%d HP\n", c.CurrentHP, c.MaxHP, enemy.CurrentHP, enemy.MaxHP)
		fmt.Println("------------------------------------------")

		if lastAction != "" {
			fmt.Println(lastAction)
			fmt.Println("------------------------------------------")
		}

		fmt.Println("1. Attaquer (Compétences)")
		fmt.Println("2. Utiliser Stimpack")
		fmt.Print("Choix : ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		turnExecuted := false

		if choice == "1" {
			fmt.Println("\n--- Compétences ---")
			for i, s := range c.Skills {
				fmt.Printf("%d. %s (%d dmg)\n", i+1, s.Name, s.Damage)
			}
			fmt.Print("Choix : ")

			skillInput, _ := reader.ReadString('\n')
			var skillIndex int
			fmt.Sscanf(strings.TrimSpace(skillInput), "%d", &skillIndex)

			if skillIndex >= 1 && skillIndex <= len(c.Skills) {
				selectedSkill := c.Skills[skillIndex-1]
				enemy.CurrentHP -= selectedSkill.Damage
				lastAction = fmt.Sprintf(" Vous attaquez avec [%s] (-%d HP) !", selectedSkill.Name, selectedSkill.Damage)
				turnExecuted = true
			} else {
				lastAction = "Choix de compétence invalide."
			}

		} else if choice == "2" {
			found := false
			for idx, item := range c.Inventory {
				if strings.EqualFold(item, "Stimpack") {
					c.Heal(idx)
					lastAction = " Stimpack utilisé (+50 HP) !"
					turnExecuted = true
					found = true
					break
				}
			}
			if !found {
				lastAction = Red + "Pas de Stimpack dans l'inventaire !" + Reset
			}
		}

		if turnExecuted && !enemy.IsDead() {
			c.CurrentHP -= enemy.Attack
			lastAction += fmt.Sprintf("\n%s riposte (-%d HP) !", enemy.Name, enemy.Attack)

			if c.CurrentHP <= 0 {
				c.Death()
				return
			}
		}
	}

	if enemy.IsDead() {
		c.Money += enemy.RewardED
		c.AddXP(enemy.RewardXP)
	}
}
