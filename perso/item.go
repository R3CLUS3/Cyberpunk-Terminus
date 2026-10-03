// fonctions etc des items, armes, etc
package perso

import "strings"

type WeaponRecipe struct {
	Name        string
	Price       int
	ReqComps    []string // Liste de TOUS les composants requis
	SkillResult Skill
}

func GetAvailableRecipes() []WeaponRecipe {
	return []WeaponRecipe{
		{
			Name:        "Monocâble Cyberware",
			Price:       80,
			ReqComps:    []string{"Composant Rang C"},
			SkillResult: Skill{Name: "Monocâble", Damage: 20},
		},
		{
			Name:        "Lames Mantis",
			Price:       150,
			ReqComps:    []string{"Composant Rang C", "Composant Rang B"},
			SkillResult: Skill{Name: "Lames Mantis", Damage: 30},
		},
		{
			Name:        "Pistolet Malorian 3516",
			Price:       250,
			ReqComps:    []string{"Composant Rang C", "Composant Rang B", "Composant Rang A"},
			SkillResult: Skill{Name: "Malorian 3516", Damage: 45},
		},
		{
			Name:        "Fusil Smart 'Yinglong'",
			Price:       400,
			ReqComps:    []string{"Composant Rang C", "Composant Rang B", "Composant Rang A", "Composant Rang S"},
			SkillResult: Skill{Name: "Yinglong Smart", Damage: 65},
		},
		{
			Name:        "Katana Thermique 'Errate'",
			Price:       600,
			ReqComps:    []string{"Composant Rang C", "Composant Rang B", "Composant Rang A", "Composant Rang S", "Composant Rang S+"},
			SkillResult: Skill{Name: "Errate Thermique", Damage: 90},
		},
	}
}

// Fonction utilitaire pour abréger l'affichage des composants (ex: [C, B, A])
func FormatReqComps(comps []string) string {
	var ranks []string
	for _, c := range comps {
		if strings.Contains(c, "Rang C") {
			ranks = append(ranks, "C")
		} else if strings.Contains(c, "Rang B") {
			ranks = append(ranks, "B")
		} else if strings.Contains(c, "Rang A") {
			ranks = append(ranks, "A")
		} else if strings.Contains(c, "Rang S+") {
			ranks = append(ranks, "S+")
		} else if strings.Contains(c, "Rang S") {
			ranks = append(ranks, "S")
		}
	}
	return "[" + strings.Join(ranks, ",") + "]"
}
