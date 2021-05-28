package src

import (
	"errors"
	"fmt"
)

type Node struct {
	next  *Node
	value int
}
type LinkedList struct {
	head *Node
	len  int
}

func (l *LinkedList) Insert(val int) {
	n := Node{}
	n.value = val
	if l.len == 0 {
		l.head = &n
		l.len++
		return
	}
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.next == nil {
			ptr.next = &n
			l.len++
			return
		}
		ptr = ptr.next
	}
}

func (l *LinkedList) GetAt(pos int) *Node {
	ptr := l.head
	if pos < 0 {
		return ptr
	}
	if pos > (l.len - 1) {
		return nil
	}
	for i := 0; i < pos; i++ {
		ptr = ptr.next
	}
	return ptr
}
func (l *LinkedList) GetHead() int {
	if l.len > 0 {
		return l.head.value
	}
	return -1
}
func (l *LinkedList) GetTail() int {
	if l.len > 0 {
		n := l.GetAt(l.len - 1)
		return n.value
	}
	return -1
}
func (l *LinkedList) InsertAt(pos int, value int) {
	newNode := Node{}
	newNode.value = value

	if pos < 0 {
		return
	}
	if pos == 0 {
		l.head = &newNode
		l.len++
		return
	}
	if pos > l.len {
		fmt.Println("can not insert data ,position not found ")
		return
	}
	n := l.GetAt(pos)
	newNode.next = n
	prevNode := l.GetAt(pos - 1)
	prevNode.next = &newNode
	l.len++
}

func (l *LinkedList) Print() {
	if l.len == 0 {
		fmt.Println("No Node in List")
	}

	ptr := l.head
	for i := 0; i < l.len; i++ {
		if i == 0 || i == l.len {
			fmt.Printf("|%d|", ptr.value)
		} else {
			fmt.Printf("--->|%d|", ptr.value)
		}
		ptr = ptr.next
	}
	fmt.Println()
}
func (l *LinkedList) Search(val int) int {
	ptr := l.head
	for i := 0; i < l.len; i++ {
		if ptr.value == val {
			return i
		}
		ptr = ptr.next
	}
	return -1
}

func (l *LinkedList) UpdateAt(pos int, val int) error {
	if pos < 0 {
		return errors.New("Not found position")
	}
	if l.len == 0 {
		return errors.New("node is empty")
	}
	prevNode := l.GetAt(pos)
	if prevNode == nil {
		return errors.New("node is empty")
	} else {
		prevNode.value = val
	}
	return nil
}

//Delete last Node
func (l *LinkedList) Delete() error {
	if l.len == 0 {
		fmt.Println("No Node in list")
		return errors.New("no node in list")
	}
	prevNode := l.GetAt(l.len - 1)
	if prevNode == nil {
		fmt.Println("Node not found")
		return errors.New("Node not found")
	}
	prevNode.next = l.GetAt(l.len - 1).next
	l.len--
	return nil
}
func (l *LinkedList) DeleteAt(pos int) error {
	if pos < 0 {
		fmt.Println("position can not be negative")
		return errors.New("position can not be negative")
	}
	if l.len == 0 {
		fmt.Println("No Node in list")
		return errors.New("no node in list")
	}
	prevNode := l.GetAt(pos - 1)
	if prevNode == nil {
		fmt.Println("Node not found")
		return errors.New("Node not found")
	}
	prevNode.next = l.GetAt(pos).next
	l.len--
	return nil
}
func (l *LinkedList) DeleteVal(val int) error {
	ptr := l.head
	if l.len == 0 {
		fmt.Println("List is empty")
		return errors.New("List is empty")
	}
	for i := 0; i < l.len; i++ {
		if ptr.value == val {
			if i > 0 {
				prevNode := l.GetAt(i - 1)
				prevNode.next = l.GetAt(i).next
			} else {
				l.head = ptr.next
			}
			l.len--
			return nil
		}
		ptr = ptr.next
	}
	fmt.Println("Node not found")
	return errors.New("Node not found")
}
