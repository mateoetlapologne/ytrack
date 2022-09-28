package piscine




func ListRemoveIf(l *List, data_ref interface{}) {
	if l.Head == nil {
		return
	}
	for l.Head.Data == data_ref {
		l.Head = l.Head.Next
	}
	for l.Head.Next != nil {
		if l.Head.Next.Data == data_ref {
			l.Head.Next = l.Head.Next.Next
		} else {
			l.Head = l.Head.Next
		}
	}

}