# 12. Interfaces

- Basics
- Composing Interfaces
- Type Conversion
  - The empty interface
  - Type switches
- Implementing vs values vs pointers
- Best Practices

---

### What is an Interface?
An interface is an abstract type in Go that is defined using a set of method signatures.
The interface defines the behavior for similar types of objects.

In contrast with `structs`, structs are data containers whereas interfaces don't describe data, they describe
**behaviors**. We store method signatures in interfaces.

> Key Takeaway: Concept of Implicit Implementation.

Any type that can have a method associated with it can implement an interface. Any type can have methods associated with
it.

### Real Life example for Interface:

- Power socket - A socket provides the functionality of plugging in an electronic device and the device's functionality
is that it is capable of drawing power from the socket

```go
package main

import "fmt"

func main() {
  m := mobile{"Apple"}
  l := laptop{"Intel i9"}
  t := toaster{4}
  k := kettle{"50%"}

  s := socket{}

  s.Plug(m, 10)
  s.Plug(l, 50)
  s.Plug(t, 30)
  s.Plug(k, 25)
}

type socket struct{}

func (s socket) Plug(device PowerDrawer, power int) {
  device.Draw(power)
}

type mobile struct {
  brand string
}

func (m mobile) Draw(power int) {
  fmt.Printf("%T -> brand: %s, power: %d\n", m, m.brand, power)
}

type laptop struct {
  cpu string
}

func (l laptop) Draw(power int) {
  fmt.Printf("%T -> cpu: %s, power: %d\n", l, l.cpu, power)
}

type toaster struct {
  amount int
}

func (t toaster) Draw(power int) {
  fmt.Printf("%T -> amount: %d, power: %d\n", t, t.amount, power)
}

type kettle struct {
  quantity string
}

func (k kettle) Draw(power int) {
  fmt.Printf("%T -> quantity: %s, power: %d\n", k, k.quantity, power)
}

type PowerDrawer interface {
  Draw(power int)
}
```