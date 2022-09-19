package main

type IDirector interface {
	setBuilder(b IBuilder)
	buildHouse() House
}
