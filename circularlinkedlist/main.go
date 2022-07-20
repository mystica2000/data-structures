package main

import "fmt"

type Node struct {
	val  int
	next *Node
}

type CircularLinkedList struct {
	head *Node
}

func (l *CircularLinkedList) append(val int) {

	// creating new Node
	newNode := new(Node)
	newNode.val = val;

	// if head is nil, insert at the beginning
	if(l.head == nil) {
		// head is newNode
		l.head = newNode;
		// newNode's next is head again due to circular linkedlist
		newNode.next = l.head;
	} else {
		// if head is not null, traverse to the end of the list and add it 
		temp := l.head;
		firstNode := l.head

		// if temp's next is first node, then STOP RIGHT THERE! (ONE CIRCLE COMPLETE, DONT REPEAT LOOP)
		for temp.next!=firstNode {
			temp = temp.next;
		}

		// temp has last node now :)
		
		// newNode's Next is head of the list
		// last's node next is now newNode (NEWNODE BECOME LAST NODE \o/)
		newNode.next = l.head;
		temp.next = newNode;
		
	}
}

func(l CircularLinkedList) display() {
	temp := l.head;

	firstNode := l.head.val;

	for temp!= nil {
		fmt.Printf(" %d ->",temp.val)
		temp = temp.next;
		if(firstNode == temp.val) {
			fmt.Println(" nil")
			return;
		}
	}
}

func(l *CircularLinkedList) remove(val int) {

	// if head is nil, STOP!
	if(l.head==nil) {
		return;
	} else if(l.head.val == val) { // if head's val is val then check if next node of head is null or not, 
	    // if null -> head is null
		// if not null -> head point to next node of the head (REMOVED TADA!)
		if(l.head.next!=nil){
		l.head = l.head.next;
	    } else {
			l.head = nil;
		}
	} else { // if head is not first node
		// traverse the list and find out!
		temp := l.head;
		firstNode := l.head;
		var prev *Node;
		// prev ptr helps to keep track of the previous node for deletion
		prev = firstNode;
		for temp.next!=firstNode {
			if(temp.val == val) {
				// prev's node NEXT is current node NEXT! (REMOVING CURRENT NODE CONNECTION BY POINTING PREV NODE'S NEXT TO NEXT NODE OF NODE THAT WE HAVE TO DELETE)
				prev.next = temp.next;
				return;
			}
			// prev node is updated
			prev = temp;
			temp = temp.next;
		}
		// if val is LAST NODE!
		// prev's node next points to firstnode!
		if(temp.val == val) {
			prev.next = firstNode;
		}
	}
}



func main() {
	fmt.Println("Started")

	cll := new(CircularLinkedList)

	cll.append(1);
	
	cll.append(2);
	
	cll.append(3);
	
	cll.append(4);
	
	cll.append(5);

	//cll.remove(1);

	//cll.remove(2);

	cll.remove(5);
	cll.display();
}