package strain

// Ints is a collection of ints
type Ints []int

// Lists is a collection of []int
type Lists [][]int

// Strings is a collection of strings
type Strings []string

// Keep filters the contents of i, keeping
// only those that make f return true
func (i Ints) Keep(f func(int) bool) Ints {
	if len(i) == 0 {
		return nil
	}
	res := make(Ints, len(i))
	count := 0
	for _, v := range i {
		if f(v) {
			res[count] = v
			count++
		}
	}
	return res[:count]
}

// Discard filters the contents of i, removing
// those that make f return true
func (i Ints) Discard(f func(int) bool) Ints {
	return i.Keep(func(n int) bool { return !f(n) })
}

// Keep filters the contents of i, keeping
// only those that make f return true
func (l Lists) Keep(f func([]int) bool) Lists {
	if len(l) == 0 {
		return nil
	}
	res := make(Lists, len(l))
	count := 0
	for _, v := range l {
		if f(v) {
			res[count] = v
			count++
		}
	}
	return res[:count]
}

// Keep filters the contents of i, keeping
// only those that make f return true
func (s Strings) Keep(f func(string) bool) Strings {
	if len(s) == 0 {
		return nil
	}
	res := make(Strings, len(s))
	count := 0
	for _, v := range s {
		if f(v) {
			res[count] = v
			count++
		}
	}
	return res[:count]
}
