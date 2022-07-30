package entity

type ID string

func ParseID(val string) ID {
	return ID(val)
}

func (id ID) IsEmpty() bool {
	return len(id) == 0
}

func (id ID) String() string {
	return string(id)
}
