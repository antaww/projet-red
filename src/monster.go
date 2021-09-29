package main

type Monster struct {
	name   string
	maxHP  int
	HP     int
	attack int
}

func (m *Monster) Init(name string, maxHP int, HP int, attack int) {
	m.name = name
	m.maxHP = maxHP
	m.HP = HP
	m.attack = attack
}

func InitGoblin() {
	var HP, maxHP int
	maxHP = 40
	HP = maxHP
	M1.Init(Red+"Gobelin d'entrainement"+Reset, maxHP, HP, 5)
}
