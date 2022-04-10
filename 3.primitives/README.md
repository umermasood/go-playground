# 3. Primitives

- Boolean
- Numeric
  - Integers
  - Floating point
  - Complex
- Text
  - String
  - Byte Slice
  - Rune

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

What if you need to access the real part or the imaginary part?

There are two builtin functions for accessing them.

```go
var n complex64 = 1 + 2i
var m complex128 = 1 + 2i

fmt.Printf("%v, %T\n", real(n), real(n))
fmt.Printf("%v, %T\n", imag(m), imag(m))
```

Output:

```
1, float32
2, float64
```

Notice how complex 64s have float32 type and complex128s have float64 type

We can also create complex numbers using the builtin complex function

```go
var num complex128 = complex(5, 12)
```

---

## Text

### String

A string in Go stands for any utf-8 character

```go
s := "this is a string"

fmt.Printf("%v, %T\n", s, s)

// We can also treat strings sort of like array indices
fmt.Printf("%v, %T\n", s[2], s[2])
```

Output:

```
this is a string, string
105, uint8
```

But what does 105, uint8 mean?

> Strings in Go are actually aliases for bytes

```go
// Converting the byte back to string gives us the letter i
fmt.Printf("%v, %T\n", string(s[2]), s[2])
```

Output:

```
i, uint8
```

- Strings are immutable in Go
- We can perform string concatenation

```go
s1 := "This is String 1"
s2 := "This is String 2"

fmt.Printf("%v, %T\n", s1+s2, s1+s2)
```

Output:

```
This is String 1This is String 2, string
```

### Byte Slice

Another thing that we can do with strings is that we can convert them into collection of bytes. In Go, it's called a
slice of byte or a Byte Slice.

But why would we want to use this slice of bytes?

A lot of functions in Go actually work with byte slices. That makes them much more generic and flexible than hardcoded
strings.

```go
s := "This is a string"
b := []byte(s)

fmt.Printf("%v, %T\n", b, b)
```

Output:

```
[84 104 105 115 32 105 115 32 97 32 115 116 114 105 110 103], []uint8
```

### Rune

- String type represents any utf-8 character and rune represents any utf-16 character.
- When we are declaring a string we use double quotes
- When we have to declare a rune, we use single quotes.
- Runes are a type alias for int32's.
- Like strings can be converted back and forth between collection of bytes, runes are a true type alias. So when we talk about a rune in go, it is the same thing as talking about integer 32.

```go
	r1 := 'a'
	var r2 rune = 'b'

	fmt.Printf("%v, %T\n", r1, r1)
	fmt.Printf("%v, %T\n", r2, r2)
```

Output:

```
97, int32
98, int32
```
