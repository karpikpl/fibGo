package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int64 {
	pp := int64(0);
	p := int64(1);
	return func() int64 {
		tmp:= pp;
		pp = p;
		p = p + tmp;
		return tmp;
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 60; i++ {
		fmt.Println(f())
	}
}
