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
			fmt.Println(Red + "\n Choix invalide." + Reset)
			fmt.Print("\nAppuyez sur Entrée pour continuer...")
			reader.ReadString('\n')
			continue
		}

		selectedItem := c.Inventory[itemIdx-1]
		itemLower := strings.ToLower(selectedItem)

		switch {
		// --- SOINS & CONSUMMABLES ---
		case strings.Contains(itemLower, "stimpack"):
			c.Heal(itemIdx - 1)

		case strings.Contains(itemLower, "grenade neurotoxique"):
			fmt.Println(Red + "\n C'est une arme de combat ! Utilise-la durant un affrontement." + Reset)

		// --- ARMURES & PROTECTIONS ---
		case strings.Contains(itemLower, "veste de david"):
			c.EquipArmor("Veste de David Martinez", 77)
			c.RemoveInventory(itemIdx - 1)

		case strings.Contains(itemLower, "plastron arasaka") || strings.Contains(itemLower, "pare-balle arasaka"):
			c.EquipArmor("Plastron Arasaka", 35)
			c.RemoveInventory(itemIdx - 1)

		case strings.Contains(itemLower, "plastron mercenaire"):
			c.EquipArmor("Plastron Mercenaire", 25)
			c.RemoveInventory(itemIdx - 1)

		// --- UPGRADES & SACS ---
		case strings.Contains(itemLower, "extension sac"):
			c.MaxInventory += 5
			fmt.Printf("\n Extension appliquée ! Capacité d'inventaire augmentée à %d slots.\n", c.MaxInventory)
			c.RemoveInventory(itemIdx - 1)

		case strings.Contains(itemLower, "puce de combat"):
			fmt.Println("\n Puce d'amélioration détectée ! Rendez-vous à l'Atelier Cyberware pour l'installer.")

		// --- COMPOSANTS DE CRAFT ---
		case strings.Contains(itemLower, "composant"):
			fmt.Println("\n Matériau d'artisanat. Utilise-le chez le Forgeron / Atelier Cyberware.")

		default:
			fmt.Printf(Red+"\nImpossible d'utiliser [%s] directement depuis le menu.\n"+Reset, selectedItem)
		}

		fmt.Print("\nAppuyez sur Entrée pour continuer...")
		reader.ReadString('\n')
	}
}
