# 5. Arrays and Slices

Arrays and slices are collection types. Arrays form the basis of slices.

- Arrays

  - Creation
  - Built-in functions
  - Working with Arrays

- Slices
  - Creation
  - Built-in functions
  - Working with Slices

## Arrays

This is how we declare arrays in Go.

```go
package main

import "fmt"

func main() {
	// Different syntax for creating arrays in Go:

	// Create an array that's just large enough to hold the data that is passed
	idk := [...]int{5, 4, 3, 2, 1}
	arr1 := [5]byte{}
	var arr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%v, %T\n", idk, idk)
	fmt.Printf("%v, %T\n", arr1, arr1)
	fmt.Printf("%v, %T\n", arr2, arr2)
}
```

Output:

```
[5 4 3 2 1], [5]int
[0 0 0 0 0], [5]uint8
[1 2 3 4 5], [5]int
```

> Array can only hold the type of data it is declared to hold at compile time.

This is how can declare multidimensional arrays in Go:

```go
package main

import "fmt"

func main() {
	var identityMatrix [3][3]int
	identityMatrix[0] = [3]int{1, 0, 0}
	identityMatrix[1] = [3]int{0, 1, 0}
	identityMatrix[2] = [3]int{0, 0, 1}
	fmt.Println(identityMatrix)
}
```

Output:

```
[[1 0 0] [0 1 0] [0 0 1]]
```

### Built-in `len` function in Go

```go
package main

import "fmt"

func main() {
  var philosophers [3]string
  philosophers[0] = "Socrates"
  philosophers[1] = "Plato"
  philosophers[2] = "Aristotle"

  fmt.Printf("%v, %T, %v\n", philosophers, philosophers, len(philosophers))
}
```

In Go, arrays are considered values, in a lot of languages if you create an array, it is pointing to
value in array, so when we pass them around, we pass the underlying data. But this is not how things work in Go.
When we copy an array to another, it duplicates and creates that array at different location in memory.

```go
package main

import "fmt"

func main() {
  arr1 := [...]int{1, 2, 3}
  arr2 := arr1
  arr2[1] = -10
  fmt.Println(arr1)
  fmt.Println(arr2)
}
```

Output:

```
[1 2 3]
[1 -10 3]
```

> To work with large size arrays, to change this default behavior of array copying, we can use pointers

```go
package main

import "fmt"

func main() {
  arr1 := [...]int{1, 2, 3}
  // Now arr2 points to the same array arr1
  arr2 := &arr1
  // If we make changes using arr2, they will reflect arr1
  arr2[1] = -10
  fmt.Println(arr1)
  fmt.Println(arr2)
}
```

Output:

```
[1 -10 3]
&[1 -10 3]
```

> The fact that arrays have a fixed size that has to be defined at compile time, limits their usefulness.

Array common use case:

## Slices

The number of elements in a slice doesn't necessarily match the size of the backing array.

```go
package main

import "fmt"

func main() {
  // Array of fixed size
  arr := [3]byte{1, 2, 3}
  fmt.Printf("%v, %T, %v, %v\n", arr, arr, len(arr), cap(arr))

  // Slice
  slc := []byte{3, 2, 1}
  fmt.Printf("%v, %T, %v, %v\n", slc, slc, len(slc), cap(slc))
}
```

We also have the cap function on arrays and slices, which basically tells us the size of the underlying array

Slices are naturally reference type. Which means that they refer to the same underlying data.

```go
package main

import "fmt"

func main() {
  // Slice of length 3
  x := []int{1, 2, 3}   // Literal syntax for creating a slice

  // Now x and y point to the same underlying array
  y := x

  // If we make changes using y, they will reflect x
  y[1] = -10
  fmt.Println(x)
  fmt.Println(y)
}
```

Output:

```
[1 -10 3]
[1 -10 3]
```

There are several other ways of creating slices.

```go
package main

import "fmt"

func main() {
  slcA := []byte{1, 2, 3, 4, 5}
  slcB := slcA[:]   // All elements of slcA
  slcC := slcA[1:]  // From index 1 to end
  slcD := slcA[1:3] // From index 1 to 3
  slcE := slcA[0]   // Notice, this doesn't create a slice

  fmt.Printf("%v, %T\n", slcA, slcA)
  fmt.Printf("%v, %T\n", slcB, slcB)
  fmt.Printf("%v, %T\n", slcC, slcC)
  fmt.Printf("%v, %T\n", slcD, slcD)
  fmt.Printf("%v, %T\n", slcE, slcE)
}
```

Output:

```
[1 2 3 4 5], []uint8
[1 2 3 4 5], []uint8
[2 3 4 5], []uint8
[2 3], []uint8
1, uint8
```

> Notice that in [1:3] the first number is inclusive and the second number is exclusive

Another example, that shows we can also create a slice from an array

```go
package main

import "fmt"

func main() {
  arr := [5]byte{1, 2, 3, 4, 5}

  // Create a slice from arr
  slc := arr[1:4]

  // Changes in slc will affect arr, since slice is of reference type
  slc[2] = 0

  fmt.Printf("%v, %T\n", arr, arr)
  fmt.Printf("%v, %T, %v\n", slc, slc, len(slc))
}
```

Output:

```
[1 2 3 0 5], [5]uint8
[2 3 0], []uint8, 3
```

Another method of creating slices in Go

```go
package main

import "fmt"

func main() {
  // Another method to create slices in Go
  slc := make([]byte, 2, 5)

  fmt.Printf("Slc length : %v, Slice Capacity : %v\n", len(slc), cap(slc))
  // append method takes variadic arguments
  slc = append(slc, 1, 2, 4, 5) // Appending 3rd, 4th, 5th and 6th values
  fmt.Printf("Slc length : %v, Slice Capacity : %v\n", len(slc), cap(slc))
}
```

Output:

```
Slc length : 2, Slice Capacity : 5
Slc length : 6, Slice Capacity : 16
```

How slices work in go

```go
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
```

Output:

```
slice : [ ]
slice length : 2
slc Capacity : 3
slice length : 2
slc[0] address : 0xc00007e150
slc Capacity : 3
slice length : 3
slc[0] address : 0xc00007e150
slc Capacity : 6 Capacity doubled
slice length : 4
slc[0] address : 0xc00006a180 Memory Address Changed
```

Concat two slices

```go
package main

import "fmt"

func main() {
  var slc1 = []byte{1, 2, 3, 4, 5}
  var slc2 = []byte{6, 7, 8, 9, 10}
  var concat []byte

  concat = append(slc2, slc1...)
  concat = append(concat, 0, 0, 0)

  fmt.Println(slc1)
  fmt.Println(slc2)
  fmt.Println(concat)
}
```

Output:

```
[1 2 3 4 5]
[6 7 8 9 10]
[6 7 8 9 10 1 2 3 4 5 0 0 0]
```

Kinda special case

```go
package main

import "fmt"

func main() {
  var slc = []byte{1, 2, 3, 4, 5}
  var newSlice []byte

  fmt.Println(slc)

  // This is what we do if we want to split the slice into two
  newSlice = append(slc[:2], slc[3:]...) // get rid of the 3

  fmt.Println(newSlice)
  fmt.Println(slc) // But slc is also updated, this behavior is because slices are reference type
}
```

Output:

```
[1 2 3 4 5]
[1 2 4 5]
[1 2 4 5 5]
```

## Summary

### Arrays

- Collection of items with same type
- Fixed size
- Declaration styles:
  - `a := [3]int{1, 2, 3}`
  - `a := [….]int{1, 2, 3}`
  - `var a [3]int`
- Access via zero-based index
  - `a := [3]int {1, 3, 5} // a[1] == 3`
- `len` function returns size of array
- Copies refer to different underlying data

### Slices

- Backed by underlying array
- Creation Styles:
  - Slice existing array or slice
  - Literal style
  - via `make` function
    - a := make([]int, 10) // create a slice with length and capacity of 10
    - a := make([]int, 10, 100) // slice with length of 10 and capacity of 100
- `len` function returns the length of the slice
- `cap` function returns the capacity of the underlying array
- `append` function to add elements to end of the slice or array
  - May cause expensive copy operations in case of slices, if underlying array is too small and you run out of initial capacity
- Copies refer to same underlying array; slices are of reference type
