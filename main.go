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
	if newRoom, ok := p.CurrentRoom.Exits[direction]; ok {
		p.CurrentRoom = newRoom
	}
}

func main() {
}