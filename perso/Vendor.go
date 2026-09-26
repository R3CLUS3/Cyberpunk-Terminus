package perso

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	stimpackStock = 10
	stimpackPrice = 199

	poisonStock = 6
	poisonPrice = 299

	puceStock = 1   // Disponible en 1 seul exemplaire
	pucePrice = 500 // Coût 500 $ED
)

func (c *Character) DisplayVendeur() {
	reader := bufio.NewReader(os.Stdin)

	for {
		soldeStr := fmt.Sprintf("[VOTRE SOLDE] : %d $ED", c.Money)
		stimpackStr := fmt.Sprintf("1. Stimpack (+50 HP)        -- %d $ED (Stock: %d)", stimpackPrice, stimpackStock)
		poisonStr := fmt.Sprintf("2. Grenade Neurotoxique    -- %d $ED (Stock: %d)", poisonPrice, poisonStock)
		puceStr := fmt.Sprintf("3. Puce de combat (Upgrade) -- %d $ED (Stock: %d)", pucePrice, puceStock)

		fmt.Println("┌──────────────────────────────────────────────────┐")
		fmt.Println("│           SYS.NET // MARCHE NOIR (VENDEUR)       │")
		fmt.Printf("│  %-48s│\n", soldeStr)
		fmt.Println("├──────────────────────────────────────────────────┤")
		fmt.Printf("│  %-48s│\n", stimpackStr)
		fmt.Printf("│  %-48s│\n", poisonStr)
		fmt.Printf("│  %-48s│\n", puceStr)
		fmt.Println("├──────────────────────────────────────────────────┤")
		fmt.Println("│  0. Retour au menu principal                     │")
		fmt.Println("└──────────────────────────────────────────────────┘")
		fmt.Print("Choisissez un article à acheter (0-3) : ")

		choiceInput, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(choiceInput)

		switch choice {
		case "1":
			if stimpackStock <= 0 {
				fmt.Println("\n Rupture de stock pour les Stimpacks !")
			} else if c.Money < stimpackPrice {
				fmt.Println("\n Fonds insuffisants !")
			} else {
				c.Money -= stimpackPrice
				stimpackStock--
				c.AddInventory("Stimpack")
			}

		case "2":
			if poisonStock <= 0 {
				fmt.Println("\n Rupture de stock pour les Grenades Neurotoxiques !")
			} else if c.Money < poisonPrice {
				fmt.Println("\n Fonds insuffisants !")
			} else {
				c.Money -= poisonPrice
				poisonStock--
				c.AddInventory("Grenade Neurotoxique")
			}

		case "3":
			// Vérification si déjà possédée dans les compétences
			hasPuce := false
			for _, s := range c.Skills {
				if s.Name == "Puce de combat" {
					hasPuce = true
					break
				}
			}

			if hasPuce {
				fmt.Println("\n Vous possédez déjà cette puce de combat !")
			} else if puceStock <= 0 {
				fmt.Println("\n Rupture de stock !")
			} else if c.Money < pucePrice {
				fmt.Println("\n Fonds insuffisants !")
			} else {
				c.Money -= pucePrice
				puceStock--
				c.AddInventory("Puce de combat")
			}

		case "0":
			return

		default:
			fmt.Println("\n Choix invalide.")
		}
	}
}
