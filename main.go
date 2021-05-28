package main

import "github.com/kriangkrai/LinkedList/src"

func main() {

	l := src.LinkedList{}
	l.Insert(1)
	l.Insert(2)
	l.Insert(3)
	l.UpdateAt(1, 11)
	l.Print()
}
