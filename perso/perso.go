package perso

import (
	"fmt"
	"strings"
)

type Character struct {
	Name          string
	Class         string
	Level         int
	XP            int
	MaxXP         int
	MaxHP         int
	CurrentHP     int
	Money         int
	Inventory     []string
	MaxInventory  int
	Skills        []Skill
	Implants      []string
	EquippedArmor string
	ArmorBonus    int
}

// CanAddInventory vérifie si l'inventaire n'est pas plein
func (c *Character) CanAddInventory() bool {
	return len(c.Inventory) < c.MaxInventory
}

// AddInventory ajoute un objet à l'inventaire si de la place est disponible
func (c *Character) AddInventory(item string) {
	if c.CanAddInventory() {
		c.Inventory = append(c.Inventory, item)
	}
}

// RemoveInventory retire un objet de l'inventaire selon son index
func (c *Character) RemoveInventory(index int) {
	if index >= 0 && index < len(c.Inventory) {
		c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)
	}
}

// InitCharacter crée un nouveau personnage avec la classe choisie
func InitCharacter(name string, classChoice string) Character {
	var className string
	var maxHP int
	var startingMoney int
	var baseSkill Skill

	// Accepte aussi bien les chiffres ("1", "2", "3") que les noms complets ("Netrunner", "Corpo", etc.)
	switch strings.ToLower(classChoice) {
	case "1", "gosse des rues":
		className = Red + "Gosse des rues" + Reset
		maxHP = 120
		startingMoney = 1000
		baseSkill = Skill{Name: "Coup de poing", Damage: 7}

	case "2", "netrunner", "netruner":
		className = Yellow + "Netrunner" + Reset
		maxHP = 80
		startingMoney = 1000
		baseSkill = Skill{Name: "Hacking", Damage: 12}

	case "3", "corpo":
		className = Blue + "Corpo" + Reset
		maxHP = 100
		startingMoney = 1500
		baseSkill = Skill{Name: "Coup de matraque", Damage: 8}

	default:
		className = Red + "Gosse des rues" + Reset
		maxHP = 120
		startingMoney = 1000
		baseSkill = Skill{Name: "Coup de poing", Damage: 7}
	}

	return Character{
		Name:          name,
		Class:         className,
		Level:         1,
		XP:            0,
		MaxXP:         100,
		MaxHP:         maxHP,
		CurrentHP:     maxHP,
		Money:         startingMoney,
		Inventory:     []string{"Stimpack"},
		MaxInventory:  10,
		Skills:        []Skill{baseSkill},
		Implants:      []string{},
		EquippedArmor: "Aucun",
		ArmorBonus:    0,
	}
}

// AddXP gère le gain d'expérience et le passage au niveau supérieur
func (c *Character) AddXP(amount int) {
	if c.Level >= 50 {
		return // Niveau Max atteint
	}

	c.XP += amount
	fmt.Printf("T'as gagné +%d XP ! (%d/%d XP)\n", amount, c.XP, c.MaxXP)

	for c.XP >= c.MaxXP && c.Level < 50 {
		c.XP -= c.MaxXP
		c.Level++
		c.MaxHP += 5
		c.CurrentHP += 5                      // Soigne de +5 HP lors du level up
		c.MaxXP = int(float64(c.MaxXP) * 1.2) // Augmente le seuil d'XP requis

		fmt.Printf("\n [LEVEL UP !] Vous passez Niveau %d !\n", c.Level)
		fmt.Printf(" Vos PV Maximaux augmentent à %d HP (+5 HP) !\n", c.MaxHP)

		if c.Level == 50 {
			fmt.Println(" NIVEAU MAXIMUM 50 ATTEINT !\n T'es une Légende eh!")
			c.XP = 0
			break
		}
	}
}

// IsDead vérifie si le personnage est K.O.
func (c *Character) IsDead() bool {
	return c.CurrentHP <= 0
}

// DisplayInfo affiche la fiche complète du sujet
func (c Character) DisplayInfo() {
	totalBlocks := 10
	filledBlocks := (c.CurrentHP * totalBlocks) / c.MaxHP
	if filledBlocks < 0 {
		filledBlocks = 0
	}
	emptyBlocks := totalBlocks - filledBlocks

	hpBar := strings.Repeat("█", filledBlocks) + strings.Repeat("░", emptyBlocks)

	nameLine := fmt.Sprintf("[IDENTIFIANT] : %s", c.Name)
	classLine := fmt.Sprintf("[ORIGINE]     : %s", c.Class)

	levelLine := fmt.Sprintf("[NIVEAU]      : %d (XP: %d/%d)", c.Level, c.XP, c.MaxXP)
	if c.Level >= 50 {
		levelLine = fmt.Sprintf("[NIVEAU]      : 50 (MAX)", c.Level)
	}

	moneyLine := fmt.Sprintf("[ARGENT]      : %d $ED", c.Money)
	hpLine := fmt.Sprintf("[SANTÉ]       : %d / %d [%s]", c.CurrentHP, c.MaxHP, hpBar)

	armorLine := fmt.Sprintf("[PROTECTION]  : %s (+%d HP)", c.EquippedArmor, c.ArmorBonus)
	if c.EquippedArmor == "Aucun" || c.EquippedArmor == "" {
		armorLine = "[PROTECTION]  : Aucune"
	}

	fmt.Println("┌──────────────────────────────────────────────────┐")
	fmt.Println("│           SYS.NET // FICHE SUJET                 │")
	fmt.Println("├──────────────────────────────────────────────────┤")
	fmt.Printf("│  %-48s│\n", nameLine)
	fmt.Printf("│  %-48s│\n", classLine)
	fmt.Printf("│  %-48s│\n", levelLine)
	fmt.Printf("│  %-48s│\n", moneyLine)
	fmt.Println("├──────────────────────────────────────────────────┤")
	fmt.Printf("│  %-48s│\n", hpLine)
	fmt.Printf("│  %-48s│\n", armorLine)
	fmt.Println("├──────────────────────────────────────────────────┤")

	if len(c.Skills) == 0 {
		fmt.Printf("│  %-48s│\n", "[COMPÉTENCES] : Aucune")
	} else {
		for i, s := range c.Skills {
			var skillLine string
			skillText := fmt.Sprintf("%s (%d dmg)", s.Name, s.Damage)
			if i == 0 {
				skillLine = fmt.Sprintf("[COMPÉTENCES] : • %s", skillText)
			} else {
				skillLine = fmt.Sprintf("              : • %s", skillText)
			}
			fmt.Printf("│  %-48s│\n", skillLine)
		}
	}

	fmt.Println("├──────────────────────────────────────────────────┤")

	if len(c.Implants) == 0 {
		fmt.Printf("│  %-48s│\n", "[IMPLANTS]    : Aucun")
	} else {
		for i, imp := range c.Implants {
			var implantLine string
			if i == 0 {
				implantLine = fmt.Sprintf("[IMPLANTS]    : • %s", imp)
			} else {
				implantLine = fmt.Sprintf("              : • %s", imp)
			}
			fmt.Printf("│  %-48s│\n", implantLine)
		}
	}

	fmt.Println("└──────────────────────────────────────────────────┘")
}
