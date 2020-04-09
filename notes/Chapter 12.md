# Reflection
A mechanism Go provides to update variables and inspect their variables at runtime, call their mehtods and apply the operations intrinsic to their representation without knowing their types at compile time.

## Why Reflection?
Without a way to inspect the representation of values of unknown types, we quickly get stuck <--- Reflection to the rescue!

## reflect.Type and reflect.Value
Reflection is provided by the reflect package. 

Type ex:
```
t := reflect.TypeOf(3) //a reflect.Type
fmt.Println(t.String()) // "int"
fmt.Println(t) //"int"

```

Value ex:
```
v := reflect.ValueOf(3) // a reflect value
fmt.Println(v)      // "3"
fmt.Printf("%v\n", v)      //"3"
fmt.Println(v.String())    // Note: "<int Value>"
```

reflect.Value.Interface method returns an interface{} holding the same concrete value as the reflect.Value:
```
v := reflect.ValueOf(3) // a reflect value
x := v.Interface()  //an interface{}
i := x.(int)    // an int
fmt.Printf("%d\n", i)   // "3"
```

## Display, a Recursive Value Printer
Display is a debugging utility function that prints the complete structure of that value, labeling each element with the path by which it was found.

```
func display(path string, v reflect.Value) {
   switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path,
				formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}
```

## Setting Variables with reflect.Value
use CanAddr to see whether a value is addressable
ex:
```
fmt.Println(a.CanAddr())    //"false"
```
3 steps to recover a variable from an addressable reflect.Value:
```
x := 2
d := reflect.ValueOf(&x).Elem() // d refers to the variable x
px := d.Addr().Interface().(*int)   // px := &x
*px = 3     // x = 3
fmt.Println(x)  //"3"
```

## Accessing Struct Field Tags
The Unpack function populates the struct from the request so that the parameters can be accessed conveniently and with an appropriate type. 
It does 3 things:
- calls req.ParseForm() to parse the request
- builds a mapping from the effective name of each field to the variable for that field
- Iterates over the name/value pairs of the HTTP parameters and updates the corresponding struct fields. 

## 3 reasons why reflection should be used with care
- reflection-based code can be fragile
- since types serve as a form of documentation and the operations of reflection cannot be subject to static type checking, heavily reflective code is often hard to understand
- reflection-based functions may be one or two orders of magnitude slower than code specialized for a particular type 
