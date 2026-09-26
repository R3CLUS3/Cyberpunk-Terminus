package perso

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// DisplayCharcudoc ouvre la boutique du Charcudoc local
func (c *Character) DisplayCharcudoc() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("┌──────────────────────────────────────────────────┐")
		fmt.Println("│       Vendeur du marché noir // CLINIC-NET       │")
		fmt.Println("├──────────────────────────────────────────────────┤")
		fmt.Println("│  1. Stimpack (Soin +50 HP)  --  GRATUIT          │")
		fmt.Println("├──────────────────────────────────────────────────┤")
		fmt.Println("│  0. Quitter la clinique                          │")
		fmt.Println("└──────────────────────────────────────────────────┘")
		fmt.Print("Choisi ce que tu veux acheter... j'ai pas ton temps (0-1) : ")

		choiceInput, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(choiceInput)

		switch choice {
		case "1":
			itemBought := "Stimpack"
			c.AddInventory(itemBought)
			fmt.Printf("\n [Vendeur] : \"%s récupéré avec succès !\"\n\n", itemBought)
		case "0":
			fmt.Println("\n [Vendeur] : \"Reviens vite et meurs pas dans la rue.\"")
			return
		default:
			fmt.Println("\n Saisie invalide dans le terminal du Vendeur.")
		}
	}
}
