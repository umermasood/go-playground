# 4. Constants

- Naming convention
- Typed constants
- Untyped constants
- Enumerated constants
- Enumeration expressions

---

## Naming Convention

Same naming convention applies to constants too. If we want our constant to be exported then we should use PascalCase otherwise we have to use the camelCase style.

---

## Typed constants

We declare typed constants just like we would declare typed variables in Go.

```go
const myConst int = 69
```

> Another charactersitic of a constant is that it has to be assignable at compile time.

i.e. If we want to have a constant whose value is to be determined at runtime, that is not allowed by Go. The value for constants have to be determined at compile time.

> We can declare constants for all Primitive Types.

```go
package main

import (
        "fmt"
)

func main() {
        const a int = 14
        const b string = "foo"
        const c float32 = 3.14
        const d bool = true

        fmt.Printf("%v\n", a)
        fmt.Printf("%v\n", b)
        fmt.Printf("%V\n", c)
        fmt.Printf("%v\n", d)
}
```

> Collection types are inherently mutable.

Which basically means that we can't declare an array to be of a constant type.

> Constants can also be shadowed

```go
package main

import (
    "fmt"
)

const a int16 = 27

func main() {
    const a int = 14
    fmt.Printf("%v\n", a)
}
```

Output:

```
14
```

```go
package main

import (
    "fmt"
)

func main() {
    const a int = 42
    var b int = 27

    fmt.Printf("%v, %T\n", a + b, a + b)
}
```

Output:

```
69, int
```

> But we cannot perform arithmetic operations on mismatched types.

---

## Untyped Constants

In this example go is inferring the type of constant a, since it is not explicitly declared.

```go
package main

import (
    "fmt"
)

func main() {
    const a = 42
    var b int16 = 27
    fmt.Printf("%V, %T\n", a + b, a + b)
}
```

---

## Enumerated Constants

> iota is a counter that we can use when we're creating enumerated constants.

```go
import (
    "fmt"
)

const (
    a = iota
    b
    c
)

const (
    a2 = iota
)

func main() {
        fmt.Printf("%V\n", a)
        fmt.Printf("%V\n", b)
        fmt.Printf("%V\n", c)
        fmt.Printf("%v\n", a2)
}
```

Output:

```
0
1
2
0
```

What this allows us to do is that we can create different constant blocks to ensure that they have different values.

```go
package main

import (
    "fmt"
)

const (
    catSpecialist = iota
    dogSpecialist
    snakeSpecialist
)

func main () {
    var specialistType int = catSpecialist
    fmt.Printf("%v\n", specialistType == catSpecialist)
}
```

This is a very valuable approach if you wanna check if a value hasn't been assigned to a constant yet.

```go
package main

import (
        "fmt"
)

const (
        _ = iota // here this is a write only variable
        catSpecialist
        dogSpecialist
        snakeSpecialist
)
                            I
func main() {
        var specialistType int
        fmt.Printf("%v\n", specialistType == catSpecialist)
}
```

## Summary

### Typed and Untyped Constants

- Constants are immutable, can be shadowed
- Replaced by compiler at compile time
  - Value must be calculable at compile time
- Named like variables
- Typed constants work like immutable variables
  - Can only interoperate with the same types
- Untyped constants work like literals
  - Can interoperate with same types

## Enumerated Constants

- Special symbol iota allows related constants to be created easily
- Iota starts at 0 in each const block and increments by one
- Watch out of constant values that match zero values for variables

## Enumerated Expressions

- Operations that can be determined at compile time are allowed
  - Arithmetic
  - Bitwise operations
  - Bitshifting
