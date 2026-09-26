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
										 														made by Rob1
`
	fmt.Print(Yellow + Cyber + Reset)
	fmt.Println("=== INITIALISATION DU SYSTÈME CYBERPUNK ===")

	fmt.Print("Entrez votre identifiant : ")
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
	player := perso.InitCharacter(name, class)
	fmt.Println("\n[PROFIL CRÉÉ AVEC SUCCÈS]")

	// --- BOUCLE DU MENU PRINCIPAL ---
	for {
		perso.ClearScreen() // Nettoie l'écran avant d'afficher le menu principal
		fmt.Println("┌──────────────────────────────────────────────────┐")
		fmt.Println("│             SYS.NET // MENU PRINCIPAL            │")
		fmt.Println("├──────────────────────────────────────────────────┤")
		fmt.Println("│  H. Aide & Guide du jeu                          │")
		fmt.Println("│  1. Afficher les informations du personnage      │")
		fmt.Println("│  2. Accéder au contenu de l'inventaire           │")
		fmt.Println("│  3. Visiter le Marchand                          │")
		fmt.Println("│  4. Atelier Cyberware (Forgeron)                 │")
		fmt.Println("│  5. Lancer un combat d'entraînement              │")
		fmt.Println("│  6. Raid Complexe Militech (Mini-carte)          │")
		fmt.Println("│  7. Raid Arasaka Tower (Mini-carte)              │")
		fmt.Println("│  8. Duel contre les Légendes de Night City       │")
		fmt.Println("│  9. Quitter                                      │")
		fmt.Println("└──────────────────────────────────────────────────┘")
		fmt.Print("Entrez votre choix (1-9) : ")

		choiceInput, _ := reader.ReadString('\n')
		choice := strings.ToLower(strings.TrimSpace(choiceInput))

		switch choice {
		case "h":
			perso.DisplayHelp(reader)
		case "1":
			perso.ClearScreen()
			player.DisplayInfo()
			promptReturn(reader)
		case "2":
			perso.ClearScreen()
			player.AccessInventory()
		case "3":
			perso.ClearScreen()
			player.DisplayVendeur()
		case "4":
			perso.ClearScreen()
			player.DisplayForgeron()
		case "5":
			perso.ClearScreen()
			player.StartTrainingFight()
			promptReturn(reader)
		case "6":
			perso.ClearScreen()
			player.StartRaidMilitech()
			promptReturn(reader)
		case "7":
			perso.ClearScreen()
			player.StartRaidArasaka()
			promptReturn(reader)
		case "8":
			perso.ClearScreen()
			player.StartLegendDuel()

		case "9":
			fmt.Println(Red + "┌──────────────────────────────────────────────────┐")
			fmt.Println("│          DÉCONNEXION DU SYSTÈME... BYE.          │")
			fmt.Println("└──────────────────────────────────────────────────┘" + Reset)
			return
		default:
			fmt.Println(" Choix invalide. Veuillez saisir un numéro entre 1 et 7.")
		}
	}
}

// Function utilitaire pour faire une pause et attendre l'entrée utilisateur ("Retour")
func promptReturn(reader *bufio.Reader) {
	fmt.Print("\n[Appuyez sur Entrée pour revenir au menu principal...]")
	reader.ReadString('\n')
}
