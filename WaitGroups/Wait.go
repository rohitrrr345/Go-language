package main

import (
	"fmt"
	// "sync"
)

// func task(id int, W *sync.WaitGroup) {
// 	defer W.Done()//it is like a cleanUp function over there
// 	fmt.Println("doing task", id)
// }
func printMe[T any](nums[]T){
	fmt.Println(nums)
}

func main() {
	// var wg sync.WaitGroup
	// for i := 0; i <= 10; i++ {
	// 	wg.Add(1)
	// 	go task(i, &wg)
	// }
	// wg.Wait()
      nums:=[]int{1,23,44};
    printMe(nums)
	
}
