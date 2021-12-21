package DS

import "fmt"

type LinkNode struct {
	Data int
	Next *LinkNode
}
type Linklist struct {
	Head   *LinkNode
	Length int
}

type LinklistMethod interface {
	Init(arr []int)
	Show()
	Reverse()
}

func (l Linklist) Show() {
	p := l.Head.Next
	for i := p; i != nil; i = i.Next {
		fmt.Print(i.Data, "->")
	}
	fmt.Println()
}

func (l *Linklist) Init(arr []int) {
	l.Head = new(LinkNode)
	p := l.Head
	for i := 0; i < len(arr); i++ {
		node := new(LinkNode)
		node.Data = arr[i]
		p.Next = node
		p = node
		l.Length++
	}
}

func (l *Linklist) Reverse() {
	p := l.Head.Next
	t := p.Next
	for t != nil {
		q := t.Next
		t.Next = l.Head.Next
		l.Head.Next = t
		t = q
	}
	p.Next = nil
}
