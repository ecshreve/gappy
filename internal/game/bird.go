package game

type Birdy struct {
	PosY int
}

func (b *Birdy) Jump() {
	b.PosY -= 3
}
