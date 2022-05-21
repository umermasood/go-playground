# 7. If and Switch (Conditionals in Go)

Conditionals in Go are pretty much similar to other languages with some minor differences.

## If-else

```go
package main

import "fmt"

func main() {

	number := 50
	guess := 50

	if guess < 1 || guess > 100 {
		fmt.Println("Guess number must be in 1-100 range")
	}

	if guess >= 1 && guess <= 100 {
		if guess < number {
			fmt.Println("Low")
		}
		if guess > number {
			fmt.Println("High")
		}
		if guess == number {
			fmt.Println("You Got It")
		}
		fmt.Println(guess < number, guess > number, guess == number)
	}
}
```

Output:

```
You Got It
false false true
```

> In use case of logical operators, we have one concept called Short-circuiting.
Go lazily evaluates logical operators.

```go
package main

import "fmt"

func main() {

	number := 50
	guess := 50

	if guess < 1 || returnTrue() || returnTrue() || guess > 100 {
		fmt.Println("Guess number must be in 1-100 range")
	}

	if guess >= 1 && guess <= 100 {
		if guess < number {
			fmt.Println("Low")
		}
		if guess > number {
			fmt.Println("High")
		}
		if guess == number {
			fmt.Println("You Got It")
		}
		fmt.Println(guess < number, guess > number, guess == number)
	}
}

func returnTrue() bool {
	fmt.Println("Returning true")
	return true
}
```

Output:

```
Returning true
Guess number must be in 1-100 range
You Got It
false false true
```

Go executes `returnTrue` only once since OR expression only needs one expression to be true, so what happens is
called Short-circuiting. In this case the second call of `returnTrue` and the `guess > 100` are ignored by Go.
---

## Switch

Switch is a special type of conditional. It allows for writing conditional cases in a very neat way.

```go
package main

import "fmt"

func main() {
	num := 10

	switch num {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2, 3, 4, 5, 6, 7, 8, 9, 10:
		fmt.Println("Two to Ten")
	default:
		fmt.Println("Some Other Number")
	}
}
```

Output:

```
Two to Ten
```

Just like `if` statements supports initializer syntax, we can also do the same thing with switch

```go
package main

import "fmt"

func main() {
	switch num := 0; num {
	case 0:
		fmt.Println("Zero")
	case 1:
		fmt.Println("One")
	case 2, 3, 4, 5, 6, 7, 8, 9, 10:
		fmt.Println("Two to Ten")
	default:
		fmt.Println("Some Other Number")
	}
}
```

Output:

```
Two to Ten
```

We also have a tag-less switch syntax which allows for writing expressions as cases

```go
package main

import "fmt"

func main() {
	num := 0

	switch {
	case num < 0:
		fmt.Println("Num is less than zero")
	case num > 0:
		fmt.Println("Num is greater than zero")
	default:
		fmt.Println("Num is literally equal to zero")
	}
}
```

Output:

```
Num is literally equal to zero
```

> Go's language design has implicit breaks in switch cases, so we don't have to write `break` explicitly after every case
> But this disables the traditional fallthrough style of switch

So if we want to enable fallthrough for switch cases, this is how we can do that:

```go
package main

import "fmt"

func main() {
	num := 10

	switch {
	case num > 0:
		fmt.Println("Num is greater than zero")
		fallthrough
	case num != 0:
		fmt.Println("Num is not equal to zero")
		fallthrough
	case num < 0:
		fmt.Println("Num is less than zero")
	default:
		fmt.Println("I don't know man")
	}
}
```

Output:

```
Num is greater than zero
Num is not equal to zero
Num is less than zero
```

Notice that the case `num < 0` executes because of the `fallthrough` keyword, fallthrough in Go is logic less.
Even though the case `num < 0` is doesn't satisfy based on `num`'s value, Go still executes that case.

---

### Type Switch

This is how we write a type switch in Go:

```go
package main

import "fmt"

func main() {
	var i interface{} = 666

	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
	case byte:
		fmt.Println("i is a byte")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("i is of some other type")
	}
}
```

Output:

```
i is an integer
```

> We can also use `break` explicitly inside switch cases to exit a case on some logic defined

```go
package main

import "fmt"

func main() {
	var i interface{} = 666

	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
		break
		fmt.Println("This is will also print, but not if we use break in previous line")
	case byte:
		fmt.Println("i is a byte")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("i is of some other type")
	}
}
```

Output:

```
i is an integer
```
