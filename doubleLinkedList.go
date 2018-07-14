package main

import "fmt"

type Node struct {
	next, prev *Node
	list       *List
	Value      interface{}
}

type List struct {
	root Node
	len  int
}

// 初始化或者清空链表
func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

// 新建链表
func New() *List {
	return new(List).Init()
}

func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// 获取链表长度
func (l *List) Len() int {
	return l.len
}

func (l *List) Front() *Node {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// 链表插入节点
func (l *List) insert(e, at *Node) *Node {
	n := at.next
	at.next = e
	e.prev = at
	e.next = n
	n.prev = e
	e.list = l
	l.len ++
	return e
}

// 链表插入具体值
func (l *List) insertValue(v interface{}, at *Node) *Node {
	return l.insert(&Node{Value: v}, at)
}

// 元素插入链表尾
func (l *List) PushBack(v interface{}) *Node {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

// 元素插入链表头
func (l *List) PushFront(v interface{}) *Node {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}

// 访问Node的后继
func (e *Node) Next() *Node {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

// 访问Node的前驱
func (e *Node) Prev() *Node {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

func main() {
	l := new(List).Init()
	l.PushBack(100)
	l.PushBack(101)
	l.PushBack(102)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
