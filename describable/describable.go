package describable

type Describable interface {
	SetDescription(description string)
	GetDescription() string
}

func UpdateDescription(d Describable, newDescription string) {
	d.SetDescription(newDescription)
}
