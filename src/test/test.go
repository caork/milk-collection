package test

import (
	"fmt"
	"milk-collection/src/routeStructure"
)

func main() {
	line := routeStructure.New()
	line.PushBack(1)
	fmt.Println(line.ToString())
	line.PushBack(2)
	fmt.Println(line.ToString())
	line.PushBack(3)
	fmt.Println(line.ToString())
	line.PushBack(4)
	fmt.Println(line.ToString())
	fmt.Println(line.ToString())
	for i := line.Front(); i != nil; i = i.Next() {
		if i.Value == 3 {
			line.InsertBefore(9, i)
			fmt.Println(line.ToString())
			line.Remove(i.Prev())
			fmt.Println(line.ToString())
		}
	}
	a := line.ToSlice()
	fmt.Println(a)
}
