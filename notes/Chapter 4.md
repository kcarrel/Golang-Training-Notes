# Composite Types
Composite types 
- arrays
- slices
- maps
- structs

Arrays and structs are aggregate types: values are concatenations of other values in memory. 

Slices and maps are dynamic data structures that grow as values are added. 

## Arrays
An array is a fixed-length sequence of zero or more elements of a particular type. 
Arrays *are rarely used in Go* because they are fixed in length. (use a slice!)
ex:
```
var a4 [4]int           // An array of 4 ints, initialized to all 0.

a5 := [...]int{3, 1, 5, 10, 100} // An array initialized with a fixed size of five

```
## Slices
Slices have dynamic size which makes it preferrable over arrays. A slice is a lightweight data struct that gives access to subsequence of the elements of an array, which is known as the slice's underlying array. 

3 components of a slice:
- pointer
- length
- capacity 

ex:
```
summer = months[6:9]
data : "June", "July", "August"
len : 3
cap : 7
```

Note: Slices are not comparable, so we cannot use == to test whether two slices contain the same elements. We must implement the comparison ourselves.

The built-in function *make* creates a slice of a specified element type, length and capcity. The cap arg may be omitted. 
ex:
```
make([]T, len)
make([]T, len, cap)     // same as make([]T, cap, [:len])
```

## Append
To append elements to a slice, the built-in append() function is used.
```
    s := []int{1, 2, 3}     // Result is a slice of length 3.
    s = append(s, 4, 5, 6)  // Added 3 elements. Slice now has length of 6.
    fmt.Println(s) // Updated slice is now [1 2 3 4 5 6]
```

## Maps
In Go a *map* is a reference to a hash table, and a map type is written map[K]V, where K and V are the types of its keys and values. 

The built in function make can be used to create a map:
```
ages := make(map[string]int)        // mapping from strings to ints
```

You can also use a *map literal* to create a new map populated with some initial key/value pairs:
```
ages := make(map[string]int)
ages["alice"] = 31
ages["charlie"] = 34
```

If a key is not present in the map --> returns 0

Order of a map is unspecified/random -> to enumerate in order we must sort keys explicitly. 
ex:

```
    import "sort"

    var names []string
    for name := range ages {
        names = append(names, name)
    }

    sort.Strings(names)
    for _, name := range names {
        fmt.Printf("%s\t%d\n", name, ages[name])
    }
```
Slices cannot be compared to each other; only nil
