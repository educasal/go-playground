package main

import (
	"fmt"
	"regexp"
)

func CountDeafRats(town string) int {
	piperFound := false
	count := 0

	pattern := regexp.MustCompile(`P|~O|O~|\s*`)
	split := pattern.FindAllString(town, -1)

	for idx, s := range split {
		fmt.Println("Substring:", s, idx)
		switch s {
		case "P":
			piperFound = true
		case "~O":
			if piperFound {
				count++
			}
		case "O~":
			if !piperFound {
				count++
			}
		}
	}
	return count
}

func rats() {

	fmt.Println("~O~O~O~O P", CountDeafRats("~O~O~O~O P"))
	fmt.Println("P O~ O~ ~O O~", CountDeafRats("P O~ O~ ~O O~"))
	fmt.Println("~O~O~O~OP~O~OO~", CountDeafRats("~O~O~O~OP~O~OO~"))
}

func rats2() {

	pattern := regexp.MustCompile(" |boat|one")
	description := "Kayak. A boat for one person."
	split := pattern.Split(description, -1)
	for _, s := range split {
		if s != "" {
			fmt.Println("Substring:", s)
		}
	}

}
