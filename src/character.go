package main

import (
	"fmt"
	"strings"
)

type Personnage struct {
	nom        string
	classe     string
	niveau     int
	maxHP      int
	HP         int
	inventaire map[string]int
}

func (p *Personnage) Init(nom string, classe string, niveau int, maxHP int, HP int, inventaire map[string]int) {
	p.nom = nom
	p.classe = classe
	p.niveau = niveau
	p.maxHP = maxHP
	p.HP = HP
	p.inventaire = inventaire
}

func InitCharacter() {
	P1.Init("Antoine", "Elfe", 1, 100, 40, map[string]int{"Potion": 3})
}

func (p *Personnage) takePot() {
	for range p.inventaire {
		if p.inventaire["potion"] >= 1 {
			if p.HP == p.maxHP {
				fmt.Println("Vous êtes full, ne pas utiliser la potion")
			} else if p.HP+50 > p.maxHP {
				p.HP += +(p.maxHP - p.HP)
				p.inventaire["potion"] -= 1
				fmt.Println(p.nom, "a utilisé une potion")
			} else {
				p.HP += +50
				p.inventaire["potion"] -= 1
				fmt.Println(p.nom, "a utilisé une potion")
			}
			break
		}
	}
}

func (p *Personnage) FreeHealPotion() {
	p.inventaire["Potion"] += 1
}

func (p *Personnage) accesInventory() {
	for key, val := range p.inventaire {
		fmt.Println(key, ":", val)
	}
	if len(p.inventaire) == 0 {
		fmt.Println("Votre inventaire est vide")
	}
}

func (p *Personnage) Marchand() {
	//Affichage choix menu
	fmt.Println("1 >> Acheter potion de soin (gratuite)")
	fmt.Println("2 >> Retour au menu principal")
	o, _ := BR.ReadString('\n') //lire input joueur quand "entrée"
	o = strings.Replace(o, "\r\n", "", -1)
	switch o {
	case "1":
		fmt.Println(">> Une potion de soin a été achetée par", P1.nom, "<<")
		P1.FreeHealPotion()
	}
}

func (p Personnage) DisplayInfo() {
	fmt.Println("Nom :", p.nom)
	fmt.Println("Classe :", p.classe)
	fmt.Println("Niveau :", p.niveau)
	fmt.Println("HP max :", p.maxHP)
	fmt.Println("HP actuels :", p.HP)
}
