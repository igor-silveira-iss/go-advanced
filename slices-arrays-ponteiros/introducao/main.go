package main

import "fmt"

func main() {
	var (
		slice = []int{1, 2, 3, 4, 5}
		array = [5]int{1, 2, 3, 4, 5}
	)

	fmt.Println("Tamanho do array:", len(array))
	fmt.Println("Capacidade do array:", cap(array))
	fmt.Println("\n")
	fmt.Println("Tamanho do slice:", len(slice))
	fmt.Println("Capacidade do slice:", cap(slice))

	slice = append(slice, 6)
	fmt.Println("\n")
	fmt.Println("Tamanho do slice:", len(slice))
	fmt.Println("Capacidade do slice:", cap(slice))
}
