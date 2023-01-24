package main

import (
	"fmt"
	"os"
)

// Point d'entrée (le premier code qui sera exécuté en lançant le programme).
// Ici le point d'entrée est la func main
func main() {
	// Ecriture dans la sortie standard
	fmt.Print("Go")
	fmt.Println("Hello")
	fmt.Println("Hello")
	x, y := 10, 20
	message := "World"
	// %d -> Injection d'une vriable de type nombre entier
	fmt.Printf("La valeur de x est %d et y est %d. %s", x, y, message)
	// Dans l'écosystème UNIX / Linux. La sortie standard est considérée comme un fichier
	// os (Operating System - Système d'exploitation) comme Windows, Linux, Android, iOS
	// os: Permet d'interagir avec l'os -> les fichiers, le processus, etc.
	// os.Create permet de créer un fichier
	f, err := os.Create("test.txt")
	if err == nil {
		fmt.Fprintf(f, "La valeur de x est %d et y est %d. %s", x, y, message)
	}
	defer f.Close()
	fmt.Fprintf(os.Stdout, "La valeur de x est %d et y est %d. %s", x, y, message)
}
