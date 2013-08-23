package main

import "github.com/AndrewVos/o"

type SomeStruct struct {
	Something     string
	SomethingElse int
}

func main() {
	a := SomeStruct{
		Something:     "mferf",
		SomethingElse: 23131,
	}
	o.O(a)

	s := map[string]string{
		"hello": "o.O",
	}
	o.O(s)
	o.O("strings!")
}
