// menu aide
package perso

import (
	"bufio"
	"fmt"
)

// DisplayHelp affiche un guide complet du jeu et des mécaniques
func DisplayHelp(reader *bufio.Reader) {
	ClearScreen()
	fmt.Println(Green + "┌──────────────────────────────────────────────────┐")
	fmt.Println("│           SYS.NET // GUIDE D'UTILISATION         │")
	fmt.Println("├──────────────────────────────────────────────────┤")
	fmt.Println("│  1. FICHE SUJET                                  │")
	fmt.Println("│     Consultez vos PV, $ED, XP, niveau et votre   │")
	fmt.Println("│     équipement actuel (compétences, armure).     │")
	fmt.Println("│                                                  │")
	fmt.Println("│  2. INVENTAIRE                                   │")
	fmt.Println("│     Gérez vos consommables (Stimpacks) et        │")
	fmt.Println("│     équipez vos plastrons/vestes.                │")
	fmt.Println("│                                                  │")
	fmt.Println("│  3. MARCHAND & FORGERON                          │")
	fmt.Println("│     Achetez des objets de soins ou fabriquer des │")
	fmt.Println("│     armes légendaires (Mantis, Malorian...) grâce│")
	fmt.Println("│     aux composants de rang C à S+.               │")
	fmt.Println("│                                                  │")
	fmt.Println("│  4. COMBAT D'ENTRAÎNEMENT                        │")
	fmt.Println("│     Arène infinie pour farmer des $ED et de l'XP.│")
	fmt.Println("│     Chaque niveau (+5 HP Max) vous rend plus fort│")
	fmt.Println("│                                                  │")
	fmt.Println("│  5. RAIDS (Militech / Arasaka)                   │")
	fmt.Println("│     Explorez des cartes 2D avec ZQSD. Éliminez   │")
	fmt.Println("│     les soldats (S) pour atteindre le Boss (B).  │")
	fmt.Println("│     Le Boss Militech drop le Cybersquelette !    │")
	fmt.Println("│                                                  │")
	fmt.Println("│  6. DUEL DES LÉGENDES                            │")
	fmt.Println("│     Affrontez directement au choix : Rebecca,    │")
	fmt.Println("│     Lucy, David, V ou Adam Smasher.              │")
	fmt.Println("│                                                  │")
	fmt.Println("│  7. Easter egg                                   │")
	fmt.Println("│     ce n'est pas n'importe qui, qui peut porter  │")
	fmt.Println("│     la veste de David Martinez !                 │")
	fmt.Println("└──────────────────────────────────────────────────┘" + Reset)
	fmt.Print("\nAppuyez sur Entrée pour revenir au menu principal...")
	reader.ReadString('\n')
}
