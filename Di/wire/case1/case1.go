package main

type Monster struct {
	Name string
}

func NewMonster() Monster {
	return Monster{Name: "Monster"}
}

type Player struct {
	Name string
}

func NewPlayer(name string) Player {
	return Player{Name: name}
}

type Game struct {
	player Player
	monster Monster
}

type Mission struct {
	Player  Player
	Monster Monster
}

func NewMission(p Player, m Monster) Mission {

	return Mission{p, m}
}

func NewGame(player Player, monster Monster) Game {
	return Game{player: player, monster: monster}
}









