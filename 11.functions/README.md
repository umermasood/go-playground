# Functions

- Basic Syntax
- Parameters
- Return Values
- Anonymous Functions
- Functions as Types
- Methods

---

## Basic Syntax

```go
package main

import "fmt"

func main() {
	greeting, name := "こんにちは", "Go"
	greet(&greeting, &name)
	fmt.Println(greeting, name)
}

func greet(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*greeting = "Hello"
	fmt.Println(*greeting)
}
```

Output:

```
こんにちは Go
Hello
Hello Go
```


## Return Values

### Variadic Parameters

```go
package main

import "fmt"

func main() {
	sum(1, 2, 3, 4, 5)
}

func sum(nums ...int) {
	fmt.Printf("%v, %T\n", nums, nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	fmt.Println(sum)
}
```

Output:

```
[1 2 3 4 5], []int
15
```

> Variadic parameter has to be only one and the last parameter of the function along with other params before it.

---


## Return Values

We can also return values from functions.

```go
package main

import "fmt"

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
}

func sum(nums ...int) int {
	fmt.Printf("%v, %T\n", nums, nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
```

Output:

```
[1 2 3 4 5], []int
15
```

### Returning local value as a pointer

```go
package main

import "fmt"

func main() {
	fmt.Println(*sum(1, 2, 3, 4, 5))
}

func sum(nums ...int) *int {
	fmt.Printf("%v, %T\n", nums, nums)
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return &sum
}
```

Output:

```
[1 2 3 4 5], []int
15
```

### Named returned values

```go
package main

import "fmt"

func main() {
fmt.Println(sum(1, 2, 3, 4, 5))
}

func sum(nums ...int) (sum int) {
fmt.Printf("%v, %T\n", nums, nums)
for _, v := range nums {
sum += v
}
return
}
```

Output:

```
[1 2 3 4 5], []int
15
```

### Multiple return values

```go
package main

import "fmt"

func main() {
	res, err := divide(3.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}

func divide(a, b float64) (float64, error) {
	if b == 0.0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}
```

Output:

```
cannot divide by zero
```


> Functions are really powerful in Go. They can be passed around as variables / types into functions, they can be
> used as return values.


## Anonymous Function

This is the most basic example of an anonymous function in Go:

```go
package main

import "fmt"

func main() {
	func() {
		fmt.Println("Hello, World!")
	}()
}
```

Output:

```
Hello, World!
```

> Anonymous functions have many uses, one such use case is for creating an isolated scope.

```go
package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
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

`func()` type

Functions can be stored in variables and passed around to other functions and can also be used as return values.

```go
package main

import "fmt"

func main() {
	f := func() func() {
		return func() {}
	}
	fmt.Printf("%T", f)
}
```

Output:

```
func() func()
```

---


## Methods

```go
package main

import "fmt"

func main() {
	g := greeter{
		greeting: "hello",
		name:     "go",
	}
	g.greet()
}

type greeter struct {
	greeting string
	name     string
}

// Method
func (g greeter) greet() {
	fmt.Println(g.greeting, g.name)
}
```


Output:

```
hello go
```

This `(g greeter)` is what makes this function a method. Method is basically a function that is executing in a known
context.

> In this the greeter object `g` received by the method `greet()` is of value type, it is a copy of the object passed to
the method. Changes to object `g` inside the method won't reflect the original `g`.


We can also have pointer receivers for methods.

```go
package main

import "fmt"

func main() {
	g := greeter{
		greeting: "hello",
		name:     "go",
	}
	g.greet()
	fmt.Println(g.greeting)
}

type greeter struct {
	greeting string
	name     string
}

// Method
func (g *greeter) greet() {
	fmt.Println(g.greeting, g.name)
	g.greeting = "こんにちは"
}
```

Output:

```
hello go
こんにちは
```

---