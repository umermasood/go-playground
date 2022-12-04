# 10. Pointers

- Creating Pointers
- De-referencing Pointers
- The `new` function
- Working with `nil`
- Types with internal pointers

---

## Creating Pointers

```go
package main

import "fmt"

func main() {
	var a int = 10
	var b *int = &a
	fmt.Println(a, b)
	fmt.Println(a, *b)

	a = 20
	fmt.Println(a, *b)

	*b = 30
	fmt.Println(a, *b)
}
```

Output:

```
10 0xc00001a100
10 10
20 20
30 30
```

When we use the `&` (Address of operator) to access the address of a variable, it is called _referencing_. When we use
the `*` to access the value stored in the address, it is called _de-referencing_
---

We can also create variables with pointers like this

```go
package main

import "fmt"

func main() {
	var ms *myStruct
	ms = &myStruct{foo: 42}
	fmt.Println(ms)
}

type myStruct struct {
	foo int
}
```

Output:

```
&{42}
```


## `new`
We can also use `new` to create new objects and variables

```go
package main

import "fmt"

func main() {
	var ms *myStruct
	fmt.Println(ms)
	ms = new(myStruct)
	fmt.Println(ms)
}

type myStruct struct {
	foo int
}
```

Output:

```
<nil>
&{0}
```

How do we set the value to the `foo` property in the ms struct?

```go
package main

import "fmt"

func main() {
	var ms *myStruct = new(myStruct)
	// setting the value of the foo property
	(*ms).foo = 10
	fmt.Println(ms)
	// Same thing as (*ms).foo, just syntactic sugar for better readability
	ms.foo = 20
	fmt.Println(ms)
}

type myStruct struct {
	foo int
}
```

> Maps and slices have pointer reference to the underlying data since they are of the _reference type_
