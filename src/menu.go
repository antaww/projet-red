package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func CloseGame() {
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println(">> Fermeture du jeu. <<")
	time.Sleep(200 * time.Millisecond)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println(">> Fermeture du jeu.. <<")
	time.Sleep(200 * time.Millisecond)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println(">> Fermeture du jeu... <<")
	time.Sleep(200 * time.Millisecond)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println(">> Fermeture du jeu. <<")
	time.Sleep(200 * time.Millisecond)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println(">> Fermeture du jeu.. <<")
	time.Sleep(200 * time.Millisecond)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	fmt.Println(">> Fermeture du jeu... <<")
	time.Sleep(200 * time.Millisecond)
	os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
	os.Exit(0)
}

func Menu() {
	//Affichage choix menu
	fmt.Println("1 >> Afficher les infos du personnage")
	fmt.Println("2 >> Afficher l'inventaire du personnage")
	fmt.Println("3 >> Aller chez le marchand")
	fmt.Println("4 >> Aller chez le forgeron")
	fmt.Println("5 >> Quitter le jeu")
	m, _ := BR.ReadString('\n') //lire input joueur quand "entrée"
	m = strings.Replace(m, "\r\n", "", -1)
	switch m {
	case "1":
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		P1.DisplayInfo()
		Options()
	case "2":
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		P1.accesInventory()
		P1.useInventory()
		Options()
	case "3":
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Println(">> Marchand <<")
		P1.Marchand()
	case "4":
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		fmt.Println(">> Forgeron <<")
		P1.Forgeron()
	case "5":
		CloseGame()
	}
}

func Options() {
	fmt.Println()
	fmt.Println(">> Options <<")
	fmt.Println("1 >> Retour au menu principal")
	fmt.Println("2 >> Quitter le jeu")
	c, _ := BR.ReadString('\n') //lire input joueur quand "entrée"
	c = strings.Replace(c, "\r\n", "", -1)
	switch c {
	case "1":
		os.Stdout.WriteString("\x1b[3;J\x1b[H\x1b[2J")
		time.Sleep(200 * time.Millisecond)
		Menu()
	case "2":
		time.Sleep(200 * time.Millisecond)
		CloseGame()
	}
}
