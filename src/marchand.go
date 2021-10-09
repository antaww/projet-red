package main

import (
	"fmt"
	"strings"
)

func (p *Personnage) MoneyCheck(prix int) bool { //vérifier si le joueur a de quoi payer puis payer
	if p.money < prix {
		fmt.Println(">> Vous n'avez pas assez d'argent ! <<")
	}
	if p.money >= 0 {
		if (p.money - prix) >= 0 {
			p.money -= prix
			return true
		}
	}
	return false
}

func (p *Personnage) Marchand() {
	//Affichage choix menu
	ClearLog()
	fmt.Println(BIWhite + ">> Karine Le Marchand <<" + Reset)
	fmt.Println(">> Vous avez", p.money, "coins ! <<")
	fmt.Println("1 >> Acheter potion de soin (3 coins)")
	fmt.Println("2 >> Acheter potion de poison (6 coins)")
	fmt.Println("3 >> Acheter livre de sort : boule de feu (25 coins)")
	fmt.Println("4 >> Acheter fourrure de loup (4 coins)")
	fmt.Println("5 >> Acheter peau de troll (7 coins)")
	fmt.Println("6 >> Acheter cuir de sanglier (3 coins)")
	fmt.Println("7 >> Acheter plume de corbeau (1 coin)")
	fmt.Println("8 >> Augmenter la taille de l'inventaire de 10 (30 coins)")
	fmt.Println(BIWhite + "e >> Options" + Reset)
	o, _ := BR.ReadString('\n') //lire input joueur quand "entrée"
	o = strings.Replace(o, "\r\n", "", -1)
	switch o {
	case "1":
		if p.CheckSize() == true {
			if p.MoneyCheck(3) {
				p.GiveItem(potionSoin)
				fmt.Println(">> Une potion de soin a été acheté par", p.nom, "<<")
			}
		}

	case "2":
		if p.MoneyCheck(6) {
			p.GiveItem("Potion de poison")
			fmt.Println(">> Une potion de poison a été acheté par", p.nom, "<<")
		}

	case "3":
		if p.CheckSize() == true {
			if p.MoneyCheck(25) {
				p.GiveItem("Livre de sort : boule de feu")
				fmt.Println(">> Le livre de sort boule de feu a été acheté par", p.nom, "<<")
			}
		}

	case "4":
		if p.CheckSize() == true {
			if p.MoneyCheck(4) {
				p.GiveItem("Fourrure de loup")
				fmt.Println(">> La fourrure du loup a été achetée par", p.nom, "<<")
			}
		}

	case "5":
		if p.CheckSize() == true {
			if p.MoneyCheck(7) {
				p.GiveItem("Peau de troll")
				fmt.Println(">> La peau de troll a été achetée par", p.nom, "<<")
			}
		}

	case "6":
		if p.CheckSize() == true {
			if p.MoneyCheck(3) {
				p.GiveItem("Cuir de sanglier")
				fmt.Println(">> Le cuir de sanglier a été acheté par", p.nom, "<<")
			}
		}

	case "7":
		if p.CheckSize() == true {
			if p.MoneyCheck(1) {
				p.GiveItem("Plume de corbeau")
				fmt.Println(">> La plume de corbeau a été achetée par", p.nom, "<<")
			}
		}

	case "8":
		if p.MoneyCheck(30) {
			p.upgradeInventorySlot()
		}

	case "e":
		Options()
	}
	fmt.Println(BIWhite + ">> Appuyez sur entrée pour continuer... " + Reset)
	Input()
	p.Marchand()
}
