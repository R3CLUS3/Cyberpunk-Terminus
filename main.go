package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"Cyberpunk_Terminus/perso"
)

const (
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Reset  = "\033[0m"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	Cyber := `

 ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄    ▄▄▄  ░▄▄▄▄▄▄▄▄▄▄▄▄      ▄▄▄▄▄▄▄▄▄▄▄ ░▄▄▄▄▄▄▄▄▄     ░▄▄▄▄▄▄▄▄▄     ▄▄▄    ▄▄▄   ▄▄▄▄▄▄▄▄▄▄   ▄▄▄    ▄▄▄ 
▒          █ ▒   █  █   █ ▓           ▀▀▄   ▒          █ ▓         ▀▄   ▓         ▀▄  ▒   █  █   █ ▒          █ ▒   █  █   █
▀▀▀▀▀▀▀▀▀▀▀▀ ▓   █  █   █ ▀▀▀▀▀▀▀▀▄▄     ▀▄ ▀▀▀▀▀▀▀▀▀▀▀▀ ▀▀▀▀▀▀▀▄    █  ▀▀▀▀▀▀▀▄    █ ▀   █  █   █ ▀▀▀▀▀▀▀█   █ ▀▀▀▀▀  █   █
 ▄▄▄         █   ▀▄▄▀   █  ▄▄▄▄▄    ▀▄    █  ▄▄▄▄▄▄▄▄▄▄▄  ▄▄▄▄▄▄▀    █   ▄▄▄▄▄▄▀    █ █   █  ▒   █ █▀▀▀█  ▒   █ █▀▀▀█▄▀    █
▒   █        ▒          █ ▒    █ ▄▀▀▀   ▄▀  ▒          █ ▒          ▄▀  ▒          ▄▀ ▒   █  ▓   █ ▒   █  ▓   █ ▒        ▄▀ 
▓   █        ▀▄▄▄▄▄▄▄   █ ▓    █ ▀▄▄▄▄   ▀▄ ▓   ▄▄▄▄▄▄▄▀ ▓    ▄▄   ▀▄   ▓    ▄▄▄▄▄▀   ▓   ▀  ▓   █ ▓   █  ▓   █ ▓   ▄▄    ▀▄
░   █         ▄▄▄   █   █ ░    █  ▄▄▀     █ ░   ▄        ░    ▀ ▀▄   ▀▄ ░    ▀        ░   ▄  █   █ ░   ▄  █   █ ░   ▄ ▀▄   █
▒   ▀▀▀▀▀▀▀█ ▒   █▄▄▀   █ ▒    ▀▀▀       ▄▀ ▒   ▀▀▀▀▀▀▀█ ▒    ▄  ▓    █ ▒    ▄        ▒   █▄▄▀   █ ▒   █  ▀   █ ▒   █  ▒   █
▓          █ ▓          █ ▓           ▄▄▀   ▓          █ ▓    █  █    █ ▓    █        ▓          █ ▓   █  █   █ ▓   █  ▓   █
▀▀▀▀▀▀▀▀▀▀▀▀ ▀▀▀▀▀▀▀▀▀▀▀▀ ▀▀▀▀▀▀▀▀▀▀▀▀      ▀▀▀▀▀▀▀▀▀▀▀▀ ▀▀▀▀▀▀  ▀▀▀▀▀▀ ▀▀▀▀▀▀         ▀▀▀▀▀▀▀▀▀▀   ▀▀▀   ▀▀▀▀  ▀▀▀▀▀  ▀▀▀▀▀
                                                            
										[ SYSTEM CYBERPUNK TERMINUS // v1.0 ]
`
	fmt.Print(Yellow + Cyber + Reset)
	fmt.Println("=== INITIALISATION DU SYSTÈME CYBERPUNK ===")

	fmt.Print("Entrez votre identifiant (Nom) : ")
	nameInput, _ := reader.ReadString('\n')
	name := strings.TrimSpace(nameInput)

	// Choix de la classe
	fmt.Println("\nChoisissez votre origine :")
	fmt.Println("1. Gosse des rues ")
	fmt.Println("2. Netrunner      ")
	fmt.Println("3. Corpo          ")
	fmt.Print("Votre choix (1-3) : ")

	classInput, _ := reader.ReadString('\n')
	classChoice := strings.TrimSpace(classInput)

	var class string
	switch classChoice {
	case "1":
		class = "Gosse des rues"
	case "2":
		class = "Netrunner"
	case "3":
		class = "Corpo"
	default:
		class = "Gosse des rues"
	}

	// Création du personnage avec inventaire de départ
	player := perso.InitCharacter(name, class, []string{"Cyberdeck v1", "Stimpack"})

	fmt.Println("\n[PROFIL CRÉÉ AVEC SUCCÈS]")

	// --- BOUCLE DU MENU PRINCIPAL ---
	for {
		fmt.Println(Blue + " \n┌──────────────────────────────────────────────────┐")
		fmt.Println("│             SYS.NET // MENU PRINCIPAL            │")
		fmt.Println("├──────────────────────────────────────────────────┤")
		fmt.Println("│  1. Afficher les informations du personnage      │")
		fmt.Println("│  2. Accéder au contenu de l'inventaire           │")
		fmt.Println("│  3. Marché noir                                  │")
		fmt.Println("│  4. Quitter                                      │")
		fmt.Println("└──────────────────────────────────────────────────┘" + Reset)
		fmt.Print(Green + "Entrez votre choix (1-4) : " + Reset)

		choiceInput, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(choiceInput)

		fmt.Println() // Ligne d'espacement pour la lisibilité

		switch choice {
		case "1":
			// Informations du personnage
			player.DisplayInfo()
			promptReturn(reader)

		case "2":
			// Inventaire
			player.AccessInventory()
			promptReturn(reader)

		case "3":
			// Visite chez le Charcudoc
			player.DisplayVendeur()

		case "4":
			//  Quitter
			fmt.Println(Red + "┌──────────────────────────────────────────────────┐")
			fmt.Println("│          DÉCONNEXION DU SYSTÈME... BYE BYE.      │")
			fmt.Println("└──────────────────────────────────────────────────┘" + Reset)
			return // Arrête le programme

		default:
			fmt.Println(Red + "Choix invalide. Veuillez saisir 1, 2, 3 ou 4." + Reset)
		}
	}
}

// Function utilitaire pour faire une pause et attendre l'entrée utilisateur ("Retour")
func promptReturn(reader *bufio.Reader) {
	fmt.Print("\n[Appuyez sur Entrée pour revenir au menu principal...]")
	reader.ReadString('\n')
}
