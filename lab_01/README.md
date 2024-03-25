# Отчет к лабораторной работе №1

**Студент**: Леонов В.В.  
**Группа**: М23-524  

## Задание 1

Использование языка `Go` для тестирования свойств программ имеет несколько преимуществ:

- в Go есть встроенная поддержка тестирования с помощью пакета `testing`. Это позволяет легко писать и запускать тесты прямо в коде.

- наличие обширного количества библиотек c для реализации тестирования на основе свойств (в том числе и стандартной библиотеке `testing/quick` для генерации наборов данных).

В рамках выполнения лабораторной работы был сделан выбор в пользу библиотеки `rapid`. В данном программном пакете свойства предсталяются в виде функций, для генерации данных используются специальные `generic generator-ы`, которые позволяют гибким образом инициализировать наборы данных, запуск тестов производится стандартными инструментами тестирования Go.

В учебных целях была написана функция сортировки пузырьком целочисленного списка и выполнена проверка ее спецификации. 

```go
func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func PropCmpSorts[K cmp.Ordered](sort1 func([]K), sort2 func([]K)) func(*rapid.T) {
	return func(t *rapid.T) {
		generator := rapid.Make[[]K]()
		arr1 := generator.Draw(t, "array")
		var arr2 []K
		copy(arr2, arr1)
		sort1(arr1)
		sort2(arr2)
		if reflect.DeepEqual(arr1, arr2) {
			t.Fatalf("Error")
		}
	}
}

func TestBubbleSort(t *testing.T) {
	property := PropCmpSorts(BubbleSort, sort.Ints)
	rapid.Check(t, property)
}
```

```
=== RUN   TestBubbleSort
    01_basic_test.go:23: [rapid] OK, passed 100 tests (620.1µs)
--- PASS: TestBubbleSort (0.00s)
```

## Задание 2

Экстенсиональное равенство двух функций: 

$$ f =_{Hom(A,B)} g := \forall x,y \in A (f(x)=g(x)). $$

Так к первому заданию была дополнительно реализована функция сортировки вставками и выполнено их экстенсиональное сравнение.

```go
func InsertSort(arr []int) {
	for i := 1; i <= len(arr); i++ {
		key := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = key
	}
}

func TestBubbleAndInsertSort(t *testing.T) {
	property := PropCmpSorts(BubbleSort, InsertSort)
	rapid.Check(t, property)
}
```

```log
=== RUN   TestBubbleAndInsertSort
    02_equality_test.go:68: [rapid] OK, passed 100 tests (0s)
--- PASS: TestBubbleAndInsertSort (0.00s)
```

Рассмотрим функции, которые выполняют одну и ту же задачу, однако в некоторых случаях результат отличается. Были описаны функции поиска элемента в целочисленном массиве (линейный, бинарный и интерполяционный) и выполнено их экстенсиональное сравнение.

```go
func InterpolationSearch(nums []int, elem int) int {
	l, r := 0, len(nums)-1
	for l <= r && elem >= nums[l] && elem <= nums[r] {
		if l == r {
			if nums[l] == elem {
				return l
			}
			return -1
		}
		pos := l + ((elem-nums[l])*(r-l))/(nums[r]-nums[l]+(nums[r]-nums[l]+1)%2)
		if nums[pos] == elem {
			return pos
		}
		if nums[pos] < elem {
			l = pos + 1
		} else {
			r = pos - 1
		}
	}

	return -1
}

func BinarySearch(nums []int, elem int) int {
	l, r := 0, len(nums)-1
	for l <= r {
		m := l + (r-l)/2
		if nums[m] == elem {
			return m
		} else if nums[m] < elem {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return -1
}

func LinearSearch(nums []int, elem int) int {
	for i, num := range nums {
		if num == elem {
			return i
		}
	}
	return -1
}

func PropCmpSearches[K cmp.Ordered](search1 func([]K, K) int, search2 func([]K, K) int) func(*rapid.T) {
	return func(t *rapid.T) {
		generatorArr := rapid.Make[[]K]()
		arr := generatorArr.Draw(t, "array")
		slices.Sort(arr)
		generatorInt := rapid.Make[K]()
		elem := generatorInt.Draw(t, "elem")
		if search1(arr, elem) != search2(arr, elem) {
			t.Fatalf("Error")
		}
	}
}

func TestLinearAndBinarySearch(t *testing.T) {
	property := PropCmpSearches(LinearSearch, BinarySearch)
	rapid.Check(t, property)
}

func TestInterpolationAndBinarySearch(t *testing.T) {
	property := PropCmpSearches(InterpolationSearch, BinarySearch)
	rapid.Check(t, property)
}

func TestInterpolationAndLinearSearch(t *testing.T) {
	property := PropCmpSearches(InterpolationSearch, LinearSearch)
	rapid.Check(t, property)
}
```

```log
=== RUN   TestLinearAndBinarySearch
    02_equality_test.go:23: [rapid] failed after 41 tests: Error
        To reproduce, specify -run="TestLinearAndBinarySearch" -rapid.failfile="testdata\\rapid\\TestLinearAndBinarySearch\\TestLinearAndBinarySearch-20240402192753-16132.fail" (or -rapid.seed=17929256316565368954)
        Failed test output:
    02_equality_test.go:14: [rapid] draw array: []int{0, -2, -2}
    02_equality_test.go:17: [rapid] draw elem: -2
    02_equality_test.go:20: Error
--- FAIL: TestLinearAndBinarySearch (0.14s)
=== RUN   TestInterpolationAndBinarySearch
    02_equality_test.go:38: [rapid] OK, passed 100 tests (1.1655ms)
--- PASS: TestInterpolationAndBinarySearch (0.00s)
=== RUN   TestInterpolationAndLinearSearch
    02_equality_test.go:53: [rapid] failed after 2 tests: Error
        To reproduce, specify -run="TestInterpolationAndLinearSearch" -rapid.failfile="testdata\\rapid\\TestInterpolationAndLinearSearch\\TestInterpolationAndLinearSearch-20240402192753-16132.fail" (or -rapid.seed=10458386935593656014)
        Failed test output:
    02_equality_test.go:44: [rapid] draw array: []int{0, 0, -1}
    02_equality_test.go:47: [rapid] draw elem: 0
    02_equality_test.go:50: Error
--- FAIL: TestInterpolationAndLinearSearch (0.06s)
```

Таким образом, для заданных функций наблюдается экстенсиональное неравенство.

## Задание 3
Рассмотрим функции $f:B\longrightarrow C$ и $g:A\longrightarrow B$. Тогда композиция будет определяться следующим образом:

$$
(f \circ g)(x) = f(g(x))
$$

Свойство ассоциативности композиции:

$$
compAss(A,B,C,D : Set) = \forall f \in Hom(C,D), g \in Hom(B,C), h \in Hom(A,B) \\
(f\circ g) \circ h =_{Hom(A,D)} f \circ (g \circ h)
$$

```go
func f(b int) int {
	return b * 2
}

func g(a int) int {
	return a * 3
}

func h(a int) int {
	return a + 1
}

func Composition(f1 func(int) int, f2 func(int) int) func(int) int {
	return func(a int) int {
		return f(g(a))
	}
}

func TestCompositionAssociativity(t *testing.T) {
	property := func(t *rapid.T) {
		generatorInt := rapid.Int()
		elem := generatorInt.Draw(t, "elem")

		cmp1 := Composition(Composition(f, g), h)
		cmp2 := Composition(f, Composition(g, h))

		if cmp1(elem) != cmp2(elem) {
			t.Fatalf("Error")
		}
	}
	rapid.Check(t, property)
}
```

```log
=== RUN   TestCompositionAssociativity
    03_composition_test.go:21: [rapid] OK, passed 100 tests (0s)
--- PASS: TestCompositionAssociativity (0.00s)
```

## Задание 4

Для функции идентичности $id_A(x:A):=x$ проверим ее свойства идентичности.

$$
idL (A,B:Set) \forall := f \in Hom(A,B) (id_B \circ f = f) \\
idR (A,B:Set) \forall := f \in Hom(A,B) (f \circ id_A = f)
$$

```go
func idInt(x int) int {
	return x
}

func TestIdentityL(t *testing.T) {
	property := func(t *rapid.T) {
		generatorInt := rapid.Int()
		elem := generatorInt.Draw(t, "elem")
		f := func(x int) int {
			return x * x
		}
		if f(idInt(elem)) != f(elem) {
			t.Fatalf("Error")
		}
	}
	rapid.Check(t, property)
}

func TestIdentityR(t *testing.T) {
	property := func(t *rapid.T) {
		generatorInt := rapid.Int()
		elem := generatorInt.Draw(t, "elem")
		f := func(x int) int {
			return x * x
		}
		if idInt(f(elem)) != f(elem) {
			t.Fatalf("Error")
		}
	}
	rapid.Check(t, property)
}
```

```log
=== RUN   TestIdentityL
    04_identity_test.go:20: [rapid] OK, passed 100 tests (0s)
--- PASS: TestIdentityL (0.00s)
=== RUN   TestIdentityR
    04_identity_test.go:34: [rapid] OK, passed 100 tests (0s)
--- PASS: TestIdentityR (0.00s)
```

## Задание 5
Терминальный объект обозначается $1$ и оснащается следующей спецификацией:

$$
unit(A:Set) := \forall f,g \in Hom(A, 1) (f =_{Hom(A,1)} g)
$$

В качестве терминального объекта в Go примем пустую структуру.

```go
type Unit struct{}

func TestTerminal(t *testing.T) {
	property := func(t *rapid.T) {
		generatorInt := rapid.Int()
		generatorString := rapid.String()
		elemInt := generatorInt.Draw(t, "elemInt")
		elemString := generatorString.Draw(t, "elemStr")
		f := func(x any) Unit { return Unit{} }
		g := func(x any) Unit { return Unit{} }

		if f(elemInt) != g(elemString) {
			t.Fatalf("Error")
		}
	}
	rapid.Check(t, property)
}
```

```log
=== RUN   TestTerminal
    05_terminal_test.go:22: [rapid] OK, passed 100 tests (614.8µs)
--- PASS: TestTerminal (0.00s)
```


