## Ways to run a Go app
1. `go run <filename>`

    If the `main` function references other functions in the same package but in a different file, the command explicitly is `go run *.go` or `go run <file of main>.go <file of other function>.go`. 

    For example, `conditionals.go` refers to function in `switch.go`; so `go run *.go` or `go run conditionals.go switch.go` will work.

2. `go  build`

    Used to build an executable. Going into the package folder and running `go build` will create an executable with the same name as the package.

3. `go install` 
    
    Creates an executable in the `bin` folder of the `go` folder. Location can be changed using `GOPATH` and `GOBIN`, which are set using 
    
    `go env -w GOPATH="<path>"` and 
    
    `go env -w GOBIN="<path>"`

## Working with constants and variables

A variable can be declared at the package level and remain unused. But within a function, unused variables will result in errors.

Go passes arguments by values, not reference. It can be overriden by using pointers. When a value without the `&` is passed to a function, a copy of the input value is made, operated upon, and returned. The original value remains unchanged.

However, when the address to a variable is passed with the `&variableName` notation, the value at the location is changed directly.

### Environment Variables
These are settings on your working system. Can be obtained using `os` package.

## Functions

`func functionName(input <inputType>) returnType {}` is called the message signature. A function requires an input in the form of `<inputName inputType>`. 

Multiple inputs, also called parameters, are comma-separated. `returnType` is the type of the output returned by the function. 

More than one return value is possible, and the types will be enclosed in parantheses and comma-separated as 
`(type1, type2)`

`return` keyword is used to end the function.

### Variadic function
can be called with a varying number of trailing arguments, used when not sure how many inputs will be received.

Declared using 
```
func functionName(paramName ...inputType)
```

The `...` ellipses indicate that the function should expect any number of inputs of the declared data type.

## Arrays and Slices

Arrays are numbered lists of the **same type**.

Slices are used a lot more in Go. 

Arrays have a fixed length. 

Slices are resizable. They're built on top of arrays. No actual data gets stored, they're just pointers to the underlying array. 

They're cheap because they're pointers. However, as a result, changing value of underlying array changes values of all slices referencing it. 

Slices are references, so they're passed as references to functions. Slice itself is passed as a value, but the values are references.

