package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func NewName() string {
	nom := Input()
	if nom == "" {
		ClearLog()
		fmt.Println(">> Entrez votre nom... <<")
		nom = NewName()
	}
	return nom
}

func charCreation() {
	ClearLog()
	fmt.Println(BIWhite + "███╗░░░███╗██╗███╗░░██╗███████╗░█████╗░██████╗░░█████╗░███████╗████████╗" + Reset)
	fmt.Println(BIWhite + "████╗░████║██║████╗░██║██╔════╝██╔══██╗██╔══██╗██╔══██╗██╔════╝╚══██╔══╝" + Reset)
	fmt.Println(BIWhite + "██╔████╔██║██║██╔██╗██║█████╗░░██║░░╚═╝██████╔╝███████║█████╗░░░░░██║░░░" + Reset)
	fmt.Println(BIWhite + "██║╚██╔╝██║██║██║╚████║██╔══╝░░██║░░██╗██╔══██╗██╔══██║██╔══╝░░░░░██║░░░" + Reset)
	fmt.Println(BIWhite + "██║░╚═╝░██║██║██║░╚███║███████╗╚█████╔╝██║░░██║██║░░██║██║░░░░░░░░██║░░░" + Reset)
	fmt.Println(BIWhite + "╚═╝░░░░░╚═╝╚═╝╚═╝░░╚══╝╚══════╝░╚════╝░╚═╝░░╚═╝╚═╝░░╚═╝╚═╝░░░░░░░░╚═╝░░░" + Reset)
	Input()
	ClearLog()
	fmt.Println(BIWhite + "░░░" + Reset)
	fmt.Println(BIWhite + "░░░" + Reset)
	fmt.Println(BIWhite + "░░░" + Reset)
	fmt.Println(BIWhite + "░░░" + Reset)
	fmt.Println(BIWhite + "██╗" + Reset)
	fmt.Println(BIWhite + "╚═╝" + Reset)
	time.Sleep(500 * time.Millisecond)
	ClearLog()
	fmt.Println(BIWhite + "░░░░░░" + Reset)
	fmt.Println(BIWhite + "░░░░░░" + Reset)
	fmt.Println(BIWhite + "░░░░░░" + Reset)
	fmt.Println(BIWhite + "░░░░░░" + Reset)
	fmt.Println(BIWhite + "██╗██╗" + Reset)
	fmt.Println(BIWhite + "╚═╝╚═╝" + Reset)
	time.Sleep(500 * time.Millisecond)
	ClearLog()
	fmt.Println(BIWhite + "░░░░░░░░░" + Reset)
	fmt.Println(BIWhite + "░░░░░░░░░" + Reset)
	fmt.Println(BIWhite + "░░░░░░░░░" + Reset)
	fmt.Println(BIWhite + "░░░░░░░░░" + Reset)
	fmt.Println(BIWhite + "██╗██╗██╗" + Reset)
	fmt.Println(BIWhite + "╚═╝╚═╝╚═╝" + Reset)
	time.Sleep(500 * time.Millisecond)
	ClearLog()
	name := getName()
	ClassMenu(name)
}

func getName() string {
	fmt.Println(">> Entrez votre nom... <<")
	nom := NewName()
	nom = Capitalize(nom)
	rand.Seed(time.Now().UnixNano())
	nbr := (rand.Intn(5))
	if nom == "Alan" {
		nbr = 6
		fmt.Print(">> Bienvenue, ", BIYellow+"Dieu"+Reset, ". <<\n")
	}
	if nom == "Lucas" {
		nom = "quequette87"
		nbr = 7
		fmt.Print(">> Oh ! C'est toi ", BIPurple+nom+Reset, " ! Je me souviens de toi ! <<\n")
	}
	if nbr == 0 {
		fmt.Print(">> Oh, ", BIPurple+nom+Reset, " te va à merveille ! <<\n")
	}
	if nbr == 1 {
		fmt.Print(">> Ah... ", BIPurple+nom+Reset, ", chacun ses goûts... <<\n")
	}
	if nbr == 2 {
		fmt.Print(">> ", BIPurple+nom+Reset, " t'es sûr que tu t'es pas trompé ? <<\n")
	}
	if nbr == 3 {
		nom = "Lucas"
		fmt.Print(">> D'accord, ", BIPurple+nom+Reset, ".<<\n")
	}
	if nbr == 4 {
		nom = nom[:len(nom)/2]
		fmt.Print(">> Bref... ", BIPurple+nom+Reset, ".<<\n")
	}
	fmt.Println(BIWhite + ">> Appuyez sur entrée pour continuer... " + Reset)
	Input()
	return nom
}

func confirmClass(nom string) bool {
	fmt.Println("1 >> Confirmer votre choix")
	fmt.Println("Autre >> Retour aux choix des classes")
	if Input() == "1" {
		return true
	}
	return false
}

func ClassMenu(nom string) {
	var classe string
	var HP, maxHP int
	var skill map[string]int
	loopClass := true
	for loopClass {
		if nom == "Alan" {
			classe = "Dieu"
			maxHP = 1000
			HP = maxHP
			skill = map[string]int{"Gear Four": 1}
			break
		}
		ClearLog()
		fmt.Println(BIWhite + ">> Quelle classe souhaitez-vous jouer ? <<" + Reset)
		fmt.Println("1 >> Humain")
		fmt.Println("2 >> Elfe")
		fmt.Println("3 >> Nain")
		fmt.Println(BIGreen+"4 >>"+Reset, "Esprit de la forêt")
		fmt.Println(BIBlue+"5 >>"+Reset, "Sirène")
		fmt.Println("6 >> Quitter le jeu")
		c, _ := BR.ReadString('\n') //lire input joueur quand "entrée"
		c = strings.Replace(c, "\r\n", "", -1)
		switch c {
		case "1":
			classe = "Humain"
			maxHP = 100
			HP = maxHP / 2
			skill = map[string]int{"Coup de poing": 1}
			ClearLog()
			fmt.Println(BIWhite+"Classe :", classe+Reset)
			fmt.Println("HP max :", maxHP)
			fmt.Println("HP actuels :", HP)
			for key := range skill {
				fmt.Println("Compétence :", key)
			}
			if confirmClass(nom) {
				loopClass = false
			}

		case "2":
			classe = "Elfe"
			maxHP = 80
			HP = maxHP / 2
			skill = map[string]int{"Coup de poing": 1}
			ClearLog()
			fmt.Println(BIWhite+"Classe :", classe+Reset)
			fmt.Println("HP max :", maxHP)
			fmt.Println("HP actuels :", HP)
			for key := range skill {
				fmt.Println("Compétence :", key)
			}
			if confirmClass(nom) {
				loopClass = false
			}
		case "3":
			classe = "Nain"
			maxHP = 120
			HP = maxHP / 2
			skill = map[string]int{"Coup de poing": 1}
			ClearLog()
			fmt.Println(BIWhite+"Classe :", classe+Reset)
			fmt.Println("HP max :", maxHP)
			fmt.Println("HP actuels :", HP)
			for key := range skill {
				fmt.Println("Compétence :", key)
			}
			if confirmClass(nom) {
				loopClass = false
			}
		case "4":
			classe = "Esprit de la forêt"
			maxHP = 170
			HP = 140
			skill = map[string]int{"Tempête verte": 1}
			ClearLog()
			fmt.Println(BIWhite+"Classe :", classe+Reset)
			fmt.Println("HP max :", maxHP)
			fmt.Println("HP actuels :", HP)
			for key := range skill {
				fmt.Println("Compétence :", key)
			}
			if confirmClass(nom) {
				loopClass = false
			}
		case "5":
			classe = "Sirène"
			maxHP = 160
			HP = 160
			skill = map[string]int{"Danse pluie": 1}
			ClearLog()
			fmt.Println(BIWhite+"Classe :", classe+Reset)
			fmt.Println("HP max :", maxHP)
			fmt.Println("HP actuels :", HP)
			for key := range skill {
				fmt.Println("Compétence :", key)
			}
			if confirmClass(nom) {
				loopClass = false
			}
		case "6":
			CloseGame()
		}
	}
	if nom != "Alan" {
		if classe != "Sirène" {
			fmt.Println(">> Vous êtes désormais un", classe, "! <<")
		} else {
			fmt.Println(">> Vous êtes désormais une", classe, "! <<")
		}
		fmt.Println(BIWhite + ">> Appuyez sur entrée pour continuer... " + Reset)
		Input()
		time.Sleep(1 * time.Nanosecond)
	}
	P1.Init(BIPurple+nom+Reset, classe, 1, maxHP, HP, 5, 10, map[string]int{"Potion de soin": 3}, true, skill, 100,
		"Casque", "Plastron", "Bottes", 0)
}
