package entities

type Interaction struct {
	ItemName   string
	EntityName string
	Event      *Event
}

var ValidInteractions = []*Interaction{}

func ValidInteraction() {
	ValidInteractions = []*Interaction{
		{
			ItemName:   "tea",
			EntityName: "rosie",
			Event:      &Event{Description: "get-your-lanyard", Outcome: "Cheers! I needed that... by the way, where is your lanyard? I must have forgotten to give it to you.\nYou'll need that to move between rooms, here it is.\n\n(lanyard can now be found in the room).\n", Triggered: false},
		},
		{
			ItemName:   "first-plate",
			EntityName: "dishwasher",
			Event:      &Event{Description: "first-plate-loaded", Outcome: "You loaded the first plate into the dishwasher.", Triggered: false},
		},
		{
			ItemName:   "second-plate",
			EntityName: "dishwasher",
			Event:      &Event{Description: "second-plate-loaded", Outcome: "You loaded the second plate into the dishwasher.", Triggered: false},
		},
		{
			ItemName:   "third-plate",
			EntityName: "dishwasher",
			Event:      &Event{Description: "third-plate-loaded", Outcome: "You loaded the third plate into the dishwasher.", Triggered: false},
		},
		{
			ItemName:   "fourth-plate",
			EntityName: "dishwasher",
			Event:      &Event{Description: "fourth-plate-loaded", Outcome: "You loaded the fourth plate into the dishwasher.", Triggered: false},
		},
		{
			ItemName:   "fifth-plate",
			EntityName: "dishwasher",
			Event:      &Event{Description: "fifth-plate-loaded", Outcome: "You loaded the fifth plate into the dishwasher.", Triggered: false},
		},
		{
			ItemName:   "sixth-plate",
			EntityName: "dishwasher",
			Event:      &Event{Description: "sixth-plate-loaded", Outcome: "You loaded the sixth plate into the dishwasher.", Triggered: false},
		},
	}
}
