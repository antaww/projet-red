package main

import (
	"fmt"
	"strings"
)

func (p *Personnage) accesInventory() {
	runeName := []rune(p.nom)
	if runeName[0] == 'A' || runeName[0] == 'E' || runeName[0] == 'I' || runeName[0] == 'O' || runeName[0] == 'U' {
		fmt.Print(">> Inventaire d'", P1.nom, " <<\n")
	} else {
		fmt.Print(">> Inventaire de ", P1.nom, " <<\n")
	}
	for key, val := range p.inventaire {
		fmt.Println(">>", key, ":", val)
	}
	if len(p.inventaire) == 0 {
		fmt.Println("Votre inventaire est vide")
	}
}

func (p *Personnage) useInventory(inFight bool) {
	//Affichage choix menu
	fmt.Println("1 >> Utiliser potion de soin")
	fmt.Println("2 >> Utiliser potion de poison")
	fmt.Println("3 >> Utiliser livre de sort : boule de feu")
	fmt.Println("4 >> Equiper le chapeau de l'aventurier")
	fmt.Println("5 >> Equiper la tunique de l'aventurier")
	fmt.Println("6 >> Equiper les bottes de l'aventurier")
	fmt.Println("7 >> Options")
	i, _ := BR.ReadString('\n') //lire input joueur quand "entrée"
	i = strings.Replace(i, "\r\n", "", -1)
	switch i {
	case "1":
		p.takePot()
	case "2":
		if p.inventaire["Potion de poison"] >= 1 {
			fmt.Println(">> Une potion de poison a été utilisée par", p.nom, "<<")
			p.poisonPot()
		} else {
			fmt.Println(">> Vous n'avez aucune potion de poison ! <<")
		}
	case "3":
		if p.inventaire["Livre de sort : boule de feu"] >= 1 && p.skill["Boule de feu"] == 1 {
			fmt.Println(">> Impossible d'avoir plusieurs boules de feu ! <<")
		}
		if p.inventaire["Livre de sort : boule de feu"] >= 1 && p.skill["Boule de feu"] == 0 {
			fmt.Println(">> Le livre de sort boule de feu a été utilisé par", p.nom, "<<")
			fmt.Println(">>", p.nom, "a appris boule de feu ! <<")
			p.spellBook()
		} else if p.inventaire["Livre de sort : boule de feu"] == 0 {
			fmt.Println(">> Vous n'avez pas de livre de sort ! <<")
		}
	case "4":
		p.EquipArmor("Chapeau de l'aventurier", 10)
	case "5":
		p.EquipArmor("Tunique de l'aventurier", 25)
	case "6":
		p.EquipArmor("Bottes de l'aventurier", 15)
	case "7":
		Options()
	}
	fmt.Println(">> Appuyez sur entrée pour continuer... ")
	Input()
	ClearLog()
	if inFight == false {
		P1.accesInventory()
		P1.useInventory(false)
	}
}
