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

> it is conventional to use `ok` in go programs for checking existence of a key-val pair in a map

```go
_, ok := m["k3"]     // if we don't need the value we can use write only operator
val, ok := m["k3"]   // if we want to store the value
```

> **Important:** Maps are reference type, just like slices in go, manipulating a map in one place is going to have impact
> on every other place that maps is used

---

## Structs

_Struct is a collection type consisting of fields. Structs are useful for grouping data together to form records. Unlike Maps, structs are value types_

This is how to use Structs in Go:

```go
package main

import "fmt"

type Superhero struct {
  id         byte
  name       string
  superpower string
  species    string
  companions []string
}

func main() {
  supe0 := Superhero{
    id:         0,
    name:       "Batman",
    superpower: "Rich",
    species:    "Human",
    companions: []string{"Alfred", "Robin", "Batwoman"},
  }

  fmt.Printf("%v, %T", supe0, supe0)
}
```

Output:

```
{0 Batman Rich Human [Alfred Robin Batwoman]}, main.Superhero
```

### Anonymous Structs

Usually used for short life structs. Anonymous structs are not reusable. This is how we use anonymous structs:

```go
package main

import "fmt"

func main() {
  supe0 := struct {
    id         byte
    name       string
    superpower string
    companions []string
  }{
    id:         0,
    name:       "Batman",
    superpower: "Rich",
    companions: []string{"Alfred", "Robin", "Batwoman"},
  }

  fmt.Printf("%v, %T", supe0, supe0)
}
```

Output:

```
{0 Batman Rich [Alfred Robin Batwoman]}, struct { id uint8; name string; superpower string; companions []string }
```

> **Important:** Unlike Maps and Slices, Structs are value type, but we can change this behavior by using pointers

This is how we change the default behavior of the structs in Go:

```go
package main

import "fmt"

type Person struct {
  name string
  age  int
}

func newPerson(name string) *Person {
  p := Person{name: name}
  p.age = 69
  return &p
}

func main() {
  // This becomes reference type
  p1 := newPerson("Rick Sanchez")
  // Since p1 is already a reference type variable p2 will point to the same underlying data
  p2 := p1

  fmt.Println(p1)

  // changes p1 reflect p2
  p1.name = "Morty"

  fmt.Println(p2) // Prints &{Morty 69}
}
```

Output:

```
&{Rick Sanchez 69}
&{Morty 69}
```

### Embedding

Struct embedding is the practice of embedding one struct in another

Embedding might be confused with inheritance, but it uses a compositional model approach rather than classic OOP Inheritance model.

Key takeaways from embedding in structs are:

- Inheritance is IS-A relationship
- Embedding is HAS-A relationship

Here's how we embed one struct in another:

```go
package main

import "fmt"

type Animal struct {
  name   bool
  origin bool
}

type Bird struct {
  Animal
  speedKPH float32
  canFly   bool
}

func main() {
  b := Bird{
    speedKPH: 0,
    canFly:   false,
    Animal:   Animal{name: true, origin: false},
  }
  fmt.Printf("%v, %T", b, b)
}
```

Output:

```
{{true false} 0 false}, main.Bird
```

In the above example, `Bird` has no relationship to the struct `Animal` other than the fact that it embeds it.
Another important thing to note is that in the above example we have used initializer-syntax for the bird struct, in this
case we have to be aware of the embedding if we want to use Animal properties.
