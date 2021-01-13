package game

import "math"

type Birdy struct {
	PosY int
}

func (b *Birdy) Jump() {
	b.PosY = int(math.Max(float64(b.PosY-3), 1))
}
