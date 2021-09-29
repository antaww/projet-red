package main

import (
	"fmt"
	"strings"
)

func NewName() string {
	nom := Input()
	if nom == "" {
		ClearLog()
		fmt.Println(">> Minecraft <<")
		fmt.Println(">> Entrer votre nom... <<")
		nom = NewName()
	}
	return nom
}

func charCreation() {
	var classe string
	var HP, maxHP int
	ClearLog()
	fmt.Println(">> Minecraft <<")
	fmt.Println(">> Entrer votre nom... <<")
	nom := NewName()
	nom = Capitalize(nom)
	fmt.Println(">> Oh,", nom, "te va à merveille ! <<")
	fmt.Println(">> Appuyez sur entrée pour continuer... ")
	Input()
	ClearLog()
	fmt.Println(">> Quelle classe souhaitez-vous jouer ? <<")
	fmt.Println("1 >> Humain")
	fmt.Println("2 >> Elfe")
	fmt.Println("3 >> Nain")
	fmt.Println("4 >> Quitter le jeu")
	c, _ := BR.ReadString('\n') //lire input joueur quand "entrée"
	c = strings.Replace(c, "\r\n", "", -1)
	switch c {
	case "1":
		classe = "Humain"
		maxHP = 100
		HP = maxHP / 2
	case "2":
		classe = "Elfe"
		maxHP = 80
		HP = maxHP / 2
	case "3":
		classe = "Nain"
		maxHP = 120
		HP = maxHP / 2
	case "4":
		CloseGame()
	}
	fmt.Println(">> Vous êtes désormais un", classe, "! <<")
	fmt.Println(">> Appuyez sur entrée pour continuer... ")
	Input()
	P1.Init(Purple+nom+Reset, classe, 1, maxHP, HP, 5, map[string]int{"Potion de soin": 3}, true, map[string]int{"Coup de poing": 1}, 100,
		"Casque", "Plastron", "Bottes")

}
