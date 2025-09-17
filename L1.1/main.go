package main

import (
	"fmt"
	"main/actionStruct"
	"main/humanStruct"
)

func main() {

	act := actionStruct.Action{
		Sleep: "always",
		Human: humanStruct.Human{
			Id:     2,
			Name:   "test",
			Family: "testovich",
			Age:    666,
		},
	}

	fmt.Printf("%s", act.Human.ToString())

}
