package main

import "fmt"

func Âge() (int, string) {
	var (
		âge  int
		noms string
	)
	âge = 25
	noms = "Alice"
	fmt.Printf("Nom: %s\n", noms)
	fmt.Printf("Âge: %d ans\n", âge)
	return âge, noms
}

func Verifmajeur(âge int, nom string) {
	if âge >= 18 {
		fmt.Printf("%s est majeur\n", nom)
	} else {
		fmt.Printf("%s est mineur\n", nom)
	}
}

func main() {
	âge, nom := Âge()
	Verifmajeur(âge, nom)
}
