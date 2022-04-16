# 4. Arrays and Slices

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
