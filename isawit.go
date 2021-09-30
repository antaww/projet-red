package main

type Equipment struct { //met le au début après la structure perso
	Equiptete  string
	Equiptorse string
	Equippieds string
}

type Personnage struct {
	nom        string
	classe     string
	niveau     int
	maxHP      int
	HP         int
	inventaire map[string]int
	alive      bool
	money      int
	equipment  Equipment
}

func (p *Personnage) Init(Equiptete string, Equiptorse string, Equippieds string) { // fonction
	p.equipment.Equiptete = Equiptete
	p.equipment.Equiptorse = Equiptorse
	p.equipment.Equippieds = Equippieds
}
