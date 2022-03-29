package main

import "fmt"

func main() {
	// Creates a string slice of length 2 and capacity 3
	slc := make([]string, 2, 3)

	// Printing an empty slice of length 2
	fmt.Println("slice :", slc)             // slice : [ ]
	fmt.Println("slice length :", len(slc)) // slice length : 2

	// Initializing the slice, using up all the available length
	slc[0] = "rick"
	slc[1] = "morty"

	// Now we can't do slc[2] = "summer", because it would throw an error
	// because we ran out of available length 2
	// So we comment the next line
	// slc[2] = "summer"
	// try running the above line by uncommenting

	fmt.Println("slc Capacity :", cap(slc))  //slc Capacity : 3
	fmt.Println("slice length :", len(slc))  // slice length : 2
	fmt.Println("slc[0] address :", &slc[0]) // Memory address of slc[0]

	// We still have capacity to add new elements. We can do that by using the
	// builtin append() method

	slc = append(slc, "summer")
	fmt.Println("slc Capacity :", cap(slc))  //slc Capacity : 3
	fmt.Println("slice length :", len(slc))  // slice length : 3
	fmt.Println("slc[0] address :", &slc[0]) // Memory address of slc[0]

	// Now we have run out of capacity because we have used up all the length
	// and available capacity. It would be reasonable to expect that we won't
	// be able to add more elements to slc now. But go allows you to add more
	// elements even if you have run out of capacity.

	slc = append(slc, "new_element")
	fmt.Println("slc Capacity :", cap(slc), "Capacity doubled") //slc Capacity : 6
	fmt.Println("slice length :", len(slc))                     // slice length : 3
	// Memory address of slc[0] is different now
	fmt.Println("slc[0] address :", &slc[0], "Memory Address Changed")

	// The above chunk of code appends "new_element" to the slc by creating a
	// new copy of the previous slice which has double capacity compared to
	// previous slice and has different memory
}
