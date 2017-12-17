# Go Notes
## Table of Contents




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

## Operators


## Types

Array: fixed length list of things, must be defined with the data type it contains
Slice: an array that can grow and srink, must be defined with the data type it contains


``` go

// create a new type of 'deck' which is a slice of strings
type deck []string
```

### Slice
- Does this return a new slice or modify the original?
- Slices are 'zero indexed'
- Byte Slice: a slice of ascii values.  aka a string of characters

``` go
    sliceName := []string{"First", "Second"}
    sliceName[1] // return second element
    sliceName[0:2] // sliceName[startIndexIncluding:upToNotIncluding]
    sliceName[:2] // sliceName[:upToNotIncluding]
    sliceName[2:] // sliceName[:includingTilEndOfSlice]
```

## Variables
- Variables can be initialized outside a function but not assigned a value

``` go
    var varName string = "Hello friend"
    var varName := "Hello neighbor"

```

## Functions
Receiver: the variable name and type that a function can be called on, every variable of this type will gain access to this function.  Identify using a one or two letter abbreviation
Function Signature: The declarative line of a function



``` go
// a reciever t of type typeName 
func (t typeName) methodName(v1 typeName, size int) (returnType, returnType2) {
    return 1, 2
}

firstVal, secondVal := calledOn.methodName(inputOfTypeA, 6)

func newCard() string {
    return "Five of Diamonds"
}
```

## Loops

- if you declare the index or the value they MUST be used in the return statement

``` go
    for index, value := range itterateOverVar {
        fmt.Println(i, card)
    }
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

## Testing
- How do we decide what to test?  What are we actually testing?  What do I really care about?
- Create a new file in the project directory that ends in _test.go
- Test function names start with `Test` and then the name of the function or functions you are testing i.e. `func TestNewDeck(t *testing.T) {}` or `func TestSaveToDeckandNewDeckFromFile(t *testing.T) {}`
- `$go test // executes all test files in the directory`
- You have to take care of cleanup yourself There is no file cleanup so you have to manually do it (i.e. it writes some file in a test you have to delete it)

``` go
```

## Lexicon
- Method: a function that belongs to an instance
- Type conversion: take one type and turn it into another

## Questions
- Do we have to explicitly declare return value types?
- What is the difference between a function and a method?
- On a high level what are we trying to do?
- What are the edge cases?
- What could be a problem down the road?

# Language Quirks
- Go's random number generator always starts with the same seed, you will always get the same squence of random numbers unless you set your own seed.
- Not O.O. there is no idea of classes, instead we take a type and extend it by adding some extra functionality (functions).  A function with a receiver is like a method.

 # General Learning
 - Briefly explain your approach and possible challenges you overcame.
 - How did you do on this exercise?
- Take a moment to reflect on what you learned from this exercise
 - What went well?  How could be improved?