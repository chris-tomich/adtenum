# ADT Enumerations for Go

[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

This project provides Rust-style enumerations, also known as algebraic data types (ADTs), for Go. It aims to provide the power and flexibility of pattern matching with enum types to Go. At the moment, to use the types and functions in this library to create enums, some boiler plate code is required but there is a plan to create a go generate utility to help reduce manual creation of that boiler plate code.

## Example

Using `adtenum` package, you're able to create enumerations that contain tuples which can then be used in a type switch as in the following example.

```go
func inspect(event WebEvent) string {
	// We can now perform a type switch on the WebEvent type to determine which enum value we have.
	// We can then extract the value from the enum value similar to what could be done in Rust.
	switch vals := event.(type) {
	case PageLoadValue:
		return fmt.Sprint(vals())
	case PageUnloadValue:
		return fmt.Sprint(vals())
	case KeyPressValue:
		return fmt.Sprintf("%c", vals())
	case PasteValue:
		return fmt.Sprint(vals())
	case ClickValue:
		return fmt.Sprint(vals())
	default:
		return "Unknown"
	}
}
```

For a more thorough example look at `enum_test.go` which contains a Go port of the WebEvent example written to explain [Rust enumerations in Rust By Example](https://doc.rust-lang.org/rust-by-example/custom_types/enum.html).

## Quick Start

First import the package using `go get`.

```bash
go get github.com/chris-tomich/adtenum
```

Declare the new enumeration type using the `adtenum.Enum` as the underlying type. In the following example we create a new `WebEvent` enumeration type which we also pass as the first type parameter to `adtenum.Enum`

```go
type WebEvent adtenum.Enum[WebEvent]
```

Once you have your enumeration type, you can begin creating the values for the enumeration. All enumeration values use one of the `adtenum` Value types as their underlying type. In the following example we create a constant value which will always have the value `PageLoad`.

```go
type PageLoadValue adtenum.ConstValue[string]
```

In this next example, we create an enumeration value called `Click` that contains two integers.

```go
type ClickValue adtenum.TwoVariantValue[int, int]
```

Once we've added our values, we now need to create constructors for creating new instances of our enumeration values. In the future, a go generate utility is planned to help with creation of this boiler plate. For our constant value, we need to provide the constant value that will be used for every value of the enumeration value to the `CreateConstValueConstructor`.

```go
var NewPageLoad func() PageLoadValue = adtenum.CreateConstValueConstructor[PageLoadValue]("PageLoad")
```

For our enumeration value `Click` we create our constructor using `CreateTwoVariantValueConstructor` as in the example below.

```go
var NewClick func(int, int) ClickValue = adtenum.CreateTwoVariantValueConstructor[ClickValue]()
```

Once we've created constructors to create new enumeration values for us, we need to complete the last piece which makes our enumeration value types legitimate values of our enumeration type `WebEvent`. The following code is effectively implementing the interface of our `WebEvent` enumeration type.

```go
func (val PageLoadValue) EnumType() WebEvent {
	return val
}

func (val ClickValue) EnumType() WebEvent {
	return val
}
```

Once this is completed, we can now use the enumeration type and it's values. We also now have similar pattern matching capabilities as found with Rust enumerations. In the following example, we can use a Go type switch to elicit the enumeration value type into a variable `vals` and then get the values for that tuple.

```go
func inspect(event WebEvent) string {
	switch vals := event.(type) {
	case PageLoadValue:
		return fmt.Sprint(vals())
	case ClickValue:
        x, y := vals()
		return fmt.Sprint(x, y)
	default:
		return "Unknown"
	}
}

func main() {
    load := NewPageLoad()
    click := NewClick(20, 80)

    // writes 'PageLoad' to output
    fmt.Println(inspect(load))

    // writes '20 80' to output
    fmt.Println(inspect(click))
}
```