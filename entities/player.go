package entities

import "fmt"

type Player struct {
	CurrentRoom     *Room
	Inventory       map[string]*Item
	CurrentEntity   *Entity
	CarriedWeight   int
	AvailableWeight int
}

func (p *Player) Move(direction string) {
	if p.CurrentEntity != nil {
		p.CurrentEntity = nil
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
	case !ok || item.Hidden:
		fmt.Printf("You can't take %s\n", itemName)
		return
	case p.AvailableWeight < item.Weight:
		fmt.Println("Weight limit reached! Please drop an item before taking more.")
		return
	case isPlate(itemName):
		if itemName == plateOrder[currentPlateIndex] {
			p.Inventory[item.Name] = item
			p.ChangeCarriedWeight(item, "increase")
			delete(p.CurrentRoom.Items, item.Name)
			currentPlateIndex++

			fmt.Printf("%s has been added to your inventory.\n", item.Name)
		} else {
			fmt.Println("As you attempt to grab the greasy plates without removing the ones stacked above them, they slip from your grasp and shatter, creating a chaotic mess.\n\nNow Rosie is very grumpy.")
			gameOver = true
		}

	default:
		p.Inventory[item.Name] = item
		p.ChangeCarriedWeight(item, "increase")
		delete(p.CurrentRoom.Items, item.Name)

		fmt.Printf("%s has been added to your inventory.\n", item.Name)
	}
}
