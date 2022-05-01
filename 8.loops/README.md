# 8. Loops

- For Statements
  - Simple Loops
  - Exiting Early
  - Looping through collections


## Simple Loops

The most basic form of `for` loop:

```go
package main

import "fmt"

func main() {
  for i := 0; i < 5; i++ {
    fmt.Println(i)
  }
}
```

Output:

```
0
1
2
3
4
```

A simple for loop with multiple variables:

```go
package main

import "fmt"

func main() {
  // Printing evens and odds using for loop
  fmt.Printf("Evens\tOdds\n")
  for i, j := 0, 1; i <= 10; i, j = i+2, j+2 {
    fmt.Printf("%v\t%v\n", i, j)
  }
}
```

Output:

```
Evens   Odds
0       1
2       3
4       5
6       7
8       9
10      11
```

## Exiting Early from Loops:

```go
package main

import "fmt"

func main() {
  for i := 0; i < 10; i++ {
    if i == 5 {
      break
    }
    fmt.Println(i)
  }
}
```

Output:

```
0
1
2
3
4
```

A do-while alternative in Go, since there is no do-while in Go:

```go

package main

import "fmt"

func main() {
  var input int
  for ok := true; ok; ok = input != 2 {
    _, _ = fmt.Scanln(&input)
    switch input {
    case 1:
      fmt.Println("Hello Friend!")
    case 2:
      // Do something here yo
    default:
      fmt.Println("Default")
    }
  }
}
```

We can omit the initializer and expression part of for loop for some special cases:

```go
package main

import "fmt"

func main() {
  i := 0
  for i < 5 {
    fmt.Printf("%v ", i)
    i++
  }
}
```

Output:

```
0 1 2 3 4
```

`continue` in loops

```go
package main

import "fmt"

func main() {
  // Printing odd numbers using continue
  for i := 0; i < 10; i++ {
    if i%2 == 0 {
      continue
    }
    fmt.Println(i)
  }
}
```

Output:

```
1
3
5
7
9
```

Infinite loops in Go:

```go
package main

import "fmt"

func main() {
  i := 0
  for true {
    fmt.Println("ayo sup B")
    i++
    if i == 3 {
      break
    }
  }
  // OR
  i = 0
  for {
    fmt.Println("ayo sup!?")
    i++
    if i == 3 {
      break
    }
  }
}
```

```
ayo sup B
ayo sup B
ayo sup B
ayo sup!?
ayo sup!?
ayo sup!?
```

Exiting early using `Loop Label`

```go
package main

func main() {
Loop:
  for i := 1; i <= 3; i++ {
    for j := 1; j <= 3; j++ {
      if i*j <= 3 {
        break Loop
        // break
        // If we used only break, then only the inner loop would've exited
      }
    }
  }
}
```

Output:

```
1
2
```

## Looping through collections

In Go, we can loop through collections such as slices and arrays using `for-range` loops.

```go
package main

import "fmt"

func main() {
  // range over a slice
  fmt.Println("Slice range")
  for _, i := range []byte{1, 2, 3, 4, 5} {
    fmt.Println(i)
  }

  // range over an array
  fmt.Println("Array range")
  for index, val := range [5]int{1, 2, 3, 4, 5} {
    fmt.Println(index, val)
  }
}
```

Output:

```
Slice range
1
2
3
4
5
Array range
0 1
1 2
2 3
3 4
4 5
```

Another example: For-range over a map:

```go
package main

import "fmt"

func main() {
  // this method fetches both keys and values
  for k, v := range map[byte]string {0: "zero", 1: "one", 2: "two", 3: "three", 4: "four"} {
    fmt.Println(k, v)
  }
}
```

Output:

```
2 two
3 three
4 four
0 zero
1 one
```

> Accessing map returns values in arbitrary order

Another use case, when we only need the keys:

```go
package main

import "fmt"

func main() {
	// this method only fetches the keys from the map
	m := map[byte]string{0: "zero", 1: "one", 2: "two"}
	for k := range m {
		fmt.Println(k, m[k])
	}
}
```

Output:

```
2 two
0 zero
1 one
```

Ranging over a string in Go:

```go
package main

import "fmt"

func main() {
  str := "This is a string"
  for k, v := range str {
    fmt.Println(k, v)
  }
}
```

Output:

```
0 84
1 104
2 105
3 115
4 32
5 105
6 115
7 32
8 97
9 32
10 115
11 116
12 114
13 105
14 110
15 103
```

We get digits along with indices because we get strings are basically represented as unicode in Go.
