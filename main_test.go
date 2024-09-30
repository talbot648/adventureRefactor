package main

import (
	"testing"
)

func TestPlayerMovement(t *testing.T) {
	//Arrange
	room1 := Room{Name: "Room 1", Description: "This is room 1.", Exits: make(map[string]*Room)}
    room2 := Room{Name: "Room 2", Description: "This is room 2.", Exits: make(map[string]*Room)}
    room1.Exits["north"] = &room2
    room2.Exits["south"] = &room1

    player := Player{CurrentRoom: &room1}

    // Act
    player.Move("north")

    // Assert
    if player.CurrentRoom.Name != "Room 2" {
        t.Errorf("Expected Room 2, got %s", player.CurrentRoom.Name)
    }
}

func TestPlayerMovementInvalidDirection(t *testing.T) {
	//Arrange
	room1 := Room{Name: "Room 1", Description: "This is room 1.", Exits: make(map[string]*Room)}
	room2 := Room{Name: "Room 2", Description: "This is room 2.", Exits: make(map[string]*Room)}
	room1.Exits["north"] = &room2
	room2.Exits["south"] = &room1

	player := Player{CurrentRoom: &room1}

	//Act
	player.Move("east")

	//Assert
	if player.CurrentRoom.Name != "Room 1" {
		t.Errorf("Expected Room1, got %s", player.CurrentRoom.Name)
	}
}

func TestPickUpItem(t *testing.T) {
	//Arrange
	room := Room{Items: make(map[string]*Item)}
	
	item := Item{Name: "Item", Description: "This is an item."}

	room.Items[item.Name] = &item
	
	player := Player{CurrentRoom: &room, Inventory: make(map[string]*Item)}
	
	//Act
	player.Take(item.Name)


	//Assert
	if _, ok := player.Inventory[item.Name]; !ok {
		t.Errorf("Expected true for item present in the inventory, got false")
	}
	
	if _, ok := room.Items[item.Name]; ok {
		t.Errorf("Expected false for item missing from the room, got true")
	}
}

func TestPickUpAbsentItem(t *testing.T) {
	//Arrange
	room1 := Room{Items: make(map[string]*Item)}
	room2 := Room{Items: make(map[string]*Item)}
	
	item := Item{Name: "Item", Description: "This is an item."}

	room1.Items[item.Name] = &item
	
	player := Player{CurrentRoom: &room2, Inventory: make(map[string]*Item)}
	
	//Act
	player.Take(item.Name)


	//Assert
	if _, ok := player.Inventory[item.Name]; ok {
		t.Errorf("Expected false for picking up absent item, got true")
	}
}

func TestPickUpNonexistentItem(t *testing.T) {
	//Arrange
	room2 := Room{Items: make(map[string]*Item)}

	player := Player{CurrentRoom: &room2, Inventory: make(map[string]*Item)}
	
	//Act
	player.Take("Item")


	//Assert
	if _, ok := player.Inventory["Item"]; ok {
		t.Errorf("Expected false for picking up nonexistent item, got true")
	}
}

func TestDropItem(t *testing.T) {
	//Arrange
	room := Room{Items: make(map[string]*Item)}
	
	item := Item{Name: "Item", Description: "This is an item."}

	room.Items[item.Name] = &item
	
	player := Player{CurrentRoom: &room, Inventory: make(map[string]*Item)}
	
	//Act
	player.Take(item.Name)

	player.Drop(item.Name)

	//Assert
	if _, ok := player.Inventory[item.Name]; ok {
		t.Errorf("Expected false for item absent from the inventory, got true")
	}
	if _, ok := room.Items[item.Name]; !ok {
		t.Errorf("Expected true for item present in the room, got false")
	}
}

func TestDropAbsentItem(t *testing.T) {
	//Arrange
	room1 := Room{Name: "Room 1", Description: "This is room 1.", Exits: make(map[string]*Room), Items: make(map[string]*Item)}
	room2 := Room{Name: "Room 2", Description: "This is room 2.", Exits: make(map[string]*Room), Items: make(map[string]*Item)}
	room1.Exits["north"] = &room2
	room2.Exits["south"] = &room1

	item := Item{Name: "Item", Description: "This is an item."}

	room1.Items[item.Name] = &item
	
	player := Player{CurrentRoom: &room1, Inventory: make(map[string]*Item)}
	
	//Act
	player.Move("north")
	player.Drop(item.Name)

	//Assert

	if _, ok := room2.Items[item.Name]; ok {
		t.Errorf("Expected false for item absent from the room, got true")
	}
}

func TestDropNonexistentItem(t *testing.T) {
	//Arrange
	room := Room{Items: make(map[string]*Item)}
	
	player := Player{CurrentRoom: &room, Inventory: make(map[string]*Item)}
	
	//Act
	player.Drop("Item")

	//Assert
	if _, ok := player.Inventory["Item"]; ok {
		t.Errorf("Expected false for item absent from the inventory, got true")
	}
	if _, ok := room.Items["Item"]; ok {
		t.Errorf("Expected false for item absent from the room, got true")
	}
}