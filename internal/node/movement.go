package node

import (
	"lab1/pkg/utils"
	"math"
)

// Энергия убывает на 1% каждые N передач
// Передвижение в пределах области
// Шлюзы стоят на месте
// Радиус уменьшается на 1%

func (s *Node) RandomMove(AreaX, AreaY float64) {

	//fmt.Printf("Node %s (%v;%v) | Goto: x: %v, y: %v\n", s.Name, s.x, s.y, x, y)

	leftBorderX := math.Max(s.x-s.MovingSpeed, 0)
	rightBorderX := math.Min(s.x+s.MovingSpeed, AreaX)

	leftBorderY := math.Max(s.y-s.MovingSpeed, 0)
	rightBorderY := math.Min(s.y+s.MovingSpeed, AreaY)

	newX := utils.GenerateRandom(leftBorderX, rightBorderX)
	newY := utils.GenerateRandom(leftBorderY, rightBorderY)

	s.x = newX
	s.y = newY
}
