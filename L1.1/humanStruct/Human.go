package humanStruct

import (
	"fmt"
)

type Human struct {
	Id     int
	Name   string
	Family string
	Age    int
}

func (h *Human) ToString() string {
	return fmt.Sprintf("id:\t%d\nName:\t%s\nFamily:\t%s\nAge:\t%d\n", h.Id, h.Name, h.Family, h.Age)
}
