package main

import (
	"fmt"
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
	} else {
		fmt.Println("You can't go that way!")
	}
}

func main() {
}