package perso

import (
	"fmt"
	"time"
)

// Death déclenche l'alerte de mort, attend 3 secondes, puis reboot
func (c *Character) Death() {
	c.CurrentHP = 0
	ClearScreen()

	fmt.Println(Red + "┌──────────────────────────────────────────────────┐")
	fmt.Println("│ ⚠️   ALERTE CRITIQUE : SIGNAL VITAL PERDU   ⚠️ │")
	fmt.Println("├──────────────────────────────────────────────────┤")
	fmt.Println("│  Sujet neutralisé... Réinitialisation du système. │")
	fmt.Println("└──────────────────────────────────────────────────┘" + Reset)

	// Pause de 3 secondes pour laisser le temps de lire
	time.Sleep(3 * time.Second)

	// Réanimation à 50% des PV max
	c.CurrentHP = c.MaxHP / 2

	ClearScreen()
	fmt.Printf("⚡ [PROTOCOL REBOOT] Réinitialisation réussie (%d/%d HP).\n\n", c.CurrentHP, c.MaxHP)
	c.DisplayInfo()
}
