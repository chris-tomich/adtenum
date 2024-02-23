package rustyenums

import (
	"fmt"
	"testing"
)

type WebEvent Enum[WebEvent]

type PageLoadValue StaticValue[string]

var NewPageLoad func() PageLoadValue = CreateStaticValueConstructor[PageLoadValue]("PageLoad")

func (val PageLoadValue) EnumType() WebEvent {
	return val
}

type PageUnloadValue StaticValue[string]

var NewPageUnload func() PageUnloadValue = CreateStaticValueConstructor[PageUnloadValue]("PageUnload")

func (val PageUnloadValue) EnumType() WebEvent {
	return val
}

type KeyPressValue OneVariantValue[rune]

var NewKeyPress func(rune) KeyPressValue = CreateOneVariantValueConstructor[KeyPressValue]()

func (val KeyPressValue) EnumType() WebEvent {
	return val
}

type PasteValue OneVariantValue[string]

var NewPaste func(string) PasteValue = CreateOneVariantValueConstructor[PasteValue]()

func (val PasteValue) EnumType() WebEvent {
	return val
}

type ClickValue TwoVariantValue[int, int]

var NewClick func(int, int) ClickValue = CreateTwoVariantValueConstructor[ClickValue]()

func (val ClickValue) EnumType() WebEvent {
	return val
}

func inspect(event WebEvent) string {
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
