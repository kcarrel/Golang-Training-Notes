# Low-Level Programming
Safety net for Go:
- type checking during compilation
- strict rules for type conversions 
- dynamic checks
- automatic memory management

Inaccessible implementation details:
- no way to discover the memory layout of an aggregate type

- a pointer identifies a variable without revealing the variable's numeric address
- pointers are transparently updated addresses may change as the garbace collector moves variables


## unsafe.Sizeof, Alignof, and Offsetof
unsafe.Sizeof - reports the size in bytes of the representation of its operand, which may be an expression of any type; the expression is not evaluated.

Sizeof - reports only the size of the fixed part of each data structure, like th epointer and length of a string, but not indirect parts like the contents of the string. 

Holes are unused space added by the compiler to ensure that the following field or element is properly aligned relative to the start of the struct of array.

unsafe.Alignof - reports the required alignment of its argument's type.

unsafe.Offsetof - computes the offset of field f relative to the start of its enclosing struct x. 

## unsafe.Pointer
unsafe.Pointer - a special kind of pointer that can hold the address of any variable.They can be compared.

In general, unsafe.Pointer conversions let us write arbitrary values to memory and thus subvert the type system. 


## Caution!
High level languages insulate programs/programmers from the arcane specifics of computer instruction sets, dependence on irrelevancies, how big a data type is, the details of structure layout, and a host of other implementation details. 

The unsafe package lets programmers see crucial but otherwise inaccessible feature to achieve higher performance at the expense of portability/safety. 