package rustyenums

type StaticValue[A any] func() A
type OneVariantValue[A any] func() A
type TwoVariantValue[A any, B any] func() (A, B)

func CreateStaticValueConstructor[V ~func() A, A any](val A) func() V {
	return func() V {
		return func() A {
			return val
		}
	}
}

func CreateOneVariantValueConstructor[V ~func() A, A any]() func(A) V {
	return func(a A) V {
		return func() A {
			return a
		}
	}
}

func CreateTwoVariantValueConstructor[V ~func() (A, B), A any, B any]() func(A, B) V {
	return func(a A, b B) V {
		return func() (A, B) {
			return a, b
		}
	}
}

type Enum[E any] interface {
	EnumType() E
}
