// cyberforgeron etc..
package perso

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func (c *Character) DisplayForgeron() {
	reader := bufio.NewReader(os.Stdin)
	recipes := GetAvailableRecipes()

	for {
		soldeStr := fmt.Sprintf("[SOLDE ATELIER] : %d $ED", c.Money)

		fmt.Println("┌──────────────────────────────────────────────────┐")
		fmt.Println("│       SYS.NET // ATELIER CYBERWARE & FORGE       │")
		fmt.Printf("│  %-48s│\n", soldeStr)
		fmt.Println("├──────────────────────────────────────────────────┤")

		for i, r := range recipes {
			reqFormatted := FormatReqComps(r.ReqComps)
			// Ligne formatée proprement avec une largeur constante
			recipeLine := fmt.Sprintf("%d. %-22s (%d dmg) %s", i+1, r.Name, r.SkillResult.Damage, reqFormatted)
			fmt.Printf("│  %-48s│\n", recipeLine)
		}

		fmt.Println("├──────────────────────────────────────────────────┤")
		fmt.Println("│  0. Retour au menu principal                     │")
		fmt.Println("└──────────────────────────────────────────────────┘")
		fmt.Print("Choisissez une arme à fabriquer (0-5) : ")

		choiceInput, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(choiceInput)

		if choice == "0" {
			return
		}

		var selected int
		_, err := fmt.Sscanf(choice, "%d", &selected)
		if err != nil || selected < 1 || selected > len(recipes) {
			fmt.Println(Red + "\n Choix invalide." + Reset)
			continue
		}

		recipe := recipes[selected-1]

		// Vérification si déjà possédée
		alreadyOwned := false
		for _, s := range c.Skills {
			if s.Name == recipe.SkillResult.Name {
				alreadyOwned = true
				break
			}
		}

		if alreadyOwned {
			fmt.Printf(Red+"\n Vous possédez déjà l'arme %s !\n"+Reset, recipe.Name)
			continue
		}

		if c.Money < recipe.Price {
			fmt.Printf(Red+"\n Fonds insuffisants (%d $ED requis) !\n"+Reset, recipe.Price)
			continue
		}

		// Vérification qu'ON A TOUS les composants requis dans l'inventaire
		missingComp := ""
		indicesToRemove := []int{}

		// Copie temporaire de l'inventaire pour simuler la suppression
		tempInventory := make([]string, len(c.Inventory))
		copy(tempInventory, c.Inventory)

		for _, req := range recipe.ReqComps {
			foundIndex := -1
			for idx, item := range tempInventory {
				if strings.EqualFold(item, req) {
					foundIndex = idx
					break
				}
			}

			if foundIndex == -1 {
				missingComp = req
				break
			} else {
				// Marque l'élément comme trouvé pour ne pas le réutiliser deux fois
				tempInventory[foundIndex] = ""
				indicesToRemove = append(indicesToRemove, foundIndex)
			}
		}

		if missingComp != "" {
			fmt.Printf(Red+"\n Composant manquant : Vous devez posséder [%s] !\n"+Reset, missingComp)
			continue
		}

		// Validation du craft : Déduction de l'argent
		c.Money -= recipe.Price

		// Suppression effective des composants de l'inventaire réels
		// (en supprimant de la fin vers le début pour ne pas altérer les index)
		for i := len(indicesToRemove) - 1; i >= 0; i-- {
			c.RemoveInventory(indicesToRemove[i])
		}

		// Ajout de l'arme aux compétences
		c.Skills = append(c.Skills, recipe.SkillResult)

		fmt.Printf(Blue+"\n [FABRICATION RÉUSSIE] %s assemblé !\n"+Reset, recipe.Name)
		fmt.Printf(Yellow+" Nouvelle compétence : %s (%d dmg).\n\n"+Reset, recipe.SkillResult.Name, recipe.SkillResult.Damage)
	}
}
