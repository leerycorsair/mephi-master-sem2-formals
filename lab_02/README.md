# Отчет к лабораторной работе №2

**Студент**:Леонов В.В.  
**Группа**: М23-524  

## Цель работы

Изучение свойств категорных произведений, сумм (копроизведений) и двойственности этих понятий. Изучение проявления данных свойств в программировании на примере типов пар и альтернативных значений.

## Задание №1. Реализация произведения 

Декартово произведение объектов $A \times B$ оснащается стрелками-проекциями $\pi_1:A \times B \longrightarrow A$ и $\pi_2:A \times B \longrightarrow B$. Для каждого объекта $X$ и пары совместимых $f:X\longrightarrow A$ и $g:X\longrightarrow B$ определена стрелка спаривания $<f,g>$.

Таким образом, выполняются следующие спецификации: 

$$ \text{pi}_1(X,A,B:\text{Set}):= f:\text{Hom}(X,A), g:\text{Hom}(X,B) \vdash \pi_1 \circ <f,g> = f$$

$$ \text{pi}_2(X,A,B:\text{Set}):= f:\text{Hom}(X,A), g:\text{Hom}(X,B) \vdash \pi_2 \circ <f,g> = g$$

Рассматриваемая задача: для матрицы получить ее определитель и ее ранг.

Для представления декартова произведения в `Go` предлагается использовать структуры. 


```go
type Pair[A, B any] struct {
	value1 A
	value2 B
}
```

Операции проекции ($\pi_1$ и $\pi_2$) будут выглядеть следующим образом:

```go
func First[A, B any](p Pair[A, B]) A {
	return p.value1
}

func Second[A, B any](p Pair[A, B]) B {
	return p.value2
}
```

Общая функция ($<f,g>$) будет выглядеть следующим образом:

```go
func BuildPair[SrcT, A, B any](
	s SrcT, f1 func(SrcT) A, f2 func(SrcT) B) Pair[A, B] {
	return Pair[A, B]{
		value1: f1(s),
		value2: f2(s),
	}
}
```

Функция вычисления определителя (функция $f$) для матрицы будет выглядеть следующим образом:

```go
func Determinant(mtr Matrix) float64 {
	if len(mtr) != len(mtr[0]) {
		panic(errors.New("Matrix is not square"))
	}

	if len(mtr) == 1 {
		return mtr[0][0]
	}

	var det float64
	for i, coeff := range mtr[0] {
		minor := minorMtr(mtr, 0, i)
		det += coeff * math.Pow(-1, float64(i)) * Determinant(minor)
	}
	return det
}

func minorMtr(mtr Matrix, row, col int) Matrix {
	size := len(mtr)
	minor := make(Matrix, size-1)
	for i := range minor {
		minor[i] = make([]float64, size-1)
	}

	for i := 0; i < size; i++ {
		if i == row {
			continue
		}
		for j := 0; j < size; j++ {
			if j == col {
				continue
			}
			minorRow := i
			if i > row {
				minorRow--
			}
			minorCol := j
			if j > col {
				minorCol--
			}
			minor[minorRow][minorCol] = mtr[i][j]
		}
	}
	return minor
}
```
Функция вычисления ранга (функция $g$) для матрицы будет выглядеть следующим образом:

```go
func Rank(mtr Matrix) int {
	var tmpMtr Matrix = copyMatrix(mtr)
	rows := len(tmpMtr)
	if rows == 0 {
		return 0
	}
	cols := len(tmpMtr[0])
	rank := 0
	for i := 0; i < rows; i++ {
		if allZero(tmpMtr[i]) {
			continue
		}
		pivotRow := i
		pivotCol := 0
		for pivotCol < cols && tmpMtr[pivotRow][pivotCol] == 0 {
			pivotCol++
		}
		if pivotCol == cols {
			continue
		}
		tmpMtr[pivotRow], tmpMtr[i] = tmpMtr[i], tmpMtr[pivotRow]
		for j := 0; j < rows; j++ {
			if j != i && tmpMtr[j][pivotCol] != 0 {
				scale := tmpMtr[j][pivotCol] / tmpMtr[i][pivotCol]
				for k := 0; k < cols; k++ {
					tmpMtr[j][k] -= scale * tmpMtr[i][k]
				}
			}
		}
		rank++
	}
	return rank
}

func allZero(row []float64) bool {
	for _, val := range row {
		if val < PRECISION {
			return false
		}
	}
	return true
}

func copyMatrix(src Matrix) Matrix {
	dst := make(Matrix, len(src))
	for i := range dst {
		dst[i] = make([]float64, len(src[i]))
	}
	for i := range src {
		copy(dst[i], src[i])
	}

	return dst
}
```

Для проверки спецификаций определим свойство:

```go
func PropProd[Src, R1, R2 any, P Pair[R1, R2]](
	f1 func(Src) R1,
	f2 func(Src) R2,
	cmp1 func(R1, R1) bool,
	cmp2 func(R2, R2) bool,
	generator *rapid.Generator[Src]) func(*rapid.T) {
	return func(t *rapid.T) {
		src := generator.Draw(t, "Src")
		resultF1 := f1(src)
		resultF2 := f2(src)
		resultF12 := BuildPair(src, f1, f2)
		if !cmp1(resultF1, First(resultF12)) || 
			!cmp2(resultF2, Second(resultF12)) {
			t.Fatalf("Error")
		}
	}
}
```

```go
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func TestPropProd(t *testing.T) {
	gen := rapid.Custom(func(t *rapid.T) Matrix {
		size := rapid.Uint16Range(1, 10).Draw(t, "mtr")
		mtr := make(Matrix, size)
		for i := 0; i < len(mtr); i++ {
			mtr[i] = make([]float64, size)
			for j := 0; j < len(mtr[i]); j++ {
				mtr[i][j] = roundFloat(rapid.Float64Range(-10, 10).Draw(t, "mtr value"), 9)
			}
		}
		return mtr
	})
	cmp1 := func(a, b float64) bool {
		return math.Abs(a-b) < PRECISION
	}
	cmp2 := func(a, b int) bool {
		return a == b
	}

	property := PropProd(Determinant, Rank, cmp1, cmp2, gen)
	rapid.Check(t, property)
}
```

В результате проведенного тестирования - приведенные спецификации были подтверждены.


## Задание №2. Реализация копроизведения


Копроизведение объектов $A+B$ оснащается стрелками-инъекциями $i_1:A\longrightarrow A+B$ и $i_2:B\longrightarrow A+B$, а также операцией ветвления и разбора случаев $\text{case}<f,g>$.

Для копроизведения определены следующие спецификации:

$$\text{inj}_1(A,B,X:\text{Set}) := f:\text{Hom}, g:\text{Hom}(B,X)\vdash(\text{case}<f,g>) \circ i_1 =f$$

$$\text{inj}_2(A,B,X:\text{Set}) := f:\text{Hom}, g:\text{Hom}(B,X)\vdash(\text{case}<f,g>) \circ i_2 =g$$

Рассматриваемая задача: от некоторого публичного API приходит ответ в виде хэш-таблицы или сообщения об ошибке - необходимо получить набор ключей хэш-таблицы в порядке возрастания соответствующих значений.

### Способ 1
Реализация типов $A$ и $B$, а также их копроизведения.

```go
type Either[A, B any] struct {
	isLeft bool
	left   A
	right  B
}
```
Реализация операций инъекции.

```go
func Left[A, B any](value A) Either[A, B] {
	return Either[A, B]{
		isLeft: true,
		left:   value,
	}
}

func Right[A, B any](value B) Either[A, B] {
	return Either[A, B]{
		isLeft: false,
		right:  value,
	}
}
```

Реализация функции ветвления ($\text{case}<f,g>$).

```go
func IsLeft[A, B any](e Either[A, B]) bool {
	return e.isLeft
}

func IsRight[A, B any](e Either[A, B]) bool {
	return !e.isLeft
}

func CaseFunction[A, B, R any](e Either[A, B], leftFunc func(a A) R, rightFunc func(b B) R) R {
	if IsLeft(e) {
		return leftFunc(e.left)
	} else {
		return rightFunc(e.right)
	}
}
```


Реализация обработки ошибки (функция $f$).

```go
func (e CustomErrorT) ParseError() ResultT {
	return ResultT{}
}
```

Реализация обработки хэш-таблицы (функция $g$).

```go
type byValue struct {
	keys   ResultT
	values HashTableT
}

func (bv byValue) Len() int {
	return len(bv.keys)
}

func (bv byValue) Less(i, j int) bool {
	if bv.values[bv.keys[i]] < bv.values[bv.keys[j]] {
		return true
	} else if bv.values[bv.keys[i]] == bv.values[bv.keys[j]] {
		return bv.keys[i] < bv.keys[j]
	} else {
		return false
	}
}

func (bv byValue) Swap(i, j int) {
	bv.keys[i], bv.keys[j] = bv.keys[j], bv.keys[i]
}

func (t HashTableT) ParseHash() ResultT {
	keys := make(ResultT, 0, len(t))
	for key := range t {
		keys = append(keys, key)
	}
	bv := byValue{keys: keys, values: t}
	sort.Sort(bv)
	return bv.keys
}
```

Для проверки спецификаций определим свойство:
```go
func PropSum[T1 any, T2 any, EType Either[T1, T2], ResT any](
	f1 func(T1) ResT,
	f2 func(T2) ResT,
	gen1 *rapid.Generator[T1],
	gen2 *rapid.Generator[T2],
	resCmp func(ResT, ResT) bool) func(*rapid.T) {
	return func(t *rapid.T) {
		val1 := gen1.Draw(t, "val1")
		val2 := gen2.Draw(t, "val2")

		if !resCmp(CaseFunction[T1, T2](Left[T1, T2](val1), f1, f2), f1(val1)) || 
			!resCmp(CaseFunction[T1, T2](Right[T1, T2](val2), f1, f2), f2(val2)) {
			t.Fatalf("Error")
		}
	}
}
```

```go
func TestPropSum(t *testing.T) {
	gen1 := rapid.Custom(func(t *rapid.T) HashTableT {
		ht := rapid.MapOfN(rapid.StringMatching(`[a-z]+`), rapid.IntRange(1, 10), 1, 100).Draw(t, "hashTable")
		return HashTableT(ht)
	})
	gen2 := rapid.Custom(func(t *rapid.T) CustomErrorT {
		ce := rapid.StringMatching(`[a-z]+`).Draw(t, "customError")
		return CustomErrorT(ce)
	})
	resCmp := func(a, b ResultT) bool {
		if len(a) != len(b) {
			return false
		}
		for i := 0; i < len(a); i++ {
			if a[i] != b[i] {
				return false
			}
		}
		return true
	}
	property := PropSum(HashTableT.ParseHash, CustomErrorT.ParseError, gen1, gen2, resCmp)
	rapid.Check(t, property)
}
```
В результате проведенного тестирования - приведенные спецификации были подтверждены.

### Способ 2

```go
type Either2[A, B any] interface {
	isEither2()
}

type Left2[T any] struct {
	value T
}

func (Left2[T]) isEither2() {}

type Right2[T any] struct {
	value T
}

func (Right2[T]) isEither2() {}

type FakeType[T any] struct {
	value T
}

func (FakeType[T]) isEither2() {}

func CaseFunction2[A, B, R any](obj Either2[A, B], leftFunc func(o A) R, rightFunc func(b B) R) R {
	switch v := obj.(type) {
	case Left2[A]:
		return leftFunc(v.value)
	case Right2[B]:
		return rightFunc(v.value)
	default:
		panic("unknown type")
	}
}
```

### Сравнение реализаций


Подход 1 (структура Either) обычно лучше подходит для большинства случаев использования благодаря своей типовой безопасности, простоте и производительности. Он обеспечивает обнаружение несоответствий типов на этапе компиляции, что имеет решающее значение для поддержания надежного и безопасного кода. Этот подход прост и эффективен, что делает его подходящим для ситуаций, когда нужно обрабатывать только два возможных типа.

Подход 2 (интерфейс Either2) более гибок и расширяем, но ценой этого являются типовая безопасность и производительность. Он лучше подходит для сценариев, когда необходима гибкость для работы с более широким диапазоном типов или когда иерархия типов может расти.

Учитывая типичную потребность в типовой безопасности и простоте, Подход 1 обычно является предпочтительным выбором.
