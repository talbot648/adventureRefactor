package main

import (
	"fmt"
)

type Item struct {
	Name string
	Description string
}

type Room struct {
	Name        string
	Description string
	Exits       map[string]*Room
	Items       map[string]*Item
}

type Player struct {
	CurrentRoom *Room
	Inventory   map[string]*Item
}

func (p *Player) Move(direction string) {
	if newRoom, ok := p.CurrentRoom.Exits[direction]; ok {
		p.CurrentRoom = newRoom
	} else {
		fmt.Println("You can't go that way!")
	}
}

func (p *Player) Take(itemName string) {
	if item, ok := p.CurrentRoom.Items[itemName]; ok{
		p.Inventory[item.Name] = item

		delete(p.CurrentRoom.Items, itemName)
	}
}

func main() {
}