// stimpack
package perso

import "fmt"

const (
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Reset   = "\033[0m"
	Cyan    = "\033[36m"
	magenta = "\033[35m"
)

// Heal consomme un objet de soin et soigne le personnage de 50 HP (max HP respecté)
func (c *Character) Heal(itemIndex int) {
	if c.CurrentHP >= c.MaxHP {
		fmt.Println(Red + "\n Vos points de vie sont déjà au maximum !" + Reset)
		return
	}

	c.CurrentHP += 50
	if c.CurrentHP > c.MaxHP {
		c.CurrentHP = c.MaxHP
	}

	consumedItem := c.Inventory[itemIndex]
	c.Inventory = append(c.Inventory[:itemIndex], c.Inventory[itemIndex+1:]...)

	fmt.Printf("\n[+] %s utilisé ! +50 HP réinjectés.\n\n", consumedItem)

	c.DisplayInfo()
}
