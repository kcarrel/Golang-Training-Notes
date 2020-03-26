# Functions

## Function Declarations
```
func name(parameter-list) (result-list) {
    body
}
```
Results may be named (like args)!

**4 ways to declare a function**
```
func add(x int, y int) int
func sub(x, y int) (z int)
func first(x int, _ int) int
func zero(int, int) int
```

Note: Go does not have default parameter values

Typical Go implementations use variable-size stacks that start small and grow as needed up to a limit on the order of a gigabyte. This allows Go to use recursion safely and without worrying about overflow. 

## Multiple Return Values
A function can return more than one result in Golang.

A bare return can be used to return each of the named result variables in order and avoid code duplication. Downside: rarely make code easier to understand! Use them sparingly. 

## Errors
A function for which failure is an expected behavior returns an additional result, conventionally the last one. If the failure has only one possible cause, the result is a boolean, usually called ok. 
ex:
```
value, ok := cache.Lookup(key)
if !ok {
    // ...cache[key] does not exist...
}
```
The built in type error is an interface type.

Go's approach sets it apart from many other languages in which failures are reported using exceptions, not ordinary values. Exception mechanisms is only used for reporting truly unexpected errors that indicate a bug, not the routine errors that a robust program should be built to expect. 

Go programs use ordinary control-flow mechanisms like if and return to respond to errors.

## Error-Handling Strategies
5 possible error handling strategies:
- Propagate the error, so that a failure in a subroutine becomes a failure of the calling routine.
- For errors that represent transient or unpredictable problems, it may make sense to retry the failed operation, possible with a delay between tries and perhaps with a limit on the number of attempts or the time spent trying before giving up entirely.
- If progress is impossible, the caller can print the error and stop the program gracefully (note: this should be reserved for the main package of a program)
- In some cases, it's sufficient just to log the error then continue, perhaps with reduced functionality.
- In rare cases, we can safely ignore an error entirely

## End of File(EOF)
io.EOF has a fixed error message "EOF". 
ex:

```
package io
import "errors"

in := bufio.NewReader(os.Stdin)
for {
    r, _, err := in.ReadRune()
    if err == io.EOF {
        break // finished reading
    }
}
```

## Function Values
Functions are *first-class values* in Go. Function values have types and can be assigned to variables or passed to or returned from functions. A function value may br called like any other function. Function values may be compared with nil. Function values let us parameterize our functions over data and behavior too. 

## Anonymous Functions
A *function literal* can be used to denote a function value within any expression. A function literal is written like a function declaration but without a name following the func keyword. The value is an *anonymous function*.

**Recursion in Anonymous functions**: must first declare a variable, and then assign the anonymous function to the variable. 


## Caveat: Capturing Iteration Variables
A consequence of the scope rules for loop variables - All function values created by a loop "capture" and share the same variable - an addressable storage location, not its value at that particular moment. 

## Variadic Functions
A variadic function is one that can be called with varying numbers of arguments. (ex: fmt.Printf)

To declare a variadic function: the type of the final parameter is preceded by an ellipsis, '...' which indicates that the function may be called with any number of arguments of this type.
ex:
```
func sum(vals ...int) int {
    total := 0
    for _, val := range vals {
        total += val
    }
    return total
}
```

Variadic functions are often used for string formatting. 

## Deferred Function Calls
Syntatically - a *defer* statement is an ordinary function or method call prefixed by the keyword *defer*. 

The function and argument expressions are evaluated when the statement is executed, but the actual call is deferred until the function that contains the defer statement has finished. Any number of calls can be deferred - they are executed in the reverse of the order in which they are deferred. 

Defer statements are often paired with open/close/connect/disconnect/lock/unlock to ensure resources are released in all cases. 

Can be used to pair "on entry" and "on exit" actions when debugging a complex function. 

Deferred functions run after return statements have updated the function's result variables. 

## Panic
When the Go runtime detects some mistakes, like out of bounds array access or nil pointer dereference it *panics*.

Typical panic:
- normal execution stops
- all deferred function calls in that goroutine are executed
- the program crashes with a log message 
The log message includes
- a panic value (error message)
- for each goroutine a stack trace

Note: not all panics come from runtime. It can also be called directly! A panic is often the best thing to do when an "impossible situation" happens!

Panics are reserved for grave errors (logical inconsistencys).

## Recover 
Recover ends the current state of panic and returns the panic value. 

Note: don't recover from panics indiscriminately. It is safest to recover selectively if at all. 