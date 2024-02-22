package rustyenums

import "testing"

type Enum[E any] any

func NewSingleVariantEnum[E Enum[any], A any, V ~func(a E) A]() func(A) V {
	return func(a A) V {
		return func(e E) A {
			return a
		}
	}
}

type MyEnum Enum[MyEnum]

type MySingleValueEnum[E MyEnum, A any] func(a E) A

var S1 = NewSingleVariantEnum[MyEnum, int, MySingleValueEnum[MyEnum, int]]()

type MySecondEnum Enum[MySecondEnum]

var S2 = NewSingleVariantEnum[MySecondEnum, int, MySingleValueEnum[MySecondEnum, int]]()

func TestEnums(t *testing.T) {
	var genericEnumValue MyEnum = S1(42)
	var genericEnumValue2 MySecondEnum = S2(53)

	switch e := genericEnumValue.(type) {
	case MySingleValueEnum[MyEnum, int]:
		t.Log("1:", e(genericEnumValue))
	case MySingleValueEnum[MySecondEnum, int]:
		t.Log("2:", e(genericEnumValue))
	default:
		t.Log("Unknown enum type")
	}

	switch e := genericEnumValue2.(type) {
	case MySingleValueEnum[MyEnum, int]:
		t.Log("1:", e(genericEnumValue2))
	case MySingleValueEnum[MySecondEnum, int]:
		t.Log("2:", e(genericEnumValue2))
	default:
		t.Log("Unknown enum type")
	}

	t.Log("Done!")
}
