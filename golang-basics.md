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
| Cannot grow with `append` | Can grow with `append` |\n\n### How to make an \"unlimited sized array\"\n\nIn Go, there is **no truly unlimited array**.\n\nIf you want something that can grow, use a **slice**:\n\n```go\nnumbers := []int{}\nnumbers = append(numbers, 10)\nnumbers = append(numbers, 20)\nnumbers = append(numbers, 30)\n```\n\nYou can also create a slice with `make`:\n\n```go\nnumbers := make([]int, 0)\nnumbers = append(numbers, 10)\n```\n\nImportant:\n\n- Arrays are always fixed-size\n- Slices are the Go way to handle changing-sized collections\n- A slice is not truly unlimited; it can grow until your program runs out of memory\n\n### Appending another slice (`...`)\n\nIn Go, `append` is designed to append **elements**, not a slice as a single value.\n\nThat’s why this works:\n\n```go\ns := []int32{1, 2, 3}\ns = append(s, 4, 5, 6)\n```\n\nIf you already have another slice, you use `...` to “unpack” its elements:\n\n```go\ns := []int32{1, 2, 3}\nmore := []int32{4, 5, 6}\ns = append(s, more...)\n```\n\nWithout the `...`, Go sees `more` as one value of type `[]int32`, which does not match the element type `int32`, so it won’t compile:\n\n```go\ns = append(s, more)\n```\n\nIf you actually want to append slices as items (a slice of slices), then the element type must be a slice:\n\n```go\ngroups := [][]int32{}\ngroups = append(groups, []int32{1, 2, 3})\ngroups = append(groups, []int32{4, 5, 6})\n```\n\n### Useful Note: inferred array size\n\nGo can count the array size for you:\n\n```go\narr := [...]int{1, 2, 3, 4}\n```\n\nThis is still an **array**, not a slice. Go just infers the size as 4.\n\n---\n\n## Maps\n\nA map stores **key-value pairs**.\n\n### Creating a map\n\n```go\nscores := make(map[string]int32)\nscores[\"math\"] = 90\nscores[\"english\"] = 85\n```\n\nYou can also create a map with values immediately:\n\n```go\nscores := map[string]int32{\n    \"math\":    90,\n    \"english\": 85,\n}\n```\n\n### Reading from a map\n\n```go\nfmt.Println(scores[\"math\"])\n```\n\nBe careful: if the key does not exist, Go returns the **zero value** of the map's value type.\n\n```go\nvalue := scores[\"science\"] // 0 if the key does not exist\n```\n\nTo safely check whether the key exists, use the `value, ok` form:\n\n```go\nvalue, ok := scores[\"science\"]\nfmt.Println(value, ok) // 0 false\n```\n\n### Updating and deleting\n\n```go\nscores[\"math\"] = 95\ndelete(scores, \"english\")\n```\n\n### Looping through a map\n\nUse `range` with `for`:\n\n```go\nfor key, value := range scores {\n    fmt.Println(key, value)\n}\n```\n\nImportant:\n\n- `key` is the map key\n- `value` is the value stored at that key\n- Map iteration order is **not guaranteed** in Go\n\n---\n\n## Constants\n\nConstants are declared with `const` and cannot be changed:\n\n```go\nconst Pi = 3.14159\nconst Greeting = \"Hello\"\nconst (\n    StatusOK = 200\n    StatusNotFound = 404\n)\n```\n\n---\n\n## Functions\n\n### Basic Function\n\n```go\nfunc functionName(parameter type) returnType {\n    // code\n    return value\n}\n```\n\nExample:\n```go\nfunc add(a int, b int) int {\n    return a + b\n}\n```\n\n### Multiple Return Values\n\nGo supports returning multiple values:\n\n```go\nfunc divide(a, b int) (int, error) {\n    if b == 0 {\n        return 0, errors.New(\"division by zero\")\n    }\n    return a / b, nil\n}\n```\n\n### Named Return Values\n\n```go\nfunc split(sum int) (x, y int) {\n    x = sum * 4 / 9\n    y = sum - x\n    return  // \"naked\" return\n}\n```\n\n---\n\n## Control Flow\n\n### If/Else\n\n```go\nif condition {\n    // code\n} else if anotherCondition {\n    // code\n} else {\n    // code\n}\n```\n\nShort variable declaration in if:\n```go\nif err := doSomething(); err != nil {\n    // handle error\n}\n```\n\n### For Loop\n\nIn Go, `for` is the **only loop keyword**. Go does not have a separate `while` loop. The different loop styles all use `for`.\n\n```go\n// Classic for loop\nfor i := 0; i < 10; i++ {\n    fmt.Println(i)\n}\n\n// While-style loop\nfor i < 10 {\n    fmt.Println(i)\n    i++\n}\n\n// Infinite loop\nfor {\n    // code\n    break  // to exit\n}\n```\n\n### `for` with `range`\n\n`range` is very common in Go. It is used to loop through arrays, slices, strings, and maps.\n\n#### Loop through a slice\n\n```go\nnumbers := []int{10, 20, 30}\nfor index, value := range numbers {\n    fmt.Println(index, value)\n}\n```\n\nHere:\n\n- `index` is the position: `0`, `1`, `2`\n- `value` is the actual element: `10`, `20`, `30`\n\nIf you only need the value, ignore the index with `_`:\n\n```go\nfor _, value := range numbers {\n    fmt.Println(value)\n}\n```\n\nIf you only need the index:\n\n```go\nfor index := range numbers {\n    fmt.Println(index)\n}\n```\n\n#### Loop through an array\n\n```go\narr := [3]int{1, 2, 3}\nfor index, value := range arr {\n    fmt.Println(index, value)\n}\n```\n\n#### Loop through a map\n\nThis matches your `tutorial_4` example:\n\n```go\nmyMap := map[string]uint8{\n    \"a\": 1,\n    \"b\": 2,\n}\n\nfor key, value := range myMap {\n    fmt.Println(key, value)\n}\n```\n\nThis means:\n\n- `key` gets `\"a\"` or `\"b\"`\n- `value` gets the value stored at that key\n\n#### Loop through a string\n\n```go\nfor index, char := range \"Go\" {\n    fmt.Println(index, char)\n}\n```\n\nThe `char` value is a `rune`.\n\n### `break` and `continue`\n\nUse `break` to stop the loop completely:\n\n```go\nfor i := 0; i < 10; i++ {\n    if i == 5 {\n        break\n    }\n    fmt.Println(i)\n}\n```\n\nUse `continue` to skip the current iteration:\n\n```go\nfor i := 0; i < 5; i++ {\n    if i == 2 {\n        continue\n    }\n    fmt.Println(i)\n}\n```\n\n### Common beginner note\n\nThis line:\n\n```go\nfor key, value := range myMap2 {\n    fmt.Println(key, value)\n}\n```\n\nmeans:\n\n- start a loop\n- read each key-value pair from `myMap2`\n- store them in `key` and `value`\n- run the code block for each pair\n\n### Switch\n\n```go\nswitch value {\ncase 1:\n    fmt.Println(\"One\")\ncase 2:\n    fmt.Println(\"Two\")\ndefault:\n    fmt.Println(\"Other\")\n}\n```\n\nSwitch without expression (like multiple if-else):\n```go\nswitch {\ncase age < 18:\n    fmt.Println(\"Minor\")\ncase age >= 18:\n    fmt.Println(\"Adult\")\n}\n```\n\n---\n\n## Error Handling\n\nGo uses explicit error handling with the `error` type:\n\n```go\nresult, err := someFunction()\nif err != nil {\n    fmt.Println(\"Error:\", err)\n    return\n}\n// use result\n```\n\nCreating errors:\n```go\nimport \"errors\"\n\nerr := errors.New(\"something went wrong\")\n```\n\n---\n\n## Useful Shortcuts\n\n| Shortcut | Description |\n|----------|-------------|\n| `:=` | Short variable declaration (inside functions only) |\n| `var()` | Block variable declaration |\n| `const()` | Block constant declaration |\n| `x, y := a, b` | Multiple variable declaration |\n| `val, err := func()` | Capture multiple return values |\n| `len(s)` | Get length of string/array/slice |\n| `fmt.Println()` | Print with newline |\n| `fmt.Printf()` | Formatted print |\n\n---\n\n## Further Reading\n\nFor topics that build on these basics:\n\n- [Go Structs and Interfaces](./golang-structs-and-interfaces.md)\n- [Go File Structure and Scope](./golang-structure-and-scope.md)\n- [Go Goroutines and Mutex](./golang-goroutines-and-mutex.md)\n- [Go Channels](./golang-channels.md)\n- [Go Generics](./golang-generics.md)\n\n---\n\n## Quick Reference\n\n```go\npackage main\n\nimport \"fmt\"\n\nfunc main() {\n    // Short variable declaration (most used!)\n    name := \"Go\"\n    version := 1.21\n    \n    // Multiple declaration\n    x, y := 10, 20\n    \n    fmt.Println(\"Hello,\", name)\n    fmt.Println(\"Version:\", version)\n    fmt.Println(\"Sum:\", x+y)\n}\n```\n\n---\n\n## Best Practices\n\n1. Use `:=` for variable declaration inside functions (most common and concise)\n2. Use `var` for package-level variables or when you need explicit type\n3. Always handle errors explicitly\n4. Use meaningful variable names\n5. Format your code with `go fmt` (automatically!)\n