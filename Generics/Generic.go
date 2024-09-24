package main
 import "fmt"
   func printSlice[T interface {}](nums[]T){
	for i:=0;i<len(nums);i++{
		fmt.Println(nums[i])
	}
   }
   func main(){
	nums:=[]int{12,3,5,5};
	printSlice(nums)
   }