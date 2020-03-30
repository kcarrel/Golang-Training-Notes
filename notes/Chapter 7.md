# Interfaces
Interface types allow us to write functions that are more flexible and adaptable because they are not tied to the details of one particular implementation. 

Go has *special interfaces* because they are satisfied implicitly - there is no need to declare all interfaces that a given concrete type satisfies. 

## Interfaces as Contracts
An *interface type* is an abstract type - it doesn't expose it's values or basic operations. Only reveals some of their methods and what they do. 

The freedom to substitute one type for another that satifies the same interface is called *substitutability* and is a hallmark of OOP.

## Interface Types
An *interface type* specifies a set of methods that a concrete type must posses to be considered an instance of that interface. 

Ther order in which the methods appear doesn't matter - all that matters is the set of methods. 

## Interface Satisfaction
A type satisfied an interface if it has all the methods the interface requires. 

A concrete type is a particular interface type. 
An expression may be assigned to an interface only if its type satsfies the interface.
ex:
```

var w io.Writer
w = os.Stdout //OK! os.File has Write method
w = time.Second //Not Ok! does not have Write method
```

type interface {}, which is called the empty interface, is highly useful because it places no demand on the types that satisfy it meaning that any value can be assigned to an empty interface.

ex:
```
var any interface{}
    any = true
    any = 12.34
    any = "hello"
```

It is occasionally useful to document and asser the relationship between a concrete type and the interfaces it satisfies.
ex:
```
var _ io.Writer = (*bytes.Buffer)(nil)
```
## Parsing Flags with flag.Value
It's easy to define new flag notations for our own data types - we only need to define a type that satisfies the flag.Value interface (below)
package flag

type Value interface {
    String() string
    Set(string) error
} 
## Interface Values
An interface value has 2 components:
- concrete value (dynamic type)
- value of that type (dynamic value)

Interface values can be compared and may be used as keys of a map or the operand of a switch statement. 

## Caveat: An Interface Containing a Nil Pointer is Non-Nil
A nil interface value, which contains no value at all, is not the same as an interface value containing a pointer that happens to be nil. This is a trap for defensive programming.
Solution - change the type  to avoid the assignment of the dysfunctional value to the interface in the first place. 
ex:
```
var buf io.Writer
if debug {
    buf = new(bytes.Buffer) //enable collection of output
}
f(buf) // OK
```

## Sorting with sort.Interface
In Go - sort.Sort function assumes nothing about the representation of either the sequence or its elements. Instead, it uses an interface, sort.Interface to specify the contract between the generic sort algorithm and each sequence type that may be sorted. 

## The error Interface
The simplest way to create an error is by calling errors.New which returns a new error for a given error message. 

## Type Assertions
A *type assertion* is an operation applied to an interface value that checks that the dynamic type of its operand matches the asserted type. 

ex:
```
var w io.Writer
w = os.Stdout
f := w.(*os.File) // success: f == os.Stdout
c := w.(*bytes.Buffer) //panic: interface holds *os.File, not *bytes.Buffer

```
## Discriminating Errors with Type Assertions
3 kinds of I/O failure:
- file already exists
- file not found
- permission denied 

3 helper functions to classify the failure by a given error value:
func IsExit(err error) bool
func IsNotExist(err error) bool
func IsPermission(err error) bool

## Querying Behaviors with Interface Type Assertions
We can define a new interface that has just the method in question and use a type assertion to test whether the dynamic type of w satisfies this new interface. 
ex:

```
package fmt

func formatOneValue(x interface{} string) {
    if err, ok := x.(error); ok {
        return err. Error()
    }
    if str, ok := x.(Stringer); ok {
        return str.String()
    }
}
```

## Type Switches 
A type switch looks like an ordinary switch statement that enables a multi-way branch based on the interface value's dynamic type.
ex:
```
switch x.(type) {
    case nil:
    case int, uint:
    case bool:
    case string:
    default:
}
```

## A Few Words of Advice
Avoid the novice Go Programmer practie of starting by creating a set of interfaces and only later define the concrete types that satisfy them.
Instead - restrict which methods of a type or fields of a struct are visible outside a package using the export mechanism. Interfaces are only need when there are two or more concrete types that must be dealt with. 

Good rule of thumb: When designing an interface ask only for what you need. 
