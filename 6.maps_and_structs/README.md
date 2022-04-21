# Maps and Structs

These are the two remaining collection types in Go.

- Maps
  - What are Maps?
  - How to create Maps?
  - How to manipulate Maps?

- Structs
  - What are Structs?
  - How to create Structs?
  - Naming conventions
  - Embedding
  - Tags

---

## Maps

Maps are basically key-value pairs like dictionary in python or like a JSON

> Constraint on key of a map in Go is they have to be able to be tested for equality, all the primitives can be tested
> for equality, so they can be used for a map key. There are other data types which can be used for key in maps.
> However there are some data types which cannot be checked for equivalency and those are slices maps and other
> functions

This is how we create Maps in Go

```go
package main

import "fmt"

func main() {
  // Two different ways to create maps in Go
  m1 := map[byte]bool{}     // Declare and init in same line
  m2 := make(map[byte]bool) // Declare and initialize later

  fmt.Println(m1, m2) // prints two empty maps: map[] map[]

  // We can also use an array as a key for a map
  // An array is a valid key-type but a slice is not
  m3 := map[[3]byte]bool{
    [3]byte{}:        false, // the key [0 0 0], since arrays are 0-valued by default
    [3]byte{0, 0, 1}: true,
    [3]byte{0, 0, 0}: true, // value for the key [0 0 0] is overwritten
  }

  fmt.Println(m3)
}
```

Output:

```
map[] map[]
map[[0 0 0]:true [0 0 1]:true]
```

> **Important**: The return order of a Map is not guaranteed

We can manipulate key values in Maps like this:

To delete an entry (key-value pair) from a map, we can use the builtin `delete` function

```go
package main

import "fmt"

func main() {
  m := map[string]int{
    "k1": 10,
    "k2": 20,
    "k3": 30,
  }

  // Before manipulation of map m
  fmt.Println(m)

  // deleting k3
  delete(m, "k3")

  // modifying k2's value
  m["k2"] = 69

  // Checking whether k3 exists in map m after deleting it
  _, exists := m["k3"]     // if we don't need the value we can use write only operator
  val, exists := m["k3"]   // if we want to store the value
  fmt.Println(val, exists) // k3 doesn't exist

  // Adding new entry k4 in the map
  m["k4"] = 99

  // After manipulations in m
  fmt.Println(m)

}
```

Output:

```
map[k1:10 k2:20 k3:30]
0 false
map[k1:10 k2:69 k4:99]
```
