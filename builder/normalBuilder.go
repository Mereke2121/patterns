package main

type NormalBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func getNormalBuilder() *NormalBuilder {
	return &NormalBuilder{}
}

func (b *NormalBuilder) setWindowType() {
	b.windowType = "window with woods"
}
func (b *NormalBuilder) setDoorType() {
	b.doorType = "door with woods"
}
func (b *NormalBuilder) setNumFloor() {
	b.floor = 12
}
func (b *NormalBuilder) getHouse() House {
	return House{
		windowType: b.windowType,
		doorType:   b.doorType,
		floor:      b.floor,
	}
}
