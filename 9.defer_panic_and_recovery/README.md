# 9. Defer Panic & Recovery

## Defer

Defer literally means postpone / delay; to put off an action / event to a later point in time.

Here's an example of using `defer` in a Go program:

```go
package main

import "fmt"

func main() {
	defer fmt.Println("hello, world")
	fmt.Println("hello, friend")
	fmt.Println("hello, gophers")
}
```

Output:

```
hello, friend
hello, gophers
hello, world
```

As we see that hello, world is printed at the very end instead of beginning since it's execution was delayed.
Although it might seem that `defer` moved that statement to the very end of the program and then executed it, but that
is not the actual case here. `defer` didn't move it to the end of the `main()` function, it actually moves it after the
main function but before the main function returned.

> Important:
> The deferred functions are executed in the LIFO order or FILO which is the execution order of a stack.

This behavior can be observed by this example:

```go
package main

import "fmt"

func main() {
	defer fmt.Println("hello, world")
	defer fmt.Println("hello, friend")
	defer fmt.Println("hello, gophers")
}
```

Output:

```
hello, gophers
hello, friend
hello, world
```

One of the most common use cases of `defer` is that it allows you to associate the opening of a resource and the closing
of the resource right next to each other. A more practical example:

```go
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)
}
```

As we can see that we are opening and closing the resource right next to each others, to avoid bugs and resource leaks
but the resource is actually closed after the main function ends but before it returns something

Another thing to keep in mind about the `defer` keyword is that it takes the value of the argument at the time Go goes
through that statement.

```go
package main

import "fmt"

func main() {
	name := "Rick"
	defer fmt.Println(name)
	name = "Genius"
}
```

Output:

```
Rick
```

## Panic

In Go, we don't have exceptions like a lot of other languages.

When a Go application can no longer run, it is in what's known as panic mode.

```go
package main

import "fmt"

func main() {
	x, y := 1, 0
	fmt.Println(x / y)
}
```

Output:
```
panic: runtime error: integer divide by zero

goroutine 1 [running]:
main.main()
```

We see that our program panicked and threw an error, we can also make our program panic based on some situation:

```go
package main

import "fmt"

func main() {
	fmt.Println("Line lumber 1")
	panic("i paniccc cus something bad happened")
	fmt.Println("Line lumber 2")
}
```

Output:

```go
Line lumber 1
panic: i paniccc

goroutine 1 [running]:
main.main()
```

We are causing our program to panic, then it is no longer able to run anymore, therefore Line lumber 2 is not executed at all.

> Go is rarely going to set an opinion about whether an error is something that should be panicked over or not, therefore
> allowing more control for the developer to decide whether that error is a serious problem or not.

> **Important:** Panics happen after deferred statements are executed.

```go
package main

import "fmt"

func main() {
	fmt.Println("start")
	defer fmt.Println("this was deferred")
	panic("something really bad happened")
	fmt.Println("end")
}
```

Output:

```
start
this was deferred
panic: something really bad happened

goroutine 1 [running]:
main.main()
```

The order of execution is really important here:

1. `main()` gets called
2. `defer` statements are executed
3. `panic` statements are executed
4. `return` value is handled

> Any `defer` calls to closed resources would still work even if a function panics

```go
package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("something bad happened")
	fmt.Println("end")
}
```

Output:

```go
package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("something bad happened")
	fmt.Println("end")
}
```

Output:

```
start
2022/05/03 13:05:33 Error: something bad happened
```

Another example:

```go
package main

import (
	"fmt"
	"log"
)

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()
	panic("something bad happened")
	fmt.Println("done panicking")
}

func main() {
	fmt.Println("start")
	panicker()
	fmt.Println("end")
}
```

Output:

```
start
about to panic
2022/05/03 13:44:45 Error: something bad happened
end
```
