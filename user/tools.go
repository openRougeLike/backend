package user

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func (coords Coords) String() string {
	return fmt.Sprintf("%v,%v", coords[0], coords[1])
}

// Random integer between min and max, inclusive of both
func RandInt(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max + 1 - min) + min
}

func AbsInt(num int) int {
	return int(math.Abs(float64(num)))
}

func Chance(chance int) bool {
	return RandInt(0, 100) <= chance
}

// Note, you have to keep them going in ascending order!
type ChanceDistribution[T any] []ChanceDistributionV[T]

type ChanceDistributionV[T any] struct {
	Value T
	Max int
}

func (chance ChanceDistribution[T]) Fetch() T {
	prob := RandInt(0, chance[len(chance)-1].Max)

	for _, val := range chance {
		if prob <= val.Max {
			return val.Value
		}
	}
	// Should not reach here
	panic("OOPSY!!!!!!!!!!!!!!!!!!!! Prob: " + fmt.Sprint(prob))
}

func (c Coords) isEqual(coords Coords) bool {
	return c[0] == coords[0] && c[1] == coords[1]
}

func (tile TileType) IsBuildWalkable() bool {
	return tile == MapEmptySpace || tile == MapMonster || tile == MapArena
}

func (tile TileType) IsWalkWalkable() bool {
	return tile == MapEmptySpace || tile == MapArena
}
