package main

import (
	"fmt"
	"time"
)

type Personnage struct { //structure personnage
	nom        string
	classe     string
	niveau     int
	maxHP      int
	HP         int
	attack     int
	invLimit   int
	inventaire map[string]int
	alive      bool
	skill      map[string]int
	money      int
	equipment  Equipment
	experience int
}

//fonction d'initialisation des caractéristiques du personnage
func (p *Personnage) Init(nom string, classe string, niveau int, maxHP int, HP int, attack int, invLimit int, inventaire map[string]int, alive bool,
	skill map[string]int, money int, equiptete string, equiptorse string, equippieds string, experience int) {
	p.nom = nom
	p.classe = classe
	p.niveau = niveau
	p.maxHP = maxHP
	p.HP = HP
	p.attack = attack
	p.invLimit = invLimit
	p.inventaire = inventaire
	p.alive = alive
	p.skill = skill
	p.money = money
	p.equipment.equiptete = equiptete
	p.equipment.equiptorse = equiptorse
	p.equipment.equippieds = equippieds
	p.experience = experience
}

//fonction d'initialisation d'un personnage pré-créé pour effectuer des tests
func InitCharacter() {
	P1.Init(BIPurple+"charTester"+Reset, "Humain", 1, 100, 40, 5, 10, map[string]int{potionSoin: 3}, true, map[string]int{"Coup de poing": 1, "Boule de feu": 5}, 100,
		"Casque", "Plastron", "Bottes", 0)
}

//fonction mort
func (p *Personnage) dead() {
	if p.HP == 0 { //verification des HP du personnage
		p.alive = false
		fmt.Println(p.nom, "est mort...")
		fmt.Println("")
		fmt.Println(">> Résurrection dans 2 secondes...")
		time.Sleep(1 * time.Second)
		fmt.Println(">> Résurrection dans 1 seconde...")
		time.Sleep(1 * time.Second)
		p.alive = true
		p.HP = p.maxHP / 2 //remise des HP du personnage à 50% de sa vie maximale une fois en vie
		fmt.Println(p.nom, "est de retour parmis nous, il n'a malheureusement plus que la moitié de sa vie.")
	}
}

//fonction d'affichage des stats du personnage
func (p Personnage) DisplayInfo() {
	runeName := []rune(p.nom)                                                                                       //création rune pour tester le premier caractère du pseudo
	if runeName[0] == 'A' || runeName[0] == 'E' || runeName[0] == 'I' || runeName[0] == 'O' || runeName[0] == 'U' { //esthétique selon le pseudo
		fmt.Print(BIWhite+">> Stats d'"+Reset, P1.nom, BIWhite+" <<\n"+Reset)
	} else {
		fmt.Print(BIWhite+">> Stats de "+Reset, P1.nom, BIWhite+" <<\n"+Reset)
	}
	fmt.Println("Nom :", p.nom)
	fmt.Println("Classe :", p.classe)
	fmt.Println("Niveau :", p.niveau)
	fmt.Println("Expérience :", p.experience)
	fmt.Println("HP max :", p.maxHP)
	fmt.Println("HP actuels :", p.HP)
	fmt.Println("Argent :", p.money)
	fmt.Println("Compétences :")
	for key := range p.skill { //affichage des compétences dans la map skill
		fmt.Println("\t", key)
	}
}
