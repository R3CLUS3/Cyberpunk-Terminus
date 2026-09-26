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
	Inventory []string
}

func InitCharacter(name string, class string, level int, maxHP int, currentHP int, inventory []string) Character {
	return Character{
		Name:      name,
		Class:     class,
		Level:     level,
		MaxHP:     maxHP,
		CurrentHP: currentHP,
		Inventory: inventory,
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

	fmt.Println("┌──────────────────────────────────────────────────┐")
	fmt.Println("│           SYS.NET // FICHE SUJET                 │")
	fmt.Println("├──────────────────────────────────────────────────┤")
	fmt.Printf("│  [IDENTIFIANT] : %-31s │\n", c.Name)
	fmt.Printf("│  [ORIGINE]     : %-31s │\n", c.Class)
	fmt.Printf("│  [NIVEAU]      : %-31d │\n", c.Level)
	fmt.Println("├──────────────────────────────────────────────────┤")
	fmt.Printf("│  [SANTÉ]       : %-3d / %-3d [%s]          │\n", c.CurrentHP, c.MaxHP, hpBar)
	fmt.Println("├──────────────────────────────────────────────────┤")
	fmt.Printf("│  [INVENTAIRE]  : %-31s │\n", invStr)
	fmt.Println("└──────────────────────────────────────────────────┘")
}
