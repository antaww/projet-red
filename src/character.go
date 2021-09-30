package main

import (
	"fmt"
	"time"
)

type Personnage struct {
	nom        string
	classe     string
	niveau     int
	maxHP      int
	HP         int
	attack     int
	inventaire map[string]int
	alive      bool
	skill      map[string]int
	money      int
	equipment  Equipment
}

func (p *Personnage) Init(nom string, classe string, niveau int, maxHP int, HP int, attack int, inventaire map[string]int, alive bool,
	skill map[string]int, money int, equiptete string, equiptorse string, equippieds string) {
	p.nom = nom
	p.classe = classe
	p.niveau = niveau
	p.maxHP = maxHP
	p.HP = HP
	p.attack = attack
	p.inventaire = inventaire
	p.alive = alive
	p.skill = skill
	p.money = money
	p.equipment.equiptete = equiptete
	p.equipment.equiptorse = equiptorse
	p.equipment.equippieds = equippieds
}

func InitCharacter() {
	P1.Init(BIPurple+"charTester"+Reset, "Elfe", 1, 100, 40, 5, map[string]int{potionSoin: 3}, true, map[string]int{"Coup de poing": 1, "Boule de feu": 5}, 100,
		"Casque", "Plastron", "Bottes")
}

func (p *Personnage) dead() {
	if p.HP == 0 {
		p.alive = false
		fmt.Println(p.nom, "est mort...")
		fmt.Println("")
		fmt.Println(">> Résurrection dans 2 secondes...")
		time.Sleep(1 * time.Second)
		fmt.Println(">> Résurrection dans 1 seconde...")
		time.Sleep(1 * time.Second)
		p.alive = true
		p.HP = p.maxHP / 2
		fmt.Println(p.nom, "est de retour parmis nous, il n'a malheureusement plus que la moitié de sa vie.")
	}
}

func (p Personnage) DisplayInfo() {
	runeName := []rune(p.nom)
	if runeName[0] == 'A' || runeName[0] == 'E' || runeName[0] == 'I' || runeName[0] == 'O' || runeName[0] == 'U' {
		fmt.Print(">> Stats d'", P1.nom, " <<\n")
	} else {
		fmt.Print(">> Stats de ", P1.nom, " <<\n")
	}
	fmt.Println("Nom :", p.nom)
	fmt.Println("Classe :", p.classe)
	fmt.Println("Niveau :", p.niveau)
	fmt.Println("HP max :", p.maxHP)
	fmt.Println("HP actuels :", p.HP)
	fmt.Println("Argent :", p.money)
	fmt.Println("Compétences :")
	for key := range p.skill {
		fmt.Println("\t", key)
	}
}
