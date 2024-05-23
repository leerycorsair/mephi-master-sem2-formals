

# Отчет к лабораторной работе №3

**Студент**:Леонов В.В.  
**Группа**: М23-524  

## Цель работы

Изучение свойств функторов и естественных преобразований. Изучение проявления данных свойств в программировании на примере типов необязательных значений и списков.

## Функтор

Пусть $\mathbb{C}$ и $\mathbb{D}$ - категории, функтором $f: \mathbb{C}\longrightarrow \mathbb{D}$ называется пара отображений $F_o$ (объектов $\mathbb{C}$ в объекты $\mathbb{D}$), $F_m$ (морфизмов $\mathbb{C}$ в морфизмы $\mathbb{D}$).

Пусть $А$ - объект в $\mathbb{D}$
$$F_o(A) = F(A)$$
$f$ - морфизм в $\mathbb{C}$
$$F(f) = F_m(f)$$

Пара таких отображений, что 
1) $\forall A,B,C  \in \mathbb{C} \\ \forall f: B\longrightarrow C ~ \forall g: A\longrightarrow B \\ (F(f\circ g) = F(f) \circ F(g))$

2) $F_m (id_A) = id_{F_o(A)}$

<!-- Попробовать придумать пример, где данная аксиома работать не будет. -->


Примеры функторов:
- списки;
- опциональные значения.

На примере списков: 
$$F_o(-) = list(-)$$
$$\text{Int} \longrightarrow^{F_o} list(\text{Int})$$

$$(\text{String}\longrightarrow \text{Int}) \longrightarrow^{F_m} (list(\text{String}) \longrightarrow list(\text{Int}))$$ 

$$F_m \rightsquigarrow map $$

## Естественные преобразования
Пусть $F(\mathbb{A}\longrightarrow \mathbb{B})$ и $G(\mathbb{A}\longrightarrow \mathbb{B})$ - функторы, тогда $\eta$ - естественное преобразование (из функора в функтор), если:

$$\forall f: A \longrightarrow B \\ G(f) \circ \eta_A = \eta_B \circ F(f)$$

## Ход работы

Определим функтор для опциональных значений Maybe.

```go
type Maybe[T any] struct {
	value *T
	isNil bool
}

func Just[T any](value T) Maybe[T] {
	return Maybe[T]{value: &value, isNil: false}
}

func Nothing[T any]() Maybe[T] {
	return Maybe[T]{value: nil, isNil: true}
}

func MapMaybe[T1, T2 any](m Maybe[T1], f func(T1) T2) Maybe[T2] {
	if m.isNil {
		return Nothing[T2]()
	}
	return Just(f(*m.value))
}
```

Для данного функтора следует проверить свойства идентичности и композиции.

```go
func PropMaybeId[T1, T2 any](
	f func(T1) T2,
	cmp func(T2, T2) bool,
	generator *rapid.Generator[T1]) func(*rapid.T) {
	return func(t *rapid.T) {
		switch rapid.IntRange(0, 1).Draw(t, "switch") {
		case 0:
			maybeObj := Nothing[T1]()
			if !MaybeCmp(MapMaybe(maybeObj, f), Nothing[T2](), cmp) {
				t.Fatalf("Error")
			}
		case 1:
			src := generator.Draw(t, "src")
			maybeObj := Just(src)
			if !MaybeCmp(MapMaybe(maybeObj, f), Just(f(src)), cmp) {
				t.Fatalf("Error")
			}
		}
	}
}

func PropMaybeComp[T1, T2, T3 any](
	f1 func(T1) T2,
	f2 func(T2) T3,
	cmp func(T3, T3) bool,
	generator *rapid.Generator[T1]) func(*rapid.T) {
	return func(t *rapid.T) {
		switch rapid.IntRange(0, 1).Draw(t, "switch") {
		case 0:
			maybeObj := Nothing[T1]()
			if !MaybeCmp(MapMaybe(MapMaybe(maybeObj, f1), f2), Nothing[T3](), cmp) {
				t.Fatalf("Error")
			}
		case 1:
			src := generator.Draw(t, "src")
			maybeObj := Just(src)
			if !MaybeCmp(MapMaybe(MapMaybe(maybeObj, f1), f2), Just(f2(f1(src))), cmp) {
				t.Fatalf("Error")
			}
		}
	}
}
```

Тестирование свойств описывается следующими сценариями.

```go
func TestPropMaybeId(t *testing.T) {
	gen := rapid.IntRange(0, 100)

	intCmp := func(a, b int) bool {
		return a == b
	}
	property := functors.PropMaybeId(sqr, intCmp, gen)
	rapid.Check(t, property)
}

func TestPropMaybeComp(t *testing.T) {
	gen := rapid.IntRange(0, 100)

	cmp := func(a, b float64) bool {
		return math.Abs(a-b) < PRECISION
	}
	property := functors.PropMaybeComp(sqr, double, cmp, gen)
	rapid.Check(t, property)
}
```

В результате проведенного тестирования - приведенные спецификации были подтверждены.

Также определим функтор для списка и опишем естественное преобразование из списков в опциональные значения.

```go
type List[T any] []T

func MapList[T1, T2 any](l List[T1], f func(T1) T2) List[T2] {
	newList := make(List[T2], len(l))
	for i, elem := range l {
		newList[i] = f(elem)
	}
	return newList
}

func FromListToMaybe[T any](l List[T]) Maybe[T] {
	if len(l) == 0 {
		return Nothing[T]()
	}
	return Just(l[0])
}
```

Для данного естественного преобразования следует проверить свойства идентичности и композиции.

```go
func PropNtId[T any](
	generator *rapid.Generator[List[T]]) func(*rapid.T) {
	return func(t *rapid.T) {
		src := generator.Draw(t, "src")
		dst := FromListToMaybe(src)
		if reflect.ValueOf(dst).Type() != reflect.TypeFor[Maybe[T]]() {
			t.Fatalf("Error")
		}
	}
}

func PropNtComp[T1, T2 any](
	f func(T1) T2,
	cmp func(T2, T2) bool,
	generator *rapid.Generator[List[T1]]) func(*rapid.T) {
	return func(t *rapid.T) {
		src := generator.Draw(t, "src")
		if !MaybeCmp(MapMaybe(FromListToMaybe(src), f), FromListToMaybe(MapList(src, f)), cmp) {
			t.Fatalf("Error")
		}
	}
}
```

Тестирование свойств описывается следующими сценариями.

```go
func TestPropNtId(t *testing.T) {
	gen := rapid.Custom(func(t *rapid.T) functors.List[int] {
		l := rapid.SliceOf(rapid.IntRange(0, 100)).Draw(t, "list")
		return functors.List[int](l)
	})
	property := functors.PropNtId(gen)
	rapid.Check(t, property)
}

func TestPropNtComp(t *testing.T) {
	gen := rapid.Custom(func(t *rapid.T) functors.List[int] {
		l := rapid.SliceOf(rapid.IntRange(0, 100)).Draw(t, "list")
		return functors.List[int](l)
	})

	cmp := func(a, b float64) bool {
		return math.Abs(a-b) < PRECISION
	}
	property := functors.PropNtComp(double, cmp, gen)
	rapid.Check(t, property)
}
```

В результате проведенного тестирования - приведенные спецификации были подтверждены.
