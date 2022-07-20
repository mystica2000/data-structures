package main

import (
	"fmt"
)

type Node struct{
	val int;
	next *Node;
}

type LinkedList struct {
	head *Node;
}

func(l *LinkedList) add(val int) {
	temp := &Node{val:val}
	if(l.head == nil) {
		l.head = temp
	}	else {
		iterator := l.head;
		for iterator.next!=nil {
			iterator = iterator.next;
		}
		iterator.next = temp;
	}
}


func(l *LinkedList) delete(val int) {
	
	var previous *Node;

	if(l.head.val == val) { // if first value 
		l.head = l.head.next;
		return;
	}

	for iterator := l.head;iterator!=nil;iterator=iterator.next {
		if(iterator.val == val) {
			previous.next = iterator.next;
			return;
		}
		previous = iterator
	}
}

func(l LinkedList) print() {
	iterator := l.head;

	fmt.Println()
	for iterator!=nil {
		fmt.Printf("%d -> ",iterator.val);
		iterator = iterator.next;
	}
	fmt.Printf("nil")
}

func(l LinkedList) length() int{
	iterator := l.head;

	len := 0

	for iterator.next!=nil {
		len += 1;
		iterator = iterator.next;
	}

	return len;
}

func(l *LinkedList) sort() *Node{
   return l.sortListHelper(l.head)
}


func(l *LinkedList) sortListHelper(head *Node) *Node {
	if(head == nil || head.next == nil) {
		return head;
	}

	midPointOfList := l.findMidPoint(head)

	firstHalf := midPointOfList
	secondHalf := midPointOfList.next;

	firstHalf.next = nil

	sortFirstHalf := l.sortListHelper(head)
	sortSecondHalf := l.sortListHelper(secondHalf)

	return mergeSortedLists(sortFirstHalf,sortSecondHalf)
}


func mergeSortedLists(first *Node,second *Node) *Node {
	if (first == nil) {
		return second
	}
	if (second == nil) {
		return first
	}

	temp := new(Node);

	if( first.val < second.val) {
		temp = first;
		temp.next = mergeSortedLists(first.next,second)
	} else {
		temp = second;
		temp.next = mergeSortedLists(first,second.next)
	}
	return temp;
}


func(l *LinkedList) findMidPoint(head *Node) *Node{
	
	if (head == nil || head.next == nil) { // if linked list contains one or two or empty elements
		return head;
	}

	fastPtr := head;
	slowPtr := head;

	for (fastPtr!=nil && fastPtr.next!=nil) {
		fastPtr = fastPtr.next;

		if(fastPtr.next==nil) {
			break;
		}

		fastPtr = fastPtr.next; // move two steps 
		slowPtr = slowPtr.next; // move one step
	}
	return slowPtr;
}


func main() {
	var list LinkedList;

	list.add(10);
	list.add(20);
	list.add(70);
	
	list.add(30);
	list.add(40);
	list.add(50);
	list.add(60);

	list.print()

	fmt.Println("\n Length of the LIST ",list.length())	


	list.head = list.sort()

	list.print()
}