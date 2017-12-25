package main

import (
	"math/rand"
	"sort"
)

func quickSort(f []foo) {
	var qSort func(f []foo, begin, end int)
	qSort = func(f []foo, begin, end int) {
		if begin >= end {
			return
		}
		curBegin, curEnd := begin, end
		for curBegin <= curEnd {
			for f[curBegin].A < f[begin].A {
				curBegin++
			}
			for f[curEnd].A > f[begin].A {
				curEnd--
			}
			if curBegin <= curEnd {
				f[curBegin], f[curEnd] = f[curEnd], f[curBegin]
				curBegin++
				curEnd--
			}
		}

		qSort(f, begin, curEnd)
		qSort(f, curBegin, end)
	}

	qSort(f, 0, len(f)-1)
}

type foo struct {
	A int
}

type lFoos []foo

func (l lFoos) Len() int {
	return len(l)
}

func (l lFoos) Less(i int, j int) bool {
	return l[i].A < l[j].A
}

func (l lFoos) Swap(i int, j int) {
	l[i], l[j] = l[j], l[i]
}

func sortByDefault(size int) {
	foos := make([]foo, 0, size)
	for i := 0; i < size; i++ {
		foos = append(foos, foo{A: rand.Int()})
	}
	sort.Slice(foos, func(i, j int) bool {
		return foos[i].A < foos[j].A
	})
}

func sortByInterface(size int) {
	foos := make(lFoos, 0, size)
	for i := 0; i < size; i++ {
		foos = append(foos, foo{A: rand.Int()})
	}
	sort.Sort(foos)
}

func sortByManual(size int) {
	foos := make([]foo, 0, size)
	for i := 0; i < size; i++ {
		foos = append(foos, foo{A: rand.Int()})
	}
	quickSort(foos)
}

func sortInt(size int) {
	ints := make([]int, 0, size)
	for i := 0; i < size; i++ {
		ints = append(ints, rand.Int())
	}
	sort.Ints(ints)
}

func main() {
	sortByManual(10)
}
