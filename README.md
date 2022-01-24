# PriorityQueueGo
The simplest implementation of Priority Queue in Golang.

## Usage

```go

import(
	"fmt"
	priority_queue "github.com/TadaTeruki/PriorityQueueGo/PriorityQueue"
)

func main(){
	var que priority_queue.PriorityQueue

	// 'MakeObject(value interface{}, priority float64) -> Object'
	//    - create a new object
	// '(PriorityQueue).Push(Object)'
	//    - Push object to PriorityQueue
	que.Push(priority_queue.MakeObject("1st", 0.5))
	que.Push(priority_queue.MakeObject("2nd", 2.5))
	que.Push(priority_queue.MakeObject("3rd", 8.5))
	que.Push(priority_queue.MakeObject("4th", 4.5))

	// '(PriorityQueue).GetSize()'
	//    - Query the number of objects
	fmt.Println(que.GetSize()) // -> 4

	// '(PriorityQueue).GetFront()'
	//    - Query the object which has the highest priority
	fmt.Println(que.GetFront().Value, que.GetFront().Priority) // -> "3th", 8.5

	// '(PriorityQueue).PopFront()'
	//    - Remove the object which has the highest priority
	que.PopFront()
	
	fmt.Println(que.GetFront().Value, que.GetFront().Priority) // -> "4th", 4.5
}
```

## Installation
All you need to do is run:<br>

```
 $ go get -u github.com/TadaTeruki/PriorityQueueGo/PriorityQueue
```

## LICENSE

MIT License<br>
Copyright (c) 2022 Tada Teruki (多田 瑛貴)
