package entities

type Entity struct {
	Name        string
	Description string
	Hidden      bool
}

func (e *Entity) SetDescription(description string) {
	e.Description = description
}

func (e *Entity) GetDescription() string {
	return e.Description
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
