package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func clearScreen() {
    var cmd *exec.Cmd
    if runtime.GOOS == "windows" {
        cmd = exec.Command("cmd", "/c", "cls")
    } else {
        cmd = exec.Command("clear")
    }
    cmd.Stdout = os.Stdout
    cmd.Run()
}

type Describable interface {
	SetDescription(description string)
	GetDescription() string
}

type Item struct {
	Name string
	Description string
	Weight int
	Hidden bool
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
	Hidden bool
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

var validInteractions = []*Interaction{}

func (e *Entity) SetDescription(description string) {
    e.Description = description
}

func (e *Entity) GetDescription() string {
    return e.Description
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
		fmt.Printf("Your inventory is empty.\nAvailable space: %d\n", p.AvailableWeight)
		return
	}
	fmt.Printf("Available space: %d\nYour inventory contains:\n", p.AvailableWeight)
	for itemName, item := range p.Inventory {
		fmt.Printf("- %s: %s Weight: %d\n", itemName, item.Description, item.Weight)
	}
}

func (p *Player) ShowRoom() {
    fmt.Printf("You are in %s\n\n%s\n", p.CurrentRoom.Name, p.CurrentRoom.Description)

	if p.EntitiesArePresent() {
			fmt.Println("\nYou can approach:")
			for _, entity := range p.CurrentRoom.Entities {
				switch {
				case p.CurrentEntity != nil:
					if entity.Name == p.CurrentEntity.Name {
						fmt.Printf("- %s (currently approached)\n", entity.Name)
					} else if !entity.Hidden{
						fmt.Printf("- %s\n", entity.Name)
					}
				default:
					if !entity.Hidden{
						fmt.Printf("- %s\n", entity.Name)
					}
				}
			}
		}
		
	if p.ItemsArePresent() {
			fmt.Println("\nThe room contains:")
			for itemName, item := range p.CurrentRoom.Items {
				if !item.Hidden {
					fmt.Printf("- %s: %s Weight: %d\n", itemName, item.Description, item.Weight)
				}
			}
		}
	}

	

func (p *Player) ItemsArePresent() bool {
	if len(p.CurrentRoom.Items) != 0 {
	for _, item := range p.CurrentRoom.Items {
		if !item.Hidden {
			return true
		}
	}
}
	return false
}

func (p *Player) EntitiesArePresent() bool {
	if len(p.CurrentRoom.Entities) != 0 {
	for _, entity := range p.CurrentRoom.Entities {
		if !entity.Hidden {
			return true
		}
	}
}
	return false
}


func (p *Player) Approach(entityName string) {
	if p.CurrentEntity != nil {
		p.CurrentEntity = nil
	}
	if entity, ok := p.CurrentRoom.Entities[entityName]; ok && !entity.Hidden{

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
		fmt.Println("You have not approached anything. If you wish to leave the game, use the exit command.")
	}
}

func (p *Player) ShowMap() {
	for direction, exit := range p.CurrentRoom.Exits {
		fmt.Printf("%s: %s\n", direction, exit.Name)
	}
}

func (p *Player) Use(itemName string, target string) {
	if p.CurrentEntity == nil {
		fmt.Println("Approach to use an item.")
		return
	}	
	if p.CurrentEntity.Name == target {
		if _, ok := p.Inventory[itemName]; ok {
				for _, interaction := range validInteractions {
					if interaction.ItemName == itemName && interaction.EntityName == target {
						p.TriggerEvent(interaction.Event)
						p.ChangeCarriedWeight(p.Inventory[itemName], "decrease")
						delete(p.Inventory, itemName)
						return
					}
				}
		} else {
			fmt.Printf("You don't have %s.\n", itemName)
			return
		}
	} else {
		fmt.Printf("%s not found.\n", target)
		return
	}
	fmt.Printf("You can't use %s on %s.\n", itemName, target)
}

func (p *Player) TriggerEvent(event *Event) {
	fmt.Println(event.Outcome)
	event.Triggered = true
}

func updateDescription(d Describable, newDescription string) {
	d.SetDescription(newDescription)
}

func showCommands() {
	fmt.Println("-exit -> quits the game\n\n-commands -> shows the commands\n\n-look -> shows the content of the room.\n\n-approach <entity> -> to approach an entity\n\n-leave -> to leave an entity\n\n-inventory -> shows items in the inventory\n\n-take <item> -> to take an item\n\n-drop <item> -> tro drop an item\n\n-use <item> -> to use a certain item\n\n-move <direction> -> to move to a different room\n\n-map -> shows the directions you can take")
}

func main() {
	introduction := "It's the last day at the Academy, and you and your fellow graduates are ready to take on the final hack-day challenge.\nHowever, this time, it's different. Alan and Dan, your instructors, have prepared something more intense than ever before — a true test of your problem-solving and coding skills.\nThe doors to the academy are locked, the windows sealed. The only way out is to find and solve a series of riddles that lead to the terminal in a hidden room.\nThe challenge? Crack the code on the terminal to unlock the doors. But it's not that simple.\nYou'll need to gather items, approach Alan and Dan for cryptic tips, and outsmart the obstacles they've laid out for you.\nAs the tension rises, only your wits, teamwork, and knowledge can guide you to freedom.\nAre you ready to escape? The clock is ticking...\n\nif at any point you feel lost, type 'commands' to display the list of all commands."

	gameOver := false
	introductionShown:= false

	validInteractions = []*Interaction{
		{
			ItemName:   "tea",
			EntityName: "rosie",
			Event:      &Event{Description: "get-your-lanyard", Outcome: "Cheers! I needed that... by the way, where is your lanyard? You'll need that to move between rooms, here it is. (lanyard can now be found in the room).\n", Triggered: false},
		},
	}

	grumpyRosie := &Event{Description: "rosie-is-grumpy", Outcome: "Rosie caught you in the act of swiping a lanyard from a fellow student. You have made Rosie grumpy and you've lost the game.\n", Triggered: false}

	staffRoom := Room{
		Name:        "Break Room",
		Description: "A cozy lounge designed for both academy students and tutors, offering a welcoming space to unwind and socialise. Comfortable seating invites you to relax, while the warm ambiance encourages lively conversations and friendly exchanges.",
		Items:      make(map[string]*Item),
		Entities:   make(map[string]*Entity),
		Exits:      make(map[string]*Room),
	}

	terminalRoom := Room{
		Name:        "Server Room",
		Description: "A dark room filled with server racks and a single, locked terminal.",
		Items:      make(map[string]*Item),
		Entities:   make(map[string]*Entity),
		Exits:      make(map[string]*Room),
	}

	staffRoom.Exits["north"] = &terminalRoom
	terminalRoom.Exits["south"] = &staffRoom

	rosie := Entity{Name: "rosie", Description: "Ugh, what? Sorry, I can't think straight without a brew. Get me some tea, and then we'll talk...", Hidden: false}
	kettle := Entity{Name: "kettle", Description: "You set the kettle to boil, brewing the strongest cup of tea you've ever made. A comforting aroma fills the room as the tea is now ready. (tea can now be found in the room)", Hidden: false}
	sofa := Entity{Name: "sofa", Description: "You come across one of your fellow academy students fast asleep on the sofa. Next to them, their lanyard lies carelessly within reach. You know you shouldn't take it, but the temptation lingers... (other-lanyard can now be found in the room)", Hidden: false}
	terminal := Entity{Name: "terminal", Description: "A locked terminal. It won't open without a key.", Hidden: false}
	tea := Item{Name: "tea", Description: "A steaming cup of Yorkshire tea, rich and comforting.", Weight: 2, Hidden: true}
	lanyard := Item{Name: "lanyard", Description: "Your lanyard, a key to unlocking any door within the building.", Hidden: true}
	otherLanyard := Item{Name: "other-lanyard", Description: "A lanyard, a key to unlocking any door within the building.", Hidden: true}

	staffRoom.Items[tea.Name] = &tea
	staffRoom.Items[lanyard.Name] = &lanyard
	staffRoom.Items[otherLanyard.Name] = &otherLanyard
	staffRoom.Entities[rosie.Name] = &rosie
	staffRoom.Entities[kettle.Name] = &kettle
	staffRoom.Entities[sofa.Name] = &sofa
	terminalRoom.Entities[terminal.Name] = &terminal

	player := Player{
		CurrentRoom:     &staffRoom,
		Inventory:       make(map[string]*Item),
		AvailableWeight: 30,
		CurrentEntity:   nil,
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {

		if player.CurrentEntity != nil && player.CurrentEntity.Name == "sofa" {
			otherLanyard.Hidden = false
			sofa.SetDescription("One of your fellow academy students. Still asleep on the sofa.")
		}

		if player.CurrentEntity != nil && player.CurrentEntity.Name == "kettle" {
			tea.Hidden = false
			kettle.SetDescription("A kettle — essential for survival, impossible to function without one nearby.");
		}

		for _, validInteraction := range validInteractions {
			if validInteraction.Event.Description == "get-your-lanyard" && validInteraction.Event.Triggered {
				lanyard.Hidden = false
				rosie.SetDescription("Can I help with anything else?")
			}
		}

		if _, ok := player.Inventory["other-lanyard"]; ok {
			player.TriggerEvent(grumpyRosie)
			gameOver = true
		}

		if gameOver {
			fmt.Println("Thank you for playing!")
			break
		}

		if !introductionShown {
			fmt.Println(introduction)
			introductionShown = true
		}

		fmt.Print("Enter command: ")


		if scanner.Scan() {
			input := scanner.Text()
			input = strings.TrimSpace(input)
			input = strings.ToLower(input)

			if input == "exit" {
				clearScreen()
				fmt.Println("Thank you for playing!")
				break
			}

			parts := strings.Fields(input)
			if len(parts) == 0 {
				continue
			}

			command := (parts[0])
			args := parts[1:]

			switch command {
			case "commands":
				clearScreen()
				showCommands()
			case "look":
				clearScreen()
				player.ShowRoom()
			case "take":
				clearScreen()
				if len(args) > 0 {
					player.Take(args[0])
				} else {
					fmt.Println("Specify an item to take.")
				}
			case "drop":
				clearScreen()
				if len(args) > 0 {
					player.Drop(args[0])
				} else {
					fmt.Println("Specify an item to drop.")
				}
			case "inventory":
				clearScreen()
				player.ShowInventory()
			case "approach":
				clearScreen()
				if len(args) > 0 {
					player.Approach(args[0])
				} else {
					fmt.Println("Specify an entity to approach.")
				}
			case "use":
				clearScreen()
				if len(args) > 0 {
					if player.CurrentEntity == nil {
						player.Use(args[0], "unspecified_entity")
					} else {
						player.Use(args[0], player.CurrentEntity.Name)
					}
				} else {
					fmt.Println("Specify an item to use.")
				}
			case "leave":
				clearScreen()
				player.Leave()
			case "move":
				clearScreen()
				if _, ok := player.Inventory["lanyard"]; ok {
					if len(args) > 0 {
						player.Move(args[0])
					} else {
						fmt.Println("Specify a direction to move (e.g., north).")
					}
				} else {
					fmt.Println("Doors are shut for you if you don't have a lanyard.")
				}
			case "map":
				clearScreen()
				player.ShowMap()
			default:
				clearScreen()
				fmt.Println("Unknown command:", command)
			}
		}
	}
}

