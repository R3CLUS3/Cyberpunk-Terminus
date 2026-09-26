package perso

import "fmt"

// Skill représente une capacité du personnage
type Skill struct {
	Name   string
	Damage int
}

// AddSkillAjoute une compétence si elle n'est pas déjà possédée
func (c *Character) AddSkill(newSkill Skill) bool {
	for _, s := range c.Skills {
		if s.Name == newSkill.Name {
			return false // Déjà possédée
		}
	}
	c.Skills = append(c.Skills, newSkill)
	return true
}

// PuceDeCombat applique l'amélioration du Coup de poing (+5 dégâts = 12 au total)
func (c *Character) PuceDeCombat() {
	// Vérifie si la puce de combat est déjà active
	for _, s := range c.Skills {
		if s.Name == "Puce de combat" {
			fmt.Println("\n⚠️ La Puce de combat est déjà installée dans votre système !")
			return
		}
	}

	// Ajout de la puce dans la liste des compétences
	c.AddSkill(Skill{Name: "Puce de combat", Damage: 0})

	// Renforcement du Coup de poing (7 -> 12)
	for i, s := range c.Skills {
		if s.Name == "Coup de poing" {
			c.Skills[i].Damage = 12
			break
		}
	}

	fmt.Println("\n⚙️ [IMPLANT INSTALLÉ] Puce de combat activée !")
	fmt.Println("💪 Ton Coup de poing passe à 12 points de dégâts !")
}
