package perso

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// CanAddInventory vérifie si l'inventaire n'est pas plein
func (c *Character) CanAddInventory() bool {
	return len(c.Inventory) < c.MaxInventory
}

// AddInventory ajoute un objet seulement si la limite n'est pas atteinte
func (c *Character) AddInventory(item string) bool {
	if !c.CanAddInventory() {
		fmt.Printf("\n❌ Inventaire plein (%d/%d) ! Impossible d'ajouter : %s\n", len(c.Inventory), c.MaxInventory, item)
		return false
	}

	c.Inventory = append(c.Inventory, item)
	fmt.Printf("\n[+] %s a été ajouté à votre inventaire (%d/%d) !\n", item, len(c.Inventory), c.MaxInventory)
	return true
}

// RemoveInventory (identique à avant)
func (c *Character) RemoveInventory(itemIndex int) {
	if itemIndex >= 0 && itemIndex < len(c.Inventory) {
		c.Inventory = append(c.Inventory[:itemIndex], c.Inventory[itemIndex+1:]...)
	}
}

// AccessInventory (affichage de la jauge d'inventaire)
func (c *Character) AccessInventory() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("┌──────────────────────────────────────────────────┐")
		fmt.Printf("│ SYS.NET // STOCKAGE & EQUIPEMENT [%02d/%02d]      │\n", len(c.Inventory), c.MaxInventory)
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
		fmt.Print("Entrez le numéro de l'objet à utiliser : ")

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

		if strings.EqualFold(item, "Stimpack") {
			c.Heal(selectedIndex - 1)
		} else if strings.EqualFold(item, "Grenade Neurotoxique") {
			c.RemoveInventory(selectedIndex - 1)
			c.Poison()
		} else if strings.EqualFold(item, "Puce de combat") {
			c.RemoveInventory(selectedIndex - 1)
			c.PuceDeCombat()
		} else {
			fmt.Printf("\n Impossible d'utiliser %s pour le moment.\n", item)
		}
	}
}
