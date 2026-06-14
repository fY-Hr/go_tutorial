# Go (Golang) Basic Knowledge Guide

## What is Go?

Go (also known as Golang) is an open-source programming language developed by Google. It's designed for simplicity, efficiency, and reliability. Go is statically typed, compiled, and has built-in support for concurrency.

## Table of Contents

1. [Basic Syntax](#basic-syntax)
2. [Variables](#variables)
3. [Data Types](#data-types)
4. [Type Conversion](#type-conversion)
5. [Pointers](#pointers)
6. [Arrays and Slices](#arrays-and-slices)
7. [Maps](#maps)
8. [Constants](#constants)
9. [Functions](#functions)
10. [Control Flow](#control-flow)
11. [Error Handling](#error-handling)
12. [Useful Shortcuts](#useful-shortcuts)
13. [Further Reading](#further-reading)

---

## Basic Syntax

Every Go program starts with a package declaration:

```go
package main
```

Import necessary packages:

```go
import "fmt"
// or multiple imports
import (
    "fmt"
    "errors"
)
```

The `main` function is the entry point:

```go
func main() {
    // your code here
}
```

---

## Variables

### 1. Long Form Declaration

```go
var variableName type = value
```

Example:
```go
var age int = 25
var name string = "John"
var isActive bool = true
```

### 2. Type Inference (Type can be omitted)

```go
var variableName = value
```

Example:
```go
var age = 25        // int
var name = "John"   // string
```

### 3. Short Variable Declaration (Most Common!)

**This is the shortcut you asked about!** Use `:=` inside functions:

```go
variableName := value
```

Example:
```go
age := 25
name := "John"
isActive := true
```

**Important Notes:**
- `:=` can only be used **inside functions**
- It declares AND initializes the variable
- Cannot be used to re-declare variables in the same scope
- Can be used to declare multiple variables at once:

```go
x, y := 10, 20
name, age := "Alice", 30
```

### 4. Multiple Variables Declaration

```go
var (
    name = "Bob"
    age  = 30
    city = "New York"
)
```

### 5. Default Values

If you declare a variable without initializing it, it gets a default value:

```go
var intDefault int        // 0
var floatDefault float64  // 0
var stringDefault string  // "" (empty string)
var boolDefault bool      // false
var errDefault error      // nil
```

---

## Data Types

### Numeric Types

- **Integers:** `int`, `int8`, `int16`, `int32`, `int64`, `uint`, `uint8` (byte), `uint16`, `uint32`, `uint64`
- **Floats:** `float32`, `float64`
- **Complex:** `complex64`, `complex128`

```go
var i int = 42
var f float64 = 3.14
var b byte = 255
```

### String

```go
var s string = "Hello, World!"
// or
s := "Hello, World!"
```

String concatenation:
```go
greeting := "Hello" + " " + "World"
```

String length:
```go
len("Hello")  // 5
```

### Boolean

```go
var isTrue bool = true
var isFalse bool = false
```

### Rune

Represents a Unicode code point (similar to a character):

```go
var r rune = 'a'  // ASCII value 97
```

---

## Type Conversion

Go does not automatically convert most different types for you.

You must convert explicitly:

```go
var x int = 42
var y float64 = float64(x)
```

Another example:

```go
price := 19.99
whole := int(price) // becomes 19
```

Important:

- `int(price)` converts the value
- converting from `float64` to `int` drops the decimal part
- conversion is different from formatting a value as text

String formatting example:

```go
age := 25
text := fmt.Sprintf("%d", age)
```

---

## Pointers

A pointer stores the memory address of another value.

Example:

```go
value := 10
ptr := &value
```

Here:

- `value` stores `10`
- `&value` means "address of value"
- `ptr` stores that address

To read or change the value through the pointer, use `*`:

```go
fmt.Println(*ptr) // 10
*ptr = 20
fmt.Println(value) // 20
```

Important:

- `&x` gets the address of `x`
- `*ptr` gets the value stored at that address
- pointers let functions or methods modify the original value

Example with a function:

```go
func updateAge(age *int) {
    *age = 30
}

func main() {
    age := 20
    updateAge(&age)
    fmt.Println(age) // 30
}
```

Go does not support pointer arithmetic like C.

---

## Arrays and Slices

### Array Basics

An array in Go has a **fixed size**. The size is part of the type.

```go
var intArr [3]int32 = [3]int32{1, 2, 3}
```

This line has 2 important parts:

- `[3]int32` means: "an array of 3 elements, and each element is an `int32`"
- `{1, 2, 3}` means: "put these values into the array"

### Why do we use `{}` instead of `[]` for the values?

Because in Go:

- `[]` is used in the **type**, like slice type `[]int` or array type `[3]int`
- `{}` is used to **initialize values** in a composite literal

Example:

```go
var arr [3]int = [3]int{1, 2, 3}
```

Here:

- `[3]int` = the type
- `{1, 2, 3}` = the values

So `[]` is not the container for the values. In Go, `{}` is the syntax used to fill arrays, slices, structs, and maps with data.

### Array vs Slice

#### Array

- Fixed size
- Size is part of the type
- Good when you know the exact number of elements

```go
var numbers [3]int = [3]int{10, 20, 30}
```

#### Slice

- Dynamic size
- Built on top of an array
- Much more common than arrays in real Go code
- Use it when the amount of data can grow or shrink

```go
numbers := []int{10, 20, 30}
```

### Main Differences

| Array | Slice |
|-------|-------|
| Fixed size | Dynamic size |
| Type includes length, like `[3]int` | Type does not include length, like `[]int` |
| Less commonly used | Very commonly used |
| Cannot grow with `append` | Can grow with `append` |

### How to make an "unlimited sized array"

In Go, there is **no truly unlimited array**.

If you want something that can grow, use a **slice**:

```go
numbers := []int{}
numbers = append(numbers, 10)
numbers = append(numbers, 20)
numbers = append(numbers, 30)
```

You can also create a slice with `make`:

```go
numbers := make([]int, 0)
numbers = append(numbers, 10)
```

Important:

- Arrays are always fixed-size
- Slices are the Go way to handle changing-sized collections
- A slice is not truly unlimited; it can grow until your program runs out of memory

### Appending another slice (`...`)

In Go, `append` is designed to append **elements**, not a slice as a single value.

That’s why this works:

```go
s := []int32{1, 2, 3}
s = append(s, 4, 5, 6)
```

If you already have another slice, you use `...` to “unpack” its elements:

```go
s := []int32{1, 2, 3}
more := []int32{4, 5, 6}
s = append(s, more...)
```

Without the `...`, Go sees `more` as one value of type `[]int32`, which does not match the element type `int32`, so it won’t compile:

```go
s = append(s, more)
```

If you actually want to append slices as items (a slice of slices), then the element type must be a slice:

```go
groups := [][]int32{}
groups = append(groups, []int32{1, 2, 3})
groups = append(groups, []int32{4, 5, 6})
```

### Useful Note: inferred array size

Go can count the array size for you:

```go
arr := [...]int{1, 2, 3, 4}
```

This is still an **array**, not a slice. Go just infers the size as 4.

---

## Maps

A map stores **key-value pairs**.

### Creating a map

```go
scores := make(map[string]int32)
scores["math"] = 90
scores["english"] = 85
```

You can also create a map with values immediately:

```go
scores := map[string]int32{
    "math":    90,
    "english": 85,
}
```

### Reading from a map

```go
fmt.Println(scores["math"])
```

Be careful: if the key does not exist, Go returns the **zero value** of the map's value type.

```go
value := scores["science"] // 0 if the key does not exist
```

To safely check whether the key exists, use the `value, ok` form:

```go
value, ok := scores["science"]
fmt.Println(value, ok) // 0 false
```

### Updating and deleting

```go
scores["math"] = 95
delete(scores, "english")
```

### Looping through a map

Use `range` with `for`:

```go
for key, value := range scores {
    fmt.Println(key, value)
}
```

Important:

- `key` is the map key
- `value` is the value stored at that key
- Map iteration order is **not guaranteed** in Go

---

## Constants

Constants are declared with `const` and cannot be changed:

```go
const Pi = 3.14159
const Greeting = "Hello"
const (
    StatusOK = 200
    StatusNotFound = 404
)
```

---

## Functions

### Basic Function

```go
func functionName(parameter type) returnType {
    // code
    return value
}
```

Example:
```go
func add(a int, b int) int {
    return a + b
}
```

### Multiple Return Values

Go supports returning multiple values:

```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

### Named Return Values

```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return  // "naked" return
}
```

---

## Control Flow

### If/Else

```go
if condition {
    // code
} else if anotherCondition {
    // code
} else {
    // code
}
```

Short variable declaration in if:
```go
if err := doSomething(); err != nil {
    // handle error
}
```

### For Loop

In Go, `for` is the **only loop keyword**. Go does not have a separate `while` loop. The different loop styles all use `for`.

```go
// Classic for loop
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// While-style loop
for i < 10 {
    fmt.Println(i)
    i++
}

// Infinite loop
for {
    // code
    break  // to exit
}
```

### `for` with `range`

`range` is very common in Go. It is used to loop through arrays, slices, strings, and maps.

#### Loop through a slice

```go
numbers := []int{10, 20, 30}
for index, value := range numbers {
    fmt.Println(index, value)
}
```

Here:

- `index` is the position: `0`, `1`, `2`
- `value` is the actual element: `10`, `20`, `30`

If you only need the value, ignore the index with `_`:

```go
for _, value := range numbers {
    fmt.Println(value)
}
```

If you only need the index:

```go
for index := range numbers {
    fmt.Println(index)
}
```

#### Loop through an array

```go
arr := [3]int{1, 2, 3}
for index, value := range arr {
    fmt.Println(index, value)
}
```

#### Loop through a map

This matches your `tutorial_4` example:

```go
myMap := map[string]uint8{
    "a": 1,
    "b": 2,
}

for key, value := range myMap {
    fmt.Println(key, value)
}
```

This means:

- `key` gets `"a"` or `"b"`
- `value` gets the value stored at that key

#### Loop through a string

```go
for index, char := range "Go" {
    fmt.Println(index, char)
}
```

The `char` value is a `rune`.

### `break` and `continue`

Use `break` to stop the loop completely:

```go
for i := 0; i < 10; i++ {
    if i == 5 {
        break
    }
    fmt.Println(i)
}
```

Use `continue` to skip the current iteration:

```go
for i := 0; i < 5; i++ {
    if i == 2 {
        continue
    }
    fmt.Println(i)
}
```

### Common beginner note

This line:

```go
for key, value := range myMap2 {
    fmt.Println(key, value)
}
```

means:

- start a loop
- read each key-value pair from `myMap2`
- store them in `key` and `value`
- run the code block for each pair

### Switch

```go
switch value {
case 1:
    fmt.Println("One")
case 2:
    fmt.Println("Two")
default:
    fmt.Println("Other")
}
```

Switch without expression (like multiple if-else):
```go
switch {
case age < 18:
    fmt.Println("Minor")
case age >= 18:
    fmt.Println("Adult")
}
```

---

## Error Handling

Go uses explicit error handling with the `error` type:

```go
result, err := someFunction()
if err != nil {
    fmt.Println("Error:", err)
    return
}
// use result
```

Creating errors:
```go
import "errors"

err := errors.New("something went wrong")
```

---

## Useful Shortcuts

| Shortcut | Description |
|----------|-------------|
| `:=` | Short variable declaration (inside functions only) |
| `var()` | Block variable declaration |
| `const()` | Block constant declaration |
| `x, y := a, b` | Multiple variable declaration |
| `val, err := func()` | Capture multiple return values |
| `len(s)` | Get length of string/array/slice |
| `fmt.Println()` | Print with newline |
| `fmt.Printf()` | Formatted print |

---

## Further Reading

For topics that build on these basics:

- [Go Structs and Interfaces](./golang-structs-and-interfaces.md)
- [Go File Structure and Scope](./golang-structure-and-scope.md)

---

## Quick Reference

```go
package main

import "fmt"

func main() {
    // Short variable declaration (most used!)
    name := "Go"
    version := 1.21
    
    // Multiple declaration
    x, y := 10, 20
    
    fmt.Println("Hello,", name)
    fmt.Println("Version:", version)
    fmt.Println("Sum:", x+y)
}
```

---

## Best Practices

1. Use `:=` for variable declaration inside functions (most common and concise)
2. Use `var` for package-level variables or when you need explicit type
3. Always handle errors explicitly
4. Use meaningful variable names
5. Format your code with `go fmt` (automatically!)
