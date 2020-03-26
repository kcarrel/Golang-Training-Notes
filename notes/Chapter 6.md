# Methods
An object-oriented program is one that uses methods to express the properties and operations of each data structure so that clients need not access the object's representation directly. 

## Method Declarations
A method is declared with a variant of the ordinary function declaration in which an extra paramete appears before the function name. The parameter attaches the function to the type of that parameter.

ex:
```
// traditional function
func Distance(p, q Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same but as a method 
func (p Point) Distance (q Point) float64 {
    return math.Hypot(q.X-p.X, q.Y-p.Y)
}
```

All methods of a given type must have unique names but *different types can use the same name for a method*. 

Benefit over ordinary functions: method names can be shorter.

## Methods with a Pointer Receiver
If a function needs to update a variable (or if an arg is so large that we want to avoid copying) we must pass the address of the variable using a pointer. 
ex:
```
func (p *Point) ScaleBy(factor float64) {
    p.X *= factor
    p.Y *= factor 
}
```

**Nil is a valid reciever value**

**Method Values and Expressions**
It is possible to seperate selecting and calling a method into two operations. 

## Encapsulation
A variable or method of an object is said to be *encapsulated* if it is inaccessible to clients of the object.

Encapsulation is also known as *information hiding*. 

Go has 1 mechanism to control the visibility of names: capitalized identifiers are exported from the package in which they are defined but uncapitalized names are not. 

3 benefits of encapsulation:
- clients cannot directly modify the object's variables 
- prevents clients from depending on things that might change
- prevents clients from setting an object's variables arbitrarily 

ex:
```
type Counter struct { n int }

func (c *Counter) N() int       //return c.n
func (c *Counter) Increment()   // c.n++
func (c *Counter) Reset()       //c.n = 0
```

