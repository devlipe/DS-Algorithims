package structures

type Node struct {
	Value any
	// This implementation uses a circular doubly linked list structure.
	// This way the root element is just a sentinel. That is, l.root is the prev element of the first element.
	// And also the next element of the last one.
	next *Node
	prev *Node
	list *List
}

//Methods to iterate over the elements of a list

func (n *Node) Next() *Node {
	if p := n.next; n.list != nil && p != &n.list.root {
		return p
	}
	return nil
}

func (n *Node) Prev() *Node {
	if p := n.prev; n.list != nil && p != &n.list.root {
		return p
	}
	return nil
}

type List struct {
	root Node
	len  int
}

// Initialize or empty a list
func (l *List) Init() *List {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

//Creat a new empty list
func New() *List {
	return new(List).Init()
}

//Return the number of elements in the list
func (l *List) Len() int {
	return l.len
}

//Return the first element of the list
func (l *List) Front() *Node {
	if l.len == 0 {
		return nil
	}
	return l.root.next
}

// Return the last element of the list
func (l *List) Back() *Node {
	if l.len == 0 {
		return nil
	}
	return l.root.prev
}

//Initialize lazily a list
func (l *List) lazyInit() {
	if l.root.next == nil {
		l.Init()
	}
}

// Inserts an Node e after element at, increment len, and return e
func (l *List) insert(e, at *Node) *Node {
	e.prev = at
	e.next = at.next
	at.next = e
	e.next.prev = e
	e.list = l
	l.len++
	return e
}

//Insert value v after the node at, also increment len
func (l *List) insertValue(v any, at *Node) *Node {
	return l.insert(&Node{Value: v}, at)
}

// Remove an element e from the list and decrement len
func (l *List) remove(e *Node) {
	e.prev.next = e.next
	e.next.prev = e.prev
	e.next = nil
	e.prev = nil
	e.list = nil
	l.len--
}

// move element E to be the next of at
func (l *List) move(e, at *Node) {
	if e == at {
		return
	}
	// Remove e from his place
	e.prev.next = e.next
	e.next.prev = e.prev

	e.prev = at
	e.next = at.next
	e.prev.next = e
	e.next.prev = e
}

// Remove an element e, if it is contained on the list
// Return the value of e the
func (l *List) Remove(e *Node) any {
	if e.list == l {
		// e.list was initialized when it was inserted on the List
		// if it is from another list, or e.list is nil remove will crash
		l.remove(e)
	}
	return e.Value
}

//PushFront inserts a new element at the front of the list
//It returns the inserted element
func (l *List) PushFront(v any) *Node {
	l.lazyInit() // If the list is empty, make sure that everything is initialized correctly
	return l.insertValue(v, &l.root)
}

//PushBack inserts a new element at the end of the list
//It returns the inserted element
func (l *List) PushBack(v any) *Node {
	l.lazyInit()
	return l.insertValue(v, l.root.prev)
}

//InsertBefore inserts a new element e with value v, before the mark and returns e.
//If mark is not an node of l, the list is not modified
func (l *List) InsertBefore(v any, mark *Node) *Node {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark.prev)
}

//InsertAfter inserts a new element e with value v, after the mark and returns e.
//If mark is not an node of l, the list is not modified
func (l *List) InsertAfter(v any, mark *Node) *Node {
	if mark.list != l {
		return nil
	}
	return l.insertValue(v, mark)
}

// MoveToFront moves element e to the front of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *List) MoveToFront(e *Node) {
	if e.list != l || l.root.next == e {
		return
	}
	l.move(e, &l.root)
}

// MoveToBack moves element e to the back of list l.
// If e is not an element of l, the list is not modified.
// The element must not be nil.
func (l *List) MoveToBack(e *Node) {
	if e.list != l || l.root.prev == e {
		return
	}

	l.move(e, l.root.prev)
}

// MoveBefore moves element e to its new position before mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *List) MoveBefore(e, mark *Node) {
	if e.list != l || e == mark || mark.list != l {
		return
	}

	l.move(e, mark.prev)
}

// MoveAfter moves element e to its new position after mark.
// If e or mark is not an element of l, or e == mark, the list is not modified.
// The element and mark must not be nil.
func (l *List) MoveAfter(e, mark *Node) {
	if e.list != l || e == mark || mark.list != l {
		return
	}

	l.move(e, mark)
}

// PushBackList inserts a copy of another list at the back of list l.
// The lists l and other may be the same. They must not be nil.
func (l *List) PushBackList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next() {
		l.insertValue(e.Value, l.root.prev)
	}
}

// PushFrontList inserts a copy of another list at the front of list l.
// The lists l and other may be the same. They must not be nil.
func (l *List) PushFrontList(other *List) {
	l.lazyInit()
	for i, e := other.Len(), other.Back(); i > 0; i, e = i-1, e.Prev() {
		l.insertValue(e.Value, &l.root)
	}
}
