# Go Generics

## What are Generics?

Introduced in **Go 1.18**, **Generics** allow you to write functions, structs, and interfaces that work with multiple data types while retaining complete compile-time type safety.

Before generics, if you wanted a function to operate on slices of different numeric types, you had to duplicate the function for every single type:

```go
func sumIntSlice(slice []int) int { /* ... */ }
func sumFloat32Slice(slice []float32) float32 { /* ... */ }
func sumFloat64Slice(slice []float64) float64 { /* ... */ }
```

Or you had to use `interface{}` / `any`, which sacrifices type safety and requires runtime type assertions or reflection.

With Generics, you write the logic **once**:

```go
func sumSlice[T int | float32 | float64](slice []T) T {
    var sum T
    for _, v := range slice {
        sum += v
    }
    return sum
}
```

---

## Generic Function Syntax

A generic function introduces a **type parameter list** in square brackets `[...]` before the regular parameter list:

```go
func FunctionName[T TypeConstraint](param T) T {
    // code
}
```

- `T` is the **Type Parameter** (a placeholder for the concrete type).
- `TypeConstraint` is the contract defining which types are allowed to replace `T`.

### Example from `tutorial_10`

```go
package main

import "fmt"

func sumSlice[T int | float32 | float64](slice []T) T {
    var sum T
    for _, v := range slice {
        sum += v
    }
    return sum
}

func main() {
    intSlice := []int{1, 2, 3}
    fmt.Println(sumSlice[int](intSlice)) // 6

    float32Slice := []float32{1.1, 2.2, 3.3}
    fmt.Println(sumSlice[float32](float32Slice)) // 6.6

    float64Slice := []float64{1.1, 2.2, 3.3}
    fmt.Println(sumSlice[float64](float64Slice)) // 6.6
}
```

---

## Type Inference

In many cases, the Go compiler can automatically infer the type argument from the function parameters, so you don't have to specify `[T]` manually:

```go
intSlice := []int{1, 2, 3}

// Explicit type argument:
sumSlice[int](intSlice)

// Inferred type argument (cleaner!):
sumSlice(intSlice)
```

Both produce identical compiled code.

---

## Built-in Constraints: `any` and `comparable`

Go provides two universal built-in constraints:

### 1. `any`
`any` is an alias for `interface{}`. It permits **any** type whatsoever.

```go
func PrintSlice[T any](items []T) {
    for _, item := range items {
        fmt.Println(item)
    }
}
```

### 2. `comparable`
`comparable` is a built-in constraint implemented by all types that support the `==` and `!=` comparison operators (booleans, numbers, strings, pointers, channels, and structs containing comparable fields).

```go
func FindIndex[T comparable](items []T, target T) int {
    for i, item := range items {
        if item == target { // Valid because T is comparable
            return i
        }
    }
    return -1
}
```

> [!NOTE]
> Slices, maps, and functions are **not** comparable and cannot be used with `comparable`.

---

## Custom Type Constraints with Interfaces

Instead of writing union types like `int | float32 | float64` in every function signature, you can define reusable constraint interfaces:

```go
type Number interface {
    int | int8 | int16 | int32 | int64 | float32 | float64
}

func sumSlice[T Number](slice []T) T {
    var sum T
    for _, v := range slice {
        sum += v
    }
    return sum
}
```

---

## The Approximation Operator (`~`)

Consider a defined type:

```go
type MyID int
```

If your constraint is `int`, `MyID` will **not** be allowed because its underlying type is `int`, but its actual named type is `MyID`.

To allow any type whose **underlying type** matches `int`, use the tilde `~` operator:

```go
type Integer interface {
    ~int | ~int32 | ~int64
}

func ProcessID[T Integer](id T) {
    fmt.Println(id)
}

func main() {
    var id MyID = 42
    ProcessID(id) // Works because ~int permits MyID!
}
```

---

## Generic Structs

You can also make structs generic to build reusable data containers:

```go
type Stack[T any] struct {
    elements []T
}

func (s *Stack[T]) Push(item T) {
    s.elements = append(s.elements, item)
}

func (s *Stack[T]) Pop() (T, bool) {
    var zero T
    if len(s.elements) == 0 {
        return zero, false
    }
    lastIndex := len(s.elements) - 1
    item := s.elements[lastIndex]
    s.elements = s.elements[:lastIndex]
    return item, true
}
```

Usage:

```go
// Stack of strings
stringStack := Stack[string]{}
stringStack.Push("first")
stringStack.Push("second")

// Stack of integers
intStack := Stack[int]{}
intStack.Push(100)
```

---

## Generics vs Interfaces: When to use which?

| Use Generics When | Use Interfaces When |
|---|---|
| Implementing data structures (trees, linked lists, stacks, queues) | Modeling domain behavior and capabilities (e.g. `Reader`, `Writer`, `Closer`) |
| Writing general-purpose helper functions for slices or maps | Different types implement the same method in completely different ways |
| The return type must match the parameter type exactly | You need to store collections of heterogeneous types (different types in one slice) |
| Avoiding repetitive boilerplate across numeric types | Writing decoupled, mockable unit tests |

---

## Summary

- Generics let you write type-safe code that works with multiple types.
- Type parameter syntax: `[T Constraint]`.
- Go supports type inference: you can often omit `[T]` when calling generic functions.
- `any` allows all types; `comparable` allows types supporting `==` and `!=`.
- Use union interfaces (`int | float64`) and `~` for underlying type constraints.
- Structs can also be generic, enabling flexible and type-safe data structures.

---

## Related Notes

- [Go Basic Knowledge Guide](./golang-basics.md)
- [Go Goroutines and Mutex](./golang-goroutines-and-mutex.md)
- [Go Channels](./golang-channels.md)
- [Go Structs and Interfaces](./golang-structs-and-interfaces.md)
- [Go File Structure and Scope](./golang-structure-and-scope.md)
