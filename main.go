package main

import (

)

type Room struct {
	Name        string
	Description string
	Exits       map[string]*Room
	Items       []string
}

type Player struct {
	CurrentRoom *Room
	Inventory   []string
}

func (p *Player) Move(direction string) {
	p.CurrentRoom = p.CurrentRoom.Exits[direction]
}

func main() {
}