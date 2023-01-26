package main

import "fmt"

func showArrayExamples() {
	arr1 := [...]int{1, 2, 3}
	arr2 := [3]int{1, 2, 3}
	fmt.Println(arr1)
	fmt.Println(arr2)
	cars := [...]string{"Volvo", "BMW", "Ford", "Mazda"}
	fmt.Println(cars)
	fmt.Println(len(cars)) // comme en python (length)
	cars[0] = "Tesla"
	fmt.Println(cars)
}

func showSliceExamples() {
	// Slice -> Tranche en Anglais
	//mySlice1 := []int{}
	mySlice2 := []int{1, 2, 3}
	fmt.Println(len(mySlice2))
	fmt.Println(cap(mySlice2))
	// Slice de 5 éléments avec une capacité de 10
	mySlice3 := make([]int, 10, 20)
	fmt.Printf("myslice3 = %v\n", mySlice3)
	fmt.Printf("length = %d\n", len(mySlice3))
	fmt.Printf("capacity = %d\n", cap(mySlice3))
	arr2 := [...]int{1, 2, 3, 4, 5}
	mySlice4 := arr2[1:3]     // à gauche c'est inclus, à droite c'est exclu arr2[1] et arr[2] mais arr[3] sera excly
	mySlice5 := mySlice3[1:3] // on prend [1], [2] et on prend pas [3]
	fmt.Printf("myslice4 = %v - myslice5 = %v\n", mySlice4, mySlice5)
	fmt.Printf("cap myslice4 = %d - cap myslice5 = %d\n", cap(mySlice4), cap(mySlice5))
	mySlice4[0] = 22
	fmt.Printf("After change mySlice4[0] = %d - arr2[1] = %d\n", mySlice4[0], arr2[1])
	fmt.Printf("mySlice4 = %v - arr2 = %v\n", mySlice4, arr2)
	mySlice4AfterAppend := append(mySlice4, 10)
	fmt.Printf("mySlice4AfterAppend %v - mySlice4 %v - arr2 %v", mySlice4AfterAppend, mySlice4, arr2)
}

func main() {
	showArrayExamples()
	showSliceExamples()
}
