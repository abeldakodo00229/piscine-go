package piscine

// import "fmt"

func ListSize(l *List) int {
	count := 0
	for l.Head != nil {
		count = count + 1
		l.Head = l.Head.Next
	}
	return count
}

/*
type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushFront(l *List, data interface{}) {
	Node := &NodeL{Data: data}
	if l.Head == nil {
		l.Head = Node
		l.Tail = Node
		return
	}
	Node.Next = l.Head
	l.Head = Node
}

func main() {
	link := &List{}

	ListPushFront(link, "Hello")
	ListPushFront(link, "2")
	ListPushFront(link, "you")
	ListPushFront(link, "man")
	ListPushFront(link, "yo")

	fmt.Println(ListSize(link))
}
*/
