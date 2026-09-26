package perso

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// AccessInventory affiche l'inventaire et permet d'utiliser un objet
func (c *Character) AccessInventory() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("┌──────────────────────────────────────────────────┐")
		fmt.Println("│           SYS.NET // STOCKAGE & EQUIPEMENT       │")
		fmt.Println("├──────────────────────────────────────────────────┤")

		if len(c.Inventory) == 0 {
			fmt.Println("│  [VIDE] : Aucun objet dans le stock.             │")
			fmt.Println("└──────────────────────────────────────────────────┘")
			return
		}

		for i, item := range c.Inventory {
			fmt.Printf("│  [%02d] %-40s │\n", i+1, item)
		}
		fmt.Println("├──────────────────────────────────────────────────┤")
		fmt.Println("│  [0]  Retour au menu principal                   │")
		fmt.Println("└──────────────────────────────────────────────────┘")
		fmt.Print("Entre le numéro de l'objet à utiliser : ")

		choiceInput, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(choiceInput)

		if choice == "0" {
			return
		}

		var selectedIndex int
		_, err := fmt.Sscanf(choice, "%d", &selectedIndex)
		if err != nil || selectedIndex < 1 || selectedIndex > len(c.Inventory) {
			fmt.Println("\n Choix invalide.")
			continue
		}

		item := c.Inventory[selectedIndex-1]

		// Appel de la méthode Heal située dans heal.go
		if strings.EqualFold(item, "Stimpack") {
			c.Heal(selectedIndex - 1)
		} else {
			fmt.Printf("\n Impossible d'utiliser %s pour le moment.\n", item)
		}
	}
}

func (c *Character) AddInventory(item string) {
	c.Inventory = append(c.Inventory, item)
	fmt.Printf("\n[+] %s a été ajouté à ton inventaire !\n", item)
}
