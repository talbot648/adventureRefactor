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

		fmt.Printf("You are in %s", p.CurrentRoom.Name)
	} else {
		fmt.Println("You can't go that way!")
	}
}

func (p *Player) Take(itemName string) {
	if item, ok := p.CurrentRoom.Items[itemName]; ok{
		p.Inventory[item.Name] = item

		delete(p.CurrentRoom.Items, item.Name)

		fmt.Printf("%s has been added to your inventory.", item.Name)
	} else {
		fmt.Println("Item not found in the room.")
	}
}

func (p *Player) Drop(itemName string) {
	if item, ok := p.Inventory[itemName]; ok {

		delete(p.Inventory, item.Name)

		p.CurrentRoom.Items[item.Name] = item

		fmt.Printf("You dropped %s.", item.Name)
	} else {
		fmt.Printf("You don't have %s.", itemName)
	}
}

func (p *Player) ShowInventory() {
	if len(p.Inventory) == 0 {
		fmt.Println("Your inventory is empty.")
		return
	}
	fmt.Println("Your inventory contains:")
	for itemName, item := range p.Inventory {
		fmt.Printf("- %s: %s\n", itemName, item.Description)
	}
}

func (p *Player) ShowRoom() {
	fmt.Printf("You are in %s: %s", p.CurrentRoom.Name, p.CurrentRoom.Description)
}

func main() {
}