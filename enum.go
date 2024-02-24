package adtenum

// Enum is a generic interface that all enumerations must use as the underlying type.
// It is implemented by the enumeration values to provide compile time type checking of enumeration values.
type Enum[E any] interface {
	EnumType() E
}

// ConstValue is the underlying type used for enumeration values that have a constant value.
type ConstValue[A any] func() A

// OneVariantValue is the underlying type used for enumeration values that have a single variant value.
type OneVariantValue[A any] func() A

// TwoVariantValue is the underlying type used for enumeration values that have two variant values.
type TwoVariantValue[A any, B any] func() (A, B)

// ThreeVariantValue is the underlying type used for enumeration values that have three variant values.
type ThreeVariantValue[A any, B any, C any] func() (A, B, C)

// FourVariantValue is the underlying type used for enumeration values that have four variant values.
type FourVariantValue[A any, B any, C any, D any] func() (A, B, C, D)

// FiveVariantValue is the underlying type used for enumeration values that have five variant values.
type FiveVariantValue[A any, B any, C any, D any, E any] func() (A, B, C, D, E)

// SixVariantValue is the underlying type used for enumeration values that have six variant values.
type SixVariantValue[A any, B any, C any, D any, E any, F any] func() (A, B, C, D, E, F)

// SevenVariantValue is the underlying type used for enumeration values that have seven variant values.
type SevenVariantValue[A any, B any, C any, D any, E any, F any, G any] func() (A, B, C, D, E, F, G)

// EightVariantValue is the underlying type used for enumeration values that have eight variant values.
type EightVariantValue[A any, B any, C any, D any, E any, F any, G any, H any] func() (A, B, C, D, E, F, G, H)

// NineVariantValue is the underlying type used for enumeration values that have nine variant values.
type NineVariantValue[A any, B any, C any, D any, E any, F any, G any, H any, I any] func() (A, B, C, D, E, F, G, H, I)

// TenVariantValue is the underlying type used for enumeration values that have ten variant values.
type TenVariantValue[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any] func() (A, B, C, D, E, F, G, H, I, J)

// CreateConstValueConstructor creates a constructor that can be used to create instances of enumeration values that have a constant value.
func CreateConstValueConstructor[V ~func() A, A any](val A) func() V {
	return func() V {
		return func() A {
			return val
		}
	}
}

// CreateOneVariantValueConstructor creates a constructor that can be used to create instances of enumeration values that have a single variant value.
func CreateOneVariantValueConstructor[V ~func() A, A any]() func(A) V {
	return func(a A) V {
		return func() A {
			return a
		}
	}
}

// CreateTwoVariantValueConstructor creates a constructor that can be used to create instances of enumeration values that have two variant values.
func CreateTwoVariantValueConstructor[V ~func() (A, B), A any, B any]() func(A, B) V {
	return func(a A, b B) V {
		return func() (A, B) {
			return a, b
		}
	}
}

// CreateThreeVariantValueConstructor creates a constructor that can be used to create instances of enumeration values that have three variant values.
func CreateThreeVariantValueConstructor[V ~func() (A, B, C), A any, B any, C any]() func(A, B, C) V {
	return func(a A, b B, c C) V {
		return func() (A, B, C) {
			return a, b, c
		}
	}
}

// CreateFourVariantValueConstructor creates a constructor that can be used to create instances of enumeration values that have four variant values.
func CreateFourVariantValueConstructor[V ~func() (A, B, C, D), A any, B any, C any, D any]() func(A, B, C, D) V {
	return func(a A, b B, c C, d D) V {
		return func() (A, B, C, D) {
			return a, b, c, d
		}
	}
}

// CreateFiveVariantValueConstructor creates a constructor that can be used to create instances of enumeration values that have five variant values.
func CreateFiveVariantValueConstructor[V ~func() (A, B, C, D, E), A any, B any, C any, D any, E any]() func(A, B, C, D, E) V {
	return func(a A, b B, c C, d D, e E) V {
		return func() (A, B, C, D, E) {
			return a, b, c, d, e
		}
	}
}

// CreateSixVariantValueConstructor creates a constructor that can be used to create instances of enumeration values that have six variant values.
func CreateSixVariantValueConstructor[V ~func() (A, B, C, D, E, F), A any, B any, C any, D any, E any, F any]() func(A, B, C, D, E, F) V {
	return func(a A, b B, c C, d D, e E, f F) V {
		return func() (A, B, C, D, E, F) {
			return a, b, c, d, e, f
		}
	}
}

// CreateSevenVariantValueConstructor creates a constructor that can be used to create instances of enumeration values that have seven variant values.
func CreateSevenVariantValueConstructor[V ~func() (A, B, C, D, E, F, G), A any, B any, C any, D any, E any, F any, G any]() func(A, B, C, D, E, F, G) V {
	return func(a A, b B, c C, d D, e E, f F, g G) V {
		return func() (A, B, C, D, E, F, G) {
			return a, b, c, d, e, f, g
		}
	}
}

// CreateEightVariantValueConstructor creates a constructor that can be used to create instances of enumeration values that have eight variant values.
func CreateEightVariantValueConstructor[V ~func() (A, B, C, D, E, F, G, H), A any, B any, C any, D any, E any, F any, G any, H any]() func(A, B, C, D, E, F, G, H) V {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) V {
		return func() (A, B, C, D, E, F, G, H) {
			return a, b, c, d, e, f, g, h
		}
	}
}

// CreateNineVariantValueConstructor creates a constructor that can be used to create instances of enumeration values that have nine variant values.
func CreateNineVariantValueConstructor[V ~func() (A, B, C, D, E, F, G, H, I), A any, B any, C any, D any, E any, F any, G any, H any, I any]() func(A, B, C, D, E, F, G, H, I) V {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) V {
		return func() (A, B, C, D, E, F, G, H, I) {
			return a, b, c, d, e, f, g, h, i
		}
	}
}

// CreateTenVariantValueConstructor creates a constructor that can be used to create instances of enumeration values that have ten variant values.
func CreateTenVariantValueConstructor[V ~func() (A, B, C, D, E, F, G, H, I, J), A any, B any, C any, D any, E any, F any, G any, H any, I any, J any]() func(A, B, C, D, E, F, G, H, I, J) V {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I, j J) V {
		return func() (A, B, C, D, E, F, G, H, I, J) {
			return a, b, c, d, e, f, g, h, i, j
		}
	}
}
