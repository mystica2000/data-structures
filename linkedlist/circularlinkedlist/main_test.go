package main

import (
	"testing"
)


func TestAdd(t *testing.T) {
	t.Run("Insertion #0 (Append)",func(t *testing.T) {
		var list CircularLinkedList

		list.append(10)
		list.append(20)
		list.append(30)

		res := list.display()
		expected := "10 -> 20 -> 30 -> nil"

		if(res != expected) {
			t.Error("Got : ",res,"Expected : ",expected)
		} else {
			t.Log("Sucessful")
		}
	});

	t.Run("Insertion/Deletion #0 (Empty)",func(t *testing.T) {
		var list CircularLinkedList

		list.append(10)
		list.append(20)
		list.append(30)

		// list.remove(10)
		// list.remove(20)
		list.remove(30)

		res := list.display()
		expected := "10 -> 20 -> nil"

		if(res != expected) {
			t.Error("Got : ",res,"Expected : ",expected)
		} else {
			t.Log("Sucessful")
		}
	});


	t.Run("Insertion/Deletion #1 (Empty)",func(t *testing.T) {
		var list CircularLinkedList

		list.append(10)
		list.append(20)
		list.append(30)

		list.remove(10)
		list.remove(20)
		list.remove(30)

		res := list.display()
		expected := "nil"

		if(res != expected) {
			t.Error("Got : ",res,"Expected : ",expected)
		} else {
			t.Log("Sucessful")
		}
	});


}