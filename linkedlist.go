package dataStruct

import "fmt"

type Node struct {
	Next *Node // 다음 노드를 기억한다.
	Prev *Node // 전의 노드를 기억한다.
	Val  int
}

type LinkedList struct {
	Root *Node
	Tail *Node
}

func (l *LinkedList) AddNode(Val int) {
	if l.Root == nil { // 맨처음엔 Root와 Tail 둘다 nil 이기때문에
		l.Root = &Node{Val: Val} // 새로운 노드를 만듦
		l.Tail = l.Root          // 하나만 있는 상태에서 Root와 Tail은 같음
		return
	}
	l.Tail.Next = &Node{Val: Val} // Tail 의 다음 노드를 만듦
	Prev := l.Tail                // 이전 Tail을 기억함
	l.Tail = l.Tail.Next          // Tail을 다음노드로 넣어서 제일 끝의 노드를 항상 Tail로 유지
	l.Tail.Prev = Prev            // 새로운 Tail의 전노드를 Prev로 기억함
}

func (l *LinkedList) Back() int {
	if l.Tail != nil {
		return l.Tail.Val
	}
	return 0
}

func (l *LinkedList) PopBack() {
	if l.Tail == nil {
		return
	}
	l.RemoveNode(l.Tail)
}

func (l *LinkedList) RemoveNode(node *Node) {
	if node == l.Root {
		l.Root = l.Root.Next
		l.Root.Prev = nil
		node.Next = nil
		return
	}
	Prev := node.Prev // 지우고자하는 로그의 이전으로

	if node == l.Tail {
		Prev.Next = nil
		l.Tail.Prev = nil
		l.Tail = Prev
	} else {
		node.Prev = nil
		Prev.Next = Prev.Next.Next
		Prev.Next.Prev = Prev
	}
	node.Next = nil
}

func (l *LinkedList) PrintNodes() {
	node := l.Root
	for node.Next != nil {
		fmt.Printf("%d ->", node.Val)
		node = node.Next
	}
	fmt.Printf("%d\n", node.Val)
}

func (l *LinkedList) PrintReverse() {
	node := l.Tail         // Tail을 시작점으로
	for node.Prev != nil { //Tail의 이전이 없을때까지 진행
		fmt.Printf("%d ->", node.Val) //Val값 출력
		node = node.Prev              //node를 이전 노드로 초기화
	}
	fmt.Printf("%d\n", node.Val) // 현제 노드 출력
}
