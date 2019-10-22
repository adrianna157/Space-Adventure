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

	ln := ""
	fmt.Sscanln("%v", ln)
	fmt.Println(ln)
}

// package main

// import (
// 	"fmt"
// 	"os"
// )

// func main() {

// 	reader := bufio.newReader(os.Stdin)
// 	var name string
// 	println("Welcome to The Solar System! ")
// 	println("There are 9 planets to explore.")
// 	fmt.Println("What is your name?")
// 	name := reader.readString("\n")
// 	fmt.Println(name, "My name is Eliza, I'm an old friend of Alex. Let's go on an adventure!")
// }
