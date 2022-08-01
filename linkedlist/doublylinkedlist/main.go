package main

import "fmt"

type Node struct {
	val  int
	next *Node
	prev *Node
}

type DoublyLinkedList struct {
	head *Node
	last *Node
}

func (l *DoublyLinkedList) append(val int) {

	// single node
	temp := new(Node)
	temp.val = val
	temp.next = nil
	temp.prev = nil

	if l.head == nil {
		l.head = temp
		l.last = temp
	} else {
		traverse := l.head;
		for traverse.next!=nil {
			traverse = traverse.next;
		}
		traverse.next = temp;
		temp.prev = traverse;
		l.last.next = temp;
		l.last = l.last.next;
	}
}

func (l *DoublyLinkedList) printFirst() string {
	temp := l.head

	str := ""
	for temp != nil {
		fmt.Printf("%d -> ",temp.val)

		str = str + fmt.Sprintf("%d -> ",temp.val)
		temp = temp.next;
	}
	fmt.Printf("nil")
	str = str + "nil"
	return str;
}

func (l *DoublyLinkedList) printLast() string {
	temp := l.last;

	str := ""

	for temp!=nil {
		fmt.Printf("%d -> ",temp.val)

		str = str + fmt.Sprintf("%d -> ",temp.val)
		temp = temp.prev;
	}
	fmt.Printf("nil")

	str = str + "nil"

	return str
}

func(l *DoublyLinkedList) remove(val int) {
	temp := l.head;

	if(temp!= nil && temp.val==val) {
		if(temp.next != nil) {
		nextNode := l.head.next
		l.head = nextNode;
		nextNode.prev = nil
		} else {
			l.head = nil;
			l.last = nil;
		}
	} else {
		for temp!=nil {
			if temp.val == val {
				if temp.next == nil {
					l.last = l.last.prev;
				}
				remainingNode := temp.next;
				temp.prev.next = remainingNode;
				temp.prev = nil;
				return;
			}
			temp = temp.next;
		}
	}
}

func(l *DoublyLinkedList) findLast() {
	fmt.Println("\n LAST : ",l.last.val)
}

func main() {
	dll := new(DoublyLinkedList)

	dll.append(1)

	dll.append(2)
	dll.append(3)
	dll.append(4)
	dll.append(5)


	fmt.Println("Print First")
	dll.printFirst()

	fmt.Println("\n REMOVE")
	dll.remove(1)

	dll.remove(2)

	 dll.remove(3)

	 dll.remove(5)
	dll.printFirst()

	dll.findLast()

	// fmt.Println("\n Print Last")
	// dll.printLast()
}