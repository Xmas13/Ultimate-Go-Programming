package main

import "fmt"

func main() {
	// Declare variables that are set to their zero value
	// This is better to use than short variable declaration because it is very obvious
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n\n", d, d)

	// Declaire variables and initialize.
	// Using the short variable declaration operator
	// This is a productivity operator.
	// This is not duck typing, it is struct typing.
	// You can set something like aa := 0, but that is where var should be used
	aa := 10
	bb := "hello"
	cc := 3.14159
	dd := true

	fmt.Printf("aa := 10 \t %T [%v]\n", aa, aa)
	fmt.Printf("bb := \"hello\" \t %T [%v]\n", bb, bb)
	fmt.Printf("cc := 3.14159 \t %T [%v]\n", cc, cc)
	fmt.Printf("dd := true \t %T [%v]\n\n", dd, dd)

	// Specify type and perform a conversion.
	// This is not casting, there is a memory cost in conversion
	// Conversion is safer than casting
	// This is an integrity choice
	var aaa int32

	fmt.Printf("var aaa int32 \t %T [%v]\n", aaa, aaa)

	aaa = int32(aa)

	fmt.Printf("int32(aa) is new variable aaa from conversion %T [%v]\n", aaa, aaa)
}
