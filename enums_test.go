package rustyenums

import "testing"

type StaticEnum[A any] func() A
type SingleVariantEnum[A any] func() A
type DoubleVariantEnum[A any, B any] func() (A, B)

func NewStaticEnum[V ~func() A, A any](val A) func() V {
	return func() V {
		return func() A {
			return val
		}
	}
}

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

type MyStaticValueEnum StaticEnum[string]

func (val MyStaticValueEnum) EnumType() MyEnum {
	return val
}

type MySingleValueEnum SingleVariantEnum[int]

func (val MySingleValueEnum) EnumType() MyEnum {
	return val
}

type MyDoubleValueEnum DoubleVariantEnum[int, string]

func (val MyDoubleValueEnum) EnumType() MyEnum {
	return val
}

type MySecondEnum Enum[MySecondEnum]

func TestEnums(t *testing.T) {
	var genericEnumValue MyEnum
	A := NewStaticEnum[MyStaticValueEnum]("Hello, World!")
	S := NewSingleVariantEnum[MySingleValueEnum]()
	D := NewDoubleVariantEnum[MyDoubleValueEnum]()

	genericEnumValue = A()

	switch e := genericEnumValue.(type) {
	case MyStaticValueEnum:
		t.Log("0:", e())
	case MySingleValueEnum:
		t.Log("1:", e())
	case MyDoubleValueEnum:
		a, b := e()
		t.Log("2:", a, b)
	default:
		t.Log("Unknown enum type")
	}

	genericEnumValue = S(42)

	switch e := genericEnumValue.(type) {
	case MyStaticValueEnum:
		t.Log("0:", e())
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
	case MyStaticValueEnum:
		t.Log("0:", e())
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
