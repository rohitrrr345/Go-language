package main 

import (
	"fmt"
	"time"
)
  func main(){
	for i:=0;i<=10;i++{
		go func (i int){
			fmt.Println(i)//this of type anonymous function
		}(i)//yagi pe kall kar diya
	}
	time.Sleep(time.Second*2)
  }