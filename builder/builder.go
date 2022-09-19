package main

type IBuilder interface {
	setWindowType()
	setDoorType()
	setNumFloor()
	getHouse() House
}

func getBuilder(builderType string) IBuilder {
	if builderType == "normal" {
		return getNormalBuilder()
	} else if builderType == "igloo" {
		return getIglooBuilder()
	}
	return nil
}
