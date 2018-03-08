package main

import (
	"fmt"
)

type Num1 int

func (n Num1) Add(n2 Num1) Num1 {
	return n + n2
}

// If define Num2 as `type Num2 Num1`, it will occured compile error cause Num2 is not type Num1.
type Num2 = Num1

func main() {
	n1 := Num2(1)
	n2 := Num2(2)
	fmt.Println(n1.Add(n2))
}
