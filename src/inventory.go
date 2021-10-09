package main

import (
	"fmt"
	"strings"
)

func (p *Personnage) upgradeInventorySlot() {
	if counterInv <= 3 {
		counterInv += 1
		p.invLimit += 10
		fmt.Println(">> La taille de votre inventaire est désormais de", p.invLimit, "! <<")
	} else {
		fmt.Println(">> Impossible d'augmenter la capacité de l'inventaire plus de 3 fois ! <<")
	}
}

func (p *Personnage) CheckSize() bool {
	if InvSize(p.inventaire) < p.invLimit {
		return true
	} else {
		fmt.Println(">> Votre inventaire est plein ! <<")
	}
	return false
}

func (p *Personnage) GiveItem(item string) {
	p.inventaire[item] += 1
	//ici pas de print "a été acheté" pour éviter les fautes d'orthographe
}

func InvSize(inventaire map[string]int) int {
	counter := 0
	for _, each := range inventaire {
		counter += each
	}
	return counter
}

func (p *Personnage) accesInventory() {
	runeName := []rune(p.nom)
	if runeName[0] == 'A' || runeName[0] == 'E' || runeName[0] == 'I' || runeName[0] == 'O' || runeName[0] == 'U' {
		fmt.Print(BIWhite+">> Inventaire d'"+Reset, P1.nom, BIWhite+" <<\n"+Reset)
	} else {
		fmt.Print(BIWhite+">> Inventaire de "+Reset, P1.nom, BIWhite+" <<\n"+Reset)
	}
	fmt.Println(">> Taille de l'inventaire :", p.invLimit, "<<")
	for key, val := range p.inventaire {
		fmt.Println(">>", key, ":", val)
	}
	if len(p.inventaire) == 0 {
		fmt.Println("Votre inventaire est vide")
	}
}

func (p *Personnage) useInventory(inFight bool) {
	//Affichage choix menu
	if p.inventaire[potionSoin] > 0 {
		fmt.Println("1 >> Utiliser potion de soin")
	}
	if p.inventaire["Potion de poison"] > 0 {
		fmt.Println("2 >> Utiliser potion de poison")
	}
	if p.inventaire["Livre de sort : boule de feu"] > 0 {
		fmt.Println("3 >> Utiliser livre de sort : boule de feu")
	}
	if p.inventaire["Casque en diamant"] > 0 {
		fmt.Println("4 >> Equiper le casque en diamant")
	}
	if p.inventaire["Jambières en diamant"] > 0 {
		fmt.Println("5 >> Equiper les jambières en diamant")
	}
	if p.inventaire["Bottes en diamant"] > 0 {
		fmt.Println("6 >> Equiper les bottes en diamant")
	}
	fmt.Println(BIWhite + "e >> Options" + Reset)
	if inFight == true {
		fmt.Println("r >> Retour")
	}
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
		p.EquipArmor("Casque en diamant", 10)
	case "5":
		p.EquipArmor("Jambières en diamant", 25)
	case "6":
		p.EquipArmor("Bottes en diamant", 15)
	case "e":
		Options()
	case "r":
		if inFight == true {
			if p.classe == "Humain" || p.classe == "Elfe" || p.classe == "Nain" {
				ClearLog()
				p.charTurn()
			}
			if p.classe == "Esprit de la forêt" {
				ClearLog()
				p.spiritTurn()
			}
			if p.classe == "Sirène" {
				ClearLog()
				p.mermaidTurn()
			}
		} else {
		}
	default:
		ClearLog()
		p.accesInventory()
		p.useInventory(inFight)
	}
	fmt.Println(BIWhite + ">> Appuyez sur entrée pour continuer... " + Reset)
	Input()
	ClearLog()
	if inFight == false {
		P1.accesInventory()
		P1.useInventory(false)
	}
}
