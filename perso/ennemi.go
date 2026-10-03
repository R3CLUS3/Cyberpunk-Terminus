// structure des ennemis etc..
package perso

import (
	"fmt"
	"math/rand"
)

type Monster struct {
	Name      string
	MaxHP     int
	CurrentHP int
	Attack    int
	RewardED  int
	RewardXP  int
}

// GenerateBot crée un bot adapté au niveau de la vague courante
func GenerateBot(wave int) Monster {
	names := []string{
		Yellow + "Bot de Securité Militech" + Reset,
		Red + "Drone d'Arasaka" + Reset,
		Blue + "Cyber-Trooper Kang Tao" + Reset,
		Cyan + "Cyber-Android de Combat Trauma Team" + Reset,
		magenta + "Proto-Mecha NetWatch" + Reset,
	}

	selectedName := names[rand.Intn(len(names))]
	fullName := fmt.Sprintf("%s [Vague %d]", selectedName, wave)

	// Les stats augmentent à chaque vague
	hp := 40 + (wave * 15)
	atk := 5 + (wave * 3)
	rewardED := 20 + (wave * 15)
	rewardXP := 25 + (wave * 10)

	return Monster{
		Name:      fullName,
		MaxHP:     hp,
		CurrentHP: hp,
		Attack:    atk,
		RewardED:  rewardED,
		RewardXP:  rewardXP,
	}
}

func (m *Monster) IsDead() bool {
	return m.CurrentHP <= 0
}
