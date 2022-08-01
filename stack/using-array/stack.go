package main

import (
	"fmt"
)

type Stack struct {
	array []int;
	size int;
	capacity int;
}

func(stack *Stack) init() {
	stack.capacity = 6;
	stack.size = 0;
}

func(stack *Stack) Length() int {
	return stack.size;
}

func(stack *Stack) push(val int) {
	if(stack.size >= stack.capacity) {
		// ask if want to expand!
		fmt.Println("Stack Overflow!")
	} else {
		stack.array = append(stack.array, val);
		stack.size = stack.size + 1;
	}
}


func(stack *Stack) pop() int{
	if(stack.size <= 0) {
		fmt.Println("Stack Underflow!")
		return -1;
	} else {
		index := stack.size -1;
		value := stack.array[index]; // get last index
		stack.array = stack.array[:index] // expect that index!
		stack.size = stack.size - 1;
		return value;
	}
}


func(stack Stack) display() {
	for i:=0;i<stack.size;i++ {
		fmt.Println(stack.array[i])
	}
}

func main()  {

	var s Stack;

	s.init();

	s.push(10);
	s.push(20);
	s.push(30);
	s.push(40);
	s.push(50);
	s.push(60);
	s.pop();
	s.pop();
	s.pop();


	s.display()

}