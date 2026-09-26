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

// InitCharacter crée un nouveau personnage avec la classe choisie
func InitCharacter(name string, classChoice string) Character {
	var className string
	var maxHP int
	var startingMoney int
	var baseSkill Skill

	switch classChoice {
	case "1":
		className = "Gosse des rues"
		maxHP = 100
		startingMoney = 1500
		baseSkill = Skill{Name: "Coup de poing", Damage: 10}
	case "2":
		className = "Nomade"
		maxHP = 120
		startingMoney = 1000
		baseSkill = Skill{Name: "Tir de précision", Damage: 12}
	case "3":
		className = "Corpo"
		maxHP = 90
		startingMoney = 1500
		baseSkill = Skill{Name: "Piratage rapide", Damage: 15}
	default:
		className = "Gosse des rues"
		maxHP = 100
		startingMoney = 1500
		baseSkill = Skill{Name: "Coup de poing", Damage: 10}
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
	fmt.Printf("⭐ +%d XP gagnés ! (%d/%d XP)\n", amount, c.XP, c.MaxXP)

	for c.XP >= c.MaxXP && c.Level < 50 {
		c.XP -= c.MaxXP
		c.Level++
		c.MaxHP += 5
		c.CurrentHP += 5                      // Soigne de +5 HP lors du level up
		c.MaxXP = int(float64(c.MaxXP) * 1.2) // Augmente le seuil d'XP requis

		fmt.Printf("\n🎉 [LEVEL UP !] Vous passez Niveau %d !\n", c.Level)
		fmt.Printf("❤️ Vos PV Maximaux augmentent à %d HP (+5 HP) !\n", c.MaxHP)

		if c.Level == 50 {
			fmt.Println("🏆 NIVEAU MAXIMUM 50 ATTEINT !")
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
