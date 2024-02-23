package rustyenums

import "testing"

type SingleVariantEnum[E any, A any] func() A
type DoubleVariantEnum[E any, A any, B any] func() (A, B)

func NewSingleVariantEnum[V ~func() A, A any]() func(A) V {
	return func(a A) V {
		return func() A {
			return a
		}
	}
}

func NewDoubleVariantEnum[V ~func() (A, B), A any, B any]() func(A, B) V {
	return func(a A, b B) V {
		return func() (A, B) {
			return a, b
		}
	}
}

type Enum[E any] interface {
	EnumType() E
}

type MyEnum Enum[MyEnum]

type MySingleValueEnum SingleVariantEnum[MyEnum, int]

func (val MySingleValueEnum) EnumType() MyEnum {
	return val
}

type MyDoubleValueEnum DoubleVariantEnum[MyEnum, int, string]

func (val MyDoubleValueEnum) EnumType() MyEnum {
	return val
}

type MySecondEnum Enum[MySecondEnum]

func TestEnums(t *testing.T) {
	var genericEnumValue MyEnum
	S := NewSingleVariantEnum[MySingleValueEnum]()
	D := NewDoubleVariantEnum[MyDoubleValueEnum]()
	genericEnumValue = S(42)

	switch e := genericEnumValue.(type) {
	case MySingleValueEnum:
		t.Log("1:", e())
	case MyDoubleValueEnum:
		a, b := e()
		t.Log("2:", a, b)
	default:
		t.Log("Unknown enum type")
	}

	genericEnumValue = D(42, "Hello")

	switch e := genericEnumValue.(type) {
	case MySingleValueEnum:
		t.Log("1:", e())
	case MyDoubleValueEnum:
		a, b := e()
		t.Log("2:", a, b)
	default:
		t.Log("Unknown enum type")
	}

	t.Log("Done!")
}
