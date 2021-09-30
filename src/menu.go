package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func CloseGame() {
	ClearLog()
	fmt.Println(">> Fermeture du jeu. <<")
	time.Sleep(200 * time.Millisecond)
	ClearLog()
	fmt.Println(">> Fermeture du jeu.. <<")
	time.Sleep(200 * time.Millisecond)
	ClearLog()
	fmt.Println(">> Fermeture du jeu... <<")
	time.Sleep(200 * time.Millisecond)
	ClearLog()
	fmt.Println(">> Fermeture du jeu. <<")
	time.Sleep(200 * time.Millisecond)
	ClearLog()
	fmt.Println(">> Fermeture du jeu.. <<")
	time.Sleep(200 * time.Millisecond)
	ClearLog()
	fmt.Println(">> Fermeture du jeu... <<")
	time.Sleep(200 * time.Millisecond)
	ClearLog()
	os.Exit(0)
}

func Menu() {
	//Affichage choix menu
	ClearLog()
	fmt.Println(">> Accueil <<")
	fmt.Println("1 >> Afficher les infos du personnage")
	fmt.Println("2 >> Afficher l'inventaire du personnage")
	fmt.Println("3 >> Aller chez le marchand")
	fmt.Println("4 >> Aller chez le forgeron")
	fmt.Println("5 >> Zone de combat")
	fmt.Println(BIWhite + "6 >> Quitter le jeu" + Reset)
	m, _ := BR.ReadString('\n') //lire input joueur quand "entrée"
	m = strings.Replace(m, "\r\n", "", -1)
	switch m {
	case "1":
		ClearLog()
		P1.DisplayInfo()
		Options()
	case "2":
		ClearLog()
		P1.accesInventory()
		P1.useInventory(false)
		Options()
	case "3":
		ClearLog()
		P1.Marchand()
	case "4":
		ClearLog()
		P1.Forgeron()
	case "5":
		trainingFight()
	case "6":
		CloseGame()
	}
}

func Options() {
	fmt.Println()
	fmt.Println(">> Options <<")
	fmt.Println("1 >> Retour au menu principal")
	fmt.Println(BIWhite + "2 >> Quitter le jeu" + Reset)
	c, _ := BR.ReadString('\n') //lire input joueur quand "entrée"
	c = strings.Replace(c, "\r\n", "", -1)
	switch c {
	case "1":
		ClearLog()
		time.Sleep(200 * time.Millisecond)
		Menu()
	case "2":
		time.Sleep(200 * time.Millisecond)
		CloseGame()
	}
}
