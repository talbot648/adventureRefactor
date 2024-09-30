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
	room1 := Room{Name: "Room 1", Description: "This is room 1", Exits: make(map[string]*Room)}
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
