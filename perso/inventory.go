package perso

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// AccessInventory gère l'affichage et l'équipement des objets
func (c *Character) AccessInventory() {
	reader := bufio.NewReader(os.Stdin)

	for {
		ClearScreen()
		fmt.Println("┌──────────────────────────────────────────────────┐")
		fmt.Println("│             SYS.NET // INVENTAIRE                │")
		fmt.Println("├──────────────────────────────────────────────────┤")
		if len(c.Inventory) == 0 {
			fmt.Println("│  (Inventaire vide)                               │")
		} else {
			for i, item := range c.Inventory {
				fmt.Printf("│  %2d. %-43s│\n", i+1, item)
			}
		}
		fmt.Println("└──────────────────────────────────────────────────┘")
		fmt.Println("Entrez le numéro d'un objet à utiliser / équiper (0 pour quitter) :")
		fmt.Print("Choix : ")

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		if choice == "0" {
			return
		}

		var itemIdx int
		_, err := fmt.Sscanf(choice, "%d", &itemIdx)

		if err != nil || itemIdx < 1 || itemIdx > len(c.Inventory) {
			fmt.Println("\n⚠️ Choix invalide.")
			fmt.Print("\nAppuyez sur Entrée pour continuer...")
			reader.ReadString('\n')
			continue
		}

		idx := itemIdx - 1
		selectedItem := c.Inventory[idx]

		switch {
		case strings.EqualFold(selectedItem, "Stimpack"):
			c.Heal(idx)

		case strings.EqualFold(selectedItem, "Veste de David Martinez") || strings.EqualFold(selectedItem, "Veste de David"):
			c.EquipArmor(selectedItem, 30)
			// Retire l'objet de l'inventaire après équipement
			c.Inventory = append(c.Inventory[:idx], c.Inventory[idx+1:]...)

		case strings.EqualFold(selectedItem, "Pare-balle Arasaka"):
			c.EquipArmor(selectedItem, 50)
			// Retire l'objet de l'inventaire après équipement
			c.Inventory = append(c.Inventory[:idx], c.Inventory[idx+1:]...)

		default:
			fmt.Printf("\n⚠️ Impossible d'équiper ou d'utiliser [%s] directement.\n", selectedItem)
		}

		fmt.Print("\nAppuyez sur Entrée pour continuer...")
		reader.ReadString('\n')
	}
}
