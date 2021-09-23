package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func Menu() {
	//Affichage choix menu
	fmt.Println("1 >> Afficher les infos du personnage")
	fmt.Println("2 >> Afficher l'inventaire du personnage")
	fmt.Println("3 >> Aller chez le marchand")
	fmt.Println("4 >> Quitter")
	m, _ := BR.ReadString('\n') //lire input joueur quand "entrée"
	m = strings.Replace(m, "\r\n", "", -1)
	switch m {
	case "1":
		fmt.Println(">> Stats du personnage", P1.nom, "<<")
		P1.DisplayInfo()
	case "2":
		fmt.Println(">> Inventaire d'", P1.nom, "<<")
		P1.accesInventory()
	case "3":
		fmt.Println(">> Marchand <<")
		P1.Marchand()
	case "4":
		fmt.Println(">> Fermeture du jeu... <<")
		time.Sleep(1 * time.Second)
		fmt.Println("3")
		time.Sleep(1 * time.Second)
		fmt.Println("2")
		time.Sleep(1 * time.Second)
		fmt.Println("1")
		os.Exit(0)
	}
	Menu()
}
