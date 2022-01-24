
package main

import(
	"fmt"
	priority_queue "./PriorityQueue"
)

func main(){
	var que priority_queue.PriorityQueue
	
	que.Push(priority_queue.MakeObject("1st", 0.5))
	que.Push(priority_queue.MakeObject("2nd", 2.5))
	que.Push(priority_queue.MakeObject("3rd", 8.5))
	que.Push(priority_queue.MakeObject("4th", 4.5))

	fmt.Println(que.GetSize()) 								   // -> 4
	fmt.Println(que.GetFront().Value, que.GetFront().Priority) // -> "3th", 8.5

	que.PopFront()
	fmt.Println(que.GetFront().Value, que.GetFront().Priority) // -> "4th", 4.5
}