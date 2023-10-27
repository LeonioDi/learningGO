package main

import "fmt"

// BEGIN (write your solution here)
func GenerateSelfStory(name string, age int, money float64) string {
	return fmt.Sprintf("Hello! My name is %s. I'm %d y.o. And I also have $%.2f in my wallet right now.", name, age, money)
}

// END
func main() {
	fmt.Print(GenerateSelfStory("Vlad", 25, 10.0000002))
}
