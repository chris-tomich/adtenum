package adtenum

type Enum[E any] interface {
	EnumType() E
}

type StaticValue[A any] func() A
type OneVariantValue[A any] func() A
type TwoVariantValue[A any, B any] func() (A, B)
type ThreeVariantValue[A any, B any, C any] func() (A, B, C)
type FourVariantValue[A any, B any, C any, D any] func() (A, B, C, D)
type FiveVariantValue[A any, B any, C any, D any, E any] func() (A, B, C, D, E)
type SixVariantValue[A any, B any, C any, D any, E any, F any] func() (A, B, C, D, E, F)
type SevenVariantValue[A any, B any, C any, D any, E any, F any, G any] func() (A, B, C, D, E, F, G)
type EightVariantValue[A any, B any, C any, D any, E any, F any, G any, H any] func() (A, B, C, D, E, F, G, H)
type NineVariantValue[A any, B any, C any, D any, E any, F any, G any, H any, I any] func() (A, B, C, D, E, F, G, H, I)
type TenVariantValue[A any, B any, C any, D any, E any, F any, G any, H any, I any, J any] func() (A, B, C, D, E, F, G, H, I, J)

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

func CreateThreeVariantValueConstructor[V ~func() (A, B, C), A any, B any, C any]() func(A, B, C) V {
	return func(a A, b B, c C) V {
		return func() (A, B, C) {
			return a, b, c
		}
	}
}

func CreateFourVariantValueConstructor[V ~func() (A, B, C, D), A any, B any, C any, D any]() func(A, B, C, D) V {
	return func(a A, b B, c C, d D) V {
		return func() (A, B, C, D) {
			return a, b, c, d
		}
	}
}

func CreateFiveVariantValueConstructor[V ~func() (A, B, C, D, E), A any, B any, C any, D any, E any]() func(A, B, C, D, E) V {
	return func(a A, b B, c C, d D, e E) V {
		return func() (A, B, C, D, E) {
			return a, b, c, d, e
		}
	}
}

func CreateSixVariantValueConstructor[V ~func() (A, B, C, D, E, F), A any, B any, C any, D any, E any, F any]() func(A, B, C, D, E, F) V {
	return func(a A, b B, c C, d D, e E, f F) V {
		return func() (A, B, C, D, E, F) {
			return a, b, c, d, e, f
		}
	}
}

func CreateSevenVariantValueConstructor[V ~func() (A, B, C, D, E, F, G), A any, B any, C any, D any, E any, F any, G any]() func(A, B, C, D, E, F, G) V {
	return func(a A, b B, c C, d D, e E, f F, g G) V {
		return func() (A, B, C, D, E, F, G) {
			return a, b, c, d, e, f, g
		}
	}
}

func CreateEightVariantValueConstructor[V ~func() (A, B, C, D, E, F, G, H), A any, B any, C any, D any, E any, F any, G any, H any]() func(A, B, C, D, E, F, G, H) V {
	return func(a A, b B, c C, d D, e E, f F, g G, h H) V {
		return func() (A, B, C, D, E, F, G, H) {
			return a, b, c, d, e, f, g, h
		}
	}
}

func CreateNineVariantValueConstructor[V ~func() (A, B, C, D, E, F, G, H, I), A any, B any, C any, D any, E any, F any, G any, H any, I any]() func(A, B, C, D, E, F, G, H, I) V {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I) V {
		return func() (A, B, C, D, E, F, G, H, I) {
			return a, b, c, d, e, f, g, h, i
		}
	}
}

func CreateTenVariantValueConstructor[V ~func() (A, B, C, D, E, F, G, H, I, J), A any, B any, C any, D any, E any, F any, G any, H any, I any, J any]() func(A, B, C, D, E, F, G, H, I, J) V {
	return func(a A, b B, c C, d D, e E, f F, g G, h H, i I, j J) V {
		return func() (A, B, C, D, E, F, G, H, I, J) {
			return a, b, c, d, e, f, g, h, i, j
		}
	}
}
