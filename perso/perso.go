package perso

import (
	"fmt"
	"strings"
)

type Character struct {
	Name         string
	Class        string
	Level        int
	MaxHP        int
	CurrentHP    int
	Money        int
	Inventory    []string
	MaxInventory int // Nouveau : Limite maximale de stockage
	Skills       []Skill
	Implants     []string
}

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
		startingMoney = 2500
		baseSkill = Skill{Name: "Coup de poing", Damage: 7}

	default:
		class = "Gosse des rues"
		maxHP = 120
		startingMoney = 1000
		baseSkill = Skill{Name: "Coup de poing", Damage: 7}
	}

	currentHP := (maxHP * 60) / 100

	return Character{
		Name:         name,
		Class:        class,
		Level:        1,
		MaxHP:        maxHP,
		CurrentHP:    currentHP,
		Money:        startingMoney,
		Inventory:    inventory,
		MaxInventory: 10, // Limite de départ
		Skills:       []Skill{baseSkill},
		Implants:     []string{},
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

	nameLine := fmt.Sprintf("[IDENTIFIANT] : %s", c.Name)
	classLine := fmt.Sprintf("[ORIGINE]     : %s", c.Class)
	levelLine := fmt.Sprintf("[NIVEAU]      : %d", c.Level)
	moneyLine := fmt.Sprintf("[ARGENT]      : %d $ED", c.Money)
	hpLine := fmt.Sprintf("[SANTÉ]       : %d / %d [%s]", c.CurrentHP, c.MaxHP, hpBar)

	// Cadre supérieur
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

	// Affichage des compétences ligne par ligne
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

	// Affichage des implants ligne par ligne
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
