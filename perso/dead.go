package perso

import "fmt"

// Death déclenche la séquence de mort et de réanimation du personnage
func (c *Character) Death() {
	c.CurrentHP = 0
	fmt.Println(Red + "\n┌──────────────────────────────────────────────────┐")
	fmt.Println("│ ⚠️   ALERTE CRITIQUE : SIGNAL VITALE PERDU  ⚠️ │")
	fmt.Println("├──────────────────────────────────────────────────┤")
	fmt.Println("│  Sujet neutralisé... Réinitialisation du système. │")
	fmt.Println("└──────────────────────────────────────────────────┘" + Reset)

	// Réanimation à 50% des PV max
	c.CurrentHP = c.MaxHP / 2

	fmt.Printf("\n⚡ [PROTOCOL REBOOT] Réinitialisation avec %d/%d HP (50%% des PV Max).\n\n", c.CurrentHP, c.MaxHP)

	c.DisplayInfo()
}
