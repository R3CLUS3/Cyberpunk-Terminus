package perso

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (c *Character) StartRaidArasaka() {
	reader := bufio.NewReader(os.Stdin)

	// Carte Arasaka (6x6) sur 2 niveaux
	maps := [2][6][6]rune{
		// Niveau 1 : QG Arasaka & Securité
		{
			{'P', '.', '.', '#', '.', 'E'},
			{'.', '#', '.', '#', '.', '.'},
			{'.', 'S', '.', '.', '.', '.'},
			{'#', '#', '#', '#', 'S', '.'},
			{'.', '.', '.', '.', '.', '.'},
			{'.', '#', 'S', '.', '.', '.'},
		},
		// Niveau 2 : Étage Exécutif & Héliport (Adam Smasher)
		{
			{'P', '.', '.', '.', '#', '.'},
			{'.', '#', '#', '.', '#', '.'},
			{'.', 'S', '.', '.', '.', '.'},
			{'.', '#', '#', '#', 'S', '.'},
			{'.', '.', '.', '.', '.', '.'},
			{'#', '#', '#', '.', '.', 'B'},
		},
	}

	playerX, playerY := 0, 0
	currentLevel := 0

	for {
		if c.IsDead() {
			return
		}

		ClearScreen()
		fmt.Printf(Red + " ========================================== \n")
		fmt.Printf("      SYS.NET // RAID ARASAKA TOWER - NIVEAU %d/2\n", currentLevel+1)
		fmt.Printf(" ========================================== \n\n" + Reset)

		for y := 0; y < 6; y++ {
			fmt.Print("  ")
			for x := 0; x < 6; x++ {
				fmt.Printf("%c ", maps[currentLevel][y][x])
			}
			fmt.Println()
		}

		fmt.Println(Blue + "\n------------------------------------------")
		fmt.Println("LÉGENDE : P = Joueur | S = Agent Ninja Arasaka")
		fmt.Println("          B = Adam Smasher | # = Mur | E = Escalier")
		fmt.Println("------------------------------------------")
		fmt.Print("Déplacement (Z, Q, S, D | 0 = Fuir) : " + Reset)

		input, _ := reader.ReadString('\n')
		dir := strings.ToLower(strings.TrimSpace(input))

		if dir == "0" {
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
			continue
		}

		if newX < 0 || newX >= 6 || newY < 0 || newY >= 6 {
			continue
		}

		targetTile := maps[currentLevel][newY][newX]
		if targetTile == '#' {
			continue
		}

		maps[currentLevel][playerY][playerX] = '.'
		playerX, playerY = newX, newY

		if targetTile == 'S' {
			maps[currentLevel][playerY][playerX] = 'P'
			enemy := Monster{
				Name:      "Garde d'Élite Arasaka",
				MaxHP:     110 + (currentLevel * 40),
				CurrentHP: 110 + (currentLevel * 40),
				Attack:    22,
				RewardED:  120,
				RewardXP:  90,
			}
			c.fightRaidEnemy(&enemy)

		} else if targetTile == 'B' {
			maps[currentLevel][playerY][playerX] = 'P'
			fmt.Println("\n[ALERTE ROUGE] ADAM SMASHER EST EN FACE DE VOUS ! (venge Rebecca.)")
			boss := Monster{
				Name:      "Adam Smasher",
				MaxHP:     450,
				CurrentHP: 450,
				Attack:    40,
				RewardED:  1500,
				RewardXP:  800,
			}
			c.fightRaidEnemy(&boss)

			if boss.IsDead() {
				fmt.Println(Blue + "\n[LÉGENDE DE NIGHT CITY] T'as butté Smasher !" + Reset)

				// Loot du Composant Ultime Rang S+
				c.Inventory = append(c.Inventory, "Composant Rang S+")
				fmt.Println(Yellow + "[LOOT LEGENDAIRE] +1x Composant Rang S+ ajouté à l'inventaire !" + Reset)
				return
			}
		}
		if targetTile == 'E' {
			currentLevel++
			playerX, playerY = 0, 0
			maps[currentLevel][playerY][playerX] = 'P'

		} else {
			maps[currentLevel][playerY][playerX] = 'P'
		}
	}
}
