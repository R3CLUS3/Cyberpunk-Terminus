package perso

import (
	"fmt"
	"strings"
)

type Character struct {
	Name      string
	Class     string
	Level     int
	MaxHP     int
	CurrentHP int
	Money     int
	Inventory []string
	Skills    []Skill
}

// InitCharacter initialise les attributs du personnage selon la classe choisie
func InitCharacter(name string, class string, inventory []string) Character {
	var maxHP int
	var startingMoney int
	var baseSkill Skill

	switch strings.ToLower(class) {
	case "gosse des rues":
		maxHP = 120
		startingMoney = 1000
		baseSkill = Skill{Name: "Coup de poing", Damage: 7}

	case "netrunner":
		maxHP = 80
		startingMoney = 1000
		baseSkill = Skill{Name: "Hacking", Damage: 12}

	case "corpo":
		maxHP = 100
		startingMoney = 1500 // Plus d'argent au départ
		baseSkill = Skill{Name: "Coup de poing", Damage: 7}

	default: // Sécurité par défaut (Gosse des rues)
		class = "Gosse des rues"
		maxHP = 120
		startingMoney = 1000
		baseSkill = Skill{Name: "Coup de poing", Damage: 7}
	}

	// PV actuels au démarrage = 60% des PV max
	currentHP := (maxHP * 60) / 100

	return Character{
		Name:      name,
		Class:     class,
		Level:     1, // Niveau 1 au départ
		MaxHP:     maxHP,
		CurrentHP: currentHP,
		Money:     startingMoney,
		Inventory: inventory,
		Skills:    []Skill{baseSkill},
	}
}

func (c Character) DisplayInfo() {
	totalBlocks := 10
	filledBlocks := (c.CurrentHP * totalBlocks) / c.MaxHP
	if filledBlocks < 0 {
		filledBlocks = 0
	}
	emptyBlocks := totalBlocks - filledBlocks

	hpBar := strings.Repeat("█", filledBlocks) + strings.Repeat("░", emptyBlocks)

	invStr := "Aucun équipement"
	if len(c.Inventory) > 0 {
		invStr = strings.Join(c.Inventory, " | ")
	}

	var skillNames []string
	for _, s := range c.Skills {
		skillNames = append(skillNames, fmt.Sprintf("%s (%d dmg)", s.Name, s.Damage))
	}
	skillsStr := strings.Join(skillNames, " | ")

	nameLine := fmt.Sprintf("[IDENTIFIANT] : %s", c.Name)
	classLine := fmt.Sprintf("[ORIGINE]     : %s", c.Class)
	levelLine := fmt.Sprintf("[NIVEAU]      : %d", c.Level)
	moneyLine := fmt.Sprintf("[ARGENT]      : %d $ED", c.Money)
	hpLine := fmt.Sprintf("[SANTÉ]       : %d / %d [%s]", c.CurrentHP, c.MaxHP, hpBar)
	skillsLine := fmt.Sprintf("[COMPÉTENCES] : %s", skillsStr)
	invLine := fmt.Sprintf("[INVENTAIRE]  : %s", invStr)

	fmt.Println("┌──────────────────────────────────────────────────┐")
	fmt.Println("│           SYS.NET // FICHE SUJET                 │")
	fmt.Println("├──────────────────────────────────────────────────┤")
	fmt.Printf("│  %-48s│\n", nameLine)
	fmt.Printf("│  %-48s│\n", classLine)
	fmt.Printf("│  %-48s│\n", levelLine)
	fmt.Printf("│  %-48s│\n", moneyLine)
	fmt.Println("├──────────────────────────────────────────────────┤")
	fmt.Printf("│  %-48s│\n", hpLine)
	fmt.Println("├──────────────────────────────────────────────────┤")
	fmt.Printf("│  %-48s│\n", skillsLine)
	fmt.Printf("│  %-48s│\n", invLine)
	fmt.Println("└──────────────────────────────────────────────────┘")
}
