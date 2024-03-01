package types

type Priority int

func (p Priority) String() string {
	return [3]string{"Low", "Medium", "High"}[p]
}

type Status int

func (s Status) String() string {
	return [3]string{"To Do", "In Progress", "Done"}[s]
}
