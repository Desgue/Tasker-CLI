package types

import "fmt"

type Priority int

func (p Priority) String() string {
	return [3]string{"Low", "Medium", "High"}[p]
}
func StrToPriority(p string) (Priority, error) {
	switch p {
	case "Low":
		return 0, nil
	case "Medium":
		return 1, nil
	case "High":
		return 2, nil
	default:
		return 0, fmt.Errorf("invalid priority, priority must be Low, Medium or High")
	}
}

type Status int

func (s Status) String() string {
	return [3]string{"Pending", "In Progress", "Done"}[s]
}

func StrToStatus(s string) (Status, error) {
	switch s {
	case "Pending":
		return 0, nil
	case "In Progress":
		return 1, nil
	case "Done":
		return 2, nil
	default:
		return 0, fmt.Errorf("invalid status, status must be Pending, In Progress or Done")
	}
}
