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
