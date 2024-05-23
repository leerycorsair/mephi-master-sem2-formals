# Лабораторная работа №4

**Студент**:Леонов В.В.  
**Группа**: М23-524  

## Цель работы

Изучение свойств монад. Изучение проявления данных свойств в программировании.

## Ход работы

Пусть есть композиция функторов $F\circ G$, тогда если $F$ - эндофунктор ($F:\mathbb{C} \longrightarrow \mathbb{C}$), 
то $F^n = F\circ F \circ F...$ 

$Id$ - идентичный функтор ($Id(A) = A, Id(f) = f$)

Монада - $<F, \eta, \mu>$, где $F$ - функтор, $\eta= Id\longrightarrow F$, $\mu = F^2 \longrightarrow F$

$$muComp⟨F⟩(A : Set) := \mu_F(A) \circ F(\mu_F(A)) = \mu_F(A) \circ \mu_F(F(A))$$

$$muEtaCompR⟨F⟩(A : Set) := \mu_F(A) \circ F(\eta_F(A)) = id_F(A)$$

$$muEtaCompL⟨F⟩(A : Set) := \mu_F(A) \circ \eta_F(F(A)) = id_F(A)$$

В качестве примера рассмотрим списки.


```go
type List[T any] []T

func Pure[T any](value T) List[T] {
	return List[T]{value}
}

func Join[T any](list List[List[T]]) List[T] {
	var result List[T]
	for _, innerList := range list {
		result = append(result, innerList...)
	}
	return result
}

func MapList[T1, T2 any](l List[T1], f func(T1) T2) List[T2] {
	newList := make(List[T2], len(l))
	for i, elem := range l {
		newList[i] = f(elem)
	}
	return newList
}
```

Опишем приведенные ранее свойства.

```go
func PropComp[T any](
	generator *rapid.Generator[List[List[List[T]]]]) func(*rapid.T) {
	return func(t *rapid.T) {
		src := generator.Draw(t, "src")
		left := Join(MapList(src, Join[T]))
		right := Join(Join(src))
		if !reflect.DeepEqual(left, right) {
			t.Fatalf("Error")
		}
	}
}

func PropCompR[T any](
	generator *rapid.Generator[List[T]]) func(*rapid.T) {
	return func(t *rapid.T) {
		src := generator.Draw(t, "src")
		left := Join(MapList(src, Pure[T]))
		right := src
		if !reflect.DeepEqual(left, right) {
			t.Fatalf("Error")
		}
	}
}

func PropCompL[T any](
	generator *rapid.Generator[List[T]]) func(*rapid.T) {
	return func(t *rapid.T) {
		src := generator.Draw(t, "src")
		left := Join(Pure(src))
		right := src
		if !reflect.DeepEqual(left, right) {
			t.Fatalf("Error, left %v, right %v", left, right)
		}
	}
}
```

Тестирование свойств описывается следующими сценариями.


```go
func TestPropCompL(t *testing.T) {
	gen := rapid.Custom(func(t *rapid.T) List[int] {
		l := rapid.SliceOfN(rapid.IntRange(0, 100), 1, 100).Draw(t, "list")
		return List[int](l)
	})

	property := PropCompL(gen)
	rapid.Check(t, property)
}

func TestPropCompR(t *testing.T) {
	gen := rapid.Custom(func(t *rapid.T) List[int] {
		l := rapid.SliceOfN(rapid.IntRange(0, 100), 1, 100).Draw(t, "list")
		return List[int](l)
	})

	property := PropCompR(gen)
	rapid.Check(t, property)
}

func TestPropComp(t *testing.T) {
	gen := rapid.Custom(func(t *rapid.T) List[List[List[int]]] {
		l1Gen := rapid.SliceOfN(rapid.IntRange(0, 100), 1, 100)
		l2Gen := rapid.SliceOfN(l1Gen, 1, 100)
		l3Gen := rapid.SliceOfN(l2Gen, 1, 100)

		var result List[List[List[int]]]
		input := l3Gen.Draw(t, "list")
		for _, outer := range input {
			var middleList List[List[int]]
			for _, middle := range outer {
				var innerList List[int]
				for _, inner := range middle {
					innerList = append(innerList, inner)
				}
				middleList = append(middleList, innerList)
			}
			result = append(result, middleList)
		}

		return result
	})

	property := PropComp(gen)
	rapid.Check(t, property)
}
```

В результате проведенного тестирования - приведенные спецификации были подтверждены.
