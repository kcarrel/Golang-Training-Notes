# Basic Data Types

**4 Categories**
- Basic Types (numbers, strings and bools)
- Aggregate Types (arrays and structs)
- Reference Types (pointers, slices, maps, channels and functions)
- Interface Types 

**Complex Numbers**
The built-in function *complex* creates a complex number from its real and imaginary components, and the built-in real and *imag* functions extract those components.

An *imaginary literal* represents the imaginary part of a complex constant. 
It consists of an integer or floating-point literal followed by the lower-case letter i. 
The value of an imaginary literal is the value of the respective integer or floating-point literal multiplied by the  imaginary unit i. 

Note: 2 complex numbers are equal if their real parts are equal and their imaginary parts are equal. 

**Strings**
A string is an immutable sequence of bytes.

examples:
```
fmt.Println(s[:5]) // "hello"
fmt.Println(s[7:]) // "world"
fmt.Println(s[:])  // "hello, world"
```
String values are immutable.

**String Literals**
A *string literal* is a sequence of bytes enclosed in double quotes. 

**Strings and Byte Slices**

**4 packages for manipulating strings**
- Bytes: Similar functions for manipulating slices of bytes, of type []byte which share some properties with strings.
- Strings: Provides many functions for searching, replacing, comparing, trimming, splitting and joining strings.
- Strconv: Provides functions for converting boolean, integer, and floating-point values to and from their string representations, and function for quoting and unquoting strings.
- Unicode: Provides functions like IsDigit, IsLetter, IsUpper, IsLower for classifying runes. 

**Conversions between Strings and Numbers**
- use fmt.Spritnf
- use strcvon.Itoa

FormatInt and FormatUnit can be used to format numbers in a different base. 

Parse a string representing an integer:
- use strconv.Atoi
- use strconv.ParseInt

**Constants**
Constants are expressions whose values is known to the compiler and whose evaluation is guaranteed to occur at compile time, not run time. 

Underlying type of every constant is a basic type:
- boolean
- string
- number 

**The Constant Generator iota**
A *constant generator iota* is used to create a sequence of related values without explicitly declaring. 

In a *const* declaration, the value of iota begins at 0 and increments by one for each item in the sequence.
ex:

type Weekday int

const (
    Sunday Weekday = iota
    Monday
    Tuesday
    Wednesday
    Thursday
    Friday
    Saturday
)

**Untyped Constants**
Constants are unusual in Go - Many constants are not committed to a particular type. 

6 untyped constants:
- untyped boolean
- untyped integer
- untyped runes
- untyped floating-point 
- untyped complex 
- untyped string 

Untyped constants retain their highest precision until later but they can participate in many more expressions than committed constants without requiring conversions. 

In a variable declaration without an explicity type the flavor of the untyped const implicitly determines the default type of the variable.

i := 0 // implicit int(0)

