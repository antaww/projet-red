package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func CloseGame() {
	fmt.Println(">> Fermeture du jeu. <<")
	time.Sleep(300 * time.Millisecond)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println(">> Fermeture du jeu.. <<")
	time.Sleep(300 * time.Millisecond)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println(">> Fermeture du jeu... <<")
	time.Sleep(300 * time.Millisecond)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println(">> Fermeture du jeu. <<")
	time.Sleep(300 * time.Millisecond)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println(">> Fermeture du jeu.. <<")
	time.Sleep(300 * time.Millisecond)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println(">> Fermeture du jeu... <<")
	time.Sleep(300 * time.Millisecond)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	os.Exit(0)
}

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
		P1.useInventory()
	case "3":
		fmt.Println(">> Marchand <<")
		P1.Marchand()
	case "4":
		CloseGame()
	}
	time.Sleep(1 * time.Second)
	Menu()
}
