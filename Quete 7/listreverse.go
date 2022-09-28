package piscine

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListReverse(l *List) {
	if l.Head == nil {
		return
	}
	var prev *NodeL
	current := l.Head
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.Head = prev

}
