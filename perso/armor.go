package perso

import (
	"fmt"
)

// Armor représente un équipement corporel
type Armor struct {
	Name    string
	BonusHP int
	Price   int
	Stock   int
	NameReq []string // Noms de joueurs autorisés (vide si accessible à tous)
}

// EquipArmor équipe une armure, met à jour les PV max et gère les bonus
func (c *Character) EquipArmor(armorName string, bonusHP int) {
	// Si le personnage porte déjà une armure, on retire son ancien bonus
	if c.EquippedArmor != "" {
		fmt.Printf("\n Vous retirez : %s\n", c.EquippedArmor)
		c.MaxHP -= c.ArmorBonus
		if c.CurrentHP > c.MaxHP {
			c.CurrentHP = c.MaxHP
		}
	}

	// Équipement de la nouvelle armure
	c.EquippedArmor = armorName
	c.ArmorBonus = bonusHP
	c.MaxHP += bonusHP
	c.CurrentHP += bonusHP // Donne directement le bonus en PV actuels aussi

	fmt.Printf("\n[ÉQUIPEMENT SÉCURISÉ] %s équipée !\n", armorName)
	fmt.Printf(" Vos PV maximaux passent à %d HP (+%d HP) !\n", c.MaxHP, bonusHP)
}
