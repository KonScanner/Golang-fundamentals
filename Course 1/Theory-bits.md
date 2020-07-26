# Pointers

- A pointer is an adress to data in memory.
  - `&` operator returns the adress of a variable/function.
  - `*` operator returns data at an address (dereferencing)

# New

- Alternative way to create a variable
  - `new()` function creates a variable and returns a pointer to the variable.
  - Variable is initialized to zero by default.
    ```
    ptr := new(int)
    *ptr = 3 // this is placed at the adress specified by ptr
    ```

# Variable scope

- The places in code where a variable can be accessed.
  - can either be `local` or `global`.

### Blocks

- A sequence of of declarations and statements within matching brackets, `{}`
  - Including function definitions
  - Hierarchy of **implicit blocks**
  - Universe block - all Go source
  - Package block - all source in a package
  - File block - all source in a file
  - `if`, `for`, `switch` - all code inside the statement
  - Clause in `switch` or `select` - individual clauses each get a block

### Lexical Scoping

- Go is a lexically scoped language using blocks

  - `b_i` >= `b_j` if `b_j` is defined inside `b_i`.

    - "defined inside" is transitive

    ```
    var x = 4 // b1

    func f() {
        fmt.Printf("%d", x) // b2
    }

    func g(){
        fmt.Printf("%d", x) // b3
    }
    ```

- Variables can be accessible from block `b_j` if :
  - Variable is declared in block `b_i`, and
  - `b_i` >= `b_j`

# Deallocating Space

- When a variable is no longer needed, it should be **deallocated**
  - Memory space is made available for other purposes
- Otherwise, we will eventually run out of memory
  - e.g. Each time you call `f()` it creates an integer `x` in memory
    ```
    func f(){
        var x = 4
        fmt.Printf("%d", x)
    }
    ```

### Stack vs. Heap

- `stack` is a memory area allocated to function calls ("localy allocated variables and such")
  - They are deallocated after function completes
- `heap` is a persistent region of memory, which does not go away just because the function that allocated it has completed. It has to be **explicitly** deallocated.

#### Manual Deallocation

- Data on the `heap` must be deallocated when it is done being used.
  - In most compiled languages (i.e. `C`), this is done manually.
  ```
  x = malloc(32); // allocates 32 bytes of memory to pointer x
  free(x); // to free space
  ```
  - Error-prone, but fast

### Garbage Collection

- In interpreted languages, this is done by **the interpreter**
  - Java Virtual Machine
  - Python Interpreter
- This makes it easy for the proggramer
- It is slower, since it requries the **interpreter**

#### Garbage Collection in `Go`

- `Go` is a compiled language which enables garbage collection
- Implementation is fast
  - Compiler determines `stack` vs `heap`
  - It will garbage collect appropriatelly in the background

#### Pointers and Deallocation

- Hard to determine when a variable is no longer in use

  - This example is legal in `go` but not in other languages

    ```
    func foo() *int {
        x := 1 // local variable
        return &x // address is being returned
    }

    func main(){
        var y *int
        y = foo() // returns pointer to x
        fmt.Printf("%d", *y)
    }
    ```

# Type Conversions

- Most binary operations need operands of the same type
- `int`s

  - Including assigmnents

  ```
  var x int32 = 1
  var y int16 = 2
  x = y
  ```

  - This will fail, as they are different types. You need to first re-assign with a new type.

    ```
    var x int32 = 1
    var y int16 = 2
    x = int32(y)
    ```

- `float`s
  - `float32` - ~6 digits of precision
  - `float64` - ~15 digits of precision
  ```
  var x float64 = 123.45
  var y float64 = 1.2345e2
  ```
- `complex` number `float`s

  ```
  var z = complex128 = complex(2,3)
  var w = complex256 = complex(3,2)
  ```

- `string`s

  - Sequences of arbitrary bytes

    - Read-only
    - Often meant to be printed

  - String literal - notated by double quotes
    - `x := "Hello there..."`
  - Each byte is a rune (_UTF-8_ code point)

- `Constant`s
  - Expression whose value is known at compile time
  - Type is inferred from right hand side (boolean, string, number)
  ```
  const x = 1.3
  const (
      y = 4
      z = "Hi"
  )
  ```
  - `iota`
    - Generates a aset of related but distinct constants
    - Often represents a property which has several distinct possible values
      - Days of the week
      - Months of the year
    - Constants must be different but **actual value is not important**
    - Like enumerated type in other languages like `C`
    ```
    type Grades int
    const (
        A Grades = iota
        B
        C
        D
        F
    )
    ```
    - Each constant is assigned to a unique integer
    - Starts at 1 and increments

#### ASCII and UNICODE

- Strings are _often_ made for printing
- _ASCII_
  - American Standard Code for Information Exchange
  - Character coding - each character is associated with an (7) 8-bit number
    - 'A' = `0x41`
- `Unicode`
  - There are multiple kinds of `unicode` representations, in this case we will consider `unicode32`
  - `UTF-8` is a variable length
    - 8-bit UTF codes are the same as ASCII
  - **DEFAULT IN `GO`** is _UTF-8_

#### UNICODE Package

- Runes are divided into many different categories
- Provides a set of functions to test categories of runes (`true`/`false`)
  - `IsDigit(r rune)`
  - `IsSpace(r rune)`
  - `IsLetter(r rune)`
  - `IsLower(r rune)`
  - `IsPunct(r rune)`
- Some functions perform conversions
  - `ToUpper(r rune)`
  - `ToLower(r rune)`

#### Strings Package

- Functions to manipulate encoded strings
- String search functions:
  - `Compare(a, b)` - returns an integer comparing two strings, lexicographically. 0 if `a` == `b`, -1 if `a` < `b`, and +1 if `a` > `b`.
  - `Contains(s, substr)` - returns true if substring is inside s
  - `HasPrefix(s, prefix)` - returns true if the string s begins with prefix
  - `Index(s, substr)` - returns the index of the first instance of substr in s
- String manipulation, strings are immutable, but modified strings are returned
  - `Replace(s, old, new, n)` - replace returns a copy of the string `s` with the first `n` instances of old replaced by new
  - `ToLower(s)` and `ToUpper(s)`
  - `TrimSpace(s)` - returns a new string with all leading and trailing whitespace removed
- `Strconv` Package
  - Converstions to and from string representations of basic data types
  - `Atoi(s)` - converts string `s` to `int`
  - `Itoa(n)` - converts int `n` to `s` sting
  - `FormatFloat (f, fmt, prec, bitSize)` - converts floating point number to string
  - `ParseFloat (s, bitSize)` - Converts a string to a floating point number

## Control Flow

- Describes the order in whihc statements are executed inside a program

  - Statements which alter control flow
  - If statements
    ```
    if <condition> {
        <consequent>
    }
    ```
    - Expression `<condition>` is evaluated
    - `<consequent>` statements are executed if condition is `true`
    ```
    if x > 3 {
        fmt.Printf("x is in fact greater than 3")
    }
    ```
  - For Loops
    - Iterates while a condition is `true`
    - May have an initialization and update operation
    ```
    for <init>; <cond>;
    <update>{
        <stmts>
    }
    ```
    - `<init>` executed the first time
    - `<cond>` checked on each iteration
    - `<update>` what is executed at the end of each iteration and updates some element of a state(s)
    - Most common forms of `for` loops
    ```
    for i:=0; i<10; i++ {
        fmt.Printf("hi ")
    }
    ```
    ```
    i = 0
    for i <10{
        fmt.Printf("hi ")
        i++
    }
    ```
    - Infinite loop
    ```
    for {
        fmt.Printf("hi ")
    }
    ```
  - `Switch`/`Case`
    - `Switch` is a multi-way if-statement
    - `Switch` typically contains a `tag` which is a variable to be checked
    - Tag is compared to a constant defined in each `case`
    - Whichever `case` which matches the `tag` gets executed
    ```
    switch x {
    case 1:
        fmt.Printf("case1")
    case 2:
        fmt.Printf("case2")
    default:
        fmt.Printf("nocase")
    }
    ```
  - Control Flow Scan
    - Tagless Switch
      - Case contains boolean expression to evaluate
      - First `true` case is executed
    ```
    switch {
        case x > 1:
            fmt.Printf("case1")
        case x < -1:
            fmt.Printf("case2")
        default:
            fmt.Printf("nocase")
    }
    ```
  - `Break` and `Continue`
    - `Break` exits the containing loop
    ```
    i := 0
    for i <10 {
        i++
        if i == 5 { break }
        fmt.Printf("hi ")
    }
    ```
  - `Continue` skips the rest of the current iteration
    ```
    i := 0
    for i < 10 {
        i++
        if i == 5 {continue} // jumps past that iteration
        fmt.Printf("hi ")
    }
    ```
  - `Scam`

    - reads user input
    - Takes a pointer as an argument, typed data is written to that pointer
    - Returns number of scanned items

    ```
    var appleNum int

    fmt.Printf("Number of apples?")
    num, err :=
    fmt.Scan(&appleNum)
    fmt.Printf(appleNum)
    ```

# Composite Data types

## Arrays

- Fixed-length series of elements of a chosen type
- Each element is accessible using subscript `[ index ]`

  - Indices start at `0`
  - Elements initializer to zero value
  - Example decleration:

  ```
  var x [5]int

  x[0] = 2
  fmt.Printf("x[1]") // prints 0
  ```

### Array Literal

- An array pre-defined with values:

```
var x [5]int = [5]{1, 2, 3, 4, 5}
```

- Length of literal must be length of array
- `[...]` as the `size` in array literal infers `size` from number of initializers
  ```
  x := [...]int{1, 2, 3, 4}
  ```
- Iterating through Arrays

  - Use a for loop with range keyword

    ```
    x := [3]int {1, 2, 3}

    for i, v := range x {
        fmt.Printf("ind %d, val %d", i, v)
    }
    ```

## Slices

- A "window" on an **underlying array**
- Variable size, up to the whole array
- _Pointer_ indicates the start of the slice
- _Length_ `len()` is the number of elements in the slice
- _Capacity_ `cap()` the maximum number of elements
  - From start (pointer) of slice to end of slice.

```
  a1 := [3]string("a","b","c")
  s1 := a1[0:1]
  fmt.Printf(len(s1), cap(s1))
```

### Silce examples

```
arr := [...]string{"a","b","c","d","e","f","g"}

s1 := arr[1:3]
s2 := arr[2:5]

fmt.Println(s1) // [b c]
fmt.Println(s2) // [c d e]
```

## Accessing Slices

- When writing to a slice, you are writing to the underlying array
- Overlapping slices refer to the same elements
  (see last example above, `"c"` is on both)

### Slice Literals

- Can be used to initialize a slice
- Creates the underlying array and creates a slice to reference it
- Slice points to the start of the array, length is capacity

```
sli := []int{1, 2, 3}
```

## Make

- Create a slice (and array) using `make()`
- 2-argument version, need to specify the length/capacity
- Init. to zero, length == capacity

```
sli = make([]int, 10)
```

- 3-argument version, need to specify the length and capacity **separately**

```
sli = make([]int, 10, 15)
```

## Append

- Size of a slice can be increased by `append()`
- Adds elements to the end of a slice
  - Increases the slice, _up to the capacity_ of the **underlying array**
  - But can **increase** the size of the array if necessary, it will create a bigger array and continue appending, which will continue to increase the size of the array.
    - There is a time penalty associated with that, but it is doable.
  ```
  sli := make([]int, 0, 3) // len is 0
  sli = append(sli, 100) // len is 1
  ```

## Hash tables

- Containes key/value pairs (key must be unique)
  - Social Security Number/email
  - GPS cords/address
- Each value is asssociated with a unique key
- **Hash function** is used to compute the slot for a key

### Tradeoffs of Hash Tables

Advantages:

- Faster lookup than lsits
  - Constant-time, vs linear-time
- Arbitrary keys
  - Not ints, like slices or arrays

Disadvantages:

- May have collisions
  - Two keys hash to the same slot
  - Collisions are rare.
