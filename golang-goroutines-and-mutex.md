# Go Goroutines and Mutex

## What is a Goroutine?

A **goroutine** is a lightweight thread of execution managed by the Go runtime, not directly by the operating system.

While an OS thread can take megabytes of memory and has noticeable creation overhead, a goroutine starts with only a few kilobytes of stack space that grows and shrinks as needed. You can easily run tens of thousands of goroutines concurrently in Go.

To launch a goroutine, simply add the `go` keyword before a function call:

```go
go dbCall(i, data, receptacle)
```

This starts `dbCall(...)` concurrently in the background and immediately moves to the next line of code without waiting for it to complete.

---

## The Problem: Main Thread Exits Early

Consider this simple program:

```go
func main() {
    go fmt.Println("Hello from goroutine!")
    fmt.Println("Hello from main!")
}
```

When you run this, you might only see `"Hello from main!"`.

**Why?**
When the `main` function exits, the Go runtime abruptly terminates the entire program—including all background goroutines, regardless of whether they have finished their work.

---

## Coordinating with `sync.WaitGroup`

To ensure the main function waits for all goroutines to finish, Go provides `sync.WaitGroup` in the `sync` standard library package.

Think of `sync.WaitGroup` like a counter:
1. Add to the counter when you start a task.
2. Decrease the counter when the task finishes.
3. Wait until the counter reaches zero before exiting.

### Core Methods

| Method | Purpose | Where to call |
|---|---|---|
| `wg.Add(n)` | Increments counter by `n` | In the parent function **before** starting the goroutine |
| `wg.Done()` | Decrements counter by 1 | Inside the goroutine when work is finished |
| `wg.Wait()` | Blocks until counter reaches 0 | In the parent function before exiting or reading final results |

### Example from `tutorial_8`

```go
var wg = sync.WaitGroup{}

func exampleOne(data []string, receptacle *[]string) {
    t0 := time.Now()

    for i := 0; i < len(data); i++ {
        wg.Add(1) // "I am starting 1 goroutine that you need to wait for"
        go dbCall(i, data, receptacle)
    }

    wg.Wait() // "Block here until the counter reaches 0"

    fmt.Printf("\nTotal execution time: %v\n", time.Since(t0))
    fmt.Printf("The result of the db call is: %v\n", *receptacle)
}
```

Inside the worker:

```go
func dbCall(i int, data []string, receptacle *[]string) {
    // ... work happens here ...
    wg.Done() // "Hey waitgroup, this goroutine is done"
}
```

> [!TIP]
> A common Go idiom is to use `defer wg.Done()` at the beginning of the goroutine function so it always runs even if the function returns early.

---

## Shared Memory and Data Race Conditions

When multiple goroutines access and modify the same memory variable at the same time without synchronization, you get a **Data Race** (race condition).

### Why this happens with slices

In `tutorial_8`, all goroutines attempt to append data to the same slice:

```go
*receptacle = append(*receptacle, data[i])
```

Without synchronization:
1. Goroutine A reads the current length and pointer of `receptacle`.
2. Goroutine B reads the exact same length and pointer at the same time.
3. Goroutine A writes to index `k` and updates length.
4. Goroutine B overwrites index `k` or writes past capacity, causing corrupted memory, missing items, or crashes.

Go includes a built-in race detector to help find these bugs:

```bash
go run -race .
```

---

## Protecting Shared State with `sync.Mutex`

A **Mutex** stands for **Mutual Exclusion**. It acts like a key to a locked room: only one goroutine can hold the lock at any given moment.

```go
var mutex = sync.Mutex{}
```

### Basic Mutex Pattern

```go
mutex.Lock()   // Acquire the lock (blocks if another goroutine has it)
// Critical section: only 1 goroutine can be here at a time
*receptacle = append(*receptacle, data[i])
mutex.Unlock() // Release the lock for other goroutines
```

If Goroutine A holds the lock, Goroutine B must pause and wait until Goroutine A calls `mutex.Unlock()`.

---

## Critical Pitfall: Serializing Concurrent Code

Look at the difference between `one.go` and `two.go` in `tutorial_8`.

### Bad Implementation (`one.go`)

```go
func dbCall(i int, data []string, receptacle *[]string) {
    delay := 2000
    mutex.Lock() // Locked BEFORE the slow work!
    time.Sleep(time.Duration(delay) * time.Millisecond) // 2-second sleep inside lock!
    *receptacle = append(*receptacle, data[i])
    mutex.Unlock()
    wg.Done()
}
```

**What went wrong?**
Because the 2-second delay is inside `mutex.Lock()`, every goroutine must wait for the previous one to finish sleeping before it can even start its own delay!
- 10 database calls took **~20 seconds** instead of running concurrently in **~2 seconds**.
- Locking too early turned concurrent code right back into sequential code.

### Better Implementation (`two.go`)

Keep the slow, independent work **outside** the lock, and only lock the exact line that touches shared memory:

```go
func dbCall(i int, data []string, receptacle *[]string) {
    delay := 2000
    time.Sleep(time.Duration(delay) * time.Millisecond) // Runs concurrently!

    save(data[i], receptacle) // Only lock during the append
    log(*receptacle)
    wg.Done()
}

func save(result string, receptacle *[]string) {
    mutex.Lock()
    *receptacle = append(*receptacle, result)
    mutex.Unlock()
}
```

Now all 10 calls sleep concurrently, finishing in **~2 seconds total** while still safely appending results!

---

## `sync.RWMutex`: Read-Write Mutex

Standard `sync.Mutex` does not distinguish between reading and writing: even two goroutines that only want to read a slice cannot do so at the same time.

`sync.RWMutex` solves this by offering two types of locks:

1. **Write Lock (`Lock()` / `Unlock()`)**:
   - Exclusive access for writing.
   - Only 1 writer at a time, blocks all readers and other writers.
2. **Read Lock (`RLock()` / `RUnlock()`)**:
   - Shared access for reading.
   - Many goroutines can hold read locks simultaneously, as long as no writer holds the lock.

### Example from `two.go`

```go
var mutex = sync.RWMutex{}

// Writing requires an exclusive write lock:
func save(result string, receptacle *[]string) {
    mutex.Lock()
    *receptacle = append(*receptacle, result)
    mutex.Unlock()
}

// Reading requires a shared read lock:
func log(results []string) {
    mutex.RLock()
    fmt.Printf("\nThe current results are: %v", results)
    mutex.RUnlock()
}
```

Use `sync.RWMutex` when your application has **many readers and few writers**.

---

## Why pass `*[]string` (Pointer to Slice)?

In `tutorial_8`, the function receives `receptacle *[]string`:

```go
func exampleTwo(data []string, receptacle *[]string)
```

In Go, a slice is a header struct consisting of three fields:
- Pointer to backing array
- Length (`len`)
- Capacity (`cap`)

If you pass `[]string` by value, the function receives a copy of this header:
- Calling `append()` updates the **local copy** of the length and capacity.
- If `append()` needs to allocate a new backing array, only the local copy gets the new array pointer.
- The original slice in `main` never sees the newly appended elements!

Passing `*[]string` allows functions to modify the caller's slice header directly.

---

## Summary and Best Practices

1. **Use `go` to launch goroutines** for non-blocking, concurrent execution.
2. **Always coordinate with `sync.WaitGroup`** so `main` does not terminate before background work finishes.
3. **Always call `wg.Add()` before launching the goroutine**, never inside it.
4. **Use `sync.Mutex`** to prevent data races when multiple goroutines read/write shared variables.
5. **Keep critical sections as small as possible**: never place slow operations (I/O, network requests, sleeps) inside a lock unless strictly necessary.
6. **Use `sync.RWMutex`** when you have frequent reads and infrequent writes.
7. **Run `go run -race .`** to test your concurrent code for race conditions.

---

## Related Notes

- [Go Basic Knowledge Guide](./golang-basics.md)
- [Go Channels](./golang-channels.md)
- [Go Generics](./golang-generics.md)
- [Go Structs and Interfaces](./golang-structs-and-interfaces.md)
- [Go File Structure and Scope](./golang-structure-and-scope.md)
