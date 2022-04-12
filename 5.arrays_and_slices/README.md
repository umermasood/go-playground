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
