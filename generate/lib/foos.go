package lib

type Foos []*Foo

func (fs Foos) ForEach(foreach func(*Foo)) {
	for _, f := range fs {
		foreach(f)
	}
}

func (fs Foos) Filter(filter func(*Foo) bool) Foos {
	newFoos := make(Foos, 0, len(fs))
	for _, f := range fs {
		if filter(f) {
			newFoos = append(newFoos, f)
		}
	}
	return newFoos
}

func (fs Foos) MapToString(m func(*Foo) string) []string {
	s := make([]string, 0, len(fs))
	for _, f := range fs {
		s = append(s, m(f))
	}
	return s
}

func (fs Foos) ReduceToString(reduce func(string, *Foo) string) string {
	var s string
	for _, f := range fs {
		s = reduce(s, f)
	}
	return s
}
