package adtenum

import (
	"fmt"
	"testing"
)

// First we need to create the enumeration type that all the values will belong under.
type WebEvent Enum[WebEvent]

// The following section creates the enum value for a PageLoad event.
// Declare the PageLoad enum value as a constant string value.
type PageLoad ConstValue[string]

// Here we create a constructor that will generate a value that will always be a string with the value "PageLoad".
// This provides us with a way to create new values.
var NewPageLoad func() PageLoad = CreateConstValueConstructor[PageLoad]("PageLoad")

// To relate this only to the WebEvent type, we need to implement this method so it returns the type this enum value belongs under.
// This will provide us with compile time type checking.
func (val PageLoad) EnumType() WebEvent {
	return val
}

// The following section creates the enum value for a PageUnload event.
// Declare the PageUnload enum value as a constant string value.
type PageUnload ConstValue[string]

// Here we create a constructor that will generate a value that will always be a string with the value "PageUnload".
// This provides us with a way to create new values.
var NewPageUnload func() PageUnload = CreateConstValueConstructor[PageUnload]("PageUnload")

func (val PageUnload) EnumType() WebEvent {
	return val
}

// The following section creates the enum value for a KeyPress event.
// Declare the KeyPress enum value that contains a single value which is a rune.
type KeyPress OneVariantValue[rune]

// Here we create a constructor that will accept a rune and create an enum value containing that rune.
// This provides us with a way to create new values of the KeyPress enum value.
var NewKeyPress func(rune) KeyPress = CreateOneVariantValueConstructor[KeyPress]()

func (val KeyPress) EnumType() WebEvent {
	return val
}

// The following section creates the enum value for a Paste event.
// Declare the KeyPressValue enum value that contains a single value which is a string.
type Paste OneVariantValue[string]

// Here we create a constructor that will accept a string and create an enum value containing that string.
// This provides us with a way to create new values of the Paste enum value.
var NewPaste func(string) Paste = CreateOneVariantValueConstructor[Paste]()

func (val Paste) EnumType() WebEvent {
	return val
}

// The following section creates the enum value for a Click event.
// Declare the TwoVariantValue enum value that contains two values which are both integers.
type Click TwoVariantValue[int, int]

// Here we create a constructor that will accept two strings and create an enum value containing the two integers.
// This provides us with a way to create new values of the Click enum value.
var NewClick func(int, int) Click = CreateTwoVariantValueConstructor[Click]()

func (val Click) EnumType() WebEvent {
	return val
}

// This example was ported from Rust By Example at the following link https://doc.rust-lang.org/rust-by-example/custom_types/enum.html.
func inspect(event WebEvent) string {
	// We can now perform a type switch on the WebEvent type to determine which enum value we have.
	// We can then extract the value from the enum value similar to what could be done in Rust.
	switch vals := event.(type) {
	case PageLoad:
		return fmt.Sprint(vals())
	case PageUnload:
		return fmt.Sprint(vals())
	case KeyPress:
		return fmt.Sprintf("%c", vals())
	case Paste:
		return fmt.Sprint(vals())
	case Click:
		return fmt.Sprint(vals())
	default:
		return "Unknown"
	}
}

// This example used in this tests was ported from Rust By Example at the following link https://doc.rust-lang.org/rust-by-example/custom_types/enum.html.
func TestEnums(t *testing.T) {
	pressed := NewKeyPress('x')
	pasted := NewPaste("my text")
	click := NewClick(20, 80)
	load := NewPageLoad()
	unload := NewPageUnload()

	if inspect(pressed) != "x" {
		t.Errorf("The inspect function did not return the expected value for KeyPress")
	}

	if inspect(pasted) != "my text" {
		t.Errorf("The inspect function did not return the expected value for Paste")
	}

	if inspect(click) != "20 80" {
		t.Errorf("The inspect function did not return the expected value for Click")
	}

	if inspect(load) != "PageLoad" {
		t.Errorf("The inspect function did not return the expected value for PageLoad")
	}

	if inspect(unload) != "PageUnload" {
		t.Errorf("The inspect function did not return the expected value for PageUnload")
	}
}
