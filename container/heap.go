package main

type myHeap struct {
	l []int
}

func (m *myHeap) Len() int {
	return len(m.l)
}

func (m *myHeap) Less(i int, j int) bool {
	return m.l[i] < m.l[j]
}

func (m *myHeap) Swap(i int, j int) {
	m.l[i], m.l[j] = m.l[j], m.l[i]
}

func (m *myHeap) Push(x interface{}) {
	m.l = append(m.l, x.(int))
}

func (m *myHeap) Pop() interface{} {
	v := m.l[m.Len()-1]
	m.l = m.l[:m.Len()-1]
	return v
}
