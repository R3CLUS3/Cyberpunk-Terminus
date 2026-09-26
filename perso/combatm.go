package perso

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// CombatM Raid Militech sur 3 niveaux
func (c *Character) StartRaidMilitech() {
	reader := bufio.NewReader(os.Stdin)

	// Layout des 3 cartes (5x5)
	// P = Joueur, . = Case vide, S = Soldat, B = Boss, E = Escalier/Sortie
	maps := [3][5][5]rune{
		// Niveau 1 : Entrée du complexe
		{
			{'P', '.', 'S', '.', '.'},
			{'.', '#', '#', '#', '.'},
			{'.', 'S', '.', '.', '.'},
			{'#', '#', '.', '#', '.'},
			{'.', '.', '.', 'S', 'E'},
		},
		// Niveau 2 : Centre de recherche
		{
			{'P', '#', '.', '.', 'S'},
			{'.', '#', '.', '#', '.'},
			{'.', 'S', '.', '#', '.'},
			{'.', '#', '#', '#', '.'},
			{'.', '.', '.', 'S', 'E'},
		},
		// Niveau 3 : Bureau du Commandant (Boss)
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

	for {
		// Vérification si le joueur est mort
		if c.IsDead() {
			return
		}

		// Affichage de l'entête du Raid
		fmt.Printf("\n🏢 ========================================== 🏢\n")
		fmt.Printf("      SYS.NET // RAID MILITECH - NIVEAU %d/3\n", currentLevel+1)
		fmt.Printf("🏢 ========================================== 🏢\n\n")

		// Rendu de la carte
		for y := 0; y < 5; y++ {
			fmt.Print("  ")
			for x := 0; x < 5; x++ {
				fmt.Printf("%c ", maps[currentLevel][y][x])
			}
			fmt.Println()
		}

		// Légende
		fmt.Println("\n------------------------------------------")
		fmt.Println("LÉGENDE : P = Joueur | S = Soldat | B = Boss")
		fmt.Println("          # = Mur    | E = Escalier | . = Vide")
		fmt.Println("------------------------------------------")
		fmt.Print("Déplacement (Z = Haut, Q = Gauche, S = Bas, D = Droite, 0 = Fuir) : ")

		input, _ := reader.ReadString('\n')
		dir := strings.ToLower(strings.TrimSpace(input))

		if dir == "0" {
			fmt.Println("\n🚪 Abandon du raid Militech...")
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
			fmt.Println("\n⚠️ Touche invalide (Utilisez Z, Q, S, D).")
			continue
		}

		// Vérification des limites de la carte
		if newX < 0 || newX >= 5 || newY < 0 || newY >= 5 {
			fmt.Println("\n⛔ Déplacement impossible : Hors limites.")
			continue
		}

		targetTile := maps[currentLevel][newY][newX]

		// Gestion des collisions et événements
		if targetTile == '#' {
			fmt.Println("\n🚧 Obstacle ! Un mur bloque le passage.")
			continue
		}

		// Efface l'ancienne position du joueur
		maps[currentLevel][playerY][playerX] = '.'

		// Mouvement du joueur
		playerX, playerY = newX, newY

		if targetTile == 'S' {
			maps[currentLevel][playerY][playerX] = 'P'
			fmt.Println("\n🚨 [ALERTE] Un Soldat Militech vous repère !")
			enemy := Monster{
				Name:      fmt.Sprintf("Soldat Militech (Niveau %d)", currentLevel+1),
				MaxHP:     50 + (currentLevel * 25),
				CurrentHP: 50 + (currentLevel * 25),
				Attack:    12 + (currentLevel * 5),
				RewardED:  50 + (currentLevel * 30),
				RewardXP:  40 + (currentLevel * 20),
			}
			c.fightRaidEnemy(&enemy)

		} else if targetTile == 'B' {
			maps[currentLevel][playerY][playerX] = 'P'
			boss := Monster{
				Name:      "Commandant Exécutif Militech",
				MaxHP:     220,
				CurrentHP: 220,
				Attack:    25,
				RewardED:  600,
				RewardXP:  400,
			}
			c.fightRaidEnemy(&boss)

			if boss.IsDead() {
				fmt.Println("\n🏆 [VICTOIRE] Complexe Militech neutralisé !")

				// Ajout du prototype Cybersquelette
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
						Uses:    10,
						IsFatal: true,
					})
					fmt.Println("⚠️ [LOOT EXPÉRIMENTAL] Vous récupérez le Proto-Cybersquelette Militech (120 dmg / 10 utilisations max) !")
				}
				return
			}
		}

		if targetTile == 'E' {
			currentLevel++
			playerX, playerY = 0, 0
			maps[currentLevel][playerY][playerX] = 'P'
			fmt.Printf("\n🪜 Vous prenez l'escalier et montez au Niveau %d !\n", currentLevel+1)

		} else {
			maps[currentLevel][playerY][playerX] = 'P'
		}
	}
}

// Sub-fonction pour gérer un combat individuel pendant le raid
func (c *Character) fightRaidEnemy(enemy *Monster) {
	reader := bufio.NewReader(os.Stdin)

	for !c.IsDead() && !enemy.IsDead() {
		fmt.Printf("\n--- COMBAT : %s (%d/%d HP) ---\n", enemy.Name, enemy.CurrentHP, enemy.MaxHP)
		fmt.Printf("👤 Vos PV : %d/%d HP\n", c.CurrentHP, c.MaxHP)
		fmt.Println("1. Attaquer")
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
				fmt.Printf("\n💥 Vous attaquez avec [%s] (-%d HP) !\n", selectedSkill.Name, selectedSkill.Damage)
				turnExecuted = true
			}
		} else if choice == "2" {
			for idx, item := range c.Inventory {
				if strings.EqualFold(item, "Stimpack") {
					c.Heal(idx)
					turnExecuted = true
					break
				}
			}
			if !turnExecuted {
				fmt.Println("\n⚠️ Pas de Stimpack disponible !")
			}
		}

		if turnExecuted && !enemy.IsDead() {
			c.CurrentHP -= enemy.Attack
			fmt.Printf("🥊 %s riposte et vous inflige %d dégâts !\n", enemy.Name, enemy.Attack)

			if c.CurrentHP <= 0 {
				c.Death()
				return
			}
		}
	}

	if enemy.IsDead() {
		fmt.Printf("\n🎉 Cible %s éliminée !\n", enemy.Name)
		c.Money += enemy.RewardED
		c.AddXP(enemy.RewardXP)
	}
}
