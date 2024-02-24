package adtenum

import (
	"fmt"
	"testing"
)

// First we need to create the enumeration type that all the values will belong under.
type WebEvent Enum[WebEvent]

// The following section creates the enum value for a PageLoad event.
// Declare the PageLoadValue enum value as a static string value.
type PageLoadValue StaticValue[string]

// Here we create a constructor that will generate a value that will always be a string with the value "PageLoad".
// This provides us with a way to create new values.
var NewPageLoad func() PageLoadValue = CreateStaticValueConstructor[PageLoadValue]("PageLoad")

// To relate this only to the WebEvent type, we need to implement this method so it returns the type this enum value belongs under.
// This will provide us with compile time type checking.
func (val PageLoadValue) EnumType() WebEvent {
	return val
}

// The following section creates the enum value for a PageUnload event.
// Declare the PageUnloadValue enum value as a static string value.
type PageUnloadValue StaticValue[string]

// Here we create a constructor that will generate a value that will always be a string with the value "PageUnload".
// This provides us with a way to create new values.
var NewPageUnload func() PageUnloadValue = CreateStaticValueConstructor[PageUnloadValue]("PageUnload")

func (val PageUnloadValue) EnumType() WebEvent {
	return val
}

// The following section creates the enum value for a KeyPress event.
// Declare the KeyPressValue enum value that contains a single value which is a rune.
type KeyPressValue OneVariantValue[rune]

// Here we create a constructor that will accept a rune and create an enum value containing that rune.
// This provides us with a way to create new values of the KeyPress enum value.
var NewKeyPress func(rune) KeyPressValue = CreateOneVariantValueConstructor[KeyPressValue]()

func (val KeyPressValue) EnumType() WebEvent {
	return val
}

// The following section creates the enum value for a Paste event.
// Declare the KeyPressValue enum value that contains a single value which is a string.
type PasteValue OneVariantValue[string]

// Here we create a constructor that will accept a string and create an enum value containing that string.
// This provides us with a way to create new values of the Paste enum value.
var NewPaste func(string) PasteValue = CreateOneVariantValueConstructor[PasteValue]()

func (val PasteValue) EnumType() WebEvent {
	return val
}

// The following section creates the enum value for a Click event.
// Declare the TwoVariantValue enum value that contains two values which are both integers.
type ClickValue TwoVariantValue[int, int]

// Here we create a constructor that will accept two strings and create an enum value containing the two integers.
// This provides us with a way to create new values of the Click enum value.
var NewClick func(int, int) ClickValue = CreateTwoVariantValueConstructor[ClickValue]()

func (val ClickValue) EnumType() WebEvent {
	return val
}

func inspect(event WebEvent) string {
	// We can now perform a type switch on the WebEvent type to determine which enum value we have.
	// We can then extract the value from the enum value similar to what could be done in Rust.
	switch e := event.(type) {
	case PageLoadValue:
		return fmt.Sprint(e())
	case PageUnloadValue:
		return fmt.Sprint(e())
	case KeyPressValue:
		return fmt.Sprintf("%c", e())
	case PasteValue:
		return fmt.Sprint(e())
	case ClickValue:
		return fmt.Sprint(e())
	default:
		return "Unknown"
	}
}

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
