package author

import "fmt"

type Author struct {
	FirstName, LastName, Bio string
}

func (a *Author) FullName() string {
	return fmt.Sprintf("%s %s", a.FirstName, a.LastName)
}
