package main

import "fmt"

func Temps() {
	var (
		tempsmax int
		tempsmin int
	)
	tempsmax = 13
	tempsmin = -2
	fmt.Printf("Temps maximal aujourd'hui: %d\n", tempsmax)
	fmt.Printf("Temps minimal aujourd'hui: %d\n", tempsmin)
}

func main() {
	Temps()
}
