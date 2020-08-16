package main

func increment(num int) int {
	num++
	return num
}
func main() {
	// Declare a variable of type int with a value of 10.package Lesson 2 Language Syntax
	count := 10
	var p *int = &count

	// Display the "value of" and "address of" count.
	println("count:\tValue Of[", count, "]\t Addr Of[", &count, "]")
	println("p:\tValue Of[", *p, "]\t Addr Of[", &p, "]")

	// Pass the "value of" the count.
	count = increment(count)
	count = increment(*p)

	println("count:\tValue Of[", count, "]\t Addr Of[", &count, "]")
	println("p:\tValue Of[", *p, "]\t Addr Of[", &p, "]")

}
