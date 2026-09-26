package perso

import (
	"fmt"
)

// ClearScreen efface le terminal et remet le curseur en haut à gauche
func ClearScreen() {
	fmt.Print("\033[H\033[2J")
}
