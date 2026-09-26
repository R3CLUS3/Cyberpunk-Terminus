package perso

import (
	"fmt"
	"strings"
)

type Skill struct {
	Name    string
	Damage  int
	Uses    int  // -1 = infini, >0 = nombre de charges
	IsFatal bool // true = déclenche la mort si Uses atteint 0
}

// PuceDeCombat applique l'amélioration du Coup de poing / Hacking (+5 dégâts)
func (c *Character) PuceDeCombat() {
	// Vérification dans les implants déjà équipés
	for _, imp := range c.Implants {
		if strings.Contains(imp, "Puce de combat") {
			fmt.Println("\n⚠️ La Puce de combat est déjà installée dans votre système !")
			return
		}
	}

	// Augmentation des dégâts du premier sort du joueur
	if len(c.Skills) > 0 {
		c.Skills[0].Damage += 5
	}

	// Ajout de l'implant dans la liste des implants équipés
	c.Implants = append(c.Implants, "Puce de combat (+5 dmg)")

	fmt.Println("\n[IMPLANT INSTALLÉ] Puce de combat activée !")
	fmt.Printf("Votre compétence de base passe à %d points de dégâts !\n\n", c.Skills[0].Damage)

	c.DisplayInfo()
}
