package main

import (
	"fmt"

)

func main()  {


	fmt.Println("Hello, what do you want the rat to say?")

	var input rune
	msg := ""
	n := 1

	for n > 0 {

		n, _ = fmt.Scanf("%c", &input)
		if n > 0 {
			if input == '\n' {
				n = 0
			} else {
				msg += string(input)
			}
		}

	}

	fmt.Println("The Rat said: ", msg)

}

