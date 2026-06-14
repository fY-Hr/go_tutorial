# Go Structs and Interfaces

## What is a struct?

A `struct` is a custom type that groups related fields together.

Example:

```go
type person struct {
    name string
    age  int
}
```

This says:

- `person` is a new type
- it has a `name` field
- it has an `age` field

Structs are often used to model real objects or related data in one place.

---

## Creating a struct value

### 1. Named fields

```go
user := person{
    name: "Alice",
    age:  30,
}
```

This is the safest and most readable form for beginners.

### 2. Positional form

```go
user := person{"Alice", 30}
```

This works, but it is easier to make mistakes because the order must match the struct definition exactly.

### 3. Zero value

```go
var user person
```

This creates a `person` with zero values:

- `name` becomes `""`
- `age` becomes `0`

### 4. Anonymous struct

You can also create a struct value without creating a reusable named type:

```go
user := struct {
    name string
    age  int
}{
    name: "Alice",
    age:  30,
}
```

This is useful for short, local examples, but if you want to reuse the type in many places, a named struct is usually better.

---

## Accessing and updating fields

Use `.` to access a field:

```go
fmt.Println(user.name)
user.age = 31
```

---

## Structs can contain other structs

Your `tutorial_6` example uses this pattern:

```go
type pilot struct {
    name   string
    age    uint8
    gender string
}

type gundamType struct {
    model        string
    pilot        pilot
    manufacturer string
}
```

Creating a nested value:

```go
g := gundamType{
    model: "RX-93-v2 Hi-v",
    pilot: pilot{
        name:   "Amuro Ray",
        age:    20,
        gender: "male",
    },
    manufacturer: "Anaheim Electronics",
}
```

Accessing nested fields:

```go
fmt.Println(g.pilot.name)
```

---

## Anonymous field embedding

Go also allows embedding a type directly as a field:

```go
type pilot struct {
    name string
    age  uint8
}

type gundamType struct {
    model string
    pilot
}
```

This means `gundamType` still contains a `pilot`, but Go also lets you access embedded fields more directly:

```go
g := gundamType{
    model: "Nu Gundam",
    pilot: pilot{name: "Amuro Ray", age: 20},
}

fmt.Println(g.pilot.name)
fmt.Println(g.name) // promoted from embedded pilot
```

Important:

- embedding is not inheritance
- it is a convenient way to reuse fields and methods
- the embedded type is still a normal field

---

## Methods on structs

A method is a function attached to a type.

```go
type pilot struct {
    name string
    age  uint8
}

func (p pilot) sayHello(robot string) {
    fmt.Printf("Hello, I am %s, I am %d years old, and I pilot %s\n", p.name, p.age, robot)
}
```

Usage:

```go
amuro := pilot{name: "Amuro Ray", age: 20}
amuro.sayHello("RX-93-v2 Hi-v")
```

The `(p pilot)` part is called the receiver.

More precisely:

- in `func (p pilot) sayHello(...)`, `p` is the receiver variable
- `pilot` is the receiver type
- together they attach the method to the `pilot` type

Your `tutorial_6` file uses this exact idea.

---

## Value receiver vs pointer receiver

### Value receiver

Works on a copy of the value:

```go
func (p pilot) sayHello(robot string) {
    p.age++
    fmt.Printf("%s is now %d while talking about %s\n", p.name, p.age, robot)
}
```

Even though `p.age++` changes `p`, it only changes the copy inside the method.

Example:

```go
amuro := pilot{name: "Amuro Ray", age: 20}
amuro.sayHello("Nu Gundam")
fmt.Println(amuro.age) // still 20
```

### Pointer receiver

Works on the original value:

```go
func (p *pilot) aging() {
    p.age++
}
```

Usage:

```go
amuro := pilot{name: "Amuro Ray", age: 20}
amuro.aging()
fmt.Println(amuro.age) // 21
```

Use pointer receivers when:

- you want to modify the original value
- the struct is large and copying it would be wasteful
- you want method behavior to be consistent across the type

Important beginner note:

```go
amuro.aging()
```

works even though `aging()` is defined on `*pilot`, because Go can automatically take the address of an addressable value in this kind of method call.

---

## What is an interface?

An `interface` defines behavior by listing method signatures.

Example:

```go
type mobileSuit interface {
    useWeapon() string
}
```

This means any type with a `useWeapon() string` method satisfies the interface.

Important:

- an interface defines methods
- it does not declare struct fields or "properties"
- the concrete struct can have any fields it wants

---

## Interface implementation is implicit

In Go, you do not write `implements`.

If a type has the required methods, it automatically satisfies the interface.

Your `tutorial_6` example uses this same idea with different types:

```go
type gundamType struct {
    model string
}

type zakuType struct {
    model string
}

func (g gundamType) useWeapon() string {
    return "Laser Beam"
}

func (z zakuType) useWeapon() string {
    return "Thermal Axe"
}

type mobileSuit interface {
    useWeapon() string
}
```

Both `gundamType` and `zakuType` satisfy `mobileSuit` because both types define `useWeapon() string`.

---

## Using an interface

```go
func useWeapon(robot mobileSuit) string {
    return robot.useWeapon()
}

func main() {
    g := gundamType{model: "RX-93-v2 Hi-v"}
    z := zakuType{model: "Zaku Mk II"}

    fmt.Println(useWeapon(g))
    fmt.Println(useWeapon(z))
}
```

Why this is useful:

- one function can work with many different types
- code depends on behavior, not concrete types
- it helps keep code flexible
- this is a very common use of interfaces in Go: function parameters

---

## Interface with multiple methods

```go
type shape interface {
    area() float64
    perimeter() float64
}
```

A type must implement all listed methods to satisfy the interface.

---

## Pointer receivers and interfaces

Be careful with pointer receiver methods.

```go
type counter struct {
    value int
}

func (c *counter) increment() {
    c.value++
}

type incrementer interface {
    increment()
}
```

This works:

```go
c := &counter{}
var inc incrementer = c
inc.increment()
```

This does not work:

```go
c := counter{}
var inc incrementer = c // compile error
```

Why:

- `counter` does not have the method set required by the interface
- `*counter` does

---

## Type assertion

Sometimes you want to get the concrete value back from an interface.

```go
var s speaker = dog{name: "Bolt"}

d, ok := s.(dog)
fmt.Println(d, ok)
```

If the underlying value is not `dog`, `ok` becomes `false`.

---

## Type switch

Use a type switch when you want different behavior for different concrete types:

```go
func describe(value interface{}) {
    switch v := value.(type) {
    case int:
        fmt.Println("int:", v)
    case string:
        fmt.Println("string:", v)
    default:
        fmt.Println("unknown type")
    }
}
```

In modern Go, `any` is an alias for `interface{}`:

```go
func describe(value any) {
    fmt.Println(value)
}
```

---

## Common beginner notes

- a `struct` stores data
- a method adds behavior to a type
- an `interface` describes required behavior
- an interface usually describes a small ability, such as `useWeapon()` or `Read()`
- Go interfaces are usually small and focused
- embedding is not class inheritance
- prefer simple, clear structs before trying advanced patterns

---

## Practical mental model

Think of it like this:

- `struct` = data shape
- method = behavior attached to that data
- `interface` = a contract for behavior

Example:

- `car` and `bike` can both be structs
- both can have a `start()` method
- a `starter` interface can accept either one

---

## Summary

- use `struct` to group related fields
- use methods to attach behavior to a type
- use pointer receivers when you need to modify the original value
- use interfaces to write flexible code based on behavior
- use anonymous structs for short local values when you do not need reuse
- remember that interface implementation in Go is automatic, not explicit

---

## Related Notes

- [Go Basic Knowledge Guide](./golang-basics.md)
- [Go File Structure and Scope](./golang-structure-and-scope.md)
