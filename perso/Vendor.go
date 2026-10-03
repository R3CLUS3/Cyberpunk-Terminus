// vendeur etc..
package perso

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	stimpackStock = 24
	stimpackPrice = 99

	poisonStock = 6
	poisonPrice = 299

	puceStock = 1
	pucePrice = 500

	sacStock = 2
	sacPrice = 399

	// Composants d'artisanat achetables
	compCStock = 20
	compCPrice = 99

	compBStock = 15
	compBPrice = 129

	compAStock = 2
	compAPrice = 299

	mercenaireStock = 1
	mercenairePrice = 699

	arasakaStock = 1
	arasakaPrice = 999

	davidJacketStock = 1
	davidJacketPrice = 0
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
		fmt.Printf("│  3. Puce de combat (Upgrade) -- %d $ED (Stock: %d)│\n", pucePrice, puceStock)
		fmt.Printf("│  4. Extension Sac (+5 places)-- %d $ED (Stock: %d)│\n", sacPrice, sacStock)
		fmt.Println("├──────────────────────────────────────────────────┤")
		fmt.Printf("│  5. Composant Rang C        -- %d $ED (Stock: %d)│\n", compCPrice, compCStock)
		fmt.Printf("│  6. Composant Rang B        -- %d $ED (Stock: %d)│\n", compBPrice, compBStock)
		fmt.Printf("│  7. Composant Rang A        -- %d $ED (Stock: %d)│\n", compAPrice, compAStock)
		fmt.Println("├──────────────────────────────────────────────────┤")
		fmt.Printf("│  8. Plastron Mercenaire (+40 HP)-- %d $ED (Stock: %d)│\n", mercenairePrice, mercenaireStock)
		fmt.Printf("│  9. Plastron Arasaka    (+50 HP)-- %d $ED (Stock: %d)│\n", arasakaPrice, arasakaStock)
		fmt.Printf("│ 10. Veste David Martinez        -- %d $ED (Stock: %d)│\n", davidJacketPrice, davidJacketStock)
		fmt.Println("├──────────────────────────────────────────────────┤")
		fmt.Println("│  0. Retour au menu principal                     │")
		fmt.Println("└──────────────────────────────────────────────────┘")
		fmt.Print("Choisissez un article à acheter (0-7) : ")

		choiceInput, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(choiceInput)

		switch choice {
		case "1":
			c.buyConsumable(&stimpackStock, stimpackPrice, "Stimpack")
		case "2":
			c.buyConsumable(&poisonStock, poisonPrice, "Grenade Neurotoxique")
		case "3":
			hasPuce := false
			for _, imp := range c.Implants {
				if strings.Contains(imp, "Puce de combat") {
					hasPuce = true
					break
				}
			}

			if hasPuce {
				fmt.Println(Red + "\n Puce déjà installée !" + Reset)
			} else if !c.CanAddInventory() {
				fmt.Println(Red + "\n Votre inventaire est plein !" + Reset)
			} else if puceStock <= 0 {
				fmt.Println(Red + "\n Rupture de stock !" + Reset)
			} else if c.Money < pucePrice {
				fmt.Println(Red + "\n Fonds insuffisants !" + Reset)
			} else {
				c.Money -= pucePrice
				puceStock--
				c.AddInventory("Puce de combat")
			}
		case "4":
			if sacStock <= 0 {
				fmt.Println(Red + "\n Limite d'achat de sacs atteinte !" + Reset)
			} else if c.Money < sacPrice {
				fmt.Println(Red + "\n Fonds insuffisants !" + Reset)
			} else {
				c.Money -= sacPrice
				sacStock--
				c.MaxInventory += 5
				fmt.Printf("\nCapacité d'inventaire augmentée à %d !\n", c.MaxInventory)
			}
		case "5":
			c.buyConsumable(&compCStock, compCPrice, "Composant Rang C")
		case "6":
			c.buyConsumable(&compBStock, compBPrice, "Composant Rang B")
		case "7":
			c.buyConsumable(&compAStock, compAPrice, "Composant Rang A")
		case "8":
			c.buyConsumable(&mercenaireStock, mercenairePrice, "Plastron de Mercenaire")

		case "9":
			c.buyConsumable(&arasakaStock, arasakaPrice, "Plastron de soldat Arasaka")

		case "10":
			// Vérification stricte du nom du personnage
			pName := strings.ToLower(c.Name)
			if pName != "david" && pName != "lucy" && pName != "falco" {
				fmt.Println(Red + "\n ACCÈS REFUSÉ : La signature biométrique de cette veste ne correspond pas à votre profil !" + Reset)
			} else if !c.CanAddInventory() {
				fmt.Println(Red + "\n Votre inventaire est plein !" + Reset)
			} else if davidJacketStock <= 0 {
				fmt.Println(Red + "\n Cette veste est unique   !" + Reset)
			} else if c.Money < davidJacketPrice {
				fmt.Println(Red + "\n Fonds insuffisants !" + Reset)
			} else {
				c.Money -= davidJacketPrice
				davidJacketStock--
				c.AddInventory(Yellow + "Veste de David Martinez" + Reset)
			}
		case "0":
			return
		default:
			fmt.Println(Red + "\n Choix invalide." + Reset)
		}
	}
}

// Fonction utilitaire pour effectuer l'achat de consommables
func (c *Character) buyConsumable(stock *int, price int, itemName string) {
	if !c.CanAddInventory() {
		fmt.Println(Red + "\n Votre inventaire est plein !" + Reset)
	} else if *stock <= 0 {
		fmt.Println(Red + "\n Rupture de stock !" + Reset)
	} else if c.Money < price {
		fmt.Println(Red + "\n Fonds insuffisants !" + Reset)
	} else {
		c.Money -= price
		*stock--
		c.AddInventory(itemName)
	}
}
