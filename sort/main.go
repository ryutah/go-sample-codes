package main

import (
	"fmt"
	"sort"
	"strings"
)

func Sort() {
	a, b := "0000", "0"
	fmt.Println(strings.Compare(a, b))
	strs := []string{"a-00010,csv", "a-00002.csv", "a-00000.csv", "a-00001.csv"}
	fmt.Println(strs)
	strs[0], strs[1] = strs[1], strs[0]
	fmt.Println(strs)
	sort.Strings(strs)
	fmt.Println(strs)
}

type sortHoge struct {
	Name string
}

func (s *sortHoge) String() string {
	return fmt.Sprintf("Name: %s", s.Name)
}

type sortHoges []*sortHoge

func (h sortHoges) Len() int {
	return len(h)
}

func (h sortHoges) Less(i int, j int) bool {
	return strings.Compare(h[i].Name, h[j].Name) < 0
}

func (h sortHoges) Swap(i int, j int) {
	h[i], h[j] = h[j], h[i]
}

func SortStruct() {
	hs := sortHoges{
		&sortHoge{"0100006"},
		&sortHoge{"0000004"},
		&sortHoge{"0000001"},
		&sortHoge{"0000002"},
		&sortHoge{"0000001"},
		&sortHoge{"0000000"},
		&sortHoge{"0000006"},
	}
	fmt.Println(hs)
	sort.Sort(hs)
	fmt.Println(hs)
}
