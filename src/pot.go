package main

import (
	"fmt"
	"time"
)

func (p *Personnage) takePot() {
	for range p.inventaire {
		if p.inventaire["Potion"] == 0 {
			fmt.Println(">> Vous n'avez pas de potion de soin... <<")
		}
		if p.inventaire["Potion"] >= 1 {
			if p.HP == p.maxHP {
				fmt.Println(">> Vos HP sont full, ne pas utiliser la potion <<")
			}
			if p.HP != p.maxHP {
				if p.HP+50 > p.maxHP {
					p.HP += (p.maxHP - p.HP)
					p.inventaire["Potion"] -= 1
					fmt.Println(">> Une potion de soin a été utilisée par", p.nom, "<<")
					fmt.Println("HP :", p.HP, "/", p.maxHP)
				}
			}
			if p.HP+50 < p.maxHP {
				p.HP += +50
				p.inventaire["Potion"] -= 1
				fmt.Println(">> Une potion de soin a été utilisée par", p.nom, "<<")
				fmt.Println("HP :", p.HP, "/", p.maxHP)
			}
			break
		}
	}
}

func (p *Personnage) poisonPot() {
	if p.inventaire["Potion de poison"] >= 1 {
		p.inventaire["Potion de poison"] -= 1
		for i := 1; i <= 3; i++ {
			p.HP = p.HP - 10
			fmt.Println("HP :", p.HP, "/", p.maxHP)
			time.Sleep(1 * time.Second)
			if p.HP == 0 {
				p.dead()
			}
		}
	}
}
