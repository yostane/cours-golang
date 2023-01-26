package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Générer un tableau de 20 entiers aléatoires
// Ensuite, dans une autre boucle, créer un slice qui contient uniquement les nombres pairs de ce tableau
// Ensuite créer une fonction qui prend un slice d'entiers en argument et calcule la somme de ce slice

// Tableau: collection d'éléments contigus. Le nombre d'éléments est fixe.
// Slice: c'est comme un tableau, sauf qu'on peut réduire le nombre d'éléments ou l'augmenter tant qu'on ne dépasse pas la capacité max

func main() {
	// C'est une donnée qui permet de calculer le premier terme dans la suite numérique des nombres pseudo-aléatoires
	rand.Seed(time.Now().UnixNano())

	randomNumbers := [20]int{}
	// Initialisation du tableau de nombres aléatoires
	for idx, _ := range randomNumbers {
		randomNumbers[idx] = rand.Intn(100)
	}
	fmt.Println(randomNumbers)
	pairNumbers := make([]int, 0, 20)
	// Ajout des éléments pairs
	for _, val := range randomNumbers {
		if val%2 == 0 {
			pairNumbers = append(pairNumbers, val)
		}
	}
	fmt.Println(pairNumbers)
	sum := 0
	// Calcul de la somme
	for _, val := range randomNumbers {
		sum += val
	}
	fmt.Println(sum)
}
