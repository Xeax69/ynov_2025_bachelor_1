package main

import "fmt"

func Âge() {
	var (
		âge  int
		noms string
	)
	âge = 25
	noms = "Alice"
	fmt.Printf("Nom: %s\n", noms)
	fmt.Printf("Âge: %d ans\n", âge)
}

func Verifmajeur() {
	var personne int
	if personne >= 18 {
		fmt.Println("La personne est Majeur")
	} else {
		fmt.Println("La personne est mineur")
	}
}

func main() {
	Âge()
	Verifmajeur()
}
