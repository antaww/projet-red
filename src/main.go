package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var P1 Personnage
var BR *bufio.Reader

func main() {
	BR = bufio.NewReader(os.Stdin)
	InitCharacter()
	Menu()
}

func Menu() {
	//Affichage choix menu
	fmt.Println("1 >> Afficher les infos du personnage")
	fmt.Println("2 >> Afficher l'inventaire du personnage")
	fmt.Println("3 >> Quitter")
	m, _ := BR.ReadString('\n') //lire input joueur quand "entrée"
	m = strings.Replace(m, "\r\n", "", -1)
	switch m {
	case "1":
		fmt.Println(">> Stats du personnage", P1.nom, "<<")
		P1.DisplayInfo()
	case "2":
		fmt.Println(">>>>>>>")
		fmt.Println("Affichage de l'inventaire du personnage...")
		fmt.Println(">>>>>>>")
		P1.accesInventory()
		fmt.Println(">>>>>>>")
	case "3":
		os.Exit(0)
	}
	Menu()
}
