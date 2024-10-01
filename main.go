package main

import (
	"fmt"
)

type Item struct {
	Name string
	Description string
	Weight int
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
	CarriedWeight int
	AvailableWeight int
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
	item, ok := p.CurrentRoom.Items[itemName]
	switch {
	case !ok:
		fmt.Printf("%s not found in the room.", itemName)
		return
	case p.AvailableWeight < item.Weight:
		fmt.Println("Weight limit reached! Please drop an item before taking more.")
		return
	default:
		p.Inventory[item.Name] = item
		p.ChangeCarriedWeight(item, "increase")
		delete(p.CurrentRoom.Items, item.Name)

		fmt.Printf("%s has been added to your inventory.", item.Name)
	}
}

func (p *Player) ChangeCarriedWeight(item *Item, operation string) {
	switch {
	case operation == "increase":
		p.CarriedWeight += item.Weight
		p.AvailableWeight -= item.Weight
		return
	case operation == "decrease":
		p.CarriedWeight -= item.Weight
		p.AvailableWeight += item.Weight
		return
	}
}

func (p *Player) Drop(itemName string) {
	if item, ok := p.Inventory[itemName]; ok {

		delete(p.Inventory, item.Name)
		p.ChangeCarriedWeight(item, "decrease")
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