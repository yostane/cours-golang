// Créer une fonction qui retourne un tableau de 20 nombres aléatoires pairs
// Créer une fonction qui prend une chaine de caractères en argument et affiche ses voyelles
// Créer une fonction qui prend une chaine de caractères en argument et retourne une nouvelle chaine de caractères contant ses voyelles

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomEvenNumbers() [20]int {
	numbers := [20]int{}
	for idx, _ := range numbers {
		r := rand.Intn(100)
		for ; r%2 != 0; r = rand.Intn(100) {
		}
		numbers[idx] = r
	}
	return numbers
}

func generateRandomEvenNumbersYanis(value int) []int {
	listI := make([]int, 0, value)
	for len(listI) != value {
		a := rand.Intn(1000)
		if a%2 == 0 {
			listI = append(listI, a)
		}
	}
	return listI
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(generateRandomEvenNumbers())
	fmt.Println(generateRandomEvenNumbers())
	fmt.Println(generateRandomEvenNumbersYanis(30))
}
