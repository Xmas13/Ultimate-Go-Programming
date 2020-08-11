package main

import (
	"fmt"
	"unsafe"
)

// example represents a type with different fields
// This struct is not a 7 byte value due to alignments
// This will change due to hardware
// You want your variables or structs to fall within word boundaries
// This is done automatically by the hardware for efficiency of reads and writes
// This example for us would be an 8 byte value because of the word syze of 8 bytes
// If boundaries did not exist, we would have inefficient reads and writes across memory due to it crossing over
// Instead of 1 read and 1 write, it would be 2 reads and 2 writes every operation
type example struct {
	flag bool // 1 byte
	// This would have 1 byte of padding in between
	// If you did int32 instead, it would have 3 bytes of padding
	// This is due to an int32 being 4 bytes
	// This is a micro optimization, but still important
	counter int16   // 2 bytes
	pi      float32 // 4 bytes
	// stringy string  // 16 bytes
	// 23 total bytes
	// With boundaries, it's actually 24 bytes
	// If you are memory optimizing based on a profile, you would want to go largest value to smallest value
	// This should only be done if you have a memory profile
	// A more readable way to have a struct is to group similar values together
}

type bill struct {
	flag    bool
	counter int16
	pi      float32
}
type alice struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	// Declate a variable of example to set to its zero value
	var e1 example

	// Display the value.
	fmt.Printf("%+v\n", e1)

	// Display total size
	fmt.Printf("Size of: e1: %T, %d\n\n", e1, unsafe.Sizeof(e1))

	// Declare a variable of type example and init using a struct literal
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	// Display the field values.
	fmt.Println("Flag", e2.flag)
	fmt.Println("Counter", e2.counter)
	fmt.Println("Pi", e2.pi)

	fmt.Printf("Size of: e2: %T, %d\n\n", e1, unsafe.Sizeof(e2))

	// Declare a variable of an anonymous type set to its zero value.
	// This is an unnamed struct
	var e3 struct {
		flag    bool
		counter int16
		pi      float32
	}

	// Display the value
	fmt.Printf("%+v\n", e3)
	fmt.Printf("Size of: e3: %T, %d\n\n", e3, unsafe.Sizeof(e3))

	// Declate a variable of an anonymous type and init using a struct literal
	e4 := struct {
		flag    bool
		counter int16
		pi      float32
	}{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	fmt.Printf("%+v\n", e4)
	fmt.Printf("Size of: e4: %T, %d\n\n", e4, unsafe.Sizeof(e4))

	// This will not work due to implicit conversion not existing in Go
	// Implicit conversion causes many issues in software development
	// Go does not have implicit conversion
	// Way to solve this is to signal to the compiler your intention

	// var a bill
	// var b alice

	// b = a
	// fmt.Println(b, a)

	// Proper way to get what you want from above
	var b bill
	var a alice
	b = bill(a)
	fmt.Println(b, a)

}
