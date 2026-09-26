package perso

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	stimpackStock = 5
	stimpackPrice = 20

	poisonStock = 3
	poisonPrice = 30

	sacStock = 2
	sacPrice = 40

	// Composants d'artisanat achetables
	compCStock = 5
	compCPrice = 30

	compBStock = 3
	compBPrice = 75

	compAStock = 1
	compAPrice = 150
)

func (c *Character) DisplayVendeur() {
	reader := bufio.NewReader(os.Stdin)

	for {
		soldeStr := fmt.Sprintf("[VOTRE SOLDE] : %d $ED", c.Money)

		fmt.Println("┌──────────────────────────────────────────────────┐")
		fmt.Println("│           SYS.NET // MARCHE NOIR (VENDEUR)       │")
		fmt.Printf("│  %-48s│\n", soldeStr)
		fmt.Println("├──────────────────────────────────────────────────┤")
		fmt.Printf("│  1. Stimpack (+50 HP)        -- %d $ED (Stock: %d)│\n", stimpackPrice, stimpackStock)
		fmt.Printf("│  2. Grenade Neurotoxique    -- %d $ED (Stock: %d)│\n", poisonPrice, poisonStock)
		fmt.Printf("│  3. Extension Sac (+5 places)-- %d $ED (Stock: %d)│\n", sacPrice, sacStock)
		fmt.Println("├──────────────────────────────────────────────────┤")
		fmt.Printf("│  4. Composant Rang C        -- %d $ED (Stock: %d)│\n", compCPrice, compCStock)
		fmt.Printf("│  5. Composant Rang B        -- %d $ED (Stock: %d)│\n", compBPrice, compBStock)
		fmt.Printf("│  6. Composant Rang A        -- %d $ED (Stock: %d)│\n", compAPrice, compAStock)
		fmt.Println("├──────────────────────────────────────────────────┤")
		fmt.Println("│  0. Retour au menu principal                     │")
		fmt.Println("└──────────────────────────────────────────────────┘")
		fmt.Print("Choisissez un article à acheter (0-6) : ")

		choiceInput, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(choiceInput)

		switch choice {
		case "1":
			c.buyConsumable(&stimpackStock, stimpackPrice, "Stimpack")
		case "2":
			c.buyConsumable(&poisonStock, poisonPrice, "Grenade Neurotoxique")
		case "3":
			if sacStock <= 0 {
				fmt.Println("\n❌ Limite d'achat de sacs atteinte !")
			} else if c.Money < sacPrice {
				fmt.Println("\n❌ Fonds insuffisants !")
			} else {
				c.Money -= sacPrice
				sacStock--
				c.MaxInventory += 5
				fmt.Printf("\n🎒 Capacité d'inventaire augmentée à %d !\n", c.MaxInventory)
			}
		case "4":
			c.buyConsumable(&compCStock, compCPrice, "Composant Rang C")
		case "5":
			c.buyConsumable(&compBStock, compBPrice, "Composant Rang B")
		case "6":
			c.buyConsumable(&compAStock, compAPrice, "Composant Rang A")
		case "0":
			return
		default:
			fmt.Println("\n⚠️ Choix invalide.")
		}
	}
}

// Fonction utilitaire pour éviter la répétition du code d'achat
func (c *Character) buyConsumable(stock *int, price int, itemName string) {
	if !c.CanAddInventory() {
		fmt.Println("\n❌ Votre inventaire est plein !")
	} else if *stock <= 0 {
		fmt.Println("\n❌ Rupture de stock !")
	} else if c.Money < price {
		fmt.Println("\n❌ Fonds insuffisants !")
	} else {
		c.Money -= price
		*stock--
		c.AddInventory(itemName)
	}
}
