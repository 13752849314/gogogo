package DS

import (
	"fmt"
	"os"
)

type LinkNode struct {
	Data int
	Next *LinkNode
}
type Linklist struct {
	head   *LinkNode
	length int
}

type LinklistMethod interface {
	Init(arr []int)
	Show()
	Reverse()
	GetLength() int
	GetHead() *LinkNode
	Empty() bool
	GetNode(n int) *LinkNode
	Insert(n, data int) bool
	Delete(n int) int
	Push(data int)
	Pop() int
}

func (l *Linklist) Pop() int {
	return l.Delete(1)
}

func (l *Linklist) Push(data int) {
	l.Insert(l.length+1, data)
}

func (l *Linklist) Delete(n int) int {
	if n >= 1 && n <= l.length {
		p := l.GetNode(n - 1)
		q := p.Next
		p.Next = q.Next
		l.length--
		return q.Data
	} else {
		fmt.Println("Out of index!")
		os.Exit(-1)
		return -1
	}
}

func (l *Linklist) Insert(n, data int) bool {
	if n >= 1 && n <= l.length+1 {
		p := l.GetNode(n - 1)
		node := &LinkNode{Data: data, Next: nil}
		node.Next = p.Next
		p.Next = node
		l.length++
		return true
	} else {
		fmt.Println("Out of index!")
		os.Exit(-1)
		return false
	}
}

func (l Linklist) GetNode(n int) *LinkNode {
	if n == 0 {
		return l.head
	} else if n > 0 && n <= l.GetLength() {
		p := l.GetHead()
		for i := 0; i < n; i++ {
			p = p.Next
		}
		return p
	}
	fmt.Println("Out of index!")
	os.Exit(-1)
	return nil
}

func (l Linklist) Empty() bool {
	return l.head.Next == nil
}

func (l Linklist) GetHead() *LinkNode {
	return l.head
}

func (l Linklist) GetLength() int {
	return l.length
}

func (l Linklist) Show() {
	p := l.head.Next
	for i := p; i != nil; i = i.Next {
		fmt.Print(i.Data, "->")
	}
	fmt.Println()
}

func (l *Linklist) Init(arr []int) {
	l.head = new(LinkNode)
	p := l.head
	for i := 0; i < len(arr); i++ {
		node := new(LinkNode)
		node.Data = arr[i]
		p.Next = node
		p = node
		l.length++
	}
}

func (l *Linklist) Reverse() {
	p := l.head.Next
	t := p.Next
	for t != nil {
		q := t.Next
		t.Next = l.head.Next
		l.head.Next = t
		t = q
	}
	p.Next = nil
}
