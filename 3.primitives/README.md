# 3. Primitives

- Boolean
- Numeric
  - Integers
  - Floating point
  - Complex
- Text

---

## Boolean

This is how you declare booleans in Go:

Use Cases:

- Using booleans as state flags
- Using as logical tests

```go
var t bool = true
var f bool = false

fmt.Printf("%v, %T\n", t, t)
fmt.Printf("%v, %T", f, f)
```

Output:

```
true, bool
false, bool
```

Creating bools without initializing

```go
var f bool

fmt.Printf("%v, %T", f, f)
```

Output:

```
false, bool
```

> Everytime we declare a variable in Go, by default they are zero-valued, initialized with 0.

Using as logical test

```go
t := 1 == 1 // yields true
f := 1 == 2 // yields false

fmt.Printf("%v, %T\n", t, t)
fmt.Printf("%v, %T", f, f)
```

Output:

```
true, bool
false, bool
```

---

## Numeric

### Integers

Various Integer types for storing different numbers of different sizes

Signed Integers

- int8
- int16
- int32
- int64

Unsigned Integers:

> There is an equivalent type of unsigned integer for every signed integer

- uint8
- uint16
- uint32
- uint64

#### Byte

Byte is an alias for 8-bit unsigned integer, The reason we have that is because unsigned 8 bit integer is really common. This is what a lot of data streams use to encode data.

```go
var n byte = 255

fmt.Printf("%v, %T", n, n)
```

Output:

```
255, uint8
```

Basic Arithmetic as expected:

```go
a := 10
b := 3

fmt.Println(a + b)
fmt.Println(a - b)
fmt.Println(a * b)
fmt.Println(a / b)
fmt.Println(a % b)
```

Output:

```
13
7
30
3
1
```

- We cannot perform arithmetic on two integers of different types i.e. uint8 + int16
- Division / operator doesn't give us floating point value because dividing an integer by an integer gives us an integer.
- If we want a floating point division, we have to perform explicit type conversion from int to float

> Go doesn't perform implicit data conversions

The bit operators:

```go
a := 10 // 1010
b := 3  // 0011

fmt.Println(a & b)  // AND		0010
fmt.Println(a | b)  // OR		1011
fmt.Println(a ^ b)  // XOR		1001
fmt.Println(a &^ b) // AND-NOT	0100
```

Output:

```
2
11
9
8
```

Bit Shifting

```go
a := 8 // 2^3

fmt.Println(a << 3) // 2^3 * 2^3 = 2^6 = 64
fmt.Println(a >> 3) // 2^3 / 2^3 = 2^0 = 1
```

Output:

```
64
1
```

---

### Floating Point

All of these float point declaration syntax are valid. Initializer := syntax always creates the float64 by default.

```go
n := 3.14
n = 13.e72
n = 2.1e4

fmt.Printf("%v, %T", n, n)
```

Arithmetic operations on floating point numbers

```go
a := 10.2
b := 3.7

fmt.Println(a + b)
fmt.Println(a - b)
fmt.Println(a * b)
fmt.Println(a / b)
```

Output:

```
13.899999999999999
6.499999999999999
37.74
2.7567567567567566
```

We cannot use the modulo operator on floating point values

> Invalid operation: a % b (the operator % is not defined on float64)

> We also don't have the bitwise operators and bit shifting operators on floating points

---

### Complex

There are two types of complex numbers:

1. Complex 64
2. Complex 128

Basic declarations of two complex numbers

```go
var n complex64 = 1 + 2i
var m complex64 = 2i

fmt.Printf("%v, %T\n", n, n)
fmt.Printf("%v, %T\n", m, m)
```

Output

```
(1+2i), complex64
(0+2i), complex64
```

We can also perform arithmetic operations with complex numbers.

```go
a := 1 + 2i
b := 2 + 5.2i

fmt.Println(a + b)
fmt.Println(a - b)
fmt.Println(a * b)
fmt.Println(a / b)
```

Output:

```
(3+7.2i)
(-1-3.2i)
(-8.4+9.2i)
(0.3994845360824742-0.038659793814433i)
```

---
