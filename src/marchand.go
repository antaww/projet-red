package main

import (
	"fmt"
	"strings"
	"time"
)

func (p *Personnage) Marchand() {
	//Affichage choix menu
	fmt.Println("1 >> Acheter potion de soin (3 coins)")
	fmt.Println("2 >> Acheter potion de poison")
	fmt.Println("3 >> Acheter livre de sort : boule de feu")
	fmt.Println("4 >> Retour au menu principal")
	o, _ := BR.ReadString('\n') //lire input joueur quand "entrée"
	o = strings.Replace(o, "\r\n", "", -1)
	switch o {
	case "1":
		if p.money > 0 {
			p.money -= 3
			p.FreeHealPotion()
			fmt.Println(">> Une potion de soin a été achetée par", p.nom, "<<")
		} else {
			fmt.Println(">> Vous n'avez plus d'argent ! <<")
		}

	case "2":
		fmt.Println(">> Une potion de poison a été achetée par", p.nom, "<<")
		p.FreePoisonPotion()
	case "3":
		fmt.Println(">> Le livre de sort boule de feu a été acheté par", p.nom, "<<")
		p.FreeSpeelBook()
	case "4":
		fmt.Println(">> Retour au menu en cours... <<")
		time.Sleep(1 * time.Second)
		Menu()
	}
}
