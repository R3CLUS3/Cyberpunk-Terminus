// menu principal du jeu cyberpunk. Il permet au joueur de créer son personnage, de choisir sa classe,
// et d'accéder à différentes fonctionnalités du jeu
// telles que l'inventaire, le marchand, l'atelier cyberware, les combats d'entraînement,
// les raids et les duels contre les légendes de Night City.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"Cyberpunk_Terminus/perso"
)

const (
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Reset   = "\033[0m"
	Cyan    = "\033[36m"
	magenta = "\033[35m"
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
                                                            
					[ SYSTEM CYBERPUNK TERMINUS // v1.0.2 ] 
										 				made by Rob1
`
	//v1.0.1 : maj majeur.maj mineur.fix bug
	fmt.Print(Yellow + Cyber + Reset)
	fmt.Println("=== INITIALISATION DU SYSTÈME CYBERPUNK ===")

	fmt.Print(Blue + "Entrez votre identifiant : " + Reset)
	nameInput, _ := reader.ReadString('\n')
	name := strings.TrimSpace(nameInput)

	// Choix de la classe avec boucle de validation
	var class string
	for {
		fmt.Println("\nChoisissez votre origine :")
		fmt.Println(Yellow + "1. Gosse des rues" + Reset)
		fmt.Println(Blue + "2. Netrunner" + Reset)
		fmt.Println(Red + "3. Corpo" + Reset)
		fmt.Print("Votre choix (1-3) : ")

		classInput, _ := reader.ReadString('\n')
		classChoice := strings.TrimSpace(classInput)

		if classChoice == "1" {
			class = "Gosse des rues"
			break
		} else if classChoice == "2" {
			class = "Netrunner"
			break
		} else if classChoice == "3" {
			class = "Corpo"
			break
		}

		fmt.Println(Red + "Choix invalide ! Veuillez saisir 1, 2 ou 3." + Reset)
	}

	// Création du personnage avec la classe choisie
	player := perso.InitCharacter(name, class)
	fmt.Println(Green+"\n[PROFIL CRÉÉ AVEC SUCCÈS :", class, "]"+Reset)
	promptReturn(reader)

	// --- BOUCLE DU MENU PRINCIPAL ---
	for {
		perso.ClearScreen() // Nettoie l'écran avant d'afficher le menu principal
		fmt.Println("┌──────────────────────────────────────────────────┐")
		fmt.Println("│             SYS.NET // MENU PRINCIPAL            │")
		fmt.Println("├──────────────────────────────────────────────────┤")
		fmt.Println(Green + "│  H. Aide & Guide du jeu                          │" + Reset)
		fmt.Println(Yellow + "│  1. informations du personnage                   │" + Reset)
		fmt.Println(magenta + "│  2. inventaire                                   │" + Reset)
		fmt.Println(Red + "│  3. Visiter le Marchand                          │" + Reset)
		fmt.Println(Cyan + "│  4. Atelier Cyberware                            │" + Reset)
		fmt.Println(magenta + "│  5. Lancer un combat d'entraînement              │" + Reset)
		fmt.Println(Yellow + "│  6. Raid Complexe Militech                       │" + Reset)
		fmt.Println(Red + "│  7. Raid Arasaka Tower                           │" + Reset)
		fmt.Println(Green + "│  8. Duel contre les Légendes de Night City       │" + Reset)
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
