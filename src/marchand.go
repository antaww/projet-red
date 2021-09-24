package main

import (
	"fmt"
	"strings"
	"time"
)

func (p *Personnage) Marchand() {
	//Affichage choix menu
	fmt.Println("1 >> Acheter potion de soin")
	fmt.Println("2 >> Acheter potion de poison")
	fmt.Println("3 >> Acheter livre de sort : boule de feu")
	fmt.Println("4 >> Retour au menu principal")
	o, _ := BR.ReadString('\n') //lire input joueur quand "entrée"
	o = strings.Replace(o, "\r\n", "", -1)
	switch o {
	case "1":
		fmt.Println(">> Une potion de soin a été achetée par", p.nom, "<<")
		p.FreeHealPotion()
	case "2":
		fmt.Println(">> Une potion de poison a été achetée par", p.nom, "<<")
		p.FreePoisonPotion()
	case "3":
		fmt.Println(">> Le livre de sort boule de feu a été acheté par", p.nom, "<<")
		p.spellBook()
	case "4":
		fmt.Println(">> Retour au menu en cours... <<")
		time.Sleep(1 * time.Second)
		Menu()
	}
}
