# Go Notes 
<!-- rename language guide and then make a cheat_sheet?   -->
---
---
---
---

## Table of Contents

[Resources](#resources)

[Lexicon](#lexicon)

[My Questions](#my-questions)

## Fundamentals
[Command Line](#command-line)

[**Types**](#types)
- [Slice](#slice)
- [Struct](#struct)
- [Interface](#interface)

[Variables](#variables)

[Operators](#operators)

[Functions](#functions)

[Loops](#loops)

[Routines](#routines)

[Channels](#channels)

[Errors](#errors)

[Scope](#scope)

[Files](#files)


## Data Structures

## Misc
[Language Quirks](#language-quirks)

[Tools](#tools)

[Testing](#testing)

[General Learning](#general-learning)

[Misc Notes](#misc-notes)


---
---
## Resources
[Run code snippets in your browser](https://play.golang.org/)

[Packages](https://golang.org/pkg/)

## Lexicon

Word | Definition 
-----|-------------
***Method*** | a function that belongs to an instance
***Argument*** | A value passed to a function
***Blocking Call*** | A function call that takes some time to execute and blocks the progression of the program
***Function Literal*** | an unnamed function, a block of code we can pass around and use at sometime in the future
***Concrete Type*** | a type that we can create a value from, access it, change it, and create copies of it (maps, structs, ints, strings, custom types etc.)
***Interface Type*** | We can't create a value directly out of this type
***Type conversion*** | take one type and turn it into another
***String Manipulation*** | processing strings
***Type Description*** | tells what type a value is
***Reference Type*** | A type which references another data structure.  Therefore in go we don't need to worry about pointers.  (slices, maps, channels, pointers, functions)
***Value Types*** | A type which references the value inside memory, we need to use pointers to change these types inside a function (int, float, string, bool, structs)
***Pass By Value Language*** |  A language where whenever we pass data as a reciever or as an argument that value is copied in memory, so the function defaults to working on a copy of the original data
***Referencing*** | a pointer, a value that directs the computer to the actual location of some data in memory
***Dereferencing*** | Accessing the data at a certain memory location.  Turning a pointer into the data it points at
***Pointer*** | A value that refers to another value stored elsewhere in the computer
***Functions*** | a named procedure that performs a distinct service
***Operator*** | The basic functions of a programming language

# My Questions

---

# Fundamentals

## Command Line 
`$ go run main.go // run a go package`

`$ go build main.go` // saves an executable file of your code in the current directory

`$ ./main` // run an executable file

## VSCode
- command click a go method name to open the source code

## Types

Type | Zero Value
---|---
string | ""
int | 0
float | 0
bool | false

``` go

// create a new type of 'deck' which is a slice of strings
type deck []string
```

### Array
A fixed length list of things, must be defined with the data type it contains

``` go
    var myArray [6]int

    a[0] = 14
    
    var thisA [4]int{1,2,3,4}
```

### Slice
Slice: an array that can grow and shrink, must be defined with the data type it contains
- Does this return a new slice or modify the original?
- Slices are 'zero indexed'
- Reference Type
- Byte Slice: a slice of ascii values.  aka a string of characters
- A slice is composed of an array and a structure which contains the length of the slice, the capacity of the slice, and the underlying array

``` go
    sliceName := []string{"First", "Second"}
    sliceName[1] // return second element
    sliceName[0:2] // sliceName[startIndexIncluding:upToNotIncluding]
    sliceName[:2] // sliceName[:upToNotIncluding]
    sliceName[2:] // sliceName[:includingTilEndOfSlice]

    bs := make([]byte, 99999) // make a byte slice with 99999 empty elements inside of it
```

### Struct
Struct: A data structure in go, a collection of different properties that are related together

- Structs have 'fields'
- Fields must be defined when a truct is declared and cannot be added or removed later
- What fields does this struct have?
- What methods do structs of this type have? 

``` go
// declare a struct with two fields/properties
type person struct {
    // property name, then type
    firstName string
    lastName string
    contactInfo // declares type and field name to be the same
}

type contactInfo struct {
    email string
    zipCode int
}

// declare, initialize, and assigning a variable
frank := person{firstName: "Frank", lastName: "Fredrickson"}
// create a variable of type person with zero-valued properties
var matilda person
matilda.firstName = "Matilda"
anders := person{"Anders", "Anderson"}

bilbo := person{
    firstName: "Bilbo",
    lastName: "Baggins",
    contactInfo: contactInfo{
        email: "notGolumn@myprecious.com",
        zipcode: 11111,
    },
}
```

### Interface
Interface: Used to define a function set, a group of functions bound together by an interface (this helps us reuse code)

- Which types satisfy which interfaces?
- a type 'satisfies an interface' when it has a method that has the same arguments and return values as declared in an interface
- interfaces are satisfied implicitly (i.e. you don't have to write code that explicitly links functions to an interface)
- If an interface is used as the type in a struct then you can put any type into that field as long as it fulfills the requirements of the interface
- Reader Interface - an interface defined in go that transforms different data types that might be input into our application into a common output type (byte slice) so they can be passed around and worked with.  Reader always returns a byte slice so we can write many different functions that take a byte slice and anything that implements a reader can use those functions. (every source of data into our app will implement the reader interface)
- Write Interface:  Used to take some information from inside our program and send that data outside our program

``` go
// a custom type 'bot' interfacing between other types, any types with a function 'getGreeting' that return a string can call functions that use type 'bot'
type bot interface {
    // we describe the function name and the arguments and or return types
    // to 'satisfy' the 'bot' interface the type must implement a function called   getGreeting that takes no arguments or receivers and returns a string
    getGreeting() string
    Reader // syntax for inserting an interface into an interface
}

type englishBot struct{}
type spanishBot struct{}

func main() {
    eb := englishBot{}
    sb := spanishBot{}

    printGreeting(eb)
    printGreeting(sb)
}

func printGreeting(b bot) {
    fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
    return "Hello"
}

func (spanishBot) getGreeting() string {
    return "Hola"
}

```


### Maps
* A collection of key value pairs
* All keys must be of one type, all values must be of one type
* All keys are indexed (we can itterate over them)
* You can add and remove keys at your leisure

``` go
// declare a map called colors where all the keys are of type string 
// and all the values are of type string
    colors := map[string]string{
        "red": "#ff0000",
        "green": "f77df3",
    }

// declare a zero valued map called colors
    var colors map[string]string
    colors := make(map[string]string)

// square brace syntax for access to maps
    colors["white"] = "#ffffff"
    delete(colors, "white")

// c is the 'argument name'
    func printMap(c map[string]string) {
    }
```

## Variables
- Variables can be initialized outside a function but not assigned a value

``` go
    var varName string = "Hello friend"
    var varName := "Hello neighbor"

```

## Operators
'*' operator called on a pointer which gives me the value this memory address is pointing at 

'*string' infront of a type is a type description telling you that it is a pointer to a string

'&' operator: gives the memory address of the value a variable is pointing at


## Functions
Receiver: the variable name and type that a function can be called on, every variable of this type will gain access to this function.  Identify using a one or two letter abbreviation
Function Signature: The declarative line of a function
- Associated: a function is 'associated' with a type when it has that type as a receiver
- Functions can have interfaces as the type of the arguments (aka anything that is passed to it has to extend from that interface)
- Whenever we pass a value to a function the function receives a copy of that variable


``` go
// a reciever t of type typeName 
func (t typeName) methodName(v1 typeName, size int) (returnType, returnType2) {
    return 1, 2
}
x := typeName
x.methodName()

firstVal, secondVal := calledOn.methodName(inputOfTypeA, 6)

func newCard() string {
    return "Five of Diamonds"
}

// giving reciever a variable name is optional if you are not using it inside the function
func (englishBot) getGreeting() string {
	return "Hi there"
}

// a 'function literal'
// the extra set of paranthesis at the end calls the function
func() {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}()

```

## Loops

- if you declare the index or the value they MUST be used in the return statement

``` go
    for index, value := range itterateOverVar {
        fmt.Println(i, card)
    }

    // A range with a channel
    // It waits for the channel to return a value, assigns it to the variable 'l'
    // and then runs the body of the loop
	for l := range c {
		go checkLink(l, c)
	}
```

## Routines
- structures inside of go that are used to handle concurrent programming
- 'Main routine' executes all the code in your program, when you define additional routines they are referred to as 'child' routines'.  The termination of the Main routine decides when program quits.
- 'go' keyword used infront of a function call tells go to run this function in a brand new go routine
- Go Scheduler: Monitors code running inside of go routines and starts and stops executing them. By default Go only attempts to use one CPU even if you have multiple, scheduler runs one routine until it finishes or makes a blocking call (i.e. an HTTP request).  If you have multiple cores then you can have go run multiple routines at once (one routine at a time per core)
- Never try to access the same variable from a different child routine.  They could be editing the same location in memory and cause unintended behavior.  Only share information with a child routine by passing it in as an argument or communicating with the child routine via channels

How many go routines can be spun up at once?
Does it only work for process that don't take up processing power on your own machine?

Concurrency: We can have multiple threads executing code.  If one thread blocks, another one is picked up and worked on.

Parallelism: Multiple threads executing at the exact same time.  Requires multiple CPU's

``` go
	for _, link := range links {
		go checkLink(link, c)
	}
```


## Channels
- used to communicate between different running go routines, two way messaging system with a sender and a reciever
- channels are typed
- Receiving messages from a channel is a blocking call, you have to wait for it to finish for the code to continue executing

``` go
// create a channel that can pass strings between go routines
c := make(chan string)
```

## Errors

``` go
// What could go wrong?
// If something goes wrong here what do I really want to have happen?
    bs, err := ioutil.ReadFile(filename)
    
    if err != nil {
        fmt.Println("Error:", err)
		os.Exit(1)
	}

```

## Scope
- Files in the same package can freely call functions defined in other files

## Files
- Every file must declare at the top of the file the package it belongs to `package main`

``` go
// package main will compile into an executable file
package main

// import packages at the top of a file
import (
    "fmt"
    "io/ioutil"
)

func main() {
    // main() function will be auto called when a package called main is executed
}
```





# Language Quirks
- Go's random number generator always starts with the same seed, you will always get the same squence of random numbers unless you set your own seed.
- Not O.O. there is no idea of classes, instead we take a type and extend it by adding some extra functionality (functions).  A function with a receiver is like a method.
- When we run our code go compiles it into an executable file in a temporary directory then runs that file

## Testing
- How do we decide what to test?  What are we actually testing?  What do I really care about?
- Create a new file in the project directory that ends in _test.go
- Test function names start with `Test` and then the name of the function or functions you are testing i.e. `func TestNewDeck(t *testing.T) {}` or `func TestSaveToDeckandNewDeckFromFile(t *testing.T) {}`
- `$go test // executes all test files in the directory`
- You have to take care of cleanup yourself There is no file cleanup so you have to manually do it (i.e. it writes some file in a test you have to delete it)

``` go
```

 ## General Learning
 - Briefly explain your approach and possible challenges you overcame.
 - How did you do on this exercise?
 - Take a moment to reflect on what you learned from this exercise
 - What went well?  How could be improved?
 - What are the basic questions I need answers to when learning a new programming language?

## Tools

## Misc Notes
* How does memory work on a computer? (memory address, memory value, RAM)
- Two types of variables
    1. Varaibles that point at a memory address
    2. Variables that point at values


# Quiz

What is the difference between an array and a slice in go?
What is the difference between reference and value types?
When do you want to use a map vs using a struct?
- Do we have to explicitly declare return value types?
- What is the difference between a function and a method?
- On a high level what are we trying to do?
- What are the edge cases?
- What could be a problem down the road?