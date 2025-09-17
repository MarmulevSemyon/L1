package actionStruct

import (
	"main/humanStruct"
)

type Action struct {
	Sleep string
	humanStruct.Human
}

// func (a *Action) ToString() string {
// 	return fmt.Sprintf("popa: %d\t%s", a.Popa, a.Human.ToString())
// }
