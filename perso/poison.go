package perso

import (
	"fmt"
	"time"
)

// Poison inflige 10 dégâts par seconde pendant 3 secondes
func (c *Character) Poison() {
	fmt.Println(Green + "\n  [AVERTISSEMENT SÉCURITÉ] EXPOSITION À UNE GRENADE NEUROTOXIQUE !" + Reset)

	for i := 1; i <= 3; i++ {
		time.Sleep(1 * time.Second) // Pause de 1 seconde

		c.CurrentHP -= 10
		if c.CurrentHP < 0 {
			c.CurrentHP = 0
		}

		fmt.Printf("[TICK %d/3] Dégâts neurotoxiques subis (-10 HP) | PV actuels : %d/%d HP\n", i, c.CurrentHP, c.MaxHP)

		// Vérification si le joueur succombe au poison
		if c.IsDead() {
			return
		}
	}

	fmt.Println(" [SYSTÈME] Neurotoxine dissipée.")
}
