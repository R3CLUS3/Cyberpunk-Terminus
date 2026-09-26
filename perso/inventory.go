package perso

import "fmt"

// AccessInventory affiche la liste complète des objets présents dans l'inventaire
func (c Character) AccessInventory() {
	fmt.Println("┌──────────────────────────────────────────────────┐")
	fmt.Println("│           SYS.NET // STOCKAGE & EQUIPEMENT       │")
	fmt.Println("├──────────────────────────────────────────────────┤")

	if len(c.Inventory) == 0 {
		fmt.Println("│  [VIDE] : Aucun objet dans le stock.             │")
	} else {
		for i, item := range c.Inventory {
			fmt.Printf("│  [%02d] %-40s │\n", i+1, item)
		}
	}

	fmt.Println("└──────────────────────────────────────────────────┘")
}
