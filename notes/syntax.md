Main is a special name declaring an executable rather than a library
Init runs first - the init function exists to initialize data for any packages with init function, not only for the main package.

```
package main

func init() {

}

func main() {

}
```

Packages is where functions live and one package can contain many files that you can name them freely.

Note: One important rule for constructing a package is “there must be only one package per one folder


Import declaration declares library packages referenced in this file

import (
    "fmt"
    "os"
)

Functions are defined by:
func main() {

}

**3 ways to declare a variable**

```
var first string 
first = "This is a string."

var second = "This is a string."

third := "This is a string."
```

**Variable Types**
If you declare a variable without assigning any value - the variable is set to zero value of its type
```
 a := 1    // var a int
  b := 3.14 // var b float
  c := "hi" // var c string
  d := true // var d bool
  fmt.Println(a, b, c, d)

  e := []int{1, 2, 3} // slice
  e = append(e, 4)
  fmt.Println(e, len(e), e[0], e[1:3], e[1:], e[:2])
  
  f := make(map[string]int) // map
  f["one"] = 1
  f["two"] = 2
  fmt.Println(f, len(f), f["one"], f["three"])
```
- int
- float
- string
- boolean
- slice (can appends its element into the back of it)
- array (similar to slice but fixed length)
- struct (similar to C)
- map
```

// map
import "fmt"

type Person struct{
  Name    string `json:"name"`
  Age     int    `json:"age"`
  isAdmin bool
}

func main() {
  p := Person{
    Name:    "Mike",
    Age:     16,
    isAdmin: false,
  }
  fmt.Println(p, p.Name, p.Age, p.isAdmin)
}


```

**Type conversion**
Some pairs of type can convert by casting directly (int(1.23), float(3), uint32(5)), but for some require an extra library to do a conversion such as string-int, int-string, bool-string, and string-bool.
```
example
    a := 3.14
    b := int(a)
```

**Control Flow**

```
    if/else 

    if a > a thing {
    } else if a != another thing {
    } else {
    }


    switch

    a := thing 
    switch a {
        case "Case1":
            a = "bloop"
        case "Case2"
            a = "No":
        default:
            a = "Y"
    }
```

**Loops**
```
    for i := 0; i < a; i++ {
        the thinggg
    }

    for a > thing {
        a--
    }

    a := []int{0, 2, 3, 5}
    for _, i := range a {
        salkj
    }
```
## Command Line Arguments
os.Args is a slice of strings.

In go - a *slice* is a dynamically sized sequence s of array elements where individual elements can be accessed as s[i] and a contiguous subsequence s[m:n]