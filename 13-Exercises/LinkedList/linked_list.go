package LinkedList

import "fmt"

type datatype string;

type LinkedList struct {
	head *Node
	length int
}

type Node struct {
	data datatype
	next *Node
}

func (l *LinkedList) Print() {
	current := l.head
	for {
		fmt.Printf("%s",current.data);
		if current.next == nil {
			fmt.Printf("\n")
			break;
		} else {
			current = current.next;
			fmt.Printf(" -> ")
		}
	}
}

func (l *LinkedList) Append(newData datatype) {

	node_to_be_added := Node {
			data: newData,
	};

	if l.head == nil { // Linked List is empty.
		l.head = &node_to_be_added
	} else {
		current := l.head;
		for current.next != nil {
			current = current.next
		}
		current.next = &node_to_be_added
	}

	l.length = l.length+1
}


func (l *LinkedList) Find(value datatype) bool{
	// If list is empty no point searching.
	if l.head == nil {
		return false;
	}

	current := l.head;
	found := false
	for {
		if current.data == value {
			found = true;
			break;
		} else {
			if current.next != nil {
				current = current.next;

			} else {
				break;
			}
		}
	}

	return found;
}

func (l *LinkedList) Pop() *Node {
	// If list is empty no point searching.
	if l.head == nil {
		return nil;
	}

	var current *Node = l.head;
	var prev  *Node = nil;
	for current.next != nil {
		prev = current;
		current = current.next;
	}
	prev.next = nil;
	return current;
}