# 2. Variables

1. Variable Declaration
2. Redeclaration & Shadowing
3. Visibility and Scope
4. Naming Conventions
5. Type Conversions

### 1. Variable Declaration

There are 3 different ways / formats to declare variables in Go:

i. `varkeyword variableName type`

```go
var i int

i = 666 // Assign the value later

```

ii. `varkeyword variableName type = value`

```go
var i int = 666
```

iii. `variableName := value`

```go
i := 69
```

In this case, go will infer the type of the value and assign it to the variable

---

> NOTE: When we're declaring variables at the _package level_, we can't use the `:=` (colon equals) syntax. We have to use the full declaration syntax

like the variable `i` in the example below:

```go
package main

import "fmt"

// Package-level declaration
var i float32 = 40

func main() {
    fmt.Printf("%v, %T", i, i)
}
```

> NOTE: Another thing we can do at the Package-Level is that we can create a **Block Declaration**

```go
package main

import "fmt"

// Package-level Block Declaration
var (
    name string = "Rick Sanchez"
    aka string = "Rick C-137"
    age int = 70
    catchphrase string = "Wubbalubbadubdub"
)

func main() {

}
```

---

### 2. Redeclaration & Shadowing

How do variables behave when we redeclare them in Go?

```go
package main

import "fmt"

var i int = 420

func main() {
    var i int = 69
    i := 666
    fmt.Println(i)
}
```

This example will throw error at line 9, that's because the variable `i` is already declared in line 8, so we can't declare the variable twice in the same scope. The variable with the inner most scope will take precedence, this concept is known as _Shadowing_. 69 will print to the standard output. The package level `i` is still available but it's being hidden by the declaration of `i` in the main function. We can see the difference with the following example:

```go
package main

import "fmt"

var i int = 420

func main() {
	fmt.Println(i)
	var i int = 69
	fmt.Println(i)
}
```

```
Output:

420
69
```

> NOTE: Variables in Go always have to be used, otherwise the compiler throws error and stops the program from being executed in the first place.

---

### 3. Visibility & Scope

There are three levels of visibility variables in Go:

1. If you have variable at the package level and it's lowercased, it's scoped to the package. Any file in the same package can access that variable.
2. If you have a variable at the package level and it's uppercased, it is scoped to the package as well as globally exported to the external packages
3. Variables declared inside a block scope are only visible to that block in which they're declared in

---

### 4. Naming Conventions

There are two things that we need to keep track of:

#### 1. Variable Naming effects its visibility

- Variables starting with lowercase in a package are hidden to external packages / outside world.
- Variables starting with uppercase letter in a package are exposed / visible to external packages / outside world.

#### 2. Naming conventions

- Length of the variable name should reflect the life of the variable
- camelCase (package scope / block scope)
- PascalCase (exported to external packages)
- use uppercase for Acronyms: like URL and HTTP

---

### 5. Type conversion

One thing to remember is that Go doesn't do implicit type conversions. The programmer has to explicitly perform type conversions so that they're responsible for the conversion.

```go
package main

import "fmt"

func main() {
	var i int = 69
	fmt.Printf("%v, %T\n", i, i)
	var j float32
	j = float32(i) // Explicit type conversion
	fmt.Printf("%v, %T\n", j, j)
}
```

```
Output:

69, int
69, float32
```
