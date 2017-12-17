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
func (t typeName) methodName(v1 typeName, size int) (returnValue, returnValue2) {
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
- 

## Lexicon
- Method: a function that belongs to an instance
- Type conversion: take one type and turn it into another

## Questions
- Do we have to explicitly declare return value types?
- What is the difference between a function and a method?

# Language Quirks
- Go's random number generator always starts with the same seed, you will always get the same squence of random numbers unless you set your own seed.
- Not O.O. there is no idea of classes, instead we take a type and extend it by adding some extra functionality (functions).  A function with a receiver is like a method.