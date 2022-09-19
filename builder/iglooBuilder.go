package main

type IglooBuilder struct {
	windowType string
	doorType   string
	floor      int
}

func getIglooBuilder() *IglooBuilder {
	return &IglooBuilder{}
}

func (b *IglooBuilder) setWindowType() {
	b.windowType = "igloo window with woods"
}
func (b *IglooBuilder) setDoorType() {
	b.doorType = "igloo door with woods"
}
func (b *IglooBuilder) setNumFloor() {
	b.floor = 21
}
func (b *IglooBuilder) getHouse() House {
	return House{
		windowType: b.windowType,
		doorType:   b.doorType,
		floor:      b.floor,
	}
}
