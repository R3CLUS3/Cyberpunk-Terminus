package perso

import "fmt"

// IsDead vérifie si les PV du personnage sont tombés à 0 ou moins.
// Si c'est le cas, simule la mort et la réanimation à 50% du MaxHP.
func (c *Character) IsDead() bool {
	if c.CurrentHP <= 0 {
		c.CurrentHP = 0
		fmt.Println(Red + "\n┌──────────────────────────────────────────────────┐")
		fmt.Println("│ ⚠️   ALERTE CRITIQUE : SIGNATION VITALE PERDUE  ⚠️ │")
		fmt.Println("├──────────────────────────────────────────────────┤")
		fmt.Println("│  Sujet neutralisé... Réinitialisation du système. │")
		fmt.Println("└──────────────────────────────────────────────────┘" + Reset)

		// Réanimation à 50% des PV max
		c.CurrentHP = c.MaxHP / 2

		fmt.Printf("\n⚡ [PROTOCOL REBOOT] Réinitialisation avec %d/%d HP (50%% des PV Max).\n\n", c.CurrentHP, c.MaxHP)

		c.DisplayInfo()
		return true
	}
	return false
}
