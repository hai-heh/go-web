//+build wireinject

package case2
import (
	"github.com/google/wire"
)


func InitMission(name string) Game {
	wire.Build(NewMonster, NewPlayer, NewGame)
	return Game{}
}




