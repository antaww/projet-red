package main

import (
	"fmt"
	"strings"
	"time"
)

type Personnage struct {
	nom        string
	classe     string
	niveau     int
	maxHP      int
	HP         int
	inventaire map[string]int
	alive      bool
	money      int
}

func (p *Personnage) Init(nom string, classe string, niveau int, maxHP int, HP int, inventaire map[string]int, alive bool, money int) {
	p.nom = nom
	p.classe = classe
	p.niveau = niveau
	p.maxHP = maxHP
	p.HP = HP
	p.inventaire = inventaire
	p.alive = alive
	p.money = money
}

func InitCharacter() {
	P1.Init("Antoine", "Elfe", 1, 100, 40, map[string]int{"Potion": 3}, true, 100)
}

func (p *Personnage) takePot() {
	for range p.inventaire {
		if p.inventaire["Potion"] >= 1 {
			if p.HP == p.maxHP {
				fmt.Println("Vous êtes full, ne pas utiliser la potion")
			} else if p.HP+50 > p.maxHP {
				p.HP += +(p.maxHP - p.HP)
				p.inventaire["Potion"] -= 1
				fmt.Println(p.nom, "a utilisé une potion")
			} else {
				p.HP += +50
				p.inventaire["Potion"] -= 1
				fmt.Println(p.nom, "a utilisé une potion")
			}
			break
		}
	}
}

func (p *Personnage) dead() {
	if p.HP == 0 {
		p.alive = false
		fmt.Println(p.nom, "est mort...")
		fmt.Println("Résurrection dans 2 secondes...")
		time.Sleep(1 * time.Second)
		fmt.Println("Résurrection dans 1 seconde...")
		time.Sleep(1 * time.Second)
		p.alive = true
		p.HP = p.maxHP / 2
		fmt.Println(p.nom, "est de retour parmis nous, il n'a malheureusement plus que la moitié de sa vie.")
	}
}

func (p *Personnage) poisonPot() {
	for i := 1; i <= 3; i++ {
		p.HP = p.HP - 10
		fmt.Println(p.HP, "/", p.maxHP)
		time.Sleep(1 * time.Second)
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
	fmt.Println("Argent :", p.money)
}
