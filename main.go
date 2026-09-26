package main

import (
	"Cyberpunk_Terminus/perso"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=== INITIALISATION DU SYSTÈME CYBERPUNK ===")

	fmt.Print("Entrez votre identifiant (Nom) : ")
	nameInput, _ := reader.ReadString('\n')
	name := strings.TrimSpace(nameInput)

	fmt.Println("\nChoisissez votre origine :")
	fmt.Println("1. Gosse des rues")
	fmt.Println("2. Corpo")
	fmt.Println("3. Netrunner")
	fmt.Print("Votre choix (1-3) : ")

	classInput, _ := reader.ReadString('\n')
	classChoice := strings.TrimSpace(classInput)

	var class string
	switch classChoice {
	case "1":
		class = "Gosse des rues"
	case "2":
		class = "Corpo"
	case "3":
		class = "Netrunner"
	default:
		class = "Gosse des rues"
	}

	player := perso.InitCharacter(name, class, 1, 100, 100, []string{"Cyberdeck v1", "Stimpack"})

	fmt.Println("\n[PROFIL CRÉÉ AVEC SUCCÈS]")
	player.DisplayInfo()
}
