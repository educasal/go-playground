package main

import "fmt"

func arrays() {
	dishes := [3]string{"ratatouille", "roulade", "coq au vin"}

	fmt.Println(dishes)

	fmt.Println(dishes[0])

	dishes2 := [...]string{"paella", "andouille"}

	for index, dish := range dishes2 {
		fmt.Println(index, dish)
	}
}

func slices() {
	dishes := []string{"ratatouille", "roulade", "coq au vin"}

	dishes2 := []string{"paella", "andouille"}

	dishes = append(dishes, dishes2...)

	for index, dish := range dishes {
		fmt.Println(index, dish)
	}

}
