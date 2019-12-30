package listops

// IntList type
type IntList []int

type binFunc func(x, y int) int
type predFunc func(n int) bool
type unaryFunc func(x int) int

// Length method on IntList type
func (l IntList) Length() int {
	return len(l)
}

// Foldl method on IntList type
func (l IntList) Foldl(f binFunc, acc int) int {
	for _, i := range l {
		acc = f(acc, i)
	}
	return acc
}

// Foldr method on IntList type
func (l IntList) Foldr(f binFunc, acc int) int {
	for i := len(l) - 1; i >= 0; i-- {
		acc = f(l[i], acc)
	}
	return acc
}

// Filter metthod on IntList type
func (l IntList) Filter(p predFunc) IntList {
	f := make(IntList, 0, len(l))
	for _, v := range l {
		if p(v) {
			f = append(f, v)
		}
	}
	return f
}

// Map metthod on IntList type
func (l IntList) Map(u unaryFunc) IntList {
	f := make(IntList, len(l))
	for i, v := range l {
		f[i] = u(v)
	}
	return f
}

// Reverse metthod on IntList type
func (l IntList) Reverse() IntList {
	f := make(IntList, len(l))
	for i, v := range l {
		f[len(l)-1-i] = v
	}
	return f
}

// Append metthod on IntList type
func (l IntList) Append(m IntList) IntList {
	f := make(IntList, 0, len(l)+len(m))
	f = append(f, l...)
	f = append(f, m...)
	return f
}

// Concat metthod on IntList type
func (l IntList) Concat(a []IntList) IntList {
	f := make(IntList, 0, len(l)*len(a))
	f = append(f, l...)
	for _, m := range a {
		f = append(f, m...)
	}
	return f
}
