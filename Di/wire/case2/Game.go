package case2

type Game struct {
	play Player
	moster Monster
}

func NewGame(play Player, moster Monster) Game {
	return Game{play: play, moster: moster}
}
