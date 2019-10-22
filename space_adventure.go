package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	println("Welcome to The Solar System! ")
	println("There are 9 planets to explore.")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("What is your Name?: ")
	text, _ := reader.ReadString('\n')
	fmt.Println(text)

	fmt.Println(text, "My name is Eliza, I'm an old friend of Alex. Let's go on an adventure!")
	text2 := ""
	fmt.Scanln(text2)
	fmt.Println(text2)

	fmt.Println("Shall I randomly choose a planet for you? Y or N ")
	text3, _ := reader.ReadString('\n')
	fmt.Scanln(text3)
	fmt.Println("name: Venus",
		"description: It's very cloudy here!")

	ln := ""
	fmt.Sscanln("%v", ln)
	fmt.Println(ln)
}
