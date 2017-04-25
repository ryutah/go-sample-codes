package mapreduce

import (
	"reflect"
)

type List struct {
	values []interface{}
}

func (l *List) Add(v interface{}) *List {
	l.values = append(l.values, v)
	return l
}

func (l *List) Remove(v interface{}) bool {
	for i, val := range l.values {
		if reflect.DeepEqual(val, v) {
			newValues := make([]interface{}, len(l.values)-1)
			copy(newValues, l.values[:i])
			copy(newValues, l.values[i+1:])
			return true
		}
	}
	return false
}
