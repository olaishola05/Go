# Package

Packages are go ways of organizing code

- Packages sholud be a single program or single-purpse library
- Packages can be published and imported from Go library.

- Importing all things from a package use .
- to rename a package import

```
import newName "namespace/packageName"
```

- Grouping Package imports

```
import {
    . "name"
    pk "namespace/packageName"

}
```

## Modules

- Modules are collection of packages and can be find in the root path of the project
  - Can be manage by Go CLI
- Contains informations about the project dependencies, Go version and package info
- Every Go projects has a go.mod file

## Data Types

All data in programs consist of binary numbers and data types is a way that the program interpretes the binary numbers.

Go is a statically typed language which mmeans

- The data types must be provided by the programmer
- Go uses type inference to determine what type of data it is working with.
- Which means data types only need to provided in specific circumstances
- Can always specified the type if desired
- Compiler error is wrong type is used for type safety

- Primitive Data Types

In Go all primitive data types in Go are numeric which means stream of bytes

- Signed Interger Types

- Is a number that is integer which can either be positive or negative
- int8 min: -128, Max: 127
- int16 min: -32768, Max: 32768
- int
- int32
- int64

- Unsigned integers

- Is an integer with positive integer starting from zero

- Other Data Types
- float32: 32-bit floating point
- float64: 64-bit floating point
- complex64
- complex128
- boolean

### Type Aliases

Used for creating or assigning new name to type aliases data. For providing indications of data being used.

- type UserId int
- type Direction byte
- type Speed float64

You can have type aliases upon aliases

- type Velocity Speed

- Converting between types using parentheses

- UserId(5)
- Speed(88.3)

### Strings & Runes

Text data in Go uses UTF-8 encoding

- Encoding is a way to represent thousands of different symbols using code pages

- In Go when working with strings, you work with bytes and not the letter itselves.
- Text is represented using rune type similar to char in other languages
- Rune is an alias for **int32** bit integer.
- You are actually working with numbers unless formatted properly
- Rune can represent - Letters - numbers - Emoji, etc

      Rune byte are represented in column of 32bytes divided into four columns of 8bytes

  String
  A string is the data type for storing multiple runes

- It is just an array of bytes and a string length
- No null termination with a Go string

- When Iterating over a string, the iteration occurs over bytes
- Bytes are not symbols and special iteration is required to retrive rune/symbols

## Creating Runes

- It can be created using single quote or backtick (literal)
  'a', `R`

To create string we use double quotes
Strings: "Some random $22\n"
Raw Literal: `Let's code in "Golang!"\n`. Used when you with escpe xharters

## Go CLI Tool

- To update dependencies
- Build & test projects
- Manage artifacts
- Format source code

- build will builds the project & emits an executable binary for coding sharing
  - d-race: checks for concurrency problems
- run: runs the project directly, not executable output
- mod: manages modules & dependencies
- mod tidy: updates dependencies
- test: runs the project test suite
  -fmt: formats all source files

## Variables

Variable provide a way to store & access data in progam - Data within can be anything - Alias to data in memory - String data to variable is called assignment

Variables have multiple components:

- Name
- Data
-

### Variable Creation

var example = 3
var example int = 3
var exmaple int

example = 3

For example if a variable with type int is declared without assignment, it will automatically has the default value which is 0 for integer

- Compound Creation

`var a, b, c = 1, 2, "sample"`

- Block Creation

```
var (
  a int = 1
  b int = 2
  c = "sample"
)
```

- Create & Assign

`example := 3`

`a, b, := 1, "sample"`

- This is mostly used in Go programming
- `example := 3`
- a = 3
- b := a Indicate b is a copy of a

Variables can be reassigned & assigned to other variables

- Variable names can only be used once per scope

### Defaults

Variables that are declared but not assigned will automatically have a defaukt value

String: ""
Number: 0
Other: nil

Comma Ok

Comma Ok idiom allows a variable to be reassigned in a creation statement

- Allows to reuse the second variables are frequently used when checking for errors
  `a, b := 1, 2`
  `c, b =: 3, 4`

`data, err := `

## Naming Convention

- Go variable naming convention is camelCase:
  e.g
- myVariableName := "hello"
- Use names appropriate for the data
- `totalGuests := 12`

- Constants
- For declaring some values that needs to be utilized throughout some or all the program and never changes or reassigned.
- Convention, starts with capital letter

```

const MaxSpeed = 30
const MinPurchasePrice = 7.50
const AppAuthor = "Bob"

```


## Functions

They are basic building block of Go programs
 - Allows functionality to be isolated which makes programs easier to
  - Test
  - debug
  - extend
  - modify
  - read
  - write
  - document

Functions are simple, they take data as input and return data as output. the input & output are optional


```
func name(param1 type, param2 type) type {
  <!-- Body -->
}
```

Example

```
func sum(num1, num2 int) int {
  return num1 + num2
}

```

Note: If the type of the input data are same, it can be added as written above in the function

Calling Function with arguments

`result := sum(2, 2)`


Return Multiple Values

```
func multiReturn() (int, int, int) {
  return 1, 2, 3
}
```

Calling Func

`a, b, _ := multiReturn()`

The _  means to ignore the last item from the return values.

## Operators

Operators are used to perform calculations and comparisons

  -  less than <
  - Subtraction -
  Operators works with operands

### Arithmetic Operators

 - + and +=
 - - and -=
 - * and *=
 - / and /=
 - % and %=  //remainder or modulus


 examples

 total = 3 + 3

 Arithmetic Assignment

 a := 1
 a += 3 // a = a + 3

 Increment & Decrement
 i++, i--

 - Arithmetic opertaors always return a number

### Relational Operators

- less than <
- Less than or Equal To <=
- Greater than >
- Greater than >=
- Equal To ==
- Not Equal To !=

- Relational operators always return a boolean (true/false)


### Logic Operators

- And or && both must be true
- Or or || one of them must be true
- Not or !

- Logic Operators operators always return a boolean (true/false)

## If/else

**Flow Control**

- program code executes line-by-line and flow control is a way to interrupt this process
 - Different lines of code can be executed based on conditions and conditions are defined by programmer.

 ```
 if condition {
  <!-- Run this code -->
 } else {
  <!-- Do something here -->
 }
 ```

 Example 

 ```bash

 if userName != "" {
  return "welcome", userName
 } else {
  # display error message
 }

 ```

 if..else if

 Logical Operators with if statement


 ```bash

 if hasTicket && ticketValid {
  # permit entry
 }

 if funds > cost && accountActive {
  # allow transaction
 }

 if usingDebitCard || usingCreditCard {
  # display card and entry screen
 }

 if !quizComplete {
  # display incomplete submission and warning
 }
 ```

 - Go has a feature called Statement Initialization that allows to create variables and perform comparisons at the same time 


Example

```bash

if i := 5; i < 10 { 
  do something with i 
  } else {
  do something else with i
  }

```

```bash

if userRole := getUserRole(); userRole == "admin" {
  do admin stuff here
} else if userRole == "manager" {
  do manager things here
} else {
  display error message for permission level
}

```

Early return
 - Use to stop processing within function as soon as possible
 - Maximizes performance


## Switch Statement

 - Used to easily check multiple statement
 - Alternative to using many if...else if blocks
 - It execute from top to bottom approach
 - Optionally have a **default** action. When no case pass it reult to default value
 - Expressions are allowed as a case. func calls, math and logic

  ### Conditional Cases

  - It uses case and assign


  ```bash

  switch result := calculate(5); {
    case result > 10:
      fmt.Println("10")
    case result == 6:
      fmt.Println("==6")
    case result < 10:
      fmt.Println("<10")
  }
  ```

  **Case List**

  - Allow us to check for testing multiple cases

```bash
  switch x {
    case 1, 2, 3:
      <!-- Do something here -->
    case 10, 20, 30
      <!-- Another thing here -->
  }
```

It use case: allow us to take multiple actions or check on a single case

**Fallthrough**

- To continue checking the next case even it does'nt match after a case met the condition


## Loop
 - For loop is the only key word in Go
 - intialization, condition, post statement
 - When using for while loop, the initial variable must declare first else with result in infinite loop
 - for infinite loop mostly use for servers
  - uses the break key word to exit the loop
 - The continue keyword to skip the current loop


## Structures
 - Structures allow data to be stored in groups, similar to a class in other programming language
 - Each data point in a structure is a field
 - It helps organize code and data more efficiently

Example

```bash
type Sample struct {
  field string
  a,b int
}
```

### Instantiating a Structure
  - The fields must be in order as defined in the structure when not naming the fields. see first example

```bash
data := Sample{"word", 1, 2}

data := Sample{
  field: "word",
  a: 1,
  b: 2,
}
```

### Structure Default Values

- Just like variables, any field not indicated in structure instantiation will have a default values

`data := Sample{}`

`data := Sammple{a: 5}`

- Accessing The Fields

- Fields can be access using the dot notation using the variable name and the field name

`word := data.field`
`a, b := data.a, data.b`

- Field can also be modified using the dot notation to access the field and reassigned the value

```bash

data.field = "Hello"
data.a = 500
data.b = 200

```

### Anonymous Structures

- It is possible to create anonymous/inline structures inside a func
- Useful when working with library functions or shipping data accross the network
- Can be use to define data strcuture as needed
- created with keyword var
- can be accessed using the dot notation
- has default values like the type version


```bash 
var sample struct {
  field string
  a,b int
}
```

**Create Assign**

- Shorthand version must have each field defined and cannot have default values

```bash
sample := struct {
  field string
  a,b int
}{
  "hello",
  1, 2,
}
```

## Arrays
  -  Ways to store multiple pieces of the same kind of data
  - Each piece of data is called an element
  - index is used to access data from an array
  - index starts from 0
  - Array are fixed-size and cannot be resized
  - Element noty addressed in array initialization will be set to default values

- Creating an Array

- Uninitialized 

```bash
var myArray [3]int
```

- Create & Assign

```bash
myArray := [3]int{7, 8, 9}

- Defualt Way or the Three dotted

```bash

myArray := [...]int{7, 8, 9}

```

If an array is set to certain size but the elements is not upto te size, the rest is set to zero

```bash

myArray := [4]int{7, 8, 9}
```
The fourth element will be zero

- Accessing Array elements
  - Array can be accessed and updated using the bracket notation

```bash

var myArray [3]int

myArray[0] = 7
myArray[1] = 8
myArray[2] = 9

//Reading
item1 := myArray[0]

```
### Array Iteration
  - It is recommended as good practice to assign the element to a variable during iteration. See below for example

```bash

myArray := [...]int{7, 8, 9}

for i := 0; i < len(myArray); i++{
  item := myArray[i]
  fmt.Println(item)
}
```

### Bounds
 - When you attempt to access an element outside of the bounds of an arrary will result in an error
 - Always use the **len** function or method when looping an array to avoid run time error

## Slices
 -  Slices are companion types that works with arrays
 - They enable a **view** into an array
  - Views are dynamic and not fixed in size like an array
- Functions can accept a slice as a function params
  - Any size array can be operated upon via slice
- When an array is sliced, the element on the left are chopped away
- Slice takes minimum amount of memory and only points to the original array
- Slice those not make a copy of an array but have a metadata that points to the original array

### Creating a Slice
 - Slices & underlying array can be created at the same time
 - The difference btw array and slice is that slice does not have anything inbtw the arra bracket during creation.
 - Accessing elements in a slice is the same as an array

```bash
mySlice := []int{1,2,3}

item1 := mySlice[0]
```

### Slice Syntax

- The slice syntax can create slices from specific elements in an array or other slice

**slice[a:b]**

- The **a** is the start index where the slice to begin, if omitted the slice starts from 0
- The **b** indicate the end index, where the slice is to end, if omitted the it slcied to the end. The b is exclusive, it doesnt get added but the a is inclusive.


```bash

numbers := [...]{1, 2, 3, 4}

slice1 := numbers[:] // slice everything from the array [1, 2, 3, 4]
slice2 := numbers[1:] //start from 1 and ignore 0 and to the end [2, 3, 4, 5]
slice3 := slice2[:1] //start from zero to 1 but 1 is exclusive [2]
slice4 := numbers[:2] // start from index zero to two but item on index two is exclusive [1, 2]

slice5 := numbers[1:3] // [2, 3]
```
- If slice sliced past the length of an array you get an error.

### Dynamic Array & Appending slices

Slices can be used to create arrays that can be extended (Dynamic arrays)
  - The append() function can be used to add additional elements
  - Note: The original slice must be reassign to the the append slice

e.g

```bash
numbers := []int{1, 2, 3}
numbers = append(numbers, 4, 5, 6) // [1, 2, 3, 4, 5, 6]
```

- 3 dots can be used to extend a slice with anther slice

```bash

part1 := []int{1, 2, 3}
part2 := []int{4, 5, 6}
combined := append(part1, part2...)
```
### Preallocation

In a situation where you no ahead of time how many item you need, slices can be preallocated with specific capacities.
- It saves computational time.
 - The make() function is used to preallocate a slice
- Useful when number of elements is known, but their values are still unknown

`slice := make([]int, 10)`

### Slices to Function

- funcs parameters which requires a slice can work with slices of any size

```
func iterate(slice []int) {
  for i := 0; i < len(slice); 1++ {
    // code here
  }
}

small := []int{1}
big := []int{1, 2, 3, 4, 5, 6, 7}

iterate(small)
iterate(big)

```

### Multidimensional Slices

- It uses two empty [][], type is optional, the first bracket represent row numbers and the second column numbers
- array bracket notation is used to access element in multidimensional array.

board := [][] string {
  []string{rows},

}

## Ranges

- It create an iterator automatic
- It as an easier way of creating for loop with the counters
- When range keyword is used, you get the i= index and the element back
- It makes it easier to further loop the return element
- It uses comma instead of ; to separate syntax and take care of the need for using the **len** function

## Map
 - Maps are commonly used data structure that stores data in key-value pairs
 - Extremely high performance when the key is known
 - Unordered: data is stored in random order
 - The key must be in double quoted.
 - The size of a map is dynamic
 - If you try read key that is not present in a map, it retunrs the default values of map type

 ### Map creation

 Empty map

 `myMap := make(map[string]int)`

 Non Empty

 ```bash 
 myMap := map[string]int{
  "item 1": 1,
  "item 2": 2,
  "item 3": 3,
 }
 ```

 ### Map Opertaions

 - Insert: `myMap["fav number"] = 5`
 - Read: `fav := myMap["fav number"]`
 - Delete: `delete(myMap, "fav number")`
 - Check Existence

```bash
price, found := myMap["price"]
  if !found {
    fmt.Println("price not found")
    return
  }
  ```

### Map Iteration
- It returns key, values

for key, value := range myMap{
  // keys, values
}

## Pointers
 - Pointers are variables that **point to** memory address
 - The values of the variable itself is a memory address
  - Accessing the data requires dereferencing the pointer
  - Derefrence provide access to data that is been pointed to.

### Creating Pointers

  - When an Asterisk (*) is used with a type, it indicates the value is a pointer.

  ```bash
  value := 10
  var valuePtr *int
  valuePtr = &value
  ```
  - With an Ampersand(&) creates a pointer from a variable or when you want to create pointer from something that already exist.

   ```bash
  value := 10
  valuePtr := &value
  ```

  When we are working with struct data type, we don't need to add the  asterisk to the variable or dereference the varaible because struct dot notation does that automatically