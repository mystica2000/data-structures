package main

import (
	"testing"
)

func TestAdd(t *testing.T) {

	t.Run("Insertion Test #1 (Append)", func(t *testing.T) {
		var list LinkedList;
		list.Add(10);
		list.Add(20);
		list.Add(30);

		ret := list.Print();
		expected := "10 -> 20 -> 30 -> nil"

		if(ret != expected) {
			t.Error("Expected:", expected, "Got:", ret)
		} else {
			t.Log("Sucessful")
		}
	});

	t.Run("Insert at First",func(t *testing.T) {
		var list LinkedList;

		list.Add(10);

		ret := list.Print();
		expected := "10 -> nil"

		if(ret != expected) {
			t.Error("Expected:", expected, "Got:", ret)
		} else {
			t.Log("Sucessful")
		}
	});

	t.Run("Length of the Linked List",func(t *testing.T) {
		var list LinkedList;

		list.Add(10);
		list.Add(20);
		list.Add(30);
		list.Add(40);
		list.Add(50);

		ret := list.Length();
		expected := 5


		if(ret != expected) {
			t.Error("Expected:", expected, "Got:", ret)
		} else {
			t.Log("Sucessful")
		}
	})

}

func TestDelete(t *testing.T) {
	t.Run("Deletion Test #0 (Deletion in Middle)", func(t *testing.T) {
		var list LinkedList;
		list.Add(10);
		list.Add(20);
		list.Add(30);

		list.Delete(20)

		ret := list.Print();
		expected := "10 -> 30 -> nil"

		if(ret != expected) {
			t.Error("Expected:", expected, "Got:", ret)
		} else {
			t.Log("Sucessful")
		}
	});

	t.Run("Deletion Test #1 (Delete All)", func(t *testing.T) {
		var list LinkedList;
		list.Add(10);
		list.Add(20);
		list.Add(30);

		list.Delete(20)
		list.Delete(30)
		list.Delete(10)

		ret := list.Print();
		expected := "nil"

		if(ret != expected) {
			t.Error("Expected:", expected, "Got:", ret)
		} else {
			t.Log("Sucessful")
		}
	});

	t.Run("Deletion Test #3 (Delete first and last)", func(t *testing.T) {
		var list LinkedList;
		list.Add(10);
		list.Add(20);
		list.Add(30);

		list.Delete(30)
		list.Delete(10)

		ret := list.Print();
		expected := "20 -> nil"

		if(ret != expected) {
			t.Error("Expected:", expected, "Got:", ret)
		} else {
			t.Log("Sucessful")
		}
	});

}


func TestSort(t *testing.T) {

	t.Run("Sorting Test #1 (Sort)", func(t *testing.T) {
		var list LinkedList;

		list.Add(24)
		list.Add(222)
		list.Add(10)
		list.Add(99)

		list.head = list.Sort()

		ret := list.Print()
		expected := "10 -> 24 -> 99 -> 222 -> nil"

		if(ret != expected) {
			t.Error("Expected:", expected, "Got:", ret)
			} else {
				t.Log("Sucessful")
			}
	});

}