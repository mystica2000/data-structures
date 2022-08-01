package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	t.Run("Insertion #0 (Append)",func(t *testing.T) {
		var list DoublyLinkedList

		list.append(10)
		list.append(20)
		list.append(30)

		res := list.printFirst()
		expected := "10 -> 20 -> 30 -> nil"

		if(res != expected) {
			t.Error("Got : ",res,"Expected : ",expected)
		} else {
			t.Log("Sucessful")
		}

	})
	t.Run("Insertion #1 (Empty)",func(t *testing.T) {
		var list DoublyLinkedList

		res := list.printFirst()
		expected := "nil"

		if(res != expected) {
			t.Error("Got : ",res,"Expected : ",expected)
		} else {
			t.Log("Sucessful")
		}
	});
}

func TestDelete(t *testing.T) {
	t.Run("Deletion #0 (Empty)",func(t *testing.T) {
		var list DoublyLinkedList

		list.append(10)
		list.remove(10)

		res := list.printFirst()
		expected := "nil"

		if(res != expected) {
			t.Error("Got : ",res,"Expected : ",expected)
		} else {
			t.Log("Sucessful")
		}
	});

	t.Run("Deletion #1 (Remove from last,first)",func(t *testing.T) {
		var list DoublyLinkedList

		list.append(10)
		list.remove(10)
		list.append(20)
		list.append(30)
		list.remove(30)

		res := list.printFirst()
		expected := "20 -> nil"

		if(res != expected) {
			t.Error("Got : ",res,"Expected : ",expected)
		} else {
			t.Log("Sucessful")
		}
	});

	t.Run("Deletion #1 (Remove from last,first)",func(t *testing.T) {
		var list DoublyLinkedList

		list.append(10)
		list.remove(10)
		list.append(20)
		list.append(30)
		list.remove(30)
		list.append(60)

		res := list.printLast()
		expected := "60 -> 20 -> nil"

		if(res != expected) {
			t.Error("Got : ",res,"Expected : ",expected)
		} else {
			t.Log("Sucessful")
		}
	});


}