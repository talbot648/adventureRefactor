package main

import (
	"fmt"
)

type Describable interface {
	SetDescription(description string)
	GetDescription() string
}

type Item struct {
	Name string
	Description string
	Weight int
}

func (i *Item) SetDescription(description string) {
    i.Description = description
}

func (i *Item) GetDescription() string {
    return i.Description
}

type Room struct {
	Name string
	Description string
	Exits map[string]*Room
	Items map[string]*Item
	Entities map[string]*Entity
}

func (r *Room) SetDescription(description string) {
    r.Description = description
}

func (r *Room) GetDescription() string {
    return r.Description
}

type Player struct {
	CurrentRoom *Room
	Inventory   map[string]*Item
	CurrentEntity *Entity
	CarriedWeight int
	AvailableWeight int
}

type Entity struct {
	Name string
	Description string
}

type Event struct {
	Description string
	Outcome string
	Triggered bool
}

type Interaction struct {
	ItemName string
	EntityName string
	Event *Event
}

var validInteractions = []*Interaction{
    {
        ItemName:   "key",
        EntityName: "door",
        Event:      &Event{Description: "unlock_door", Outcome: "The door unlocks with a loud click.\n", Triggered: false},
    },
    {
        ItemName:   "water",
        EntityName: "plant",
        Event:      &Event{Description: "water_plant", Outcome: "The plant looks healthier after being watered.\n", Triggered: false},
    },
}

func (e *Entity) SetDescription(description string) {
    e.Description = description
}

func (e *Entity) GetDescription() string {
    return e.Description
}

func (p *Player) Move(direction string) {
	if p.CurrentEntity != nil {
		fmt.Printf("You have to leave %s before moving!\n", p.CurrentEntity.Name)
		return
	}
	if newRoom, ok := p.CurrentRoom.Exits[direction]; ok {
		p.CurrentRoom = newRoom

		fmt.Printf("You are in %s\n", p.CurrentRoom.Name)
	} else {
		fmt.Println("You can't go that way!")
	}
}

func (p *Player) Take(itemName string) {
	item, ok := p.CurrentRoom.Items[itemName]
	switch {
	case !ok:
		fmt.Printf("%s not found in the room.\n", itemName)
		return
	case p.AvailableWeight < item.Weight:
		fmt.Println("Weight limit reached! Please drop an item before taking more.")
		return
	default:
		p.Inventory[item.Name] = item
		p.ChangeCarriedWeight(item, "increase")
		delete(p.CurrentRoom.Items, item.Name)

		fmt.Printf("%s has been added to your inventory.\n", item.Name)
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

		fmt.Printf("You dropped %s.\n", item.Name)
	} else {
		fmt.Printf("You don't have %s.\n", itemName)
	}
}

func (p *Player) ShowInventory() {
	if len(p.Inventory) == 0 {
		fmt.Println("Your inventory is empty.")
		return
	}
	fmt.Println("Your inventory contains:")
	for itemName, item := range p.Inventory {
		fmt.Printf("- %s: %s Weight: %d\n", itemName, item.Description, item.Weight)
	}
}

func (p *Player) ShowRoom() {
    fmt.Printf("You are in %s: %s\n", p.CurrentRoom.Name, p.CurrentRoom.Description)
    
	if len(p.CurrentRoom.Entities) != 0 {
		fmt.Println("You can approach:")
		for _, entity := range p.CurrentRoom.Entities {
			fmt.Printf("- %s\n", entity.Name)
		}
	}
    if len(p.CurrentRoom.Items) != 0 {
        fmt.Println("The room contains:")
        for itemName, item := range p.CurrentRoom.Items {
            fmt.Printf("- %s: %s Weight: %d\n", itemName, item.Description, item.Weight)
        }
    }
}


func (p *Player) Approach(entityName string) {
	if p.CurrentEntity != nil {
		fmt.Printf("You have to leave %s before you can do that!\n", p.CurrentEntity.Name)
		return
	}
	if entity, ok := p.CurrentRoom.Entities[entityName]; ok {

		p.CurrentEntity = entity
		fmt.Println(entity.Description)
	} else {
		fmt.Printf("%s not found in the room.\n", entityName)
	}
}

func (p *Player) Leave() {
	if p.CurrentEntity != nil {
		p.CurrentEntity = nil
		p.ShowRoom()
	} else {
		fmt.Printf("You can't leave.\n")
	}
}

func (p *Player) ShowMap() {
	for direction, exit := range p.CurrentRoom.Exits {
		fmt.Printf("%s: %s\n", direction, exit.Name)
	}
}

func (p *Player) Use(itemName string, target string) {
	for _, interaction := range validInteractions {
		if interaction.ItemName == itemName && interaction.EntityName == target {
			p.TriggerEvent(interaction.Event)
		}
	}
}

func (p *Player) TriggerEvent(event *Event) {
	fmt.Println(event.Outcome)
	event.Triggered = true
}

func updateDescription(d Describable, newDescription string) {
	d.SetDescription(newDescription)
}

func main() {
}
