# Go File Structure and Scope

## Why same function names fail in different files

In Go, files inside the **same folder** usually belong to the **same package**.

Example:

```go
// file1.go
package main

func preAllocExample() {
    fmt.Println("from file1")
}
```

```go
// file2.go
package main

func preAllocExample() {
    fmt.Println("from file2")
}
```

This does **not** work if both files are in the same folder and both use:

```go
package main
```

Why:

- Go compiles all `.go` files in that folder together
- They become one package
- Package-level names must be unique
- So `preAllocExample` can only be declared once in that package

---

## Important rule

Same folder + same `package` name = same namespace

That means these names must be unique inside that package:

- function names
- type names
- const names
- var names at package level

You also cannot have 2 `main()` functions in the same package.

---

## Package scope vs local scope

### 1. Package scope

Names declared outside functions belong to the whole package.

```go
package main

var appName = "Go App"

func printName() {
    fmt.Println(appName)
}
```

`appName` can be used by any file in the same package.

### 2. Local scope

Names declared inside a function only exist inside that function block.

```go
func main() {
    age := 20
    fmt.Println(age)
}
```

`age` cannot be used outside `main`.

This means this is invalid:

```go
func main() {
    age := 20
}

func printAge() {
    fmt.Println(age) // error
}
```

---

## When same variable names are allowed

You **can** reuse the same variable name in different local scopes.

```go
func one() {
    name := "A"
    fmt.Println(name)
}

func two() {
    name := "B"
    fmt.Println(name)
}
```

This works because each `name` belongs to a different function scope.

You can also shadow an outer variable:

```go
var name = "global"

func main() {
    name := "local"
    fmt.Println(name) // prints local
}
```

The local `name` hides the package-level `name` inside `main`.

---

## How Go organizes files

Using your project as an example:

```text
golang/
  go.mod
  golang-basics.md
  golang-structure-and-scope.md
  cmd/
    tutorial_1/
      main.go
    tutorial_2/
      main.go
    tutorial_3/
      main.go
    tutorial_4/
      main.go
```

### What `go.mod` does

`go.mod` defines your module:

```go
module github.com/fY-Hr/go_tutorial
```

Think of the module as the top-level project identity.

### What `package main` means

`package main` means that package builds into an executable program.

A file like this:

```go
package main

func main() {
}
```

is meant to run as an app.

### Why `tutorial_1` and `tutorial_4` do not clash

Even if both folders contain:

```go
package main
func main() {}
```

that is okay because they are in **different folders**, so they are different packages.

In practice:

- `cmd/tutorial_1` is one package
- `cmd/tutorial_4` is another package

The package name text may be the same (`main`), but the folder is different, so the package is different.

---

## How to organize multiple files in one folder

If you want multiple files in `tutorial_4`, that is fine.

Example:

```text
tutorial_4/
  main.go
  slices.go
  maps.go
```

As long as all files say:

```go
package main
```

they can call each other directly.

Example:

```go
// main.go
package main

func main() {
    sliceExample()
    mapExample()
}
```

```go
// slices.go
package main

func sliceExample() {
}
```

```go
// maps.go
package main

func mapExample() {
}
```

This works well because all function names are unique.

---

## How to avoid duplicate names

Use clear, specific names:

```go
func sliceExample() {}
func mapExample() {}
func preAllocExample() {}
func loopExample() {}
```

This is better than trying to reuse one generic name in many files.

---

## Exported vs unexported names

In Go, capitalization matters.

### Unexported

Starts with lowercase:

```go
func helper() {}
```

Usually used only inside the same package.

### Exported

Starts with uppercase:

```go
func Helper() {}
```

Can be accessed from other packages.

This is Go's visibility rule.

---

## Quick mental model

Think of it like this:

- folder = package
- package = shared namespace
- file = just a piece of that package
- function-local variables live only inside their block

So if two files are in the same folder, Go treats them like they are parts of one bigger code file.

---

## Practical advice for your learning project

- Keep one `main.go` in each tutorial folder
- Split helper functions into other files only when the topic gets bigger
- Use unique names like `mapExample`, `sliceExample`, and `scopeExample`
- Put related files in the same folder only if they belong to the same package
- Create a new folder when you want a separate runnable example

---

## Summary

- Same folder and same package means one namespace
- That is why duplicate function names fail
- Package-level names must be unique
- Local variables can reuse the same name in different functions
- Different folders can use the same function names because they are different packages
