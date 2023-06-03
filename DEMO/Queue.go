package main

import "fmt"

type Queue struct {
	que []interface{}
}

func Add(q *Queue, value interface{}) {
	q.que = append(q.que, value)
}

func Remove(q *Queue) interface{} {
	if IsEmpty(q) {
		fmt.Println("QueueEmptyException")
		return nil
	}
	value := q.que[0]
	q.que = q.que[1:]
	return value
}

func RemoveBack(q *Queue) interface{} {
	if IsEmpty(q) {
		fmt.Println("QueueEmptyException")
		return nil
	}
	n := len(q.que)
	value := q.que[n-1]
	q.que = q.que[:n-1]
	return value
}

func Front(q *Queue) interface{} {
	if IsEmpty(q) {
		fmt.Println("QueueEmptyException")
		return nil
	}
	return q.que[0]
}

func Back(q *Queue) interface{} {
	if IsEmpty(q) {
		fmt.Println("QueueEmptyException")
		return nil
	}
	n := len(q.que)
	return q.que[n-1]
}

func IsEmpty(q *Queue) bool {
	return len(q.que) == 0
}

func Len(q *Queue) int {
	return len(q.que)
}

func Print(q *Queue) {
	for _, value := range q.que {
		fmt.Print(value, " ")
	}
	fmt.Println()
}

func main() {
	que := &Queue{}
	Add(que, 1)
	Add(que, 2)
	Add(que, 3)
	fmt.Println("IsEmpty:", IsEmpty(que))
	fmt.Println("Length:", Len(que))
	fmt.Println("Queue remove:", Remove(que))
	fmt.Println("Queue remove:", Remove(que))
}

/*
IsEmpty: false
Length: 3
Queue remove: 1
Queue remove: 2
*/
