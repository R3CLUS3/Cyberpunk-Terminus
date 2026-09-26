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
	Skills    []Skill // Nouveau : liste des compétences
}

func InitCharacter(name string, class string, level int, maxHP int, currentHP int, inventory []string) Character {
	// Compétence de base
	defaultSkills := []Skill{
		{Name: "Coup de poing", Damage: 7},
	}

	return Character{
		Name:      name,
		Class:     class,
		Level:     level,
		MaxHP:     maxHP,
		CurrentHP: currentHP,
		Money:     1000,
		Inventory: inventory,
		Skills:    defaultSkills,
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

	// Pré-formatage de chaque ligne pour aligner la bordure droite
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
