package types

type Priority int

func (p Priority) String() string {
	return [3]string{"Low", "Medium", "High"}[p]
}
