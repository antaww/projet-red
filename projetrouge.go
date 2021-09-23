package main

import "fmt"

type personnage struct {
	nom        string
	classe     string
	niveau     int
	maxHP      int
	HP         int
	inventaire []string
}

func (p *personnage) Init(nom string, classe string, niveau int, maxHP int, HP int, inventaire []string) {
	p.nom = nom
	p.classe = classe
	p.niveau = niveau
	p.maxHP = maxHP
	p.HP = HP
	p.inventaire = inventaire
}

func main() {
	var p1 personnage
	p1.Init("Antoine", "Elfe", 1, 100, 40, []string{"potion", "potion", "potion"})
	p1.DisplayInfo()
	p1.accesInventory()
	p1.takePot()
	p1.DisplayInfo()
	p1.accesInventory()
}

func (p *personnage) takePot() {
	for i := range p.inventaire {
		if p.inventaire[i] == "potion" {
			if p.HP == p.maxHP {
				fmt.Println("Vous êtes full, ne pas utiliser la potion")
			} else if p.HP+50 > p.maxHP {
				p.HP += +(p.maxHP - p.HP)
				p.inventaire[i] = ""
				fmt.Println(p.nom, "a utilisé une potion")
			} else {
				p.HP += +50
				p.inventaire[i] = ""
				fmt.Println(p.nom, "a utilisé une potion")
			}
			break
		}
	}
}

/// Tâche 5

func (p *personnage) accesInventory() {
	for i := range p.inventaire {
		if i >= 0 {
			fmt.Println("Item", i+1, ":", p.inventaire[i])
		} else {
			fmt.Println("Votre inventaire est vide")
		}
	}
}

func (p *personnage) removeInventory(item string) {
	inv := &p.inventaire
	for i := range p.inventaire {
		if p.inventaire[i] == item {
			*inv = append(p.inventaire[:i], p.inventaire[i+1:]...)
		}
	}
}

func (p personnage) DisplayInfo() {
	fmt.Println("Nom :", p.nom)
	fmt.Println("Classe :", p.classe)
	fmt.Println("Niveau :", p.niveau)
	fmt.Println("HP max :", p.maxHP)
	fmt.Println("HP actuels :", p.HP)
	fmt.Println("Inventaire :", p.inventaire)
}
